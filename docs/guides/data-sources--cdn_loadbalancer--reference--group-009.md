---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1103033302220303-2030021302210003-0103101013312323-3110232100312310-1213202113331012-1032221020001330-3133001333203213-2100330332313221"></a>

## bot_defense.policy.protected_app_endpoints.path — path / 303320003200 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-3111013232201301-3312013312103311-0101002011200203-0323112301132013-1221310300222103-2130323121002201-3000312022302020-0321131210212133"></a>

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

<a id="canonical-0220223102020201-3133220022203231-2011233122130103-1101132002013231-0033200010202230-1100312302121223-0021231300031210-1030332003211032"></a>

## Direct properties — path / 303320003200 / 3

<a id="canonical-0213200212030031-1211122002300032-1111120200313103-0203333200232110-3002331301221021-2112113303333332-1033000012221202-2002103122213001"></a>

<a id="canonical-3233110011320123-3231112230233021-0230303001211100-3320022122130021-0133130001123003-1102210323230123-3231102112310222-2020223302130211"></a>

## path property — path / 303320003200 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-2133233003233133-0222332333032221-3320222302112321-0311211102021333-1013213131231103-2331113101230310-3003312100023033-1323330200332010"></a>

<a id="canonical-1223102022223110-2301221110100111-1333201302022323-0223322320133303-1321002320122110-3121100111130031-2022233012002322-2120302102101101"></a>

## prefix property — path / 303320003200 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-3301122122202211-0302213200233001-2332231113222302-0022111321322121-1333012020010220-0323131003301033-2022201003010231-1012312020010210"></a>

<a id="canonical-3111312031102222-3330331311002123-2210300032310300-1112233211111132-2102122022210022-3121221103132112-2320220101200001-2000210202103020"></a>

## regex property — path / 303320003200 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-0302301333023011-3232300303210200-1320300002032010-3311030120303211-3200133103003111-1030221032222002-0221130311313102-2111133001231323"></a>

## Next pages — path / 303320003200 / 7

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103333301121210-3223100010312221-2322002002113000-1023202010210213-2312030320020301-2230210200121011-0113330130003011-0300100221123112"></a>

## bot_defense.policy.protected_app_endpoints.query_params — query_params / 120220022020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-3123021012213321-1002212200313022-3311000232331330-0202323122000103-3031312001030302-1231223010222332-3220221301133210-2122232312001211"></a>

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

<a id="canonical-2120322120212331-0201222122013333-2331303021122210-3333111212320313-2233032122122330-1011123313010313-3133111313120120-2001300112022200"></a>

## Direct properties — query_params / 120220022020 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031333210221233-3223210210200001-3323230023133100-1112002000011221-2320303120003123-2301222321302321-3212112322312031-3013323201020033): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311213011012020-3033311330111230-2210123300300133-2213013012121330-0312310113232102-0003110201222011-0032201232231101-3323020003300101): complete subsection reference.

<a id="canonical-0113303012003012-1200023001301132-3011311110130021-2010030132201211-2010101102233202-3212132303330312-3013321022231012-2333223210031000"></a>

<a id="canonical-3333213001212213-2111221223020222-3213111123213103-3121202210201021-1003001232220330-1032232200000030-2123002302332223-3320313011223010"></a>

## invert_matcher property — query_params / 120220022020 / 4

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

- [item](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0213221120110131-0320013010200232-1112030300200220-2110113103331323-0210032002032032-0120203003313330-2133012001002310-1233212231213133): complete subsection reference.

<a id="canonical-2202212301222323-1003031301223131-3002013210132130-3223312030330011-2232013333033112-2222122301022230-3210131303303021-0320332113200020"></a>

<a id="canonical-0211302223113103-2110213300012232-3200330322312012-2103020032300302-2001331310010022-2332312313110300-2223123033130332-2233203111121001"></a>

## key property — query_params / 120220022020 / 5

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

<a id="canonical-0023211303203221-2112113023320020-2133231023132302-1113323313222220-3223111232200101-1133322233022102-0212303222002010-3321122003203312"></a>

## Next pages — query_params / 120220022020 / 6

- [bot_defense.policy.protected_app_endpoints.query_params.check_not_present](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031333210221233-3223210210200001-3323230023133100-1112002000011221-2320303120003123-2301222321302321-3212112322312031-3013323201020033)
- [bot_defense.policy.protected_app_endpoints.query_params.check_present](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311213011012020-3033311330111230-2210123300300133-2213013012121330-0312310113232102-0003110201222011-0032201232231101-3323020003300101)
- [bot_defense.policy.protected_app_endpoints.query_params.item](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0213221120110131-0320013010200232-1112030300200220-2110113103331323-0210032002032032-0120203003313330-2133012001002310-1233212231213133)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1031333210221233-3223210210200001-3323230023133100-1112002000011221-2320303120003123-2301222321302321-3212112322312031-3013323201020033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021032002132001-1213310210000233-1230231223220223-1213111011030203-2021233101322010-1000030123112020-1313122003323021-1120312200322120"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_not_present — check_not_present / 121031231301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-1310012302303032-2331322122311112-3300223032302333-2213311222123331-1311220211312330-3210002220301201-2133230112200003-3132020301020012"></a>

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

<a id="canonical-1001002330013003-3332333112033021-2332310122301330-1230022333302121-2313133212203201-3121301000002131-3330232020021031-0323211020322201"></a>

## Direct properties — check_not_present / 121031231301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200222100220302-0212222331212113-0332221012102223-2111233012232112-1230203123001110-1302303032320201-3330012231331201-2301312220120220"></a>

## Next pages — check_not_present / 121031231301 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1311213011012020-3033311330111230-2210123300300133-2213013012121330-0312310113232102-0003110201222011-0032201232231101-3323020003300101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103030003302133-0131210222301032-0023331013013110-0200222021230233-3311101013332222-2303133323012233-2101031223200013-1300323312211333"></a>

## bot_defense.policy.protected_app_endpoints.query_params.check_present — check_present / 323312230012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-0100031332011220-3300321231033230-0031001322303130-2030012310012203-2222021310202110-2020002133102010-1123203033321033-0133011323100122"></a>

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

<a id="canonical-1130302312333121-0012111113122030-0013110300130201-2231003011003220-1112222132231323-1110300202212230-0031020132202101-3132023321112221"></a>

## Direct properties — check_present / 323312230012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312012131232111-3232001100210001-1233331310111133-1333220100013131-2101211213201110-3032310012202132-2303131211111120-1222012222301000"></a>

## Next pages — check_present / 323312230012 / 4

- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0213221120110131-0320013010200232-1112030300200220-2110113103331323-0210032002032032-0120203003313330-2133012001002310-1233212231213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322001021113333-1312331321121133-3303310031120133-0320232101110111-3000122210211312-0230310011231100-0020113333033001-2012230301213210"></a>

## bot_defense.policy.protected_app_endpoints.query_params.item — item / 113331312020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-2102011203101321-2221001332321211-2220101032133123-0101111202313231-0202000201112233-2320201320030010-0113030303331202-2311223023130302"></a>

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

<a id="canonical-3123320221313003-3233031320331321-1110230212132001-0320301011112221-1131100012331303-2021313310003233-1103132300311220-3111030302010022"></a>

## Direct properties — item / 113331312020 / 3

<a id="canonical-3023211030222110-0322021031133223-3330213232221200-0331311212322100-3322122003032233-2110012200212013-0020321302101312-1120012022300030"></a>

<a id="canonical-3230001332231133-1212003233330330-1112211212121021-2002230030003333-3120303212123101-2102303331013321-1031101022121203-1122313022102320"></a>

## exact_values property — item / 113331312020 / 4

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

<a id="canonical-0231113020302012-0103331201132111-1020320121001110-0102030310032033-2113110311101033-1022013231131333-2202121330010313-0332323133220021"></a>

<a id="canonical-2200103011222313-2330030013030003-1202231112020022-0310211120122120-0213212122210102-2300011000201203-2312301113033111-3100132132031201"></a>

## regex_values property — item / 113331312020 / 5

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

<a id="canonical-1021300010220122-1213001001003220-0331010011022123-3110032011222200-3323103220021312-3201321003213130-3121031321001121-2232102000020113"></a>

<a id="canonical-3121123312031233-2132122012013033-0022302032232000-1301202123330302-3311323202301200-2132010000012201-3121303032111323-2123112021012321"></a>

## transformers property — item / 113331312020 / 6

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

<a id="canonical-3132032113030301-1231023032033013-2312112331003312-0203223010221112-0323113331123301-2303201110220213-3333011220320013-3320132121300030"></a>

## Next pages — item / 113331312020 / 7

- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2033213231111310-1201133200120003-2000203101330212-0313030322132332-3033031231110033-3232201331213211-1323002113020311-3120110033330200)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1003003211020113-3103312331332221-2103313103132131-1113322131300002-1031212331122122-3012103220120231-3032222232112132-2303111022032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000312321322111-1232232212321021-2110100022210201-3202023101331332-3020032122233112-1030013021033133-0200311223221033-0023322102020112"></a>

## bot_defense.policy.protected_app_endpoints.undefined_flow_label — undefined_flow_label / 333130303133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-1113102232210110-3010110210300232-2311222031131130-0321230013221003-3230100231012123-1123020301203130-2230222302220121-3130201332032300"></a>

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

<a id="canonical-2220000312112311-3321023223002023-1313221203233031-2220001311132311-0230200120113220-0331100310010231-0100302323210310-0222112232301233"></a>

## Direct properties — undefined_flow_label / 333130303133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022223230213032-1032323133023231-0233021032032000-1021300132230133-1313112102203322-2123223032000330-1123312210022230-3133032202323301"></a>

## Next pages — undefined_flow_label / 333130303133 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2222302100002000-1121300110111033-2303003223331311-0113322310023022-1123233232023213-3103232323330012-0002320230103130-1111220303320033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031212120021332-3201021110212110-0211013123010311-3000310310313013-2213320021333323-2011213030330230-2032332101321321-3232100223331200"></a>

## bot_defense.policy.protected_app_endpoints.web — web / 310200321322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-3120122012311233-3213223122132311-2111232302231311-2121223322330302-2332011113122012-3103100330333003-3312202020223112-3322011100301101"></a>

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

<a id="canonical-2133012123230231-2232111231010300-3223322310211022-0202010300303100-0230002112332130-3233310323013101-3230220322230211-0012300332122301"></a>

## Direct properties — web / 310200321322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303202303011100-3201102123101001-2103130232132031-1110300331112103-3200233100200010-1322120031213231-2231110102013323-3020231021103333"></a>

## Next pages — web / 310200321322 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2212233111032332-2322130100100201-1312110231313203-3103222223203300-1323133131321310-3310021133322023-0121033021113312-1231221220103101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131310211012112-1310332303032021-3313130031201222-3131222333212333-2231313211230231-3210101311111023-3013302002203120-1101310003300102"></a>

## bot_defense.policy.protected_app_endpoints.web_mobile — web_mobile / 010120111023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-2013231001110213-1100131303013033-0233231211002201-0003022030321332-1203222202301230-2023332203021220-0320331210001122-1110011321000020"></a>

Type: `"single"`. Computed.

Web and Mobile traffic type. Web and Mobile traffic type.

Upstream description:

Web and Mobile traffic type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0122020201330203-1112021030313011-2012021002302112-2313302022113213-1311112203023203-1201301321102320-3201202113000211-2301330320033301"></a>

## Direct properties — web_mobile / 010120111023 / 3

<a id="canonical-0131132331013011-3233031112000211-1101123101132112-3313031303322123-3031011203003332-2313103211033230-0303310123020203-2021021323200303"></a>

<a id="canonical-0130220032323231-1231123023002113-0313013011223123-1023030330030100-1010333313313133-0313130210233021-2000100123303100-3121012201031123"></a>

## mobile_identifier property — web_mobile / 010120111023 / 4

Type: `"string"`. Computed.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Upstream description:

Mobile identifier type

&#8203;- HEADERS: Headers

Headers.

Receipt-pinned upstream constraints:

```json
{
  "default": "HEADERS",
  "enum": [
    "HEADERS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212330132133332-1113231322010022-3002311000313330-3313300313323013-0023100310233120-1130312131321302-1022212011322202-0332003230310233"></a>

## Next pages — web_mobile / 010120111023 / 5

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1103002330300201-1231313031223333-0303202310013222-0023112211102302-3010003113123000-2212320100321323-2032133101121303-3112310212212213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311020212020202-3112002132033313-0001032002133111-3223000002332333-3301033120310031-1320131033210021-0102103132130033-1133112102023002"></a>

## captcha_challenge — captcha_challenge / 120013021022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- captcha_challenge

<a id="canonical-3111302002301221-0303313233020120-1031033111201122-0312010113202031-1000003011031001-0231220321010033-2020111221332231-2320103301321213"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

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

OneOf alternatives in this subsection:

- [captcha_challenge](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3111302002301221-0303313233020120-1031033111201122-0312010113202031-1000003011031001-0231220321010033-2020111221332231-2320103301321213)
- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0212033300102210-2322130010322222-2030202300113231-1111303300313333-1032231121102300-0101020111100333-3330020033210300-1322203202210323)
- [js_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3330221111033030-0012113213012002-2033023113003223-1233203023311203-1000312213200220-3232232211323130-0331012203001311-2211203302323021)
- [no_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3310301110320021-0213111113212022-3132132101000131-3200220321201123-1332220102000021-3003110110011011-3300301312033311-3223313332033121)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3320322221113222-1301022123211111-0332333001003130-3223221331301230-0032202300300213-3012213113233021-3011001102203211-3312103133001231)

Select alternatives according to the provider validators above.

<a id="canonical-2122310222032131-1030101111021210-0211122301121111-1332200100300013-3323120000132132-1121100032102200-3132313231303132-2213210032200021"></a>

## Direct properties — captcha_challenge / 120013021022 / 3

<a id="canonical-3133131001230211-0013022003013112-3303002221021330-2330323220102020-3000310300123231-2322320303002133-3331111312220123-1230011031302331"></a>

<a id="canonical-2123111311113203-0232103030123200-0302231020222312-2031220100103033-0222231112303022-1023301202002221-0122231223312322-2122021300012212"></a>

## cookie_expiry property — captcha_challenge / 120013021022 / 4

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

<a id="canonical-0331311213201131-0112311012101223-2113330032320122-1113012022023033-0210033122203200-2330313210131211-3333330032230303-3331011202033310"></a>

<a id="canonical-1030202313032002-0002003012312020-3312001012122223-0003333312210031-1303313033023021-3332111203213132-1132203121021102-2110211321323133"></a>

## custom_page property — captcha_challenge / 120013021022 / 5

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

<a id="canonical-3230100221223313-1023133010310211-1233331020110010-3121000332301333-3000311301100010-2310231332331120-0330311130333130-2100000110230021"></a>

## Next pages — captcha_challenge / 120013021022 / 6

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132022100212031-0233312311322220-3201222101300222-3221230103033001-2231203313311032-3103103202321321-3130303003222003-1321230132130212"></a>

## client_side_defense — client_side_defense / 011022001001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- client_side_defense

<a id="canonical-3003332102113322-2031011133211120-1103322331313010-1312132233332002-0003133212220023-3202120122022000-2311030020010312-2311231210012002"></a>

Type: `"single"`. Computed.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense Policy.

Receipt-pinned upstream constraints:

```json
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

- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3003332102113322-2031011133211120-1103322331313010-1312132233332002-0003133212220023-3202120122022000-2311030020010312-2311231210012002)
- [disable_client_side_defense](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2220201123022201-0023320203013210-0103023102310013-3021121032321033-0322032301001210-1303311000012110-3313332333310202-3323311313300010)

Select alternatives according to the provider validators above.

<a id="canonical-3332013330103230-1103301023210122-0310213102112122-2023103202301221-3232013331200100-0312013222022000-3001310200231010-3133223120120332"></a>

## Direct properties — client_side_defense / 011022001001 / 3

- [policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312): complete subsection reference.

<a id="canonical-2313321222023312-2022002303011221-1322313020011311-3100120121000222-1330112020131331-2021212031200303-0211230112220122-0213200132203020"></a>

## Next pages — client_side_defense / 011022001001 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103330223211222-2202330121003131-1300021201031221-1323301131300000-1301021313012203-3302220112203322-2213323323032013-0331130213122323"></a>

## client_side_defense.policy — policy / 030132230002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- client_side_defense.policy

<a id="canonical-2032100133212033-0133230121223223-0132122220003031-0233203013331120-1101321212231210-0232211220301112-3332210130232322-0201121021002310"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Client-Side Defense policy.

Upstream description:

This defines various configuration OPTIONS for Client-Side Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

<a id="canonical-2130221302231003-1310012030131101-1023131122220220-1102301022013132-1002131311113120-0130310101003200-0223123320301102-0113203211222021"></a>

## Direct properties — policy / 030132230002 / 3

- [disable_js_insert](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1120031231310013-2213331111201100-0322301022332032-1122112113123302-0110021302002323-3031011320320120-2232112322333012-3233110100012303): complete subsection reference.

- [js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0222220303323112-0121100221220210-0021322231021322-1321200031233313-1330321323133112-2012310301101131-0322102311302203-1200032231112012): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231): complete subsection reference.

<a id="canonical-3322322012010101-2102122023133233-0011202320102212-1333221300311022-2013111010333110-2031233113013302-3010120101033332-2002113221022132"></a>

## Next pages — policy / 030132230002 / 4

- [client_side_defense.policy.disable_js_insert](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1120031231310013-2213331111201100-0322301022332032-1122112113123302-0110021302002323-3031011320320120-2232112322333012-3233110100012303)
- [client_side_defense.policy.js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0222220303323112-0121100221220210-0021322231021322-1321200031233313-1330321323133112-2012310301101131-0322102311302203-1200032231112012)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1120031231310013-2213331111201100-0322301022332032-1122112113123302-0110021302002323-3031011320320120-2232112322333012-3233110100012303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223302031112311-1131211000111222-0332023313101211-2002132002310233-0323332213311312-1113013300002000-3122031212211002-2202300022000302"></a>

## client_side_defense.policy.disable_js_insert — disable_js_insert / 000203021102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- client_side_defense.policy.disable_js_insert

<a id="canonical-1223302021233201-3220301302223212-1230232331121331-0230231111223233-1033001220202022-0301133221220303-3120212323310011-3030223031200222"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2123310201232103-1202330212220002-2132221010113121-0131203020202330-0210202311323222-0132131132202030-3102001003100203-2301111232220101"></a>

## Direct properties — disable_js_insert / 000203021102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110132010212321-0122232133002020-0200222123003030-3013220313031223-2002133232120131-0212233012103120-3312111201113301-1013133211223203"></a>

## Next pages — disable_js_insert / 000203021102 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0222220303323112-0121100221220210-0021322231021322-1321200031233313-1330321323133112-2012310301101131-0322102311302203-1200032231112012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003131313013323-0301220303133223-0122021111310022-3303111212232333-2222300111200323-2300211020210021-2131133031223223-0133220211000332"></a>

## client_side_defense.policy.js_insert_all_pages — js_insert_all_pages / 211113122120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-0310231033112303-0321000130023222-0333132033322013-1002112101003203-1112320220302211-0110032322210102-1012331131010123-3130232022210230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for js insert all pages.

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

<a id="canonical-3033320323211322-3023023303231331-2030003311211211-0301323001110110-0001311121301231-2232203222110113-2003133112132223-2310310311330133"></a>

## Direct properties — js_insert_all_pages / 211113122120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100322221123213-0031123101210323-2112102023230213-0002020022102220-1210200123103032-3233023031123003-0303321100203013-3121201110100113"></a>

## Next pages — js_insert_all_pages / 211113122120 / 4

- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122210103031132-3021311000112110-1020121313330322-3130200302333213-1321112102112000-2333310202321011-1030210013113210-3210033213222331"></a>

## client_side_defense.policy.js_insert_all_pages_except — js_insert_all_pages_except / 232033320310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-1102130032203112-1100210320203322-0022232223001233-0030000213023110-3310212001132100-2123001312010212-2223200112223213-3310310131102021"></a>

Type: `"single"`. Computed.

Insert Client-Side Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1211013123010022-0022002013301002-2203020300033021-3010302223013332-0122222231202220-1210001310331230-1312001322012221-1110003230021122"></a>

## Direct properties — js_insert_all_pages_except / 232033320310 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301): complete subsection reference.

<a id="canonical-1222032200012111-0023113032310302-2122221323133001-3033133311032013-2203232003321100-2221030311032030-3311302000102200-3003002001011023"></a>

## Next pages — js_insert_all_pages_except / 232033320310 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203213212211203-3330302322133210-1021213021113213-0323121111032322-2221002122020222-3101203100113321-0030003321123330-0203111302131331"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list — exclude_list / 133302111332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-3120001101030330-0312331131211030-0100021220321022-1210113231011232-1023212302200213-3332112211332132-0203022130002333-0231133112211312"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-1200233110101001-0202333023233331-3331012321213013-1003031032113210-3113211222113021-1202310110300203-2120130111322320-2301222033212322"></a>

## Direct properties — exclude_list / 133302111332 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2110201111223110-1020222020303312-3101330021121232-3302231003003032-1031123000220102-0001330103111131-1302221021320312-1312013211000211): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1230210212203022-2210000113012131-2212030200013030-0333031203212033-2023001303133133-0300031131203310-2233233011333231-3322302020313332): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2201220100131332-1131310011202211-1321330303332332-0123222211103112-0033113301110312-3131210033203330-3222122231033131-2100310031023021): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0202230201133233-0121302210101011-1320220122120202-0312330023310230-3112130022131111-0122321032113320-3311122111031211-1122013133320133): complete subsection reference.

<a id="canonical-3102133131331020-1002111002232022-2112123133022330-2033210011221103-2003101003000010-0232320203101120-2331100030000221-1320230030022310"></a>

## Next pages — exclude_list / 133302111332 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2110201111223110-1020222020303312-3101330021121232-3302231003003032-1031123000220102-0001330103111131-1302221021320312-1312013211000211)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1230210212203022-2210000113012131-2212030200013030-0333031203212033-2023001303133133-0300031131203310-2233233011333231-3322302020313332)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2201220100131332-1131310011202211-1321330303332332-0123222211103112-0033113301110312-3131210033203330-3222122231033131-2100310031023021)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0202230201133233-0121302210101011-1320220122120202-0312330023310230-3112130022131111-0122321032113320-3311122111031211-1122013133320133)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2110201111223110-1020222020303312-3101330021121232-3302231003003032-1031123000220102-0001330103111131-1302221021320312-1312013211000211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113212310130332-0113221312000220-3011203031010222-2113212331000311-0231132131000120-0223233300301201-0012023300323100-2232023111113321"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — any_domain / 110123203330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-2312123021232121-3010002033303222-2110031210102201-2023013300221101-0321221223220310-1333130202321302-1333232013201231-2032303003112113"></a>

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

<a id="canonical-1113322111213222-2102311212021020-3201203001032331-0130110131130012-3020323330120111-2102301321230320-0333000111102031-2101203201322201"></a>

## Direct properties — any_domain / 110123203330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303110333330102-3222021302102003-0001111220030231-0303111212032130-2120211330123313-1012131013211110-3033301033112232-3033022022002222"></a>

## Next pages — any_domain / 110123203330 / 4

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1230210212203022-2210000113012131-2212030200013030-0333031203212033-2023001303133133-0300031131203310-2233233011333231-3322302020313332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000023121220133-0131131222301132-1033222031030301-2221211232203220-1210303111123031-2311110100001102-1332102331123011-1122032022231330"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain — domain / 032202300123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-0013322320033122-0001222032121010-1000121233221102-2003322201233100-1313312210223000-3221230122333230-2322213210111211-0111301310013100"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

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

<a id="canonical-0012010131211002-0321000203112032-3113321303122111-2233210211313020-1021111310303121-1220121200000321-2201332102022233-1201333030032211"></a>

## Direct properties — domain / 032202300123 / 3

<a id="canonical-2110020113312233-0203212223200032-0231333033212102-3320300000003021-1301212012003133-1230101220020300-2102300003131222-2032332202320221"></a>

<a id="canonical-3203010103230320-2030322132220310-3010213321310322-2222330031121203-0013130231102333-3023310010110012-1301310103100023-3022001133212323"></a>

## exact_value property — domain / 032202300123 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-0031033230001322-1212213020121120-3120032331233121-1012322021310220-2332111011010232-3213210201332331-3020200022200220-0312132132232021"></a>

<a id="canonical-0003001103132120-0232021321201113-1121122330133313-1202230200121122-3202212303023312-1320333001313030-0020231312330021-0331220332122101"></a>

## regex_value property — domain / 032202300123 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

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

<a id="canonical-1212323210231112-1202322030212110-2112020011201311-3230023023120030-2330102010132201-3330230113301100-0011311031230001-3103020311232310"></a>

<a id="canonical-3323223132123102-1133131223211023-3102200032101110-3231302011122223-2233023320320012-0203301221213011-2210232221002201-1300303022121023"></a>

## suffix_value property — domain / 032202300123 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-2123201003002231-2303100003112101-2210313001132302-0212130122222130-2111330332033003-3323312022101111-2113322333030012-1001212001333031"></a>

## Next pages — domain / 032202300123 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2201220100131332-1131310011202211-1321330303332332-0123222211103112-0033113301110312-3131210033203330-3222122231033131-2100310031023021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333303320110001-2210103001232202-2120122111110031-2101330203031021-1101013003021032-1331233121030330-3010110103003333-2323023032020231"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata — metadata / 221210131100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-0002213310113330-1211103212032013-1030333023032203-1201213212332212-0111033022110231-2121200323120333-1023003323220221-3100033210313302"></a>

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

<a id="canonical-3121220211210302-0131133321233102-2321303111320030-1022212212001323-2023010102221001-2202301222210303-2301301002301222-3001322320000303"></a>

## Direct properties — metadata / 221210131100 / 3

<a id="canonical-1210020322112032-1321132012221012-3012221111112120-1320102231331012-1000301200122123-1133121030332131-2303011312031000-1021132230031322"></a>

<a id="canonical-0331210302310012-3323230031133133-3032331330012131-1021333231013313-2130121312121033-3311113002332320-0013231003330123-2022100333113113"></a>

## description_spec property — metadata / 221210131100 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1202020323100032-2010120320033022-3310301201001001-0010002003031210-3010322103332211-1320201121222200-1112110233113111-0030111310301313"></a>

<a id="canonical-0222201012310330-3030120011301111-0203033110330322-0022232322112302-1201210323233212-3011232112022120-0033331313001021-0100211120130210"></a>

## name property — metadata / 221210131100 / 5

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

<a id="canonical-1311020101212300-2023320300301000-2112012003033301-0220331200132310-3002123223311120-2023322230110323-1312211300111312-0333113333111312"></a>

## Next pages — metadata / 221210131100 / 6

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0202230201133233-0121302210101011-1320220122120202-0312330023310230-3112130022131111-0122321032113320-3311122111031211-1122013133320133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303033113222312-3120211030223000-3123110223111021-2220321302230033-0031131230213300-0022032220220103-3213011122220301-1030310101123211"></a>

## client_side_defense.policy.js_insert_all_pages_except.exclude_list.path — path / 000021130001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0232232200020321-2303132103323022-1231212303013131-1300102333010103-3000031122300323-1323331021303112-3131211122212202-3031303321010232)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-2013001320231303-3330021123203002-3201101311303030-2021012130230131-3113132302130001-1022303003231223-0323111312333121-0021213103032002"></a>

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

<a id="canonical-0023220300321003-0322113100221033-1313303013332303-1230332001212120-3022132232232021-3322211033101210-1210023230003233-0220013221112311"></a>

## Direct properties — path / 000021130001 / 3

<a id="canonical-2003131110312210-0122030121113223-2131113112202202-1122211312203003-1223012111132103-3302313122233131-2300033120302311-1023130212302011"></a>

<a id="canonical-3123033212331203-0310121023112311-3000232211231113-3330231110011001-3122033321103020-0013022223233123-0002321030231312-3111111021132131"></a>

## path property — path / 000021130001 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-1220033111021303-2003130011101021-1210333231003122-0322030001233312-1000310120210010-0130033123323033-0020222212310030-0201323013101210"></a>

<a id="canonical-1001323030321021-3210321203333303-0120001120211231-3200032230022202-0310031321122323-1221013120103211-2313321112203201-0111100133003220"></a>

## prefix property — path / 000021130001 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-1210133020222311-0123021131213120-0233033132130113-2223133131223013-3122002311302231-1232232310030101-2000112021300133-2010010233023323"></a>

<a id="canonical-1312002113332123-0330213322013101-2022200103022130-1133222210011023-2030001012020133-2321310203032123-2312131023102332-2131011320021013"></a>

## regex property — path / 000021130001 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-0013010132111331-0330131303030102-1100312131101032-1021030020312022-3003123002120223-1133003232302000-2213010200010120-2233313133333131"></a>

## Next pages — path / 000021130001 / 7

- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0133013031021210-2211222112323231-1111031021312332-3303330323323011-1221220132210102-1213231102300233-1322000332013100-3031233213123301)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231230322002313-1232133112322021-2021212131001112-2333132311131301-3210130020220323-2112300013002333-3300300123103102-2200222000321122"></a>

## client_side_defense.policy.js_insertion_rules — js_insertion_rules / 011320120100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-1033100012322010-0122113233000310-2000321011101200-2013332330323212-1223123301000200-2121223030313231-1231030110103111-1301312121212121"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Client-Side Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030331110321121-0003231333022100-2022222112302131-3031200221001131-1022022310211103-3311100202201311-1112311002012303-2303111021232122"></a>

## Direct properties — js_insertion_rules / 011320120100 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101): complete subsection reference.

<a id="canonical-0113001331013101-2132320103103203-0102233330130230-2023311113103232-3322022000102221-2133033200301300-0000132231033001-2202022321322030"></a>

## Next pages — js_insertion_rules / 011320120100 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311031320022013-0303320333001310-0132301133113033-1110000220330032-0213120213233323-3222103111111331-0321133130200223-3220331230221210"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list — exclude_list / 003012032121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-2211130223020030-1231012110303222-3122333312301310-3202221030113322-3332222101330012-1011133001131113-2121100211111301-1332222321311133"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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

<a id="canonical-3101022310003301-1230233203210030-2112012313300010-3003003100301221-2301030200303212-2321133232330210-3001123301132300-0100200020020111"></a>

## Direct properties — exclude_list / 003012032121 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3012223032110033-1130120112300121-1233121222312313-0233222022101202-2200100020332031-1302133220221200-1102212111211210-2321303321322123): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3131121320121202-3221213223220000-0331022122223230-2223110103130113-1101231130221123-1323101201232101-0023211212000120-2213321211220323): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1332223322023101-3233103313120113-1012010301230212-1113312012303223-2210322002311311-1301222222220112-3002220231022011-1231023231223321): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2332022103111110-1222123202220231-3023030032100301-3101123031211231-0233101132200313-2121013323222301-0111130201132120-1311231313032201): complete subsection reference.

<a id="canonical-0101102213210210-0031232122323320-0133230333002203-2303110001200302-0111111233100313-0111133022213020-2313303030002321-0230323002112101"></a>

## Next pages — exclude_list / 003012032121 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3012223032110033-1130120112300121-1233121222312313-0233222022101202-2200100020332031-1302133220221200-1102212111211210-2321303321322123)
- [client_side_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3131121320121202-3221213223220000-0331022122223230-2223110103130113-1101231130221123-1323101201232101-0023211212000120-2213321211220323)
- [client_side_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1332223322023101-3233103313120113-1012010301230212-1113312012303223-2210322002311311-1301222222220112-3002220231022011-1231023231223321)
- [client_side_defense.policy.js_insertion_rules.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2332022103111110-1222123202220231-3023030032100301-3101123031211231-0233101132200313-2121013323222301-0111130201132120-1311231313032201)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3012223032110033-1130120112300121-1233121222312313-0233222022101202-2200100020332031-1302133220221200-1102212111211210-2321303321322123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301103200023001-3203113211030320-3133002323023323-3300331022330333-3103202202010130-1013032103110023-1031213321012130-0321311210012123"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.any_domain — any_domain / 313131213011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-1302023022322231-0003222003231102-3322123201022211-1123200202211113-3231221201113202-0131130312001302-0301211203212123-2123331030231023"></a>

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

<a id="canonical-1033320321002023-0333301003211113-0320020100213003-2301132132311233-2011313332301311-0112133320333020-1003112300330333-0202233000120302"></a>

## Direct properties — any_domain / 313131213011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3200310313012332-1123121232111121-1000211101001101-1113330300220011-1130022200323212-1231200023323332-2311301331220101-1030202012201130"></a>

## Next pages — any_domain / 313131213011 / 4

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3131121320121202-3221213223220000-0331022122223230-2223110103130113-1101231130221123-1323101201232101-0023211212000120-2213321211220323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202332100330222-1133202203111201-1201000303303330-1030103113213101-3130323211211310-0102011130001112-2120123330303001-1132221222111312"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.domain — domain / 321220201113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2023311133231002-2302301330003101-3203331131123012-0113310202330102-1203101111200333-2020232103023023-1121202303202103-1120130313302330"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

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

<a id="canonical-2300023112100200-1302221212323111-3033013300032311-0331300212232030-2132220132320230-1030331111230133-0210202233130312-1012303031230202"></a>

## Direct properties — domain / 321220201113 / 3

<a id="canonical-3222101221221000-0300221130033210-2311030303320032-2013032203012300-1122220233221001-3220111030222022-0312020330122132-3233001100021300"></a>

<a id="canonical-0013013132123101-2003011020333311-3331333312331202-3310111312112130-3031322221030111-3112002023123100-0023302200032201-1212011132210002"></a>

## exact_value property — domain / 321220201113 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-0231201332102203-1001232110322202-1002303003011131-3211123301333011-0310203223313033-2103232113132201-1001323020233011-1112322203133320"></a>

<a id="canonical-1313332211301003-0320103332310222-3303001021203122-3122033021013003-3312330113231222-0202012210132330-1003110113203123-0120220210303330"></a>

## regex_value property — domain / 321220201113 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

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

<a id="canonical-3321113011321302-3301223230101131-3032102213311133-1223132020201320-2111113330220311-0210000232122211-2232320200321200-2303123221233131"></a>

<a id="canonical-2003231231113003-0130210323023333-0302300010122120-2223222110112122-3102310111132312-3231202033331001-3002003202001002-1121333101001220"></a>

## suffix_value property — domain / 321220201113 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-2331103131333302-1201331012231130-2231112012132212-3213011010233313-2112003123123313-0312111222212012-3222022032111303-1133221303202302"></a>

## Next pages — domain / 321220201113 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1332223322023101-3233103313120113-1012010301230212-1113312012303223-2210322002311311-1301222222220112-3002220231022011-1231023231223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321012210002123-0101313233312022-2220100222232030-2201331212123311-2120311121012102-3030011201110312-0133003103212312-0023030010332130"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.metadata — metadata / 333013003031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2001222320322020-2203332221001323-3212311122000001-3200301333133012-2312331003301100-3112131130231300-2123200101210012-2031210022000332"></a>

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

<a id="canonical-2111230102223131-2330122211203113-2013232000012223-3121330121310233-3222002213010120-2230300100203110-3123102101100011-2321000110130322"></a>

## Direct properties — metadata / 333013003031 / 3

<a id="canonical-2310030131330320-1010020321021300-2021030012203023-3101020100203022-0021112013223312-1213110321202030-2220203122301131-2333120032201001"></a>

<a id="canonical-2200202122100032-2002013313320102-1133103302330331-0110202033331110-2033001112113311-3020001131201111-2213020311333322-2311120200021302"></a>

## description_spec property — metadata / 333013003031 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0321023223233132-3123223332233130-1303101032231303-2011210112332021-0202213333230322-0032102120312130-1133121310020020-3211102102122121"></a>

<a id="canonical-2102131203112322-2300321021010230-0320333312133313-3132221123112322-2213023302301112-2211012332032313-2303000300003210-3110320310132101"></a>

## name property — metadata / 333013003031 / 5

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

<a id="canonical-2233332022233113-2012103113001333-2300101232132203-3321003200212111-1221031223203211-0021002320101033-2211201123321220-2312010020033110"></a>

## Next pages — metadata / 333013003031 / 6

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2332022103111110-1222123202220231-3023030032100301-3101123031211231-0233101132200313-2121013323222301-0111130201132120-1311231313032201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032300112122210-3113130001301001-0333031101203202-1301222101203130-1332222212113132-3033230320032322-1211231011233020-3012130300203001"></a>

## client_side_defense.policy.js_insertion_rules.exclude_list.path — path / 233233312302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-1322132110031212-1101110031210123-3013211332103112-2002333001213131-3331222030032310-2033232330032112-2212102002100013-1200200023112023"></a>

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

<a id="canonical-3010132133233100-0303022200133133-3021211133202303-3233023022232310-1030231010012311-3100102030112123-2301333022200113-2201202230221103"></a>

## Direct properties — path / 233233312302 / 3

<a id="canonical-1302212100000100-0232000220310301-3321130331213111-0131301121213130-0012203020311332-3213200000201222-2122330120032100-3122012122322121"></a>

<a id="canonical-0233312302312223-3222020212103223-0001220202312133-2011333003122321-3203131012320002-2103120132122000-1203200221210032-2322011211302200"></a>

## path property — path / 233233312302 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-2113210221301102-1213120132301333-0002103121112323-3221112222232032-2031023000120003-0121003203321113-3231033322130131-3011312133322033"></a>

<a id="canonical-1332110023210023-1230113133323231-3001331003222300-1121003013223110-1012313323111211-2131100120211021-3323313011032003-3010220133111211"></a>

## prefix property — path / 233233312302 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2003120121330122-1232121312001302-0201330030100332-3313201010002000-1211231023222131-0123103001113332-1210032002231013-1303333122213211"></a>

<a id="canonical-3022001323301322-1022021203323023-2103011322020033-0132320001112013-0100033011301222-3322330302101021-2300302110300000-3300111023121121"></a>

## regex property — path / 233233312302 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-1021013000210031-3323231130322313-2022000303313332-2313113232333300-1133002100220102-2123310131022111-1211312020022000-2010002112100302"></a>

## Next pages — path / 233233312302 / 7

- [client_side_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2232103020322200-2100001030213030-3002323122001031-0301323110222023-3211013012322123-3310313233001322-0012323001233012-3023322013210122)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311023312001123-1122310230220331-2303100213323101-1320133030310320-3311301011302103-0031300213230210-3132000212100110-2110220122013301"></a>

## client_side_defense.policy.js_insertion_rules.rules — rules / 123011302331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-3200230122200032-3300301330231113-0031111221301210-0303203331322003-3022320123000303-2122011032030122-2222101330123232-1301021102210323"></a>

Type: `"list"`. Computed.

Required list of pages to insert Client-Side Defense client JavaScript.

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

<a id="canonical-1113033011320312-0313310130021312-1313330320123103-2030131023130230-1031230322223213-1201212123211330-0001102131003032-0303033031023222"></a>

## Direct properties — rules / 123011302331 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1032233310300203-0121031021221100-0200130010110113-0020131202333001-0111322332320222-1103201100110022-0010332101121003-3022120310323222): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0322112110103312-2132031001323311-0100200330331100-2003310002322221-2030030202121032-0232221323230110-3323313020203102-1113013322133011): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3211221110133300-0010101132222313-1303220022201110-0323013113201030-0033121131120201-0002231112001311-3111012100020011-0002121300101200): complete subsection reference.

<a id="canonical-1200101313021122-2323131021023130-2133023032232002-2033310320003213-1213311232120132-2021321100303102-1330132222111033-0313122021313133"></a>

## Next pages — rules / 123011302331 / 4

- [client_side_defense.policy.js_insertion_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1032233310300203-0121031021221100-0200130010110113-0020131202333001-0111322332320222-1103201100110022-0010332101121003-3022120310323222)
- [client_side_defense.policy.js_insertion_rules.rules.domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222)
- [client_side_defense.policy.js_insertion_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0322112110103312-2132031001323311-0100200330331100-2003310002322221-2030030202121032-0232221323230110-3323313020203102-1113013322133011)
- [client_side_defense.policy.js_insertion_rules.rules.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3211221110133300-0010101132222313-1303220022201110-0323013113201030-0033121131120201-0002231112001311-3111012100020011-0002121300101200)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1032233310300203-0121031021221100-0200130010110113-0020131202333001-0111322332320222-1103201100110022-0010332101121003-3022120310323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211010033203210-1022201210232112-1010012303112311-1000113120113202-1330133101313212-3112232330113303-3121130123313333-0133232323303001"></a>

## client_side_defense.policy.js_insertion_rules.rules.any_domain — any_domain / 322111312130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-1021102132132221-1013333031222012-3003100010313001-2301010222102001-2022302112130213-0033312020301103-3222011223203310-3023011220032120"></a>

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

<a id="canonical-0233032121321300-2201313321011113-3101013300020022-3311311203211101-2031210020111200-0223221032203302-3201122210330220-2112003333013022"></a>

## Direct properties — any_domain / 322111312130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030330333030130-0022031123113211-3311021322320321-1132210212031323-1321301202220100-3002031123322111-0120321021112302-2132323121121111"></a>

## Next pages — any_domain / 322111312130 / 4

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0331232032213020-3320203122130001-0201311133300031-2321200333311022-0222110302110002-2013020031210203-2213330012122123-2311030210200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013231210010233-3302123220002233-0300001101030303-2000200331230121-2132312003312230-1330322220112220-0302211320032321-2211010103122032"></a>

## client_side_defense.policy.js_insertion_rules.rules.domain — domain / 303311011302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3323133330301202-0201232102131232-0233332303300230-2311323213033323-3313233201213303-0302020212331201-1023311222230010-2020331133300133"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

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

<a id="canonical-3012033310122302-0000122211123113-2122323200330013-2212232130210211-1033003001111121-0233330100232302-0201213113103113-0213322211323103"></a>

## Direct properties — domain / 303311011302 / 3

<a id="canonical-1330100022023110-2302320310103211-1022102312200202-3210333033001312-3001200210130030-0323121023310333-3033000310120113-3021131130022021"></a>

<a id="canonical-0003013303023231-0222321312123132-3110122032332222-0021320213002113-3222202222321112-2303131232332133-1312121113111020-0212011233311013"></a>

## exact_value property — domain / 303311011302 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

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

<a id="canonical-0220223112133321-3331321201310120-1011330010233111-2203110131031300-1003321100031132-1123301330331000-1233030221132003-0312310012130133"></a>

<a id="canonical-0301313303323203-2202221012312301-1331033101302211-3230011211130323-1032302122002133-3021112320013020-2023232203333322-3033213122300132"></a>

## regex_value property — domain / 303311011302 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

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

<a id="canonical-2123010300031203-2222230320221002-0232033102133321-3002032132211202-2312111121330212-3333333122112003-1331033311311130-2222231333020322"></a>

<a id="canonical-1213112303022011-3100300021220332-1111032302302313-1220001222103101-0310130312113330-3231100330203112-3312210021230301-3001310303103132"></a>

## suffix_value property — domain / 303311011302 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
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

<a id="canonical-1033333201202000-3232112100122322-0213330212023103-2231213323022323-2121310212132303-0111311233133313-3033210223122322-1122033303313213"></a>

## Next pages — domain / 303311011302 / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0322112110103312-2132031001323311-0100200330331100-2003310002322221-2030030202121032-0232221323230110-3323313020203102-1113013322133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302311230131110-0321220320311122-2020032202022123-3200120000212232-2333010032332121-0220112332313011-3031310301110122-1121102130321332"></a>

## client_side_defense.policy.js_insertion_rules.rules.metadata — metadata / 211021303201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-0010021131003311-1313020211232133-0102112311302320-1330013030330231-1223210323032303-3012322121320233-2330123312101100-3333033331113222"></a>

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

<a id="canonical-2233130021302230-2301211210233010-2320203011100203-0103030000110030-0002233203022131-2203312132213300-3110100221321301-1001100103210321"></a>

## Direct properties — metadata / 211021303201 / 3

<a id="canonical-0130003130032311-1020011221312201-0123312112031300-0301013001203003-2122232031120003-2221011100010023-2021031330111212-0210011103013322"></a>

<a id="canonical-2020003123032322-2121330302312120-3331010101022220-1201212101123233-3122010331101021-1232010032212121-0233000100232102-1013213030311332"></a>

## description_spec property — metadata / 211021303201 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0010212330321201-2322023232300011-2003133121001221-1003103330310301-0022102210130201-3111300230113210-0302310112023021-3133311120012111"></a>

<a id="canonical-1133101233230233-1312312120102231-0121013201321310-2031113023002322-2103131230111220-0110000132032122-0322323022002013-0231312320200020"></a>

## name property — metadata / 211021303201 / 5

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

<a id="canonical-1313131230321230-0103010002210100-0023111112132020-1313211332112001-3023311012030310-0203000113323030-0101313021300122-0332303331013000"></a>

## Next pages — metadata / 211021303201 / 6

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3211221110133300-0010101132222313-1303220022201110-0323013113201030-0033121131120201-0002231112001311-3111012100020011-0002121300101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320322213323033-2013133200132030-1132002011222333-2132313200203112-0011322313220111-3122010231033113-1321002023121020-2011003121011002"></a>

## client_side_defense.policy.js_insertion_rules.rules.path — path / 201232013132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231)
- [client_side_defense.policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3120200111312023-2212013133113103-2001030110230123-0332323322011313-3113310102030020-3011202210213102-2321101113211131-0011321032213312)
- [client_side_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3213030130221101-3102222101333002-3333202300030310-0310101002303303-1303232111130030-1213022232200303-2303310023113022-1332222302101231)
- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-3110302001310002-0210231233203100-1013320212021002-1112210130102120-2210123023102231-2211032132321122-2311230323100120-1100021023300023"></a>

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

<a id="canonical-1311332012113120-1000011303032323-0021121311120130-3110322200200233-3312203220311130-3030223201001311-1123000001022231-0220331310220232"></a>

## Direct properties — path / 201232013132 / 3

<a id="canonical-3133330133311133-0000001003330122-1120133301320331-3313030010030110-2303112313101030-3333103110131212-1223112223102021-2101213033232023"></a>

<a id="canonical-0332202020113332-0002330333101230-2131223022113310-0031122211301302-3331323202233323-0233333202223232-1123211020122323-2313333222213220"></a>

## path property — path / 201232013132 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3011233113231001-2103202313130102-1222333202320222-1130330101033130-1222023301000002-3213212003302101-1000221101331012-3001133230002300"></a>

<a id="canonical-3011023110313221-0332013330212332-0210103020020100-1111201201333323-3302203132300032-3122311303012211-3220102102210113-1033130220302313"></a>

## prefix property — path / 201232013132 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-2200323010020202-3200332212121222-1303000102102123-0033201120020110-0011110333101021-1300130312013123-2030302111311003-3203310211330000"></a>

<a id="canonical-2021303103003102-0033233321230212-2103221322121022-3233031213322331-1230303010200303-2231222211331231-3303121310011023-3313123303331301"></a>

## regex property — path / 201232013132 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-3303201331113331-2311133220210110-3223213111023002-1211202033102132-0130013230312211-0322210133010232-3100022210200030-0223323200103331"></a>

## Next pages — path / 201232013132 / 7

- [client_side_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2230201130322102-2313103001002212-3103231131202123-3312132213332213-0302000332310333-3110032231331023-1210321321210233-3321012200203101)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3330100113233220-0200103010232013-0220221020032201-2012031211002200-0321000202211233-0202201031003232-2322233111100113-1212133131020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130231012230100-0310001331133100-3201332131122221-3331221002321123-0123212320032103-1113102131310102-0030203231111010-1302311030011321"></a>

## cors_policy — cors_policy / 211131002020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- cors_policy

<a id="canonical-1021113222023013-2112213003123130-2310110301101012-2230123300011202-3131230211333112-0332202330320122-1233123000232220-3302330022130303"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021132023222033-1131122031202003-1221220103322010-2110232211311023-2332220112211222-3022232023031020-3223123231033301-0002203301021201"></a>

## Direct properties — cors_policy / 211131002020 / 3

<a id="canonical-0223330203231202-2021321032120022-2221320312012322-1031032312212213-3311103312331133-2000230132332131-2101131202101032-1130200330130121"></a>

<a id="canonical-1231012232010010-0103122033111133-3200310331001310-1231122112111020-1223232320003233-2122320113211322-3230021203332332-1013122101233230"></a>

## allow_credentials property — cors_policy / 211131002020 / 4

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3031003211322112-0032211212213031-2011003122012123-0310233032222203-0102123202210212-1021320013313300-0220121230031023-1131100313212330"></a>

<a id="canonical-1332030123320033-0033011100001013-1310200222223032-1003302112212112-1111113203131311-2110022333313220-3033232112201332-1132023201213023"></a>

## allow_headers property — cors_policy / 211131002020 / 5

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

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

<a id="canonical-2132113121131012-1321312202120131-3203021313201010-0313001030100131-0332020223032320-0212103302120310-0223213230332330-2333331011121123"></a>

<a id="canonical-3020312323002200-3312213322332200-2101222203031102-0130330113100010-0223021231103030-0212133130032232-1233132323100103-3003220221322223"></a>

## allow_methods property — cors_policy / 211131002020 / 6

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-2133303301212112-2232030300112021-2310012122220203-3003311300123323-2131122121233113-2302002310101330-3100120001100323-1002320322301030"></a>

<a id="canonical-1313021200201220-3012031331322023-0222222033032131-1231103220001012-0100313232213101-1321132033011231-2300003232232132-0203133332013121"></a>

## allow_origin property — cors_policy / 211131002020 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2023230203022130-0231000102103122-2202311211123102-3123212322301023-2121230332032302-3032122032011210-1003220122200210-3131113302203211"></a>

<a id="canonical-3303222210233222-0020313220232312-3032210033311102-1200030100013223-2103313311201322-0132331321013202-0033101302120123-0120032233113220"></a>

## allow_origin_regex property — cors_policy / 211131002020 / 8

Type: `["list", "string"]`. Computed.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2023310201233322-2033122032202023-0033011131221022-2211020110323332-3213102223310220-3031201303311322-3032313310103133-2111331311033332"></a>

<a id="canonical-0312231012001013-1111223301302300-2100312200320302-0333112312323033-1021311032203220-2331002232202311-1012303121230032-1311223021223012"></a>

## disabled property — cors_policy / 211131002020 / 9

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2121020200230122-1100310020220123-2001103213103133-1123203021331303-1010102123233233-0131121311012032-0132202201112103-0120232300102030"></a>

<a id="canonical-3021322220323223-1203103303331103-1332030330111300-3201120110011121-0012322221000032-1332013022021212-0121002311331013-1311113232313010"></a>

## expose_headers property — cors_policy / 211131002020 / 10

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

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

<a id="canonical-0333020330011110-2203132300323331-3332113121031020-0121333101321220-1000000202211202-1323000111133333-1211301031012120-2021221033231222"></a>

<a id="canonical-2113122012002312-2310130300021323-1103202122322222-0032313001311302-0122321201033023-2322021022303011-3211333222201031-2223313211020220"></a>

## maximum_age property — cors_policy / 211131002020 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-0032321230331022-3212032321320111-3133323321232300-3003202312211211-0003100331312311-2220231311111212-2131212213002332-3200321322202323"></a>

## Next pages — cors_policy / 211131002020 / 12

- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130022231202333-2302120123320333-1221321023320011-0013100310023120-1012003220131221-0100322203120210-2012002200113323-3033002301320301"></a>

## csrf_policy — csrf_policy / 111031320311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- csrf_policy

<a id="canonical-2032002311312013-0000112220310102-0133332013012133-3223011322000201-3332000003333122-2201020233330002-1021230101200311-2023231002033131"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-0100012032301212-3301102200332000-3111223033202111-1221210322233010-0110311101321010-2310330221122101-0311033333022302-0103003223232210"></a>

## Direct properties — csrf_policy / 111031320311 / 3

- [all_load_balancer_domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0230332033033301-2302102231330312-1013001111131131-1220000000133201-0310232002102301-3021123321002131-3012132133211211-0132032112003312): complete subsection reference.

- [custom_domain_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231310222213201-1322012223002113-2022030123321212-2022110110303222-3023302302320311-3101102100333021-2202201101301302-2122110111112222): complete subsection reference.

- [disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3211103022323100-2122030211321332-1301332110022133-2221223112330113-1012232203020023-3121320223023320-0020020122320321-3232033311201012): complete subsection reference.

<a id="canonical-2113030032330330-3213322221331312-3133231112031030-2321212300102110-0103233131103313-0331100113233212-1011333031003112-3313103230312013"></a>

## Next pages — csrf_policy / 111031320311 / 4

- [csrf_policy.all_load_balancer_domains](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0230332033033301-2302102231330312-1013001111131131-1220000000133201-0310232002102301-3021123321002131-3012132133211211-0132032112003312)
- [csrf_policy.custom_domain_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231310222213201-1322012223002113-2022030123321212-2022110110303222-3023302302320311-3101102100333021-2202201101301302-2122110111112222)
- [csrf_policy.disabled](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3211103022323100-2122030211321332-1301332110022133-2221223112330113-1012232203020023-3121320223023320-0020020122320321-3232033311201012)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0230332033033301-2302102231330312-1013001111131131-1220000000133201-0310232002102301-3021123321002131-3012132133211211-0132032112003312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030230012212201-2323101332001003-2313313200133003-1133000222130110-2311321313313331-0201330131131011-3301011032011103-2303022103133003"></a>

## csrf_policy.all_load_balancer_domains — all_load_balancer_domains / 133313223102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.all_load_balancer_domains

<a id="canonical-1101231213132123-1303123101020131-1100132110133101-2230010302230323-1321200333121300-1220131023302010-1113131212321033-2012233011322223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

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

<a id="canonical-1321220000023123-0221003202121301-3200102310013003-0232130002211102-3112030132002022-2110203211221202-1133202201110011-3233112020333210"></a>

## Direct properties — all_load_balancer_domains / 133313223102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121213010133130-2013210300120331-0300220030302312-2200031211230320-1021201121321012-0223030302133013-3332301003003001-1212120121021320"></a>

## Next pages — all_load_balancer_domains / 133313223102 / 4

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0231310222213201-1322012223002113-2022030123321212-2022110110303222-3023302302320311-3101102100333021-2202201101301302-2122110111112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301302032103212-3001300221120101-1200323113110231-1031011231333213-0311203012331310-3320213111331001-2102302111203213-2213331201011310"></a>

## csrf_policy.custom_domain_list — custom_domain_list / 103100131111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.custom_domain_list

<a id="canonical-3320330002320003-3022223223211231-3130233032022313-2211021300201333-2220121002002302-1023030230113133-2303022023112332-3220311023223231"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1032012312203123-1203310213213211-2023212312013301-2012110312221133-3131121133020220-1123233323302020-1032132311220032-0313302100013213"></a>

## Direct properties — custom_domain_list / 103100131111 / 3

<a id="canonical-1313210322132123-1213310333320321-2312030223111333-1233223202201331-1212131303011313-2233233313020322-0303221130323010-1031323021211020"></a>

<a id="canonical-1203113322301111-2321130122131321-0312032123202312-3013110033331201-3101111210003112-3223130001210102-2311221021123220-0311031231302321"></a>

## domains property — custom_domain_list / 103100131111 / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3111330212310132-2022320013010020-3133213000103112-2212011103322100-3110000233321221-0100030300001023-3133131302201002-1333330010323123"></a>

## Next pages — custom_domain_list / 103100131111 / 5

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3211103022323100-2122030211321332-1301332110022133-2221223112330113-1012232203020023-3121320223023320-0020020122320321-3232033311201012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213200302121001-0223220101333002-3021020231210313-0132111000211310-3303000222030223-0320020203100233-1020213213001333-2031201311031102"></a>

## csrf_policy.disabled — disabled / 113311201131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- csrf_policy.disabled

<a id="canonical-2230101132002233-3022231022320032-1313031010233131-0123020221233113-3110332303102323-3201233000213231-3202131312212033-3112032211230031"></a>

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

<a id="canonical-0022112221002323-1230303102122131-0122111230110230-0320130133211230-0023021211023121-1112000122000203-2032022120223121-2201122311321102"></a>

## Direct properties — disabled / 113311201131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033330311100323-0023200131020320-2021323220101003-1323322032113302-3203312213202110-2102022330123230-0331312001023213-3130330031133001"></a>

## Next pages — disabled / 113311201131 / 4

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112213321123103-1211131331221300-3120302110332320-3301112121320122-2130322022231201-3201011322211330-1131301321333220-1101213133330110"></a>

## custom_cache_rule — custom_cache_rule / 103223303033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- custom_cache_rule

<a id="canonical-2002233220302302-0011232132023010-3313233103213113-3232311202100200-1100103302200231-0121010123331303-0330021213321332-3230021313333231"></a>

Type: `"single"`. Computed.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0311032023320332-0221132133020202-0120132013331002-0022012310302113-3020313023310233-2120100221013002-3203203012121221-1012310300111112"></a>

## Direct properties — custom_cache_rule / 103223303033 / 3

- [cdn_cache_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2021332313100111-3332313313321312-2321131102001122-1001302322303112-2101020131320022-3331233102032232-3113232311101203-1012021331003203): complete subsection reference.

<a id="canonical-3220302212221300-1133202021231121-3202210332332322-1001321132021102-3230001022033212-1011221222032330-3033221102012222-1320303102223031"></a>

## Next pages — custom_cache_rule / 103223303033 / 4

- [custom_cache_rule.cdn_cache_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-2021332313100111-3332313313321312-2321131102001122-1001302322303112-2101020131320022-3331233102032232-3113232311101203-1012021331003203)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-2021332313100111-3332313313321312-2321131102001122-1001302322303112-2101020131320022-3331233102032232-3113232311101203-1012021331003203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232121310203103-0032101132302231-2311321210110112-3303211003312031-0003020003001302-0103111231003110-1121032120211201-3022122200102230"></a>

## custom_cache_rule.cdn_cache_rules — cdn_cache_rules / 310302130203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320)
- custom_cache_rule.cdn_cache_rules

<a id="canonical-2231330000201120-0333233000033323-2101303000130312-3113212111101000-1312221232020113-2320132320333223-0000203120001021-1202331011113232"></a>

Type: `"list"`. Computed.

Reference to CDN Cache Rule configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2222100011110000-0300131002111300-3101333232212302-3220111011131130-0310000022113022-1311200110033212-2333222112123202-2112002013032003"></a>

## Direct properties — cdn_cache_rules / 310302130203 / 3

<a id="canonical-3012031112002102-2321200210221101-3232200212321330-0201013001233301-0311202130313113-1110310032322100-2210212203132330-3112311111221120"></a>

<a id="canonical-3321010100233320-3011330110000110-3131011313302030-2031010201122031-3313031231313310-2231221301011222-0111102101311133-1202211332120013"></a>

## name property — cdn_cache_rules / 310302130203 / 4

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

<a id="canonical-1131200202230202-0100303303003331-3322121221320202-0132103330322030-1310120012013012-2221132121030201-1201202220031101-3100200131311102"></a>

<a id="canonical-0201221120101312-3030112232333221-1112101220230133-3321103311312331-3100330033103213-1000130120312100-1122122130021011-0102112032012313"></a>

## namespace property — cdn_cache_rules / 310302130203 / 5

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

<a id="canonical-3131111232303033-1322230232200330-3321212113311202-3233310100203201-3110230000003202-2300231202311023-3110323101302122-0310001313331332"></a>

<a id="canonical-0001121201031121-0101331002232213-2012103100212220-0220023013022123-1233020202333100-1000112013113221-0223323300032312-2102202032330001"></a>

## tenant property — cdn_cache_rules / 310302130203 / 6

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

<a id="canonical-1000001121301123-1112333103223000-0202103210221000-2230100221311001-2022302010221000-0101133323210322-3132331012210333-0210133311010212"></a>

## Next pages — cdn_cache_rules / 310302130203 / 7

- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000202221123133-2132021323122021-3331111211011200-1330102113233300-0100310113012231-2231110303303022-3020111020031302-1121012212001313"></a>

## data_guard_rules — data_guard_rules / 233013100303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- data_guard_rules

<a id="canonical-2210301100023211-2133121113131030-2131333130020032-3223120103202302-3120123321221020-1211110320330323-3333013100121102-2131333113033333"></a>

Type: `"list"`. Computed.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

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

<a id="canonical-3303233122321300-3221321301030303-3031301133231012-3312302322001212-2113111322320322-2010103221033313-1230321110310012-0012223003001120"></a>

## Direct properties — data_guard_rules / 233013100303 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232321121012031-2231320120122001-3030021103212311-2032210323100132-1111011312203313-2310300031301000-1331012210022312-2123330020301020): complete subsection reference.

- [apply_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231200112332010-3011103033131301-2210013233011300-3000321113103121-2331221221301202-3331323321321222-1210111231110233-0301113300110221): complete subsection reference.

<a id="canonical-2131002010112210-3312232002330103-2123120033320103-2100300320133203-0132111213202112-3103100000102010-3201203021110221-2032010110021212"></a>

<a id="canonical-1132230331203103-3223213001320323-1211121322101203-0232333211023300-1032233300112201-3303123221032313-3202222212213020-0223323113031031"></a>

## exact_value property — data_guard_rules / 233013100303 / 4

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

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

- [metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1231031032112033-0132301003311100-1203321332303302-1111112322000023-3020022203332112-2313322030310210-3300021100230202-0000221300200101): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0300121020333212-3003033232111130-3313010331132123-0012103220020020-2213032013310100-3021232333000320-2222331003032323-1233120322311302): complete subsection reference.

- [skip_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0032133120001001-3220200103000200-1332011021110003-3003013102321231-1203331201010031-3012130022101031-0211222113023112-2002221331313203): complete subsection reference.

<a id="canonical-3310102131110230-3103332021130221-0213130001133322-3001300333333320-0130001032020203-1323312102012000-1123113322301323-3211010221230111"></a>

<a id="canonical-2333311301011202-3320022030223021-2222011001010030-0231210323022023-0010202121131210-3003311231021233-1202321220213321-3221132323232330"></a>

## suffix_value property — data_guard_rules / 233013100303 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-2310123020222211-0232210331032312-0030310110010313-3320113131030002-3200111311111300-1333012002113011-1132212132312230-0131132112102111"></a>

## Next pages — data_guard_rules / 233013100303 / 6

- [data_guard_rules.any_domain](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232321121012031-2231320120122001-3030021103212311-2032210323100132-1111011312203313-2310300031301000-1331012210022312-2123330020301020)
- [data_guard_rules.apply_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231200112332010-3011103033131301-2210013233011300-3000321113103121-2331221221301202-3331323321321222-1210111231110233-0301113300110221)
- [data_guard_rules.metadata](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1231031032112033-0132301003311100-1203321332303302-1111112322000023-3020022203332112-2313322030310210-3300021100230202-0000221300200101)
- [data_guard_rules.path](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0300121020333212-3003033232111130-3313010331132123-0012103220020020-2213032013310100-3021232333000320-2222331003032323-1233120322311302)
- [data_guard_rules.skip_data_guard](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0032133120001001-3220200103000200-1332011021110003-3003013102321231-1203331201010031-3012130022101031-0211222113023112-2002221331313203)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-3232321121012031-2231320120122001-3030021103212311-2032210323100132-1111011312203313-2310300031301000-1331012210022312-2123330020301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130221232021010-0301231201103130-3123311212102213-0200020322331222-0203303332121000-3301301321001222-0212301213013211-3101100111001221"></a>

## data_guard_rules.any_domain — any_domain / 130222022213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.any_domain

<a id="canonical-1021201221301331-3101302301113301-0233210001220231-1010112322202003-1010122211132333-3332310212230222-2313332232032213-1000202301321132"></a>

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

<a id="canonical-2112030131003302-3131100003230101-0001112211322311-2000231300132112-1023003013032213-3233133002111302-0130023301100213-0220102332301030"></a>

## Direct properties — any_domain / 130222022213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032322231210222-1130211010302121-3030003113231010-0222320100210102-3303110111233121-3030121101300310-2220200130320321-2031312230011013"></a>

## Next pages — any_domain / 130222022213 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0231200112332010-3011103033131301-2210013233011300-3000321113103121-2331221221301202-3331323321321222-1210111231110233-0301113300110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223011112130220-3223110130303010-3322011121101300-1331013202213333-1211233030133223-3312010131302321-3313020223230303-2322200222321122"></a>

## data_guard_rules.apply_data_guard — apply_data_guard / 203100300321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.apply_data_guard

<a id="canonical-3322210121033322-0133102310331223-2003323123320313-1111001133031121-1130330231232013-3223000103022203-2022102200233233-1111001033023132"></a>

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

<a id="canonical-3300201332313013-2223223221322302-0200300322103132-3010233010330111-2321211102130133-2111110211032231-1000223100303110-2301022303310213"></a>

## Direct properties — apply_data_guard / 203100300321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330221020112323-3321113232030123-2100111011121032-3202001032020230-1232020213121303-1201133111020000-3110302123201020-1102032111330000"></a>

## Next pages — apply_data_guard / 203100300321 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1231031032112033-0132301003311100-1203321332303302-1111112322000023-3020022203332112-2313322030310210-3300021100230202-0000221300200101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212210020012012-2313312213203021-0011231112112202-2011223222200023-0011021030021033-0210330330231330-2222000012101020-2300331002211102"></a>

## data_guard_rules.metadata — metadata / 002012303123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.metadata

<a id="canonical-1201022100012001-1312113011023202-0123221031103230-0000312311123202-1030133023303012-3333311221210200-1332212211023220-2230331313032310"></a>

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

<a id="canonical-0013213300303301-3320212300233131-3200222001131110-1010322132013001-2230313201311013-3322232000210202-3130001130031113-3102332131203211"></a>

## Direct properties — metadata / 002012303123 / 3

<a id="canonical-0300133223200023-0003121030132033-2133101012030132-0311013231232121-1321032102202302-2030011230232013-0200213233230232-0011212002213203"></a>

<a id="canonical-0310230231021333-1303120333331032-3120002011312232-1030331002001200-2021230322021212-3023102100031013-1322020312002203-3013000332002122"></a>

## description_spec property — metadata / 002012303123 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3103302300130121-0213132133321330-2332311021031310-1121213122130012-0021022023031213-0003033222131102-3033310301111123-2230330102102230"></a>

<a id="canonical-2120121033021102-2112003130031321-1201331330130113-2001020312302200-0032100000232033-3311332333020001-0120202303002231-2212130301302323"></a>

## name property — metadata / 002012303123 / 5

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

<a id="canonical-0103323131031031-2111233303330002-1220122213300301-2133000002300312-3212100320310123-2201112123030331-3320121021121121-2202303123000212"></a>

## Next pages — metadata / 002012303123 / 6

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0300121020333212-3003033232111130-3313010331132123-0012103220020020-2213032013310100-3021232333000320-2222331003032323-1233120322311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301210023222202-0133120122113201-3133031132300023-0201311212022233-1001331233231201-0202232221003233-3313203223203211-0122031121123320"></a>

## data_guard_rules.path — path / 210213113000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.path

<a id="canonical-2000230321020020-0101102023022022-3222012333230120-0300030233002202-2033123211001103-0123202113032003-2013100130311210-2300221100311220"></a>

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

<a id="canonical-2311020301212303-3213333030200303-1322300002330203-1011032030112210-2033201322213312-1301000223223233-1022321033321201-1120101202101321"></a>

## Direct properties — path / 210213113000 / 3

<a id="canonical-3100330331202100-1013111222330012-1121011021102231-3302201310311330-1232002203130211-0003003221021212-2222122113322121-0213302023030120"></a>

<a id="canonical-1031110003032032-0122000122110113-2021200223001321-2130231201021332-2323313322211030-3100232312233123-0022333110210202-1332322310312031"></a>

## path property — path / 210213113000 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

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

<a id="canonical-3121133113013112-0032000302102033-3120330223011303-1130201302003333-0321021101301032-3003011313013032-1312221221202212-3233313231201120"></a>

<a id="canonical-1113311112000300-2221323323310322-1123211330302100-2222210011322203-1133302213112302-3230123231023113-2320210333131022-1102323020312300"></a>

## prefix property — path / 210213113000 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-0133002112202231-0201011030001121-3021321013301332-3030123233000020-2212233321323021-2001132030230330-1200130010301323-0320003321220101"></a>

<a id="canonical-1330300020103313-1001211222232320-3210010333001210-2311311030330231-1100032321012210-3130011010102122-1000232003131103-3112031003221022"></a>

## regex property — path / 210213113000 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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

<a id="canonical-3202320013222112-0020232301323120-0313232231301202-2220001330320130-1222201223101122-3313110021121110-3033111222012211-0301321032223323"></a>

## Next pages — path / 210213113000 / 7

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0032133120001001-3220200103000200-1332011021110003-3003013102321231-1203331201010031-3012130022101031-0211222113023112-2002221331313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232103033220320-2233100033213010-2032200131022330-1220331302021212-0223102032011300-2120130021001013-2332110332331012-1013101002330203"></a>

## data_guard_rules.skip_data_guard — skip_data_guard / 113221022002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- data_guard_rules.skip_data_guard

<a id="canonical-1000132113110023-0020112102211122-1332013010220002-1023222230321101-0131013331001331-0331022002333200-0120133020213031-0200302011220310"></a>

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

<a id="canonical-1121311331232101-1001330213202333-0112012030230122-0131302010323021-3010030122210221-2121200103131032-3123132122211130-0301300231332313"></a>

## Direct properties — skip_data_guard / 113221022002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132011300023032-0122001210122012-1013333021320220-3130123302021200-1022131002103303-0031022310201033-0032221003132002-0112301220220311"></a>

## Next pages — skip_data_guard / 113221022002 / 4

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322333101001033-0233301002200330-1321303001311123-1100213300221110-2312102011230001-0123230330033120-3132000221001123-3211201113313331"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 332311333332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- ddos_mitigation_rules

<a id="canonical-0103311221320210-0333301313222313-0112333011330112-3233332011303303-3230120233312321-3102020213020203-0120223103312123-0313133113233022"></a>

Type: `"list"`. Computed.

Define manual mitigation rules to block L7 DDoS attacks.

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

<a id="canonical-1313131113120003-2003032011022213-1212010113132002-2330322133101211-2202210331120323-1303111120031303-1301321202300311-1100021001302222"></a>

## Direct properties — ddos_mitigation_rules / 332311333332 / 3

- [block](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1110203213331010-0211033320130001-1123132032023132-0123310200112200-1100203002231013-0131310231302113-0113333013330021-3331001323230210): complete subsection reference.

- [ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311): complete subsection reference.

<a id="canonical-1110301233212010-2232323110231201-2310312321102321-0203102310330012-3012333002030011-1311300200010213-0121230233220332-1210131111123111"></a>

<a id="canonical-2101123121023302-1000322033101112-3223201100301302-0310222202003332-1323012102021022-2331130332022112-2202103120332113-1112132022331330"></a>

## expiration_timestamp property — ddos_mitigation_rules / 332311333332 / 4

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1023131012231111-2223333211011103-0230132032323301-2011103010300003-3300112112110213-0210102133231210-1003321200100123-3202201301123130): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2100220232201032-0220103110012110-0210030130100202-3031003111313030-0111111011112011-2100122010300210-2331230311310301-0313121133211311): complete subsection reference.

<a id="canonical-3222330002322002-2103310200121220-2230212232000003-0203231233023232-3101232002132001-0113233312133012-3101312222302120-3202113111122132"></a>

## Next pages — ddos_mitigation_rules / 332311333332 / 5

- [ddos_mitigation_rules.block](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1110203213331010-0211033320130001-1123132032023132-0123310200112200-1100203002231013-0131310231302113-0113333013330021-3331001323230210)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- [ddos_mitigation_rules.ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1023131012231111-2223333211011103-0230132032323301-2011103010300003-3300112112110213-0210102133231210-1003321200100123-3202201301123130)
- [ddos_mitigation_rules.metadata](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-2100220232201032-0220103110012110-0210030130100202-3031003111313030-0111111011112011-2100122010300210-2331230311310301-0313121133211311)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1110203213331010-0211033320130001-1123132032023132-0123310200112200-1100203002231013-0131310231302113-0113333013330021-3331001323230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010012332220202-2322300122332321-0301010121333232-3220223010003031-1213030311102211-1031213123132230-2111002103020021-0013202013102313"></a>

## ddos_mitigation_rules.block — block / 030112312002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.block

<a id="canonical-1110003002100100-0313313113131312-3031332303001312-0000312323232110-1210110313222101-3110322201303322-2211123033020032-0232110101210123"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3120111031311333-0312203023200221-0000122213023101-0112331222310312-2123333201310221-3222310001010003-0210131310023303-3102101030300002"></a>

## Direct properties — block / 030112312002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301332020332202-3231200300020320-2011102131113221-2001021130132311-1003033221122133-1030201112200223-2212020220322302-2021231101312010"></a>

## Next pages — block / 030112312002 / 4

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113013010121222-2111033121302101-1313322303332313-3312113013112201-2321300330003103-1113222121202231-3032131101330310-2030220132000003"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_client_source / 132311201020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-3013102123312033-1232230212221311-3302312213302203-1321002323333212-1120022223300212-1113213321101230-0003100223133120-3112331102213231"></a>

Type: `"single"`. Computed.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233123110110121-1212303001331301-3032311200112230-3210313030102332-0232210122010301-1311030020130200-1212000311232200-2022303022322301"></a>

## Direct properties — ddos_client_source / 132311201020 / 3

- [asn_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0312133021030231-1010312022321220-1101133201222222-0230303000211031-2201120030223202-2303310220132322-3231321133011103-2200221210213003): complete subsection reference.

<a id="canonical-1301002312312020-0002111011231001-1133230230302022-2020323132103032-1303323201211203-0310132001330321-3201103220122223-1110221001123320"></a>

<a id="canonical-2301100012323210-3303203213102311-1023322130113302-3232323023302300-0122231121113203-2332000011233021-1331123212101102-0322312101221311"></a>

## country_list property — ddos_client_source / 132311201020 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1000233103010133-3310202120331122-1213031103300101-0302012100131232-2000021132101101-1233032122320200-3013202323001112-3301220130130200): complete subsection reference.

<a id="canonical-1212332320013330-0133021201303133-1011111333011321-0312330121023113-2202233321313210-2222000301021100-2220313222101130-2012332323000303"></a>

## Next pages — ddos_client_source / 132311201020 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0312133021030231-1010312022321220-1101133201222222-0230303000211031-2201120030223202-2303310220132322-3231321133011103-2200221210213003)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1000233103010133-3310202120331122-1213031103300101-0302012100131232-2000021132101101-1233032122320200-3013202323001112-3301220130130200)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-0312133021030231-1010312022321220-1101133201222222-0230303000211031-2201120030223202-2303310220132322-3231321133011103-2200221210213003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011033213201211-0220232111301000-1011231012032030-3211011333223023-2232313211313002-3131313011123112-1303230003033223-1202300003200002"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — asn_list / 131322131001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-1121003331302123-1020111102313211-2001102323133001-1211313011232300-2322223032201200-3122011103210033-2313301020022022-3012230223120310"></a>

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

<a id="canonical-2103120311202022-3222030012300212-2110332332030222-3302211200233313-1201221132132111-0023321211303201-1201002133213030-3022023233131013"></a>

## Direct properties — asn_list / 131322131001 / 3

<a id="canonical-3221123111112110-1020023010022020-1110111023121311-2211101323003120-1223102333331310-1133101310333222-3330310003031020-2332133311132101"></a>

<a id="canonical-0013320221230120-3213130233221121-1200133311101030-1123331322111032-2003210210202300-0032223322123101-1023030101010111-2213132023313320"></a>

## as_numbers property — asn_list / 131322131001 / 4

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

<a id="canonical-0110113002030033-3120001101201231-2221012132300321-1011011032213313-1111332011202002-0031210120122103-2223113333303313-0323212232021313"></a>

## Next pages — asn_list / 131322131001 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1010020312100303-2110330230121312-2302333202321113-3003032012103131-2011020302013313-2321131022313002-0021311001101200-3122032201311210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223232112232222-1002320132031032-0000112223022220-2022223133203232-1310211021031021-2130030222320120-1133123201310232-3030320331003332"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ja4_tls_fingerprint_matcher / 000322232231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-3011120122000003-2022332233203321-3011000101032003-1211111002213102-0210220201123302-1201113121210232-1230012031103020-0100231323322020"></a>

Type: `"single"`. Computed.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301203223031020-3322132010331303-1120320223223311-0022023022000132-2023110322212012-2321023013000220-2120221203210331-3230123313023013"></a>

## Direct properties — ja4_tls_fingerprint_matcher / 000322232231 / 3

<a id="canonical-3031300201103321-3121300332103222-2233112200133133-1010002100202101-2333012221321210-1202303321021331-3321021031010213-0210010211232103"></a>

<a id="canonical-1320332333001223-1203222000133102-3210113033210223-1100221113102331-1220133330100120-2310310121232232-3000220133022122-3332100201033233"></a>

## exact_values property — ja4_tls_fingerprint_matcher / 000322232231 / 4

Type: `["list", "string"]`. Computed.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131122011232301-0131200133301230-0120010330212123-3102120021322030-3203200113213012-3020133331022320-1110203011020003-3222030033200032"></a>

## Next pages — ja4_tls_fingerprint_matcher / 000322232231 / 5

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1000233103010133-3310202120331122-1213031103300101-0302012100131232-2000021132101101-1233032122320200-3013202323001112-3301220130130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312202332332332-3033123232123311-2202123033311103-1201122122113211-2330132233322321-0203223310213123-0130323222131312-3222113002200322"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — tls_fingerprint_matcher / 133233303233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-2133300310112230-1323212032221012-2102301023300021-1103031322233300-2220233210013121-2021301331130120-1132311113323303-0100132021231122"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213101223002301-3032012022200203-3311321121233301-0023310003111313-1032103003201000-1122031121323103-3310013032022210-2103320310323211"></a>

## Direct properties — tls_fingerprint_matcher / 133233303233 / 3

<a id="canonical-3320102033230331-0131010112330112-3033003112112002-1032323333120121-3201020132213111-3122211033221230-2103312313223120-0102223011212032"></a>

<a id="canonical-1302231100022210-1031032301211112-0032103221021012-3000332200211221-3301122022103212-2011010013030000-3003332312323200-1233302312121302"></a>

## classes property — tls_fingerprint_matcher / 133233303233 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0230231011233032-2131013230013233-1320231103312230-0000301230213330-1131023202011202-3213302233320121-1212303222000132-2122201030022133"></a>

<a id="canonical-2231032322032330-1122001202322022-2330320232003100-3220012210122123-0000300202311033-2032021031130232-2123000111311320-3313001310232311"></a>

## exact_values property — tls_fingerprint_matcher / 133233303233 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2130131230121120-0220323030131002-2213100323021011-0312030210012203-0212212132031211-3202012211221000-0211002330331332-0323130113111330"></a>

<a id="canonical-3333201300223232-2133112001022312-1032320020222100-3013013132233313-0331303311121312-2030020000001103-2210301111103222-1222332101002222"></a>

## excluded_values property — tls_fingerprint_matcher / 133233303233 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3310121301312110-0313031333222010-2010301200001132-1121201333102013-2123003212310332-0312000212322011-1012130113301231-3230232232030210"></a>

## Next pages — tls_fingerprint_matcher / 133233303233 / 7

- [ddos_mitigation_rules.ddos_client_source](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0231322231110013-0203001033311331-1033103211213020-3310322033003023-2300013323303032-1031133323221002-2132310332322021-1201201201201311)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)

<a id="canonical-1023131012231111-2223333211011103-0230132032323301-2011103010300003-3300112112110213-0210102133231210-1003321200100123-3202201301123130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222201110303313-3232100002001222-1320301232031202-0333133020311001-3110010303202131-1223302300202113-2110301131133232-3222010203033001"></a>

## ddos_mitigation_rules.ip_prefix_list — ip_prefix_list / 003202112023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-1331011201023012-1023123023203233-2232112011211300-0110110333111313-3001102001231221-1223213133121002-2210331200133030-2212332101212131"></a>

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

<a id="canonical-1311023333203122-0300220001302131-2302022313302111-1211100313310110-2222132130133110-0012222132111103-0001111322332332-0223310211121122"></a>

## Direct properties — ip_prefix_list / 003202112023 / 3

<a id="canonical-2203300302230020-3010021113022133-2023012001131132-2121311302113111-0213031312300211-0122110332010100-1321311303121001-0303120122020102"></a>

<a id="canonical-1332122102212321-3031322030223223-2013001021203130-1030200131031023-0212110131123013-3220013302013201-3310123313031003-3102133300111212"></a>

## invert_match property — ip_prefix_list / 003202112023 / 4

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

<a id="canonical-3232130232110222-0111311100302211-1323301032200013-1003331012021032-0223220020022101-1320111311210211-2323002023133110-0120222133310001"></a>

<a id="canonical-3131233022302113-1313202332230032-1132211110012131-0011212333011021-2000120333102130-2310202010021313-2020123102030310-3213001303312030"></a>

## ip_prefixes property — ip_prefix_list / 003202112023 / 5

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
