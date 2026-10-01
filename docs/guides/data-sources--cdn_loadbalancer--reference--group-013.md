---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1200202211000112-2011330202212030-2030011012322022-0011001011122102-2031231211302132-0131120203303111-2110110200233200-2203102032302312"></a>

## policy_based_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 021303230222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-1202102332031323-1330320321333001-0031120030210322-3301301222213322-3211233302331113-1300322100310320-1301200233133012-0111301013123211"></a>

Type: `"single"`. Computed.

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

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0230120222332101-2200031301121121-3302002112001312-0202031323333111-1113210120333033-1000031303320200-3010202302032313-3123203031223321"></a>

## Direct properties — captcha_challenge_parameters / 021303230222 / 3

<a id="canonical-3320311233023302-3331230210313032-0213102132023302-2003132000213032-3001333031101111-2020132022230300-0111213221220203-0000023022312030"></a>

<a id="canonical-2002013303021012-2110232232331302-1013120331330301-1212112212223222-0013010020322201-3211312133022002-2322012322033021-2103300132010211"></a>

## cookie_expiry property — captcha_challenge_parameters / 021303230222 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

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

<a id="canonical-0212021211312100-1100023311213113-1221131021022203-3303313200113323-1232001111322302-3101311010000111-1303133132121123-2001300023100112"></a>

<a id="canonical-0320013101021110-1012321120031101-1313213310203210-2213223001112132-0230213332300331-2101102321313011-3031223103132030-2313002020021123"></a>

## custom_page property — captcha_challenge_parameters / 021303230222 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
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

<a id="canonical-0133321111212221-2132233200031012-2001023222000032-0123222103323021-0311113111003103-0132031013222101-0002220201033233-0131110101033333"></a>

## Next pages — captcha_challenge_parameters / 021303230222 / 6

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0331321120033113-3101221012002312-3022320202232320-2310011132202011-2111232120123322-3010030202321212-2302113320000000-1203031132231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121021021030003-1311000112220120-3002032213220232-0133220212123100-3021023203310210-2013012322102013-0123203320330111-1233303023121103"></a>

## policy_based_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 201332122011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-3301003200022032-2032000102302220-2101112003311033-3303123001031213-0030330213003310-2210132213213301-0322121231131331-0033012033120132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-1200231010211003-0300302232301021-2223103311010321-3323033110021221-1231313132013302-0331201211203311-1133100320310301-1102131003222330"></a>

## Direct properties — default_captcha_challenge_parameters / 201332122011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110111122001331-3031102021111130-1303112302111132-2332021222121131-0113100030201112-2102311310322233-1300113122232223-2011131213202111"></a>

## Next pages — default_captcha_challenge_parameters / 201332122011 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3020030331320002-2323312103120222-0233230333010101-0021201002220312-1100120123133312-1223333330123232-0330210121202121-2301300222123300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210212300132211-3110330232322123-2022322300211032-1032111112001323-1302321013301021-3131133230313112-0122103110101012-2302213031120221"></a>

## policy_based_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 202101113000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-3301310220233332-0021123312313000-2110201301010023-0211321120321130-2122203301323000-0012333000301230-0001000132101010-3110323010011010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-2101023011130202-2102203312302202-3121212110000032-1101011213132011-2333102303222112-3012302113022133-2033200023303203-1133312210210223"></a>

## Direct properties — default_js_challenge_parameters / 202101113000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002322313223333-0013320311211310-1211021123010030-2130011323331033-1032330231320310-2102030033003332-3322002130002000-0031031200132102"></a>

## Next pages — default_js_challenge_parameters / 202101113000 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0013310103300131-0320102103313301-3110201221302003-0201101332013301-0110221332202122-0200223010132002-0110231220302001-0212121033102102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012113121032230-3101310113302012-1111203001101233-3300312211210300-0020202010020203-1333232130201132-2200000010232222-3310133000011213"></a>

## policy_based_challenge.default_mitigation_settings — default_mitigation_settings / 002212032210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0301320321221233-2133022123012021-0201133310221200-0001021310002132-2211002101110103-1303032313213203-0110333003003220-1210020022003000"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3022233220121113-3322321002100003-1330311112133120-2200233110213200-3321230220301312-3101333103002132-3113113230001003-2222310331303322"></a>

## Direct properties — default_mitigation_settings / 002212032210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111022003202321-1022013112030321-1130120100003131-1013030131102121-3132310123001101-2030221320133100-3230133211013222-0220200130033013"></a>

## Next pages — default_mitigation_settings / 002212032210 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2000332003010103-1332232202222102-3023131022123002-3112110001311302-3301003202012101-0233023131213231-2121322301033201-2313020222321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030223232323323-1233230101130002-1101002213110000-1030303223110213-3330312213113312-1321100221301020-1121102102323300-0213100023201021"></a>

## policy_based_challenge.default_temporary_blocking_parameters — default_temporary_blocking_parameters / 111021312221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-2032010310013323-1323020010330113-3301100320121200-1023002032023332-1130310323311101-0312132211312103-1333323101110003-1310022002321220"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0221011002210213-0200002111001111-0302333101103321-0223011012321220-2001101021200321-1200000121013333-2312032331321223-0100321012022102"></a>

## Direct properties — default_temporary_blocking_parameters / 111021312221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201212233113022-3132121301201212-3133230321301132-3121221033131310-1130001321332322-0321103332320233-2132312122031200-1133300031131030"></a>

## Next pages — default_temporary_blocking_parameters / 111021312221 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3312332033330310-2102020113221113-3101100321233023-1112230300001110-2201122230122202-3102111101333130-1332220233232332-2003002330003022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233322102330101-0002102120233211-1321000002232313-0211123031010112-3330030223331320-2321211003010213-1220311211120233-1303030121103203"></a>

## policy_based_challenge.js_challenge_parameters — js_challenge_parameters / 222111131203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-3121021032220013-0130132120022221-2303303022110331-2203112213000131-2100303321303211-2323312000120330-3302100102132311-3313120120313033"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript. With this feature enabled, only clients that are capable of executing Javascript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
Javascript.

With this feature enabled, only clients that are capable of executing Javascript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do Javascript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have Javascript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the Javascript. Javascript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid Javascript challenge for subsequent requests.

Javascript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running Javascript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either Javascript challenge or Captcha challenge on a virtual host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213331111102220-2233031012222111-0030012130332130-1111110300203103-0013022011223220-3323320231322333-3013213011311023-3132020311131021"></a>

## Direct properties — js_challenge_parameters / 222111131203 / 3

<a id="canonical-2120001301233212-0122030013212213-2101031123131333-1313000113000033-3233303223203231-3322301123303021-3101131211100131-2220010301113112"></a>

<a id="canonical-1221212201021200-1120332332030112-3000331123031110-1120320321003112-2320210101033302-0110320201321021-1212311010101320-0123211210000113"></a>

## cookie_expiry property — js_challenge_parameters / 222111131203 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

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

<a id="canonical-1123010333130110-0101300133323130-1032201011011112-2200203212103010-0333032020203332-0012300203011330-1000022130133201-1122012203210323"></a>

<a id="canonical-3031113312332203-3102323333332200-2131103033310132-3001210013211022-0302101300013223-3001211022100333-0202032320001233-3323212000220033"></a>

## custom_page property — js_challenge_parameters / 222111131203 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
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

<a id="canonical-1331012310130122-3302221222321310-2122202311001111-2222011023020211-0200233031103320-3022002130131213-1211002030131330-2302101131122032"></a>

<a id="canonical-0101120323301103-2121232321200302-1113312123233100-0211131222033110-1312330020231003-2203212012310311-3223103013120323-3212133011232210"></a>

## js_script_delay property — js_challenge_parameters / 222111131203 / 6

Type: `"number"`. Computed.

Delay introduced by Javascript, in milliseconds.

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

<a id="canonical-2132010101112223-3003100120202330-3132122113322300-3113300110012020-3333300220202312-0202011302110111-0013301233213203-0032021211302130"></a>

## Next pages — js_challenge_parameters / 222111131203 / 7

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1330002230033301-1101210211330033-3222123112310332-2120020120220110-1331221313302032-0233330300030313-2222232003232031-2302002013120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222030000113123-2231322311313110-2022222303310211-0203123031322303-2203112232021232-2131232113133203-1230001022100001-3231230010303130"></a>

## policy_based_challenge.malicious_user_mitigation — malicious_user_mitigation / 312332112022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-3023133122132021-1110222220201021-1021200213121201-1221020133203031-2221120103120231-0202012302030201-2220101100200212-1222013303132232"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2212001100310010-0010300320131023-0103330020223132-1232301301303333-0221311010010133-0302121231003002-3330210301320233-2333300011313302"></a>

## Direct properties — malicious_user_mitigation / 312332112022 / 3

<a id="canonical-0333232023202303-0023223002102012-1333121200010031-2233200322222200-0311010132012332-1023010022331330-0113101330230323-2001010301222233"></a>

<a id="canonical-2320012013002303-3012113220222001-2020313303330020-1131002310312232-3310030012231221-3222100210032001-0210022032220122-3310321010333321"></a>

## name property — malicious_user_mitigation / 312332112022 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1002301203001100-2310332101303203-2201021230211000-2231233123212300-2202211130233230-2233100231321102-3332221130011220-3010022031031003"></a>

<a id="canonical-1121212030031113-2331020021220320-2220320210320330-2303012302020001-2023201023231233-3231130121220101-3012220321020012-2233123012122111"></a>

## namespace property — malicious_user_mitigation / 312332112022 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0010310210103133-2200213233011303-2222230001033013-0321122323310230-2202100111330003-2321021210323331-1223312200122013-1210332022210011"></a>

<a id="canonical-1021021103211100-0123302312321033-2013230221002022-0211323111333321-1033130201110023-1201133000301133-2231303023330003-2310121331022302"></a>

## tenant property — malicious_user_mitigation / 312332112022 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0103212330030221-1322321332331222-0100310212111222-1202103211020120-0122302003320322-2003122311022300-3110202120203020-2130322021221233"></a>

## Next pages — malicious_user_mitigation / 312332112022 / 7

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2321003023230011-0001310133230312-1321311122223130-3311300232130110-1231301132210323-0102103110013003-3132232311112002-2223233200201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330103330010001-3232000122121122-2011033011200030-3232022001133123-0010113121132021-2310311302333100-1021220223231222-0021003122212132"></a>

## policy_based_challenge.no_challenge — no_challenge / 113023132020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.no_challenge

<a id="canonical-2322021002120111-2300110033002302-2213223302123232-3030202321013200-0002210130300231-2330202002130220-0210131023023112-0322020021232013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-3233313313231310-3021001000330202-2201112213302002-3013221110121233-3203122002203001-3333311123113331-0301203112201110-0133001100133012"></a>

## Direct properties — no_challenge / 113023132020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030121030321020-1321021221320231-1003203230001011-2221113120211020-0001122303120002-3021201200113110-1100031013332120-2230112201113311"></a>

## Next pages — no_challenge / 113023132020 / 4

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322112120201113-0020023102232233-2000303122003332-0333023130220322-2003011331020010-1020030120101133-0031230203320320-0233202123000223"></a>

## policy_based_challenge.rule_list — rule_list / 330220033101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
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

<a id="canonical-2022132222310122-3103132123113220-1101031333212020-1213031301211001-1031322002212332-2301130332020101-0100032002011231-1022013032130022"></a>

## Direct properties — rule_list / 330220033101 / 3

- [rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302): complete subsection reference.

<a id="canonical-2012032103323322-0113232321120003-1001031023002311-0120110131033220-1303322231102100-2201033212003102-3031003001333201-2200202203310131"></a>

## Next pages — rule_list / 330220033101 / 4

- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230322333322232-0331123123200203-0003101321000322-3012213210323313-3001310231131220-2320022222131302-3331202123313120-0220323201230233"></a>

## policy_based_challenge.rule_list.rules — rules / 233201210311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- policy_based_challenge.rule_list.rules

<a id="canonical-2032201021203200-2112232030123112-0321223303103001-1223332133322120-0111331112221013-2213030123122132-3232130031322130-2123210222232202"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3120033202331332-3330132112031112-1013120110032033-2312103223022120-2230030120103231-1303101012231233-2231322001111101-1322203020000320"></a>

## Direct properties — rules / 233201210311 / 3

- [metadata](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2213332232012120-0000232121012231-1202321003100000-2023033121213332-1312100012202001-3200303132103321-3203131011202132-0023232002221101): complete subsection reference.

- [spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230): complete subsection reference.

<a id="canonical-2203330212033100-2113321030012230-3000203110020312-1330102002131030-1332110100201200-3111301000312323-3303121011211032-3100123000211003"></a>

## Next pages — rules / 233201210311 / 4

- [policy_based_challenge.rule_list.rules.metadata](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2213332232012120-0000232121012231-1202321003100000-2023033121213332-1312100012202001-3200303132103321-3203131011202132-0023232002221101)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2213332232012120-0000232121012231-1202321003100000-2023033121213332-1312100012202001-3200303132103321-3203131011202132-0023232002221101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101312313333311-2103321320330201-2223102201103101-1231323033012030-1120321002302113-2210132033221300-1103122022110222-0112222200033332"></a>

## policy_based_challenge.rule_list.rules.metadata — metadata / 313201121200 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-1023020111301133-1000120113032133-0023001330320023-3000130132123213-3220112312032021-0200012212013222-3010201001221130-3333322001212100"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

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

<a id="canonical-3223100233002231-3320100331131210-3202232100011010-3011130213201123-1333300122331230-3310213111303122-1203101200322033-2120023233122221"></a>

## Direct properties — metadata / 313201121200 / 3

<a id="canonical-3301103031112033-2231100111121212-2122313033131122-1112223220011011-2230113313012022-2123300121022300-0311031112121100-1322220233132003"></a>

<a id="canonical-1121033202110200-3133112012313032-1102201223011213-2123100312203000-2322211223223012-3130022030232330-2303220231331300-1103231032110311"></a>

## description_spec property — metadata / 313201121200 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0302210100211301-2001000223102232-2320300133211233-3112300023302310-3003222130312231-1331123031113210-1323010032030202-0213213013302003"></a>

<a id="canonical-0320200201320220-0312323110230202-3003021121023323-2101201003022211-0323000213201201-1231001000201202-1032031232013122-3103121100012303"></a>

## name property — metadata / 313201121200 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

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

<a id="canonical-2230203331201132-2010002212200120-2030122122123313-3100220231221123-2303331121101012-0103203332023100-0332102002103322-0301032131121310"></a>

## Next pages — metadata / 313201121200 / 6

- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123013133221212-1103020002223223-2210023212312313-1133222221222323-3322302023210202-1330232213202100-0031101010303313-1021102213130231"></a>

## policy_based_challenge.rule_list.rules.spec — spec / 011210123213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-1022213103032223-0211022310130111-0122230103313203-0002321332030333-3000332002010300-2221301133113323-2101103223330032-1130131310222120"></a>

Type: `"single"`. Computed.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

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

<a id="canonical-3102301011323313-1103230221121323-2122022101331030-0220000123212023-0030203012020322-3220221333002321-2202220110001031-1130230210112302"></a>

## Direct properties — spec / 011210123213 / 3

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

- [enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3011132030011201-2211311232213300-0312302301232121-3330311310211332-1123332021213002-1232230220130032-1231303030102230-1132313121201323): complete subsection reference.

- [enable_javascript_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2310032320333030-1003211003002112-2220030321201321-2312212330103131-3102112121111313-0103301022112213-2332123333231022-0233120222133023): complete subsection reference.

<a id="canonical-3322130120002003-0302020013310232-2201122102313021-3022030323322023-0013132200113301-1010212321333223-1101220220021010-2010131211013323"></a>

<a id="canonical-3002331020000233-1331203022320013-1121131231130103-0113011331222000-3101130211130222-2231311212122320-0223213212212122-1300013131210112"></a>

## expiration_timestamp property — spec / 011210123213 / 4

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

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

- [headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123): complete subsection reference.

- [http_method](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0203122003302010-0311301031311231-1113123010123300-1231200102112212-0111013121233330-0102113203321103-2323133133131301-0110321113200323): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0212033022130233-1300022103102203-3221233211032110-0022200332332213-3121222133232013-0122231011020300-3311100313212031-0300232013311002): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1323200003020201-0013112330130332-0230323022322321-2231311020200133-1312320220310220-1302221121320212-2103002221230003-0302010321213122): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2311200300102223-2200011111023033-1303211002130001-2221330223002003-2133120032110001-1030331310222102-1303231310132020-0311310011103102): complete subsection reference.

<a id="canonical-2313022031313113-0101103103003130-3022020320013113-0100332122231211-1003001111220102-1120310210121312-3232112322110133-2030023133231100"></a>

## Next pages — spec / 011210123213 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1211213323333331-3123231310232023-0102332112110100-2011112333113110-1210213012301101-0233123203311333-0212000332013310-0211301122011232)
- [policy_based_challenge.rule_list.rules.spec.any_client](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3001032210121312-2232021320102331-2221130312230313-0322032201301033-0313120033310020-1120133211031201-2201030012031023-2021320321112111)
- [policy_based_challenge.rule_list.rules.spec.any_ip](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2132021231130112-2123321303233202-1003030103123331-3220221301231301-0320121202232322-3113302121103110-3122011123303333-3113322220020022)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- [policy_based_challenge.rule_list.rules.spec.asn_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3030233320003122-1011021332000221-3203021022013102-2130330003322122-1203302321002323-3211210201111313-0012132222031031-2310010203120022)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1213311021203302-2332101131201022-1122221301221222-2300200300120230-2022132311130021-0100323233103233-1000133330002011-1123331220122111)
- [policy_based_challenge.rule_list.rules.spec.client_selector](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3102120111001203-1112131330132320-2121010233221012-2320103020332010-2113223021322212-2311332300120013-1311000030211222-2132213323012211)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3333031012000322-3022131233200332-1222200332232033-2210120021231332-1023232322313002-0133223103121211-0003110221123323-0313331331022230)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1213313313333103-1013220331032213-2322322203300120-0232232023222301-1110333033030011-2120302233312122-3102011002331023-2102012213103220)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3011132030011201-2211311232213300-0312302301232121-3330311310211332-1123332021213002-1232230220130032-1231303030102230-1132313121201323)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2310032320333030-1003211003002112-2220030321201321-2312212330103131-3102112121111313-0103301022112213-2332123333231022-0233120222133023)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- [policy_based_challenge.rule_list.rules.spec.http_method](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0203122003302010-0311301031311231-1113123010123300-1231200102112212-0111013121233330-0102113203321103-2323133133131301-0110321113200323)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0212033022130233-1300022103102203-3221233211032110-0022200332332213-3121222133232013-0122231011020300-3311100313212031-0300232013311002)
- [policy_based_challenge.rule_list.rules.spec.path](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1323200003020201-0013112330130332-0230323022322321-2231311020200133-1312320220310220-1302221121320212-2103002221230003-0302010321213122)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2311200300102223-2200011111023033-1303211002130001-2221330223002003-2133120032110001-1030331310222102-1303231310132020-0311310011103102)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1211213323333331-3123231310232023-0102332112110100-2011112333113110-1210213012301101-0233123203311333-0212000332013310-0211301122011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220310122112211-0121311322030110-3111012331103201-1232320000123103-1333002101300031-1010330223002032-3021311021101110-1122101001030313"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — any_asn / 330021213131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2022321201311001-2022123020113302-3222131231102021-1311102002333213-1123131210201310-2002202123113222-2310103212232222-3200032111122113"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3322032301133221-3111033110212312-2222332320330133-2230232100332032-2133233320003032-0101200100331201-0211033132003232-0313321223320303"></a>

## Direct properties — any_asn / 330021213131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002201111111003-1313131302102033-3302102033130323-0101200001120011-1232033000320333-0313100000131310-1202021300112231-2313022301231103"></a>

## Next pages — any_asn / 330021213131 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3001032210121312-2232021320102331-2221130312230313-0322032201301033-0313120033310020-1120133211031201-2201030012031023-2021320321112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112113200012330-3332113213200131-0222030130011303-1223332131330122-2323110200030120-0012122213332310-3221021300230211-1023223231020022"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — any_client / 100333032303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-2113323021232020-3321112213211132-1033001002002323-0122101202032233-1333011303320102-3113112231101102-3331313321200021-2100302020232202"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1030000232210021-3222011223033301-3331120330210120-3302103333232023-3213001012000312-1020020233211101-3001210001220010-2130210301120221"></a>

## Direct properties — any_client / 100333032303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300110003113313-2010020210131301-1133000022333112-3302332200222032-2333201011032312-2012110121032001-2003010313011003-1220332310223223"></a>

## Next pages — any_client / 100333032303 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2132021231130112-2123321303233202-1003030103123331-3220221301231301-0320121202232322-3113302121103110-3122011123303333-3113322220020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121310201321132-1010120231133221-3032000010022300-0023032223310000-2223131002200123-0201331003231003-1203112030202233-1021310131230021"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — any_ip / 100011222311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0101310011002212-2011232231303331-2233122211023322-3303222221312010-2112211320102003-2021103320312033-0212212020003212-1230213221000303"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3132001030032223-0303301312301231-0100021021122001-1202311223232030-2021211101220012-2200220330123003-1210233211131113-0333203320002331"></a>

## Direct properties — any_ip / 100011222311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231020200021032-0222031132100332-0110002131232113-2013210011311032-2112130023033110-0212321012230322-2020212033030110-2320021300313133"></a>

## Next pages — any_ip / 100011222311 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221113321110313-1200203021321201-0000330031210231-1113122002301000-0001033332230312-0010121023023132-0131320333333012-1011110200000012"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — arg_matchers / 112331112222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1322122122110222-3110303013132013-1303330030102120-3013132233113310-1323130213222200-0230320203320222-2333303000002312-2123130000333031"></a>

Type: `"list"`. Computed.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0203033112012222-3003333303210021-1233222220303111-3320111301012220-2132211011313112-3122213023300011-2230103220003300-1032323021023012"></a>

## Direct properties — arg_matchers / 112331112222 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0011210200110321-3133122231213301-1222102311231331-1211320002312231-0001120223300232-0013112021212230-1102312123120102-1333022202101021): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1132333220020301-0313333333301101-2303331233003010-1203001113201310-0223011021030001-0030121210122100-3102120300121022-1011013211233201): complete subsection reference.

<a id="canonical-2202113131132011-0322010331121202-1310230132013023-2110003320323213-3323302033321223-3220021310313120-1021003122303001-2202210322000230"></a>

<a id="canonical-1220313122331220-0301130332300001-3002012322021332-2132032030033020-3201013303220112-0220023330033202-1231002102231200-1331310100221223"></a>

## invert_matcher property — arg_matchers / 112331112222 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1300312210010213-0220212233332333-3212110021303202-2130023133030331-2132121201122131-0021131203003031-2133110123220101-1110010200223001"></a>

## name property — arg_matchers / 112331112222 / 5

Type: `"string"`. Computed.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

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

<a id="canonical-2222220021001101-3231300123021002-0112031332002332-2123313231200112-3020001232313032-1231023202011331-2213133000031230-2302213200121130"></a>

## Next pages — arg_matchers / 112331112222 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0011210200110321-3133122231213301-1222102311231331-1211320002312231-0001120223300232-0013112021212230-1102312123120102-1333022202101021)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1132333220020301-0313333333301101-2303331233003010-1203001113201310-0223011021030001-0030121210122100-3102120300121022-1011013211233201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2222320022131003-0333230101200030-1001003332011102-2033322112101302-0231021300332122-3031123223220122-0331130110231333-3031211122333100)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0011210200110321-3133122231213301-1222102311231331-1211320002312231-0001120223300232-0013112021212230-1102312123120102-1333022202101021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102220123100311-1113321111131122-2121222000212333-1110312222021110-0301333033332303-1110112312112032-0213110332032101-3101311300220313"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — check_not_present / 133221310132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-0122102322213210-0302210001013330-3232303020001213-3223131200122111-2310012020113021-0230230001121003-3030320033020330-0111113300333202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1102313011112131-0312233110021203-2030133333020202-2333331023212013-3232003033023332-2201312122130230-1013321031110103-2323000032133321"></a>

## Direct properties — check_not_present / 133221310132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023232003023322-3011000311301002-3313323110011000-0230332322030200-3111333130101113-2303213123100313-0222203033003200-0131011303021223"></a>

## Next pages — check_not_present / 133221310132 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1132333220020301-0313333333301101-2303331233003010-1203001113201310-0223011021030001-0030121210122100-3102120300121022-1011013211233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310022100012331-2130003220230202-3022221030200313-3013313030312132-2303323020033002-0313102122010303-2210220311322211-0330311213130031"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — check_present / 213102232300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0321323202123303-1030332330031132-3332132333031201-0303220001303213-2031110230103213-0310213231212132-1110310322131211-3323121132221322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1323101120130031-1223230333033230-2002023013323133-3230112033103113-0232133002322032-0032120132010102-1333230302132013-1321222131002001"></a>

## Direct properties — check_present / 213102232300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302133031002020-0112312011300023-2303122013001320-1031021203132213-0212321010023212-0011223033312020-1303020302302201-3300333310322012"></a>

## Next pages — check_present / 213102232300 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2222320022131003-0333230101200030-1001003332011102-2033322112101302-0231021300332122-3031123223220122-0331130110231333-3031211122333100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133313011302231-1102332313002012-3322031120120032-3333330300330233-0113100302223221-1013110031023030-3201020120301332-3320223133030113"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — item / 222300200030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-3203033132011112-2310122200022302-2123300020213331-2132110322002322-2201110012010122-1130231210110233-3111013312111123-0200132010011213"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-1112222311100131-1313333221013030-1311302003223210-2211103021032221-0111003313313331-2122200331110101-3102010230020322-3222312013132303"></a>

## Direct properties — item / 222300200030 / 3

<a id="canonical-2010313320233220-0103123321312123-1321231203100301-3103021012000311-2231223101222212-2111101221103333-0012002120021021-3031232023312203"></a>

<a id="canonical-1021203030113032-0302120303113003-2033033000023221-0332001013320330-2332222022020323-2233130200220002-3331220012231020-3330020103101012"></a>

## exact_values property — item / 222300200030 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-0233103232121331-2300202330022230-0200032220320213-2120233033321203-2201321303022103-1121333113100123-3320221231002322-2123001011322311"></a>

## regex_values property — item / 222300200030 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-3030122230232200-2221101012202010-2022201223312101-2101320231101121-1013002022130012-3320131103130210-3223210233321010-3131031113302111"></a>

## transformers property — item / 222300200030 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030123020300200-0023000133003030-1322302023111332-3130121113202013-3121202303233133-3023303312123133-1323303130322120-3003103202220231"></a>

## Next pages — item / 222300200030 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3030233320003122-1011021332000221-3203021022013102-2130330003322122-1203302321002323-3211210201111313-0012132222031031-2310010203120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213311101212330-3100301002121303-1222301021222321-2113021103121032-0302201221003122-3220032213030033-1213102302023113-1332020211101200"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — asn_list / 012012331220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-0220233031230013-3202233312323011-3303301123322123-3133320311212313-3102312031011220-2001133123303121-2220322012323223-0211222111333223"></a>

Type: `"single"`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

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

<a id="canonical-3023222033133002-1300222320321031-0231201110203321-0122211021333013-3120311103131301-2223220102132013-1101113200022030-3112312230231000"></a>

## Direct properties — asn_list / 012012331220 / 3

<a id="canonical-2213320223023323-1213331230011311-2130313111130021-0212302121310021-1300221121113212-0311323210012123-1112200100302003-2233130222132313"></a>

<a id="canonical-0331303000233332-2013000011011202-3230302101212120-0110012322130103-2201013100120221-2011323103121113-1231000233220212-0001102131123111"></a>

## as_numbers property — asn_list / 012012331220 / 4

Type: `["list", "number"]`. Computed.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

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

<a id="canonical-1220112300213313-1131122323300211-3300232230233110-1202333301102010-2101100301332100-3013123323122123-1112302120330230-2132112113113031"></a>

## Next pages — asn_list / 012012331220 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302111122012022-3023221111223313-3231200231113221-0330332313333021-3233022000220112-0121231012022303-3123222313210100-2000120233232232"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — asn_matcher / 333220203113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
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

<a id="canonical-3303212202303131-0212303230003022-2233313030023011-1120200311320123-2231102120132201-0011000313020131-1223033221021302-3020130121232232"></a>

## Direct properties — asn_matcher / 333220203113 / 3

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122002213211130-3201030122031323-1113020310232001-3312130023013032-2202101133313130-3023013310012313-2123011332213113-1331021200200330): complete subsection reference.

<a id="canonical-1321220130123233-3312131320312321-1021330132033110-2003133332012023-1312100231103032-0320331230033210-2132020233221332-3231123011113330"></a>

## Next pages — asn_matcher / 333220203113 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122002213211130-3201030122031323-1113020310232001-3312130023013032-2202101133313130-3023013310012313-2123011332213113-1331021200200330)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2122002213211130-3201030122031323-1113020310232001-3312130023013032-2202101133313130-3023013310012313-2123011332213113-1331021200200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123002003123331-2121322121020003-2213001023303022-1311013211202210-1003003001011302-2010310210320020-0032332220200200-0323323233202213"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — asn_sets / 110130133312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-3120312033001211-2231322111213320-2113221333133133-1133130022312111-3323222311310012-2223022112110232-0333331101310012-3101211132231330"></a>

Type: `"list"`. Computed.

List of references to bgp\_asn\_set objects.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2002030010001311-0300201022321323-1123100030131333-0011223323232102-1001200221110330-0123201222011330-2133210031300100-3232131032323331"></a>

## Direct properties — asn_sets / 110130133312 / 3

<a id="canonical-0121330121013302-1021232302022202-0022323013201010-1233123212123012-2220212032032010-0203210220133100-3232012301322213-0323333132013231"></a>

<a id="canonical-3001003123311332-3321323221113232-1110003302221032-2202032011301103-2300312231230200-2133031333130133-0212302131301313-2232213310231330"></a>

## kind property — asn_sets / 110130133312 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1300022321333030-1012201001211033-0210131010123212-0030030103030023-0330002030200201-0301313303100110-2200023131311033-3103330102102331"></a>

<a id="canonical-3333212301101101-1132221111010331-3002331330201211-2210233210233223-0112203330232212-0322032003333321-0220222101113032-0013322331303321"></a>

## name property — asn_sets / 110130133312 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1100223223202201-1110300020112312-1130111211332223-3002330123202132-3000303031302120-2002200020101010-0132122301213220-0232210133111113"></a>

<a id="canonical-0312023322132022-0030032110300302-3120113231223021-2303012320131330-1120120022220222-1020120333320223-3030111133212313-3223031011111020"></a>

## namespace property — asn_sets / 110130133312 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0013312200122311-2333113003010212-3133100323312031-3020111232303223-0302011303213000-3332102110303111-1213112030321003-0120131031030012"></a>

<a id="canonical-0133321110132203-0211320100122020-2220302211311200-3003103031322220-3233103330302331-3231223013213012-0103311313321333-1211100331332213"></a>

## tenant property — asn_sets / 110130133312 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2300232101202000-1100301323002100-3122003013112312-3123011131011030-3032000231221013-1220032001333333-1023222210020300-2231311222020022"></a>

<a id="canonical-0233010113130022-3113112001002210-3222030133203101-0100223323231223-2121103221203301-3022101013211332-1133222322301222-2133233131200323"></a>

## uid property — asn_sets / 110130133312 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3000320021211013-2332120032023313-0230120201113210-3220000122300000-3311123221312120-0030321231322231-1200221000023103-0321002111203130"></a>

## Next pages — asn_sets / 110130133312 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1213311021203302-2332101131201022-1122221301221222-2300200300120230-2022132311130021-0100323233103233-1000133330002011-1123331220122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121303001201111-2221023203220123-1333333023330023-0202110221232322-3030102001320110-0003230001313332-1331312231311122-2102112301003110"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — body_matcher / 010312122031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-0021232132332231-3223332321120113-3201201200022220-0122000031122303-0133332031130012-1312202320200002-1101300212001312-1130112331202120"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-3302033322331100-1230100033213220-2302231102013103-0202302222221000-1001233322323321-2210100300021031-1033032323222112-1232322302130203"></a>

## Direct properties — body_matcher / 010312122031 / 3

<a id="canonical-2132032230201331-0011122332301110-1302233312203020-0101122032013022-3012123331201201-1223233223033300-1231102120220013-1301212200323012"></a>

<a id="canonical-2231212320011101-3221112032202110-2332010202331301-0002303111200312-3023310233122102-3300322330121030-0002032333033332-3022020010313213"></a>

## exact_values property — body_matcher / 010312122031 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-2311033132231320-2200110333323123-1013121301202113-0323321131210222-3102303002312222-1000222113033133-2111200001113013-3232220312122123"></a>

## regex_values property — body_matcher / 010312122031 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-2322230020020300-2310111332122310-3222020101213033-0021223120330101-0311201220202010-2013013110122002-0212313221330330-2133103312101003"></a>

## transformers property — body_matcher / 010312122031 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2200112102321031-3332122320020033-1333020321110122-0302221321310211-2231323030132120-2320232110301103-3323133012331122-2110021013023331"></a>

## Next pages — body_matcher / 010312122031 / 7

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3102120111001203-1112131330132320-2121010233221012-2320103020332010-2113223021322212-2311332300120013-1311000030211222-2132213323012211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221301222131313-3321300312011223-2122011031110131-2030220210201232-3132332033020223-3132313111200203-3203110000103313-1123322030300012"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — client_selector / 200332311221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-3100302003100121-2101031212313311-1311211200120000-1311230200122221-0012312132010120-1001030032031120-3112020230310203-0212302200232321"></a>

Type: `"single"`. Computed.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

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

<a id="canonical-2322130233320023-2001303031201130-3330210221002200-2120233323113022-0132323301320103-1232002220230023-2212012010023232-3310203000311101"></a>

## Direct properties — client_selector / 200332311221 / 3

<a id="canonical-2131123020323300-3330102312212223-0333023321332310-1332111232113223-0101100110211322-3110023321302223-1201000131021112-3132323002212030"></a>

<a id="canonical-2220212232311203-0301022310321113-0012220102120130-1121333003012021-1023320313111031-0110322232310230-2232113300311003-2003021131320032"></a>

## expressions property — client_selector / 200332311221 / 4

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

<a id="canonical-1030120302120000-1302303332011023-2231023223303001-1322002223332200-0313003030131010-1030100231321033-3111213123100121-3102102301012110"></a>

## Next pages — client_selector / 200332311221 / 5

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301320120132212-3231002032123200-2231332230213221-3132032033201132-3202323201120230-0113111012211210-2120102232301100-1001222132231203"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — cookie_matchers / 003123100022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-2031303023020310-2103200303200230-3002122313010321-2031022003310100-1301103113000112-2232311000303201-1122231010203230-1121210232211223"></a>

Type: `"list"`. Computed.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0121031201102213-2223311230212023-0133211231012201-2011001220013101-2301120113011313-1312332031301300-0032132320101322-3130200333012020"></a>

## Direct properties — cookie_matchers / 003123100022 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3112220113030003-1023031111132130-2220033031110223-1032201132133310-1330231202211113-1022011132301320-2303211202203221-3111103110113310): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2023012013122221-2232310323120110-3321321301313210-3212121020112212-3110120110022013-3223200222002302-0031001220313001-3100132033311312): complete subsection reference.

<a id="canonical-1012313202332301-2232231121213032-0212202321210203-3313232200222023-3123300022230003-0103203230331211-1313112010130122-2101120010233103"></a>

<a id="canonical-1002313331010211-0230031223222101-2021203211113122-0123301211321012-3112100203311123-1302130111323233-0100100300333012-0101302201003212"></a>

## invert_matcher property — cookie_matchers / 003123100022 / 4

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1331230000321121-3122220330123101-0032221230120000-3012322213133030-0130000301010011-2102301023321230-3220302132300113-0320022310301332"></a>

## name property — cookie_matchers / 003123100022 / 5

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2203223230333121-0130100303333232-2311121330233222-3030203311021010-2232023112100201-2110203120211223-2011223301320030-0212130303112120"></a>

## Next pages — cookie_matchers / 003123100022 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3112220113030003-1023031111132130-2220033031110223-1032201132133310-1330231202211113-1022011132301320-2303211202203221-3111103110113310)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2023012013122221-2232310323120110-3321321301313210-3212121020112212-3110120110022013-3223200222002302-0031001220313001-3100132033311312)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2330131222212101-2030300121322113-2132131311222121-1323000101322130-2123230323131122-1010323021231031-2213021200203211-3312130120323313)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3112220113030003-1023031111132130-2220033031110223-1032201132133310-1330231202211113-1022011132301320-2303211202203221-3111103110113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223223112011222-3121322111111231-2212320010313221-1231310203333002-3133131311233121-0131231123000322-1232031121132220-0001303121001020"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present — check_not_present / 133321120303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-2021002211123001-2031320102012002-1013131100103120-3311312013031120-3100232202101030-1333011031130111-1131010211123202-2010232322132102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-3221321112021200-2230312312031312-2033230100301100-3222112333313110-0333020310103310-3031313230222300-1313023211332022-3131321201200123"></a>

## Direct properties — check_not_present / 133321120303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223210102021003-2122013001120031-0332210322110013-1010130111120201-2012222001031023-1121133103032131-2312233023330121-1102310332111200"></a>

## Next pages — check_not_present / 133321120303 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2023012013122221-2232310323120110-3321321301313210-3212121020112212-3110120110022013-3223200222002302-0031001220313001-3100132033311312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112102112012232-0202221133110212-0233020113122332-2202303021320101-1331233031330300-3323333003022222-3000022021333210-1030100213101302"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present — check_present / 021121103310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2022223301220301-2322200011203323-2301131320120223-2230210000110311-3311101322112232-3003132133102311-2222320223132211-0201022003323230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-0023022012210301-1123201110013122-1031212223201301-2212233230321110-0030201233100211-1000312010313223-1010213121223302-3002121130210332"></a>

## Direct properties — check_present / 021121103310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313123110233321-0002113312330232-0123012021011110-3220102301101132-2323200303320322-2230233201111132-2310211132223111-3330030333303112"></a>

## Next pages — check_present / 021121103310 / 4

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2330131222212101-2030300121322113-2132131311222121-1323000101322130-2123230323131122-1010323021231031-2213021200203211-3312130120323313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230023303031132-3121102333331023-3300301322202302-1210311121033030-0221313232221100-2101331133023312-2013013300332131-1211331300211322"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers.item — item / 331333331330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-2012212000313113-2100132011023101-2322332211231330-1323113332302112-3022030312002312-0231131131203220-3210103011132223-0322321302200122"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-0302223130333321-3221013002202101-0001010220131223-3021312321232231-2301013211303303-1112321000313313-1120332131333013-3021033123110103"></a>

## Direct properties — item / 331333331330 / 3

<a id="canonical-2213020100230102-0023213102223031-1333021322233101-1323310000233203-2122231213220112-1001101131322130-2002000102231313-3021333133223001"></a>

<a id="canonical-0122100102300212-1332202000103102-0322033322231220-0010201111100122-1030202013203101-0201303021021010-3200033121013332-1323332200031112"></a>

## exact_values property — item / 331333331330 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-2212031011122203-3022030030203120-2013301031121102-2031011102220022-0032022232320020-1001222030000210-3303332013023211-1033032121233303"></a>

## regex_values property — item / 331333331330 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-3222302301121221-1031130001311113-3032330112002323-2313133333030210-1021031320011311-3102021130202001-0022132310132232-3113013012120000"></a>

## transformers property — item / 331333331330 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2303210201311231-1020200120111102-2220120310032333-2321121023100302-3223020103130302-3232230032002211-0003201102202002-3110111333303111"></a>

## Next pages — item / 331333331330 / 7

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3333031012000322-3022131233200332-1222200332232033-2210120021231332-1023232322313002-0133223103121211-0003110221123323-0313331331022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313122303333022-0020213012330101-1233322321112133-0123031123320212-2320112132302200-2301213320321123-2100233303033201-3330121110032103"></a>

## policy_based_challenge.rule_list.rules.spec.disable_challenge — disable_challenge / 003311101131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-2102211211013310-3131003323203000-3120032221011100-0103022333112211-0122031030333201-2233301101120131-0121030233110221-0013030011221031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable challenge.

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

<a id="canonical-2120201022313311-0101023003333010-0211113210101302-2232131310113132-0222330201331100-3112213223011330-2132121203200003-3022002323111220"></a>

## Direct properties — disable_challenge / 003311101131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302302010122311-1132021113011003-3023220203130320-2300130222201220-0112210323200333-2103133133212310-2231312112123202-0203033103123322"></a>

## Next pages — disable_challenge / 003311101131 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1213313313333103-1013220331032213-2322322203300120-0232232023222301-1110333033030011-2120302233312122-3102011002331023-2102012213103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030230023320120-3211203331332111-2221330030320201-1303112020110113-0032210221212122-2003112013021012-1212120123002303-1013022023322120"></a>

## policy_based_challenge.rule_list.rules.spec.domain_matcher — domain_matcher / 201113322212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-3202031332022010-0212311330100022-2010330303211311-3011300030103012-3223011313321221-0202223020100212-0231332003322112-0010012203311110"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-2323131110032230-2302333330031331-0300030123213223-2333012012013200-2102223122121210-3202223221102011-1111212203100333-0013322003001030"></a>

## Direct properties — domain_matcher / 201113322212 / 3

<a id="canonical-1003311002013300-0103110201233102-2101002323011021-2020311130132101-3000033030333110-3101211033223313-0331112003002022-3020330110013210"></a>

<a id="canonical-3201312331031202-0302032313311310-0022301201312300-3213133021313020-2231133322033102-1021021010100213-3131122133200100-2113122012132101"></a>

## exact_values property — domain_matcher / 201113322212 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-1122012221230221-1200101023012132-2002322013210030-3232210222232221-0232013120310100-2100303333303200-3313202113332323-0201223012012123"></a>

## regex_values property — domain_matcher / 201113322212 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-1010022132212022-2202313212103031-2230231301221101-1020000313331132-3021030112232212-3030110023222330-3313212311200312-0003022112130310"></a>

## Next pages — domain_matcher / 201113322212 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3011132030011201-2211311232213300-0312302301232121-3330311310211332-1123332021213002-1232230220130032-1231303030102230-1132313121201323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031302313120112-1300323311101330-1113011133232111-1111110023332022-2110022322102233-3022111000100121-1302130112320022-3230022333120031"></a>

## policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge — enable_captcha_challenge / 020030011130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-1010310100320332-0310000011322113-2331322010012322-3333301021233311-1033130223210120-0022001302211012-1022230010222331-2023312332001012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable captcha challenge.

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

<a id="canonical-2131103222203231-0130310231121033-3133133001030221-3320111333021022-2101330111102031-1301313001033300-3232220300313220-2112011333030203"></a>

## Direct properties — enable_captcha_challenge / 020030011130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122222310011122-1011122230013303-3032002333020010-0303220122131222-1230210300212021-0210000231103322-0010330301232312-1200013330230200"></a>

## Next pages — enable_captcha_challenge / 020030011130 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2310032320333030-1003211003002112-2220030321201321-2312212330103131-3102112121111313-0103301022112213-2332123333231022-0233120222133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213130123133130-1232010331133320-2112322313233121-0203102302002003-1331022203210232-2011000320003323-1021131002113311-2202201012300211"></a>

## policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge — enable_javascript_challenge / 201123222300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge

<a id="canonical-2130231033111113-0123322120331231-2210030123111031-1332110102113132-0332131231333223-0222132213321321-0231303122223230-2211033230130102"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0310131003332332-1122122213012123-1112303223103132-0212231031323132-3221211112332220-3201322121303231-3013013221230312-1313220230233103"></a>

## Direct properties — enable_javascript_challenge / 201123222300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002000030132203-1322130012101233-2001222022131223-0022322220010311-3103231103212230-0121133231000122-2311133330101011-3123331221003202"></a>

## Next pages — enable_javascript_challenge / 201123222300 / 4

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233222123132023-2031322222330200-3111303003223000-0013300202020102-3221121010130313-3222120303003220-3031220120000203-2320101201230033"></a>

## policy_based_challenge.rule_list.rules.spec.headers — headers / 321002121333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.headers

<a id="canonical-1000111131221313-2032121223322312-3302101320032232-1322322023111021-2031102103011011-2110133203100210-0332020203101201-3132120211103020"></a>

Type: `"list"`. Computed.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

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

<a id="canonical-1310231303302333-0303300003130112-3130103020133011-3320022331101331-3023200200222011-3233203021011032-1313300212120233-2130132011101022"></a>

## Direct properties — headers / 321002121333 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2231110013320313-1033011230133233-1333002201312222-2030201032101102-1100131203111031-0021210301313223-0111122121322213-3000120131213023): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0002311131221132-0202210312312223-3302312103220033-0033023132310321-1012322213332122-2123330222221322-0120330203103220-1303103330311113): complete subsection reference.

<a id="canonical-2030312031100210-0331003213022022-1002003203120231-1101110003020101-1011332021031022-3213010202123232-0030211003301131-3100320200130130"></a>

<a id="canonical-3003320123321133-1230113300101122-3203203011312013-0211311213201030-3302021300010000-2100001212322120-1220033123213121-3002321213112301"></a>

## invert_matcher property — headers / 321002121333 / 4

Type: `"bool"`. Computed.

Invert Header Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1120003122212032-1130103111000300-2302213221311302-0120312322211032-2320303202113233-1201111121332111-2030230130233301-1001033132002120): complete subsection reference.

<a id="canonical-2202303111003023-1303212120021200-0310320230002220-0312233031321333-1020103013312111-3201211012030300-0222232021331233-0133201001003301"></a>

<a id="canonical-2101232303333011-0203003201122312-2003211303231030-0031100320102011-3323121112231230-2330102120101231-3332302223320011-3002111230111030"></a>

## name property — headers / 321002121333 / 5

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

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
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3330200123213212-3200311312203102-2033100123131202-2101121233131310-2122020113330313-1002231121232102-1121211300323300-2030111120313003"></a>

## Next pages — headers / 321002121333 / 6

- [policy_based_challenge.rule_list.rules.spec.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2231110013320313-1033011230133233-1333002201312222-2030201032101102-1100131203111031-0021210301313223-0111122121322213-3000120131213023)
- [policy_based_challenge.rule_list.rules.spec.headers.check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0002311131221132-0202210312312223-3302312103220033-0033023132310321-1012322213332122-2123330222221322-0120330203103220-1303103330311113)
- [policy_based_challenge.rule_list.rules.spec.headers.item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1120003122212032-1130103111000300-2302213221311302-0120312322211032-2320303202113233-1201111121332111-2030230130233301-1001033132002120)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2231110013320313-1033011230133233-1333002201312222-2030201032101102-1100131203111031-0021210301313223-0111122121322213-3000120131213023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013202100101313-2011113003330133-2220021122132332-0211120220020023-0322002322013200-3312012000230201-3010112101102211-0213030122132010"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_not_present — check_not_present / 003210200211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- policy_based_challenge.rule_list.rules.spec.headers.check_not_present

<a id="canonical-1320123032200030-0021101330311023-3322203032323103-1311023103300111-1012003322302013-3301002023032133-0012121113331232-2031213033331010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-2220021100112031-2233113300210022-3122210130102113-3100032220030100-3120021213330312-3301220313131113-2233013310131211-2320133233122011"></a>

## Direct properties — check_not_present / 003210200211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321121221122113-0012132221223002-3311223323303223-1333303300212320-2231232311300332-0232213031102210-2233220030110311-2110221023022100"></a>

## Next pages — check_not_present / 003210200211 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0002311131221132-0202210312312223-3302312103220033-0033023132310321-1012322213332122-2123330222221322-0120330203103220-1303103330311113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211202331111230-2200030032230020-3120020333023223-1010101113223113-3010101220223030-3213032102010303-0301130120000320-3001330330231300"></a>

## policy_based_challenge.rule_list.rules.spec.headers.check_present — check_present / 333212213131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- policy_based_challenge.rule_list.rules.spec.headers.check_present

<a id="canonical-1033033210213212-3231222123110201-0011202002112123-0100021003113113-2301311203230131-3110310313122132-3223312303312123-3110011002320012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1123233122300312-3301131113112123-2021201010233332-0023232101302012-0101122012032223-1130301312013311-0122201233011011-2022332103121302"></a>

## Direct properties — check_present / 333212213131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332003012312001-0130322213221302-0032310210320212-1012012322312110-1132120123303332-1232310201230111-0123221302221132-1120310221231302"></a>

## Next pages — check_present / 333212213131 / 4

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1120003122212032-1130103111000300-2302213221311302-0120312322211032-2320303202113233-1201111121332111-2030230130233301-1001033132002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301302310313030-0203320300313212-1020233320220200-1121120130202032-1312232102020332-0200121211300012-1202220102121112-0033230301200001"></a>

## policy_based_challenge.rule_list.rules.spec.headers.item — item / 222123310103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- policy_based_challenge.rule_list.rules.spec.headers.item

<a id="canonical-2011231213301110-0003320021023320-3123201222110122-3203210311312011-3200320102330013-2202102330311220-1300230110110110-0230302320021212"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-1301031033312100-3233222100331312-2020133030030321-0233001331010231-1231133221322030-1112033002221230-1203323130201222-3330322211110200"></a>

## Direct properties — item / 222123310103 / 3

<a id="canonical-1333212032010231-1032002031113121-2300011022300000-3321303130033003-0221301221021321-0031301113031201-0231101121012110-0320100023132331"></a>

<a id="canonical-1310330033201211-2122213233030322-3321323120011123-2022100313230333-0231101323321110-0212002303102203-2213110332211112-3303322003020200"></a>

## exact_values property — item / 222123310103 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

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

<a id="canonical-1123313012301022-0312030312322223-3323313301213130-1312111021203101-3201103020322212-1201311331210303-3320000003032133-3033002201102130"></a>

<a id="canonical-3332233220200320-1311130011010111-0101301021101131-2020202011233210-2331222302023001-0022332301231013-0231103021311021-2312013331230223"></a>

## regex_values property — item / 222123310103 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

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

<a id="canonical-2120010000230031-3101210301111010-3232333312030300-1230001221200220-1000002320132321-2232013232013320-2031200303110112-0310031212222333"></a>

<a id="canonical-1013010103123012-3330132303311313-2321223313031313-3313203311230213-2322302120303000-1221332320300011-1030112003333300-2211333020023320"></a>

## transformers property — item / 222123310103 / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3333331113210221-3101132303330102-0133220222113113-0020133102330333-0010322011012323-3012020300320210-0123130131002320-1321111101101011"></a>

## Next pages — item / 222123310103 / 7

- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0203122003302010-0311301031311231-1113123010123300-1231200102112212-0111013121233330-0102113203321103-2323133133131301-0110321113200323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110203222031230-1210312203311030-1113202232120303-0321031122211033-0222320012210311-0113311221302303-1023321203210301-3121032132131000"></a>

## policy_based_challenge.rule_list.rules.spec.http_method — http_method / 311323001101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.http_method

<a id="canonical-0010203231313332-1300132230212003-1122103101002322-3230100002321323-2123302313232001-3302202013010222-1112233111323121-1010100001323320"></a>

Type: `"single"`. Computed.

HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Upstream description:

A HTTP method matcher specifies a list of methods to match an input HTTP method. The match is
considered successful if the input method is a member of the list. The result of the match based on
the method list is inverted if invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221331233221132-3020201120331310-1003213103221212-1121002302231313-1130003210332123-2301110001213130-0213331103233332-3121130031332331"></a>

## Direct properties — http_method / 311323001101 / 3

<a id="canonical-0012130301000001-1222112031312212-0101011302011321-3112211000023120-2013122213232121-2333230223131202-0000013010032032-0102010310122211"></a>

<a id="canonical-2302122322032111-0333101310001310-0010220030010120-3032132031002032-1302021201023003-3312031032133122-2213211022120233-1201030223002300"></a>

## invert_matcher property — http_method / 311323001101 / 4

Type: `"bool"`. Computed.

Invert Method Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002313021110222-3302103020122100-2020122020123222-1023033000303332-2313110203232213-3323212010000221-3023102223201200-2321312112322213"></a>

<a id="canonical-0201012202323303-1113332323032000-2130212301220311-2310203110213023-1321322021000223-2132030130100202-3132101023030230-2322020201322231"></a>

## methods property — http_method / 311323001101 / 5

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] List of methods values to
match against. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`,
\`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

List of methods values to match against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330010312332210-0110101021100030-0302303031111111-1010132123220313-3302222223011010-0222200113322210-0211002031011130-1023111133121231"></a>

## Next pages — http_method / 311323001101 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223111103201230-0132331101020230-2023211123211233-0032233111310113-3312010313213212-2122120301020133-1332120220021012-2302132111223113"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher — ip_matcher / 110000201212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.ip_matcher

<a id="canonical-0212213202112203-3212003033312303-3031000033133303-3121211010030211-3312302213100211-1231130300312003-0231202233300122-1212020333230230"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132101133103001-1302221033112222-2213222123233033-0330313031201222-1103101231311103-3301210020312312-0003130222223313-0122101110233002"></a>

## Direct properties — ip_matcher / 110000201212 / 3

<a id="canonical-1113223232232200-2112103122302330-0321133012211022-2210121202220122-3001300202020012-1021302122330122-3010212110331113-0321311001013210"></a>

<a id="canonical-1321313312213013-3100333000220330-0110333111223031-2102310320123210-2310220313300321-0321032131010321-3223000213000131-0210100300212213"></a>

## invert_matcher property — ip_matcher / 110000201212 / 4

Type: `"bool"`. Computed.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [prefix_sets](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201200133010331-0303200011120323-2012333010111103-1321322111032112-1332011223112331-0311211322230331-1023121101101030-2030233133013103): complete subsection reference.

<a id="canonical-1323213213333313-3000000111302201-3201030201103211-3032101102312131-1111331133330221-0122111033303203-0121122203000201-1332313122031121"></a>

## Next pages — ip_matcher / 110000201212 / 5

- [policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201200133010331-0303200011120323-2012333010111103-1321322111032112-1332011223112331-0311211322230331-1023121101101030-2030233133013103)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0201200133010331-0303200011120323-2012333010111103-1321322111032112-1332011223112331-0311211322230331-1023121101101030-2030233133013103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330303020232110-0302203232020222-1322302123312221-3012212322030121-2110132122000320-2301031112020222-3300112332232323-1112210211203113"></a>

## policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets — prefix_sets / 102032123212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221)
- policy_based_challenge.rule_list.rules.spec.ip_matcher.prefix_sets

<a id="canonical-1303323100200010-1203122231122100-3333212032010132-3210330320113322-1312311233001221-1323203130203021-1122213012113231-2102120021033102"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-0332302210013023-3330000330032000-1232202033202110-0122203122212212-3120111311121123-0123312113103321-2212130332033111-2320312303302200"></a>

## Direct properties — prefix_sets / 102032123212 / 3

<a id="canonical-3202221030120001-0210221030110332-3103131003002232-3221333023131100-3233230001021003-3012220301030131-3322312121033022-0212230123102033"></a>

<a id="canonical-2312212033013310-3131130120101212-1131131013130020-2121103021001302-3320033222321200-0302123020013010-0203120130100200-3003013202111211"></a>

## kind property — prefix_sets / 102032123212 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0323002200033331-3312130200033200-2310321321133323-1200302312313020-3032022232330113-1020220031323221-2300002131211122-2122232111122022"></a>

<a id="canonical-0003030133300112-0331032031211213-1211123300000232-3311231323022003-0310121023010231-1211320210133111-2231312322322220-2302131201013021"></a>

## name property — prefix_sets / 102032123212 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3123300203212110-2230022211110333-2132013103112233-3022302103121131-1112333320312130-0123030102313010-0311020222313323-2331002223111221"></a>

<a id="canonical-2023022030331303-0332301211131302-3331211000223103-3211033330313122-1013111323112002-1020333201221312-0323221202202212-1323203111131122"></a>

## namespace property — prefix_sets / 102032123212 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2101001202030203-1330002103202200-2112123230010001-0011332311022003-0113100320210300-2301301220002110-0000131100321300-2132123312232201"></a>

<a id="canonical-0303002232223232-0133310023023021-0333121312122013-2012222133030313-3202233213302012-1330103023013312-2010011313022001-1110231320200203"></a>

## tenant property — prefix_sets / 102032123212 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1330302220332022-0203230110132100-2121302021213220-0322011102110203-2020300100222020-1111112123233130-3321113331033033-0032220310133031"></a>

<a id="canonical-0121100101123313-3310221301103301-3213013213130210-1310002103301311-2302322102303201-2213211301303133-2322210033232330-0123202123030311"></a>

## uid property — prefix_sets / 102032123212 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1310200110012001-2323303120230103-3021211233333300-3113232212220230-2110130312331313-1120321322310300-2222021212223121-2203133131201222"></a>

## Next pages — prefix_sets / 102032123212 / 9

- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0212033022130233-1300022103102203-3221233211032110-0022200332332213-3121222133232013-0122231011020300-3311100313212031-0300232013311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100202020222201-2001120100023011-0330131031130322-0323130332133123-2310132300031121-0110330132201002-3313102210332222-1030121330012013"></a>

## policy_based_challenge.rule_list.rules.spec.ip_prefix_list — ip_prefix_list / 300222013110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.ip_prefix_list

<a id="canonical-2311032030113022-1101010211220202-3013010123212030-3231121023020231-0323323302233131-1233322000220012-1000210131130130-2233322103013212"></a>

Type: `"single"`. Computed.

List of IP Prefix strings to match against.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301220232320032-1311233000111321-3320030130233131-0001133112120013-1132322233121301-1321333121201330-1023110113233233-0212130200001231"></a>

## Direct properties — ip_prefix_list / 300222013110 / 3

<a id="canonical-2111223331221321-3133332120103332-3211120102213222-0000123322001330-0222032112223013-0110103022033230-0113331123131003-1202321331312230"></a>

<a id="canonical-3130312010303023-0013131112221021-3001132121312301-3100100022003103-1030031021023013-3032110201301000-2101013110333133-2002332201031130"></a>

## invert_match property — ip_prefix_list / 300222013110 / 4

Type: `"bool"`. Computed.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0332232021222212-2112113210213331-3300210211110231-2201211111012023-2211020002312303-3200032200302121-1232002303323032-0012321102113320"></a>

<a id="canonical-0230103310302000-2020013223333102-1122030312021000-3203212033223323-3101031123122001-2022020101312221-2023221030211032-3310320121031313"></a>

## ip_prefixes property — ip_prefix_list / 300222013110 / 5

Type: `["list", "string"]`. Computed.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1121121321223113-3030202132331113-1330133332013112-1213313213120312-2200010113320231-3222012003022030-0013222222233022-3033012320313201"></a>

## Next pages — ip_prefix_list / 300222013110 / 6

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1323200003020201-0013112330130332-0230323022322321-2231311020200133-1312320220310220-1302221121320212-2103002221230003-0302010321213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131010200003220-3110022012211000-2212300102322020-2220010133131202-3110021320022110-2333032232330213-0030131112033210-2210201332320122"></a>

## policy_based_challenge.rule_list.rules.spec.path — path / 322010020302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.path

<a id="canonical-3200301310123303-0003311131211002-3312011332013003-0003021020021302-2321021301003201-0320023003131130-0010332212003333-3113021112003212"></a>

Type: `"single"`. Computed.

Path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Upstream description:

A path matcher specifies multiple criteria for matching an HTTP path string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of path prefixes, a list of exact path values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2111220032032012-1031202200100100-2230322221001333-3222220123121032-3131301111022003-1020321300131122-3132113211122233-2133023203031213"></a>

## Direct properties — path / 322010020302 / 3

<a id="canonical-2022121112010130-1213331023200013-1330233312330110-1022320002113133-0101112103231301-0233310102121131-2000012313020133-3022222113032220"></a>

<a id="canonical-1131203202201010-1222302231123312-1230311221133100-3003132012223221-3110323213111033-1202112132033111-3102123130233200-3001022223111323"></a>

## encoded_path_matcher property — path / 322010020302 / 4

Type: `"bool"`. Computed.

Match against the encoded, escaped path.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0211001212222202-0333103010031110-0232032323110202-0133201331333313-0112033223033003-3231030103100222-2110010333112222-3313003122202333"></a>

<a id="canonical-0021200223031120-3303322312030312-1322220112222002-3330323313200300-0223001003112021-1211132112102000-3112233122301221-1330230021011332"></a>

## exact_values property — path / 322010020302 / 5

Type: `["list", "string"]`. Computed.

List of exact path values to match the input HTTP path against.

Upstream description:

A list of exact path values to match the input HTTP path against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1331131010003103-3210122131201103-1012032111301220-2322233322310130-2331033233121013-1221131000010001-2332313000330000-2310331121212033"></a>

<a id="canonical-3021032122102030-3303000233301013-1311322112133112-0112032102121003-1103131031130230-3103223103002023-2231020121121000-1302321011221203"></a>

## invert_matcher property — path / 322010020302 / 6

Type: `"bool"`. Computed.

Invert Path Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3022003301100211-2333303333131131-0223221201013332-1133011203231221-1121133000131232-1200202120330122-1012200330013003-1323103322131103"></a>

<a id="canonical-2003001211230312-1331321322132302-1030232231020032-0311013333232322-1001022322311200-0223201103300312-2211332121302102-2311133213300033"></a>

## prefix_values property — path / 322010020302 / 7

Type: `["list", "string"]`. Computed.

List of path prefix values to match the input HTTP path against.

Upstream description:

A list of path prefix values to match the input HTTP path against.

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323313100223111-0223101012010132-1303123220011330-2011323132201212-3003211333200132-1122012310233101-2230200222020301-3303232231121133"></a>

<a id="canonical-3002120202100033-3312002032033333-2323100203120023-1202312321312330-0300112322112321-1312113023310120-0020300300221301-3101303233111230"></a>

## regex_values property — path / 322010020302 / 8

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input HTTP path against.

Upstream description:

A list of regular expressions to match the input HTTP path against.

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

<a id="canonical-1011112312310110-0331002313321021-2020011312010303-3032212220310322-1020321212022221-0222013103333103-3100222323032030-0123101123213100"></a>

<a id="canonical-2300322302133202-0121233202233212-1120111133100123-3231133231222323-0013123331102211-0330101313021033-3313223310301200-0203010213100130"></a>

## suffix_values property — path / 322010020302 / 9

Type: `["list", "string"]`. Computed.

List of path suffix values to match the input HTTP path against.

Upstream description:

A list of path suffix values to match the input HTTP path against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "64",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1130010013013303-1123131203312011-1230330230130211-1212023103223101-3331001220312233-1303033031213330-1220333112333221-1323323222303310"></a>

<a id="canonical-0330132132220331-3202211001110312-2122300103331210-1211313313112023-1323320322132110-3130121000320133-3231230002302231-3231310123002211"></a>

## transformers property — path / 322010020302 / 10

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2033232111013010-1031102122332302-1230311203331122-2220002123322200-2303221122111012-2121130200003331-0312313200310022-3110102232212330"></a>

## Next pages — path / 322010020302 / 11

- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330202310221030-2331323120222230-3200123220032312-1112101010112231-0000013331112123-2012330112321303-1311012113302121-3322131120032132"></a>

## policy_based_challenge.rule_list.rules.spec.query_params — query_params / 223312011120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.query_params

<a id="canonical-3223030100001112-3300203310220111-2331013121211103-0112000200020133-2310232323301212-3333033012323332-0312132000203313-0031312002002032"></a>

Type: `"list"`. Computed.

List of predicates for all query parameters that need to be matched. The criteria for matching each
query parameter are described in individual instances of QueryParameterMatcherType. The actual query
parameter values are extracted from the request API as a list of strings for each query..

Upstream description:

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

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

<a id="canonical-3032312230230012-0333323302231001-1131133131220302-3032232011210001-3331133330312100-1211221202023100-0122311001212330-3303202311212320"></a>

## Direct properties — query_params / 223312011120 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0310221122220321-0200112023310103-1010231121120221-2132023213111221-1320131203223101-1130300300013032-2023322033221320-2113230220131102): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210033000022201-2011132322022311-3221303223200012-2021031031020112-3310030022321321-2212022310103223-3313232213111322-3313323013133032): complete subsection reference.

<a id="canonical-0132130130001313-0102302310121302-2111001302033122-1020130101331023-2333220022130332-1323011223113200-0231110313202030-2233010313222321"></a>

<a id="canonical-2032102221212000-3221300321131322-3320010322333101-3220012333033210-3231223000220022-3002130132300210-3132121212123200-3212002322320332"></a>

## invert_matcher property — query_params / 223312011120 / 4

Type: `"bool"`. Computed.

Invert Query Parameter Matcher. Invert the match result.

Upstream description:

Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1332322101003033-0121313330231203-1332123101331331-3010131300333020-2300031020011010-1032132323303313-3130130211323322-0021002331321113): complete subsection reference.

<a id="canonical-2312123132230303-1122033220120302-0033002203230230-1132132313013101-1022232021102211-3230233031233021-1222322013223302-3102101222002123"></a>

<a id="canonical-1132301001222302-1003031002002020-0322010302131201-3032000031232101-1120311331003101-0030021001100022-3012120103320333-2310220110022103"></a>

## key property — query_params / 223312011120 / 5

Type: `"string"`. Computed.

Case-sensitive HTTP query parameter name.

Upstream description:

A case-sensitive HTTP query parameter name.

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

<a id="canonical-3120100010233012-2131010000023030-1323123221012120-3011200232221131-0103132031133300-1000223203130231-1312232122010121-3011113213132131"></a>

## Next pages — query_params / 223312011120 / 6

- [policy_based_challenge.rule_list.rules.spec.query_params.check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0310221122220321-0200112023310103-1010231121120221-2132023213111221-1320131203223101-1130300300013032-2023322033221320-2113230220131102)
- [policy_based_challenge.rule_list.rules.spec.query_params.check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210033000022201-2011132322022311-3221303223200012-2021031031020112-3310030022321321-2212022310103223-3313232213111322-3313323013133032)
- [policy_based_challenge.rule_list.rules.spec.query_params.item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1332322101003033-0121313330231203-1332123101331331-3010131300333020-2300031020011010-1032132323303313-3130130211323322-0021002331321113)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0310221122220321-0200112023310103-1010231121120221-2132023213111221-1320131203223101-1130300300013032-2023322033221320-2113230220131102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121331111100210-0011333132333011-3121321222012113-1030003000103122-1320231110002003-2020330200013231-1302103103221113-3302301023120330"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_not_present — check_not_present / 301301021311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- policy_based_challenge.rule_list.rules.spec.query_params.check_not_present

<a id="canonical-0302333210213213-1201013330003221-1012032023213031-0003313110311223-2012201211222102-0003331103202230-2102201322011000-3001100323213131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-1302300012022312-2020122310301033-2312133102022320-3213113101120233-2003132030323012-0210212313230310-0332121022232222-3031022330310230"></a>

## Direct properties — check_not_present / 301301021311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221030330101300-3010111113101133-2001031203011203-3211202233321301-1302132100310022-0332012202323312-2300302022320102-1230232120333133"></a>

## Next pages — check_not_present / 301301021311 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0210033000022201-2011132322022311-3221303223200012-2021031031020112-3310030022321321-2212022310103223-3313232213111322-3313323013133032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031230001223133-2201003032001023-3210132031012032-1220231320311001-2131111210021213-1033231100332233-3002112102203032-3333022311033111"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.check_present — check_present / 320010301212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- policy_based_challenge.rule_list.rules.spec.query_params.check_present

<a id="canonical-3101211221301310-1231032220120133-3131021221002203-3303202211012303-2233303023300003-1102221133232133-3000133223312201-2223122223013212"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1020223223221200-0130203201030233-3231211201302113-0302123212021012-0303003123122030-1003313331213312-1011201223100231-0202332331222112"></a>

## Direct properties — check_present / 320010301212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121132010122132-3000210230310223-3101100210013031-1232013011100310-0103013203002000-3302233122131001-1311122103221130-2220132323020013"></a>

## Next pages — check_present / 320010301212 / 4

- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1332322101003033-0121313330231203-1332123101331331-3010131300333020-2300031020011010-1032132323303313-3130130211323322-0021002331321113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312310332202210-2101223031132203-0011300310111322-1122011323301131-0121031333233212-0310233132323103-0301113303301233-0321202010310132"></a>

## policy_based_challenge.rule_list.rules.spec.query_params.item — item / 111100013031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023)
- policy_based_challenge.rule_list.rules.spec.query_params.item

<a id="canonical-3323211122002200-0130311300033011-3010001010030320-1202101130032232-2220123133102322-2013233033131112-0030131122232003-0000201112313122"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

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

<a id="canonical-0313122230011110-0133012133103311-3231133212221030-0332023302102313-3213112210220030-2130200210010211-0301100001031230-2100232012213011"></a>

## Direct properties — item / 111100013031 / 3

<a id="canonical-3331121323203231-0222111120323321-1213233131322011-3020111100011113-1123103213112010-2021212302201311-0303203010313013-1210012333031010"></a>
