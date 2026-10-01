---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013333020203100-3012221010332013-0121133330233310-2233330202022312-3311200211210100-1122032300120121-2002320211032113-0023231010223121"></a>

## Property reference — Property reference / 300210112021 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- Property reference

<a id="canonical-3131231212102312-2213123013303232-0022020002132233-1221213020210201-2020130123322300-1200031130230211-1003130000033231-2310001203131003"></a>

## Direct properties — Property reference / 300210112021 / 3

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-2131122002322200-1211102013312210-0332313001322211-3203310002112222-1102230310021123-1302303130013203-0210310220203232-1201203310213203): complete subsection reference.

<a id="canonical-0233211010122220-2110201101200013-0131110133102321-0122332013100302-3230032332213031-3010330300130221-1132120313312112-0331303303321312"></a>

<a id="canonical-1333322203121110-2023100213200221-1133232322320133-0023121012102003-2210100311003210-1000232331121133-2222201111303133-1003313111223333"></a>

## annotations property — Property reference / 300210112021 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-3121210303101201-1231303202003123-1231320121022310-3001201012333211-2010131131103323-0230330313113000-3233023131132211-3332203002023000): complete subsection reference.

- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331): complete subsection reference.

- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310): complete subsection reference.

- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-2131022120013230-3323322131002320-3112130302320203-0132230022331312-3202130201002000-0310132300310330-1103022101033312-1203333103123122): complete subsection reference.

<a id="canonical-3312333012213211-0323230320123103-0200121311213311-3200112123103321-2201223203322313-0231111121203232-2313210313102233-2111301333220333"></a>

<a id="canonical-3020121012330230-3113310022130012-1311301021112113-2011013132302123-1112302113032313-1330222211102102-1003210130011203-1031120333301000"></a>

## description property — Property reference / 300210112021 / 5

Type: `"string"`. Computed.

Description of the ProtectedApplication.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-2221321233201022-2312300032001020-0333110031033331-0133213100320220-0313313301321000-2131300212321000-1211302300200001-2300130230311301): complete subsection reference.

<a id="canonical-2111001110230031-2101333220323113-3211002122022103-3132331302221002-1211231113133022-3331202333110120-3220321233330201-3211331232110233"></a>

<a id="canonical-1031202113200301-1031323330113100-2213122101032011-2213020131023011-1101002213313232-1131221301201200-1220203213123210-0103011123330230"></a>

## ID property — Property reference / 300210112021 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2132230300201311-3111310330002222-3120230100322021-1022303202131103-0123301030112002-2323220030013321-1123233103333313-0003121203110330"></a>

<a id="canonical-3111212130330101-3031110331313023-0113123301103103-2222333200223210-1223132203120023-2122102330102030-1020102111022022-1123310112023021"></a>

## labels property — Property reference / 300210112021 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0023123123230212-1103211322111123-0010012032112213-2323002311113300-2120302132200122-3133233233313222-1022002111200102-1001233211312221"></a>

<a id="canonical-2033033100120320-0231003200213120-3333221023002333-3030011213202303-0003113211111000-0121210033130111-1123322011310322-1303201021012033"></a>

## name property — Property reference / 300210112021 / 8

Type: `"string"`. Required.

Name of the ProtectedApplication.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

<a id="canonical-1310103102101223-0230021323212321-3133212230121222-1323131012002322-1230013200010023-1231231003032033-0123021003101203-3100001210023333"></a>

<a id="canonical-1232112012231322-2212332323120201-3033100132112012-2101113333333121-0121012311010121-3333313110011303-2023030311201210-1102231132233321"></a>

## namespace property — Property reference / 300210112021 / 9

Type: `"string"`. Required.

Namespace where the ProtectedApplication exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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

<a id="canonical-0010102001012002-1323203111213100-2032220022322220-3320220112200213-3132313211121332-0031001312231200-1031032001313322-3333310111130122"></a>

<a id="canonical-2201120133322013-3323101311231213-2313120233211101-1013030030321003-0330122110212232-0031321211223310-1130022201100230-2312201202332313"></a>

## region property — Property reference / 300210112021 / 10

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- US: US

United States of America &#8203;- EU: EU

European Union &#8203;- ASIA: ASIA

Asia &#8203;- CA: CA

Canada.

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA",
    "CA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-3300103021210020-1101231333023233-1202122312030020-1322030202123101-2122211031213030-0203233223330223-3331032100031322-3101033002131222): complete subsection reference.

<a id="canonical-1023021112003110-3012121000121333-2320213002312231-2123130301220312-0332103120201303-3021310233131330-2131131322101013-2111331312101101"></a>

## All schema paths — Property reference / 300210112021 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `adobe_commerce_connector` | [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-0011311010322211-0231123221220032-2221230102020200-1320103102222213-3102313202333231-2102220032210002-1112232123103010-3010331333301311) |
| `annotations` | [annotations](data-sources--protected_application--reference--group-001.md#canonical-0233211010122220-2110201101200013-0131110133102321-0122332013100302-3230032332213031-3010330300130221-1132120313312112-0331303303321312) |
| `big_ip_iapp` | [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-2310010302100121-2100201222033113-0121010133333222-3331301002332312-2003311113020202-2201231133202112-3300310301033320-0102232230321033) |
| `cloudflare` | [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-3023111033212211-1111122330202301-0012113103330120-3112300011213122-1111332310321102-2101102231002202-3320333023323111-0023110123110001) |
| `cloudflare.continue_mitigation_action_hdr` | [cloudflare.continue_mitigation_action_hdr](data-sources--protected_application--reference--group-001.md#canonical-2202200332333113-0321020131310023-2022311121012120-0012102313103032-3210000001332123-3331303330102300-1020222230221302-3113201132313323) |
| `cloudflare.disable_js_insert` | [cloudflare.disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-0123210122333123-1103213201222301-2032023233121210-2020032023011333-0023010021202111-1020213120212110-3131323313113032-3100031113211000) |
| `cloudflare.disable_mobile_sdk` | [cloudflare.disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-0120230020032110-1233130200333211-1223221003313221-0311031310222013-1100121003332310-1022000232021103-2010331030323100-2133302322132300) |
| `cloudflare.js_insertion_rules` | [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-3033130212321203-1220212110233133-3213232032233210-0002012100323023-0330312131223120-3101213120202031-3101313333232311-1033000221111203) |
| `cloudflare.js_insertion_rules.exclude_list` | [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2101020020021303-0213230231003300-3110030130303203-2222211322302222-1102311213023321-3203303022110111-1311331201130023-3030222231021203) |
| `cloudflare.js_insertion_rules.exclude_list.any_domain` | [cloudflare.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-001.md#canonical-2111031200022311-3011302133102211-1103311222311203-2213133002212023-0221031312223013-0023101033203100-2022112311121022-0201220223300000) |
| `cloudflare.js_insertion_rules.exclude_list.domain` | [cloudflare.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-001.md#canonical-2123011120102211-2200020120333020-0121310030011112-2212113000103311-0323012311202222-2231323200130000-2221021132201203-2001001103302213) |
| `cloudflare.js_insertion_rules.exclude_list.domain.exact_value` | [cloudflare.js_insertion_rules.exclude_list.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-2030111102032020-2022312100012201-0222120121330300-2203022031122303-2300030200320210-3120100112021211-0033110112020031-1233210112123021) |
| `cloudflare.js_insertion_rules.exclude_list.domain.regex_value` | [cloudflare.js_insertion_rules.exclude_list.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-2333203101223103-2220000202122133-0303103102111022-3232032232332033-1333222010221230-0030111110020102-0300023210002232-3003323000220121) |
| `cloudflare.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudflare.js_insertion_rules.exclude_list.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-1310223000000223-3330001200311103-2201232222113322-2310102010222233-1301231303120110-0031232131301132-0332202231232322-2022333333133100) |
| `cloudflare.js_insertion_rules.exclude_list.metadata` | [cloudflare.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-001.md#canonical-1003023133202221-0220011220231210-1202333023311312-1131321102311030-1222012101232001-1320302213022012-1001013200003301-3223231003113301) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudflare.js_insertion_rules.exclude_list.metadata.description_spec](data-sources--protected_application--reference--group-001.md#canonical-1311002112210132-1110231111100031-2102100001033010-0031132032210131-3101310332002033-0123220300033301-3312023033102020-1033233202120322) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.name` | [cloudflare.js_insertion_rules.exclude_list.metadata.name](data-sources--protected_application--reference--group-001.md#canonical-3213110102122312-1201012132221011-0131110102232300-1203323321001033-1302103213233021-0302310301310311-3321213011133000-3212102000130220) |
| `cloudflare.js_insertion_rules.exclude_list.path` | [cloudflare.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-001.md#canonical-1103223013202110-0202202100021130-2031002123120321-1011211030033312-2321023232322132-2112033233130210-3100303211223102-1012033111210100) |
| `cloudflare.js_insertion_rules.exclude_list.path.path` | [cloudflare.js_insertion_rules.exclude_list.path.path](data-sources--protected_application--reference--group-001.md#canonical-0100220122023033-0010223103130022-1232121011101121-0310033231333213-2003103110233112-0013200100203222-2303203312000032-0002231213203130) |
| `cloudflare.js_insertion_rules.exclude_list.path.prefix` | [cloudflare.js_insertion_rules.exclude_list.path.prefix](data-sources--protected_application--reference--group-001.md#canonical-0313102301032213-0120132111300300-1231100131033331-0002311212001111-2312322120013020-0211303203001222-1312220213300233-0100121223110021) |
| `cloudflare.js_insertion_rules.exclude_list.path.regex` | [cloudflare.js_insertion_rules.exclude_list.path.regex](data-sources--protected_application--reference--group-001.md#canonical-1300311323310003-0010113101101132-2303001312220100-3322000323231002-2223232231123001-0030032033301322-0023100302031303-2002331101232322) |
| `cloudflare.js_insertion_rules.javascript_location` | [cloudflare.js_insertion_rules.javascript_location](data-sources--protected_application--reference--group-001.md#canonical-2302323021011231-2012230020020330-0201021233011121-2031003123321222-3232000033131230-3211201131330222-1230102002220012-0302312203113103) |
| `cloudflare.js_insertion_rules.js_download_path` | [cloudflare.js_insertion_rules.js_download_path](data-sources--protected_application--reference--group-001.md#canonical-1010320122311211-3113322333100310-2321023213022310-2323120232320000-3200032020312012-2103321300322303-1313321123101320-2212200033332223) |
| `cloudflare.js_insertion_rules.rules` | [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-0301100233133322-0323212333103212-3023210333331133-1132310322210320-3022123232111223-0333301320000032-2333231033011132-2013232220020013) |
| `cloudflare.js_insertion_rules.rules.any_domain` | [cloudflare.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-001.md#canonical-2033023231011110-1332213221203103-3313330330113202-3200010313201210-3022133300121100-2210232302323233-2232101013201031-2013131030213320) |
| `cloudflare.js_insertion_rules.rules.domain` | [cloudflare.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-001.md#canonical-0112311321131311-0201101101131033-0331302222031032-2020303230320011-3330130012033310-3030000030231001-1333233330200202-2223203032101321) |
| `cloudflare.js_insertion_rules.rules.domain.exact_value` | [cloudflare.js_insertion_rules.rules.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-2020221223131012-2131033123232102-3230223232033302-3310321013100203-2302202221312310-0201333022110210-3332000012110130-3120201010300301) |
| `cloudflare.js_insertion_rules.rules.domain.regex_value` | [cloudflare.js_insertion_rules.rules.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-2011002330003321-3031033001111333-2100301023331330-2322032121301120-0200013200112131-0313010230223102-0030032321202002-0220201120122100) |
| `cloudflare.js_insertion_rules.rules.domain.suffix_value` | [cloudflare.js_insertion_rules.rules.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-2313212302023333-3322000231031333-3111003322131103-2332131131033110-0131120133213031-1210211102310200-2032002123013103-0332033102012001) |
| `cloudflare.js_insertion_rules.rules.exact_path` | [cloudflare.js_insertion_rules.rules.exact_path](data-sources--protected_application--reference--group-001.md#canonical-2223310001121233-2012221213002202-0302000021201111-2203023132232010-0020012331131033-3223211310222023-3330131320213022-2130121123332231) |
| `cloudflare.js_insertion_rules.rules.glob` | [cloudflare.js_insertion_rules.rules.glob](data-sources--protected_application--reference--group-001.md#canonical-1013211320002131-3022210331330311-2210201011203032-0121021102331011-2021230332322103-0001021013230111-1212313310310232-0023111013301311) |
| `cloudflare.js_insertion_rules.rules.metadata` | [cloudflare.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-001.md#canonical-1231322113013122-1221103321130330-1003223213221302-3002103232310111-3111232331311123-2120210121021210-3131212120100300-0201301320033120) |
| `cloudflare.js_insertion_rules.rules.metadata.description_spec` | [cloudflare.js_insertion_rules.rules.metadata.description_spec](data-sources--protected_application--reference--group-001.md#canonical-1232201330233211-0122102010311001-0111122021231333-1101133020302002-1303023230321223-1300030333210003-3120130023221113-3000200101131322) |
| `cloudflare.js_insertion_rules.rules.metadata.name` | [cloudflare.js_insertion_rules.rules.metadata.name](data-sources--protected_application--reference--group-001.md#canonical-1230322312130000-3122000310211211-3312022332111110-0230031221330303-2321121221111031-2223121133303021-0132133313100103-3232121011131231) |
| `cloudflare.js_insertion_rules.rules.prefix` | [cloudflare.js_insertion_rules.rules.prefix](data-sources--protected_application--reference--group-001.md#canonical-3123010210031302-3313331033313013-2122222112100033-2320213211122331-2001011003033020-2100102120133322-2232010232111002-2101202121013102) |
| `cloudflare.loglevel` | [cloudflare.loglevel](data-sources--protected_application--reference--group-001.md#canonical-2232322200300300-2012300233103122-1133113222312110-2132000310233200-3233220032223202-3322100303013000-3130332312330322-1002033213311112) |
| `cloudflare.manual_js_insert` | [cloudflare.manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-3010110321202310-3102122010021121-1212203013132010-2011313220101131-2102122332220000-2330201223031323-2112031322331300-3133301313321010) |
| `cloudflare.manual_js_insert.js_download_path` | [cloudflare.manual_js_insert.js_download_path](data-sources--protected_application--reference--group-001.md#canonical-3233311301103321-1310323121001010-2011230033213021-3010101003311232-1320122212032103-0212201332103000-1321002302222020-2223130123311331) |
| `cloudflare.mobile_sdk_config` | [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1312323003331123-2221012232232111-1020102330032132-0011220123020302-0030122103020233-2310331102131012-3303123321301231-3332300103022002) |
| `cloudflare.mobile_sdk_config.mobile_identifier` | [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-0030003221330113-0312012032323011-1003200130310031-2233331123331232-1023121232312320-1231022101310303-2110303113131203-1322221313032212) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers` | [cloudflare.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-001.md#canonical-1203330130313101-3012123102323230-3133223322013232-2021310000333211-0203200233202033-0121312111332313-2012101200020001-0322203010013033) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.exact](data-sources--protected_application--reference--group-001.md#canonical-1000032003130210-3013031231031133-3131113020220301-1030103222011102-2212310300023130-0012013231211002-1210332310200012-1022103332320312) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.name` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.name](data-sources--protected_application--reference--group-001.md#canonical-0220002132031232-0022233220202113-2201301310010020-1312311233022323-2032000222322101-2312020300103320-3202013233032211-1310023032100232) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.regex](data-sources--protected_application--reference--group-001.md#canonical-0201113000303101-2311121022222213-0311120101300232-2032110212213212-3032301022032312-2120212132000333-3111030302103321-1212232031122303) |
| `cloudflare.protected_endpoints` | [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-2300121303010201-0332332210221112-2333212312230100-3011131213021130-2010333220313020-1111221332112102-1021331120331300-0123033012021000) |
| `cloudflare.protected_endpoints.any_domain` | [cloudflare.protected_endpoints.any_domain](data-sources--protected_application--reference--group-001.md#canonical-3223223133330120-3332331232213112-1101313310221321-1301313100233103-1203332212123212-3033200232101331-0203303332003331-1133233203201003) |
| `cloudflare.protected_endpoints.domain` | [cloudflare.protected_endpoints.domain](data-sources--protected_application--reference--group-001.md#canonical-0223212010233220-0012100021130123-1213312331203303-1022332002213030-0103310322022021-3030102113333200-3213133303210031-3232231320000330) |
| `cloudflare.protected_endpoints.domain.exact_value` | [cloudflare.protected_endpoints.domain.exact_value](data-sources--protected_application--reference--group-001.md#canonical-2011011323331210-0032113323110030-1110332222020011-0330032201211310-2111111033103001-1232232303112120-3001200022311230-1200122110212132) |
| `cloudflare.protected_endpoints.domain.regex_value` | [cloudflare.protected_endpoints.domain.regex_value](data-sources--protected_application--reference--group-001.md#canonical-2112333002330001-0233013001223330-3013331113012032-2122322013312031-0321131021311032-0203123321233230-3310330231231202-3222300010312221) |
| `cloudflare.protected_endpoints.domain.suffix_value` | [cloudflare.protected_endpoints.domain.suffix_value](data-sources--protected_application--reference--group-001.md#canonical-0033133302321322-3133320021003021-3012132123202123-0003100120300010-1110021121121323-0203033103013113-0232001221113231-0210103132120000) |
| `cloudflare.protected_endpoints.http_methods` | [cloudflare.protected_endpoints.http_methods](data-sources--protected_application--reference--group-001.md#canonical-1300001221100003-3322211312030311-2220311233330010-1200011012013221-0113230030113232-1222312033313101-3301232011033223-0011123313022120) |
| `cloudflare.protected_endpoints.metadata` | [cloudflare.protected_endpoints.metadata](data-sources--protected_application--reference--group-002.md#canonical-3221022120313012-1303120330313233-0020321102321130-0111022310321011-2323201323110120-1012031131222321-3200101033123013-2031130333101212) |
| `cloudflare.protected_endpoints.metadata.description_spec` | [cloudflare.protected_endpoints.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-0232220100232332-2332111110331120-2101103203100001-3013213121133323-2000221001130302-2031231222203113-1030332010001200-2023231311000330) |
| `cloudflare.protected_endpoints.metadata.name` | [cloudflare.protected_endpoints.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-1322213010200031-0123000022122010-0321222011132203-3232100110301301-0323300111313021-2200100323133133-0113011023021310-1110013232013221) |
| `cloudflare.protected_endpoints.mobile_client` | [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-1010220333030321-3011112023023130-2000012003112221-3201230323313001-1023022010333321-1223132233320133-1302031320211301-0001321020130330) |
| `cloudflare.protected_endpoints.mobile_client.block` | [cloudflare.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-002.md#canonical-0011030321232033-3112330300003222-2203313000102322-0002121322210020-2212321201102030-3211133011132213-3300013021011003-2333212102321001) |
| `cloudflare.protected_endpoints.mobile_client.block.body` | [cloudflare.protected_endpoints.mobile_client.block.body](data-sources--protected_application--reference--group-002.md#canonical-3002212002121322-0331311301302111-0012010022113031-2230210221011222-3003200120222332-1322010330021320-3332030300032112-0303010102310223) |
| `cloudflare.protected_endpoints.mobile_client.block.content_type` | [cloudflare.protected_endpoints.mobile_client.block.content_type](data-sources--protected_application--reference--group-002.md#canonical-1132100231013330-3023221030031000-2100323111010232-1011033311333212-3323321023312112-2220201332223303-3232002233220031-3000023120031022) |
| `cloudflare.protected_endpoints.mobile_client.block.status` | [cloudflare.protected_endpoints.mobile_client.block.status](data-sources--protected_application--reference--group-002.md#canonical-0330033333010320-3100002011320022-2130100230032300-1222332212012123-0033010303330201-3211233131301203-1000100310233232-1231111111020322) |
| `cloudflare.protected_endpoints.mobile_client.continue` | [cloudflare.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-002.md#canonical-1023223321310210-3331200330112313-2012330003311121-3331213123320200-1133031020111031-3003121231301002-1313100030131010-0232210303022310) |
| `cloudflare.protected_endpoints.mobile_client.continue.add_header` | [cloudflare.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-2233131220133101-3320211231301300-0012002002200121-1321030012023211-1311203032023033-3020032301233210-3302020022232133-3101023103230002) |
| `cloudflare.protected_endpoints.mobile_client.continue.no_header` | [cloudflare.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-3123322100331231-2200302121210320-0222112133223112-3011323032221213-2011211113231233-3021111323231201-2120221131321112-3330002100121031) |
| `cloudflare.protected_endpoints.path` | [cloudflare.protected_endpoints.path](data-sources--protected_application--reference--group-002.md#canonical-1000132322033112-3231322312102012-1222330212022121-0300021310033210-3013021100123112-1303103332231231-2133103032313212-0030112230320003) |
| `cloudflare.protected_endpoints.path.caseinsensitive` | [cloudflare.protected_endpoints.path.caseinsensitive](data-sources--protected_application--reference--group-002.md#canonical-0312232213302331-2320312320032131-1120002000222233-0313233322213231-1031223322233202-3030333132113130-1200231023301011-3120122300231301) |
| `cloudflare.protected_endpoints.path.path` | [cloudflare.protected_endpoints.path.path](data-sources--protected_application--reference--group-002.md#canonical-1330213122201030-0033211301103231-1310031230220323-1021210322203233-0123013212023301-2011320003110202-1103000031222100-2321330210213331) |
| `cloudflare.protected_endpoints.query` | [cloudflare.protected_endpoints.query](data-sources--protected_application--reference--group-001.md#canonical-0001132201102300-1200112033001221-3000020233021113-0332220233202032-2000330103330031-0212020121031130-2010323102322323-3203012301101131) |
| `cloudflare.protected_endpoints.web_client` | [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-2123131313302132-3212102211121033-1132332332002121-0020011313131111-0212321310321313-1331132313301321-3123223312012210-2332101133222332) |
| `cloudflare.protected_endpoints.web_client.block` | [cloudflare.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-002.md#canonical-0110321102221221-3003230120031200-0222302001211022-3333202323222100-3222132032133102-0321021203223010-3022220331210000-2131230301320203) |
| `cloudflare.protected_endpoints.web_client.block.body` | [cloudflare.protected_endpoints.web_client.block.body](data-sources--protected_application--reference--group-002.md#canonical-1111001213002301-3311013200120122-3213203220031012-1132032333122110-3122300211101211-3132202212330301-3003033020210300-1211313211232010) |
| `cloudflare.protected_endpoints.web_client.block.content_type` | [cloudflare.protected_endpoints.web_client.block.content_type](data-sources--protected_application--reference--group-002.md#canonical-1013000030303220-3332220321011232-1110332012321313-2210203332231303-3003331022120111-0210101221301220-1232122000203220-2020111303011202) |
| `cloudflare.protected_endpoints.web_client.block.status` | [cloudflare.protected_endpoints.web_client.block.status](data-sources--protected_application--reference--group-002.md#canonical-2323330210003111-0011030331120113-1011320311233122-2201123121222031-0131130120323303-0213131203122202-0132201300131112-0131232133220310) |
| `cloudflare.protected_endpoints.web_client.continue` | [cloudflare.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-002.md#canonical-3033223111030212-1323010303013003-3200331221323221-1101033300312210-3302233103010013-2232211211310000-2001200132311213-1320131123112022) |
| `cloudflare.protected_endpoints.web_client.continue.add_header` | [cloudflare.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-002.md#canonical-2302222211301220-3031310102131031-1103310221031332-1130212312222213-0333223112031333-0023310202113310-0223233013112121-0320132230303001) |
| `cloudflare.protected_endpoints.web_client.continue.no_header` | [cloudflare.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-002.md#canonical-3222300131021023-1031102203233311-3100323311330011-0330113023203022-2323130203313302-0033233220130131-2002201000313131-3222223120023013) |
| `cloudflare.protected_endpoints.web_client.redirect` | [cloudflare.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-002.md#canonical-3213203111103102-1320033023102031-3130012122012100-1301232221020200-2320012310200010-0213123101313001-3120013103121110-0122031033012130) |
| `cloudflare.protected_endpoints.web_client.redirect.location` | [cloudflare.protected_endpoints.web_client.redirect.location](data-sources--protected_application--reference--group-002.md#canonical-3003312021332122-3331321302030003-3302303013222020-1011310110030003-2010023212102222-2313202123332022-2130023322111230-1131222332222120) |
| `cloudflare.protected_endpoints.web_client.redirect.status` | [cloudflare.protected_endpoints.web_client.redirect.status](data-sources--protected_application--reference--group-002.md#canonical-0331233013321231-2230122321133001-0221100133213321-3331232020031102-2331322023100113-0031000133013333-1110003102001203-2030320213332032) |
| `cloudflare.protected_endpoints.web_mobile_client` | [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2112112301200233-2231312101110100-3002121230330310-3230312132111031-0333120023323121-2332301222310030-0212132023212001-2221112130100311) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-002.md#canonical-2001020110002110-3201232002303102-2011333110113121-1123301032220230-1330100230120122-3333001001110132-1120023322210212-2301020011023121) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.body](data-sources--protected_application--reference--group-002.md#canonical-3322102211021001-3011132022130031-1320020313121333-2130020131121021-2030323310213112-0113010010233203-0022123103012212-3013100032123311) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type](data-sources--protected_application--reference--group-002.md#canonical-3003131003003022-0100012310020321-2123302232130010-1110321132131011-3022123213331321-2303301312003021-1212010122322110-0131210302220310) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.status](data-sources--protected_application--reference--group-002.md#canonical-1312131222031321-3310001230202232-3012101132320322-1332002330033212-0023200123002102-1201333331333303-3121300030331031-3021013300210301) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web` | [cloudflare.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-002.md#canonical-2113232300131302-2003103232210213-3013023012201131-1313102222300031-0201303011031022-3223232033332021-2221022110001312-0332231113300023) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.body` | [cloudflare.protected_endpoints.web_mobile_client.block_web.body](data-sources--protected_application--reference--group-002.md#canonical-2100100030101302-3220201313312030-1233110231033130-1113032010320000-1110211310131123-3112132003310132-3231010232203000-3231213223323300) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_web.content_type](data-sources--protected_application--reference--group-002.md#canonical-0023123233002320-3033113200013000-3330312310120330-1133230121010021-1231110310010011-0213200332201310-2013210202020101-0110303103221103) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.status` | [cloudflare.protected_endpoints.web_mobile_client.block_web.status](data-sources--protected_application--reference--group-002.md#canonical-1223111310332000-0222331210231123-1301010200232330-0212030300020320-3133323122302200-0120331211310223-2131130102122023-0102120212331103) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-002.md#canonical-0221210010023111-2231022330022130-1101303212302322-3033231332230123-2100210110010111-2001100023011022-3121113032120002-3230200001223103) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-002.md#canonical-0331302300003201-2010200122202231-0110030123230012-3332310023113300-3313020212223222-0022233113302023-1230331122310022-3312323333010112) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-002.md#canonical-0320302001212000-2003310030032222-0003112231202312-0030102103232320-3022001332322222-1121232202330002-0020100012312213-2133211103230020) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web` | [cloudflare.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-002.md#canonical-1220113120203121-0321130233111033-0200011103020021-2231333210202112-0212000020112012-1300200223321030-2000210110000023-2322021010031131) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-002.md#canonical-1303313102200013-0222132122330330-2022330302031311-3111023022320113-0011021303030010-0323011332113210-0220202300303121-1213102112133200) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-002.md#canonical-2323311222112121-2113101131132201-1310312131330012-1033210130203302-0220103103320330-3023110113000232-3213023010230200-2010030011002130) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-002.md#canonical-3322310100322132-2011033011100300-2203103213030121-0002023312111320-1323131230112332-0222100030120332-2033002211123313-0130003003220032) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.location](data-sources--protected_application--reference--group-002.md#canonical-0223103100301103-1130220121012220-3333123111231203-3130222111220233-3202000133122030-3201000031033221-3112021230231200-2102230313112110) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.status](data-sources--protected_application--reference--group-002.md#canonical-1321313331110232-2103303012201332-0232302211121023-2101131131223113-3223230021121011-0332220202320011-0220332011201133-3020210331211233) |
| `cloudflare.timeout` | [cloudflare.timeout](data-sources--protected_application--reference--group-001.md#canonical-1222221300223120-3113133302202013-1230032212131120-1032233100202033-2220002130020120-2310202303111132-3132023112133010-0303231322120302) |
| `cloudflare.trusted_clients` | [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-0133102322111220-3311232001000011-0230021001031010-1022212123321122-1123130233203113-1323330210311020-3110323313222301-1120130113013233) |
| `cloudflare.trusted_clients.http_header` | [cloudflare.trusted_clients.http_header](data-sources--protected_application--reference--group-002.md#canonical-3331130021002200-1121230312011231-0331100323123333-2031103323110322-1223013033300130-2001233130111002-1021201032222331-2001222332221103) |
| `cloudflare.trusted_clients.http_header.headers` | [cloudflare.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-002.md#canonical-3201332011001112-3200213100211003-1320021103130310-0320110301030203-0313313100313313-2113212330310012-0013122201333103-1221233210200310) |
| `cloudflare.trusted_clients.http_header.headers.exact` | [cloudflare.trusted_clients.http_header.headers.exact](data-sources--protected_application--reference--group-002.md#canonical-2031001310010111-2123222002333301-2222322302232033-2130020203033300-1231233123120211-2010032010310102-3330320201203010-1222223021122033) |
| `cloudflare.trusted_clients.http_header.headers.name` | [cloudflare.trusted_clients.http_header.headers.name](data-sources--protected_application--reference--group-002.md#canonical-1022303210231330-0300112130233100-1223323233212202-1231311101120110-2010230220033013-3032301100320230-2112211222200021-2123120320231131) |
| `cloudflare.trusted_clients.http_header.headers.regex` | [cloudflare.trusted_clients.http_header.headers.regex](data-sources--protected_application--reference--group-002.md#canonical-1111103330312232-3023222333111313-3013222121220132-3232031200211323-0100023312312223-2303022201212313-0113323303021213-2202032023200231) |
| `cloudflare.trusted_clients.ip_prefix` | [cloudflare.trusted_clients.ip_prefix](data-sources--protected_application--reference--group-002.md#canonical-3133011202313222-2231321120132330-2100330312012112-1320100332033113-1102130100303332-1212011011121320-1011331231023333-0320101202301023) |
| `cloudflare.trusted_clients.metadata` | [cloudflare.trusted_clients.metadata](data-sources--protected_application--reference--group-002.md#canonical-0210013011001033-2103113201332303-1013210202313003-3213000112121213-3232331223112212-1312310233212232-0113002220312202-0020011031110311) |
| `cloudflare.trusted_clients.metadata.description_spec` | [cloudflare.trusted_clients.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-3102200112001100-0302320111131300-3001023202330231-2101332330202102-2020011030303312-2031202020320100-3201023210031300-3303233103133210) |
| `cloudflare.trusted_clients.metadata.name` | [cloudflare.trusted_clients.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-3020101202300232-0103110132023222-0111132213330013-0132002333331313-2010031132201301-0032222003302211-1303130200300312-3002010032032001) |
| `cloudfront` | [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1001230210132003-1312122030130213-3323113320202103-0212010010230110-2113003310332323-2200230323320021-0021000333232102-1313302122221002) |
| `cloudfront.aws_configuration_id_selector` | [cloudfront.aws_configuration_id_selector](data-sources--protected_application--reference--group-002.md#canonical-0220011020023322-2010100312322220-1301222102301210-3122103330202030-2203100121303200-0102303311122322-1220201031201110-2132332121312113) |
| `cloudfront.aws_configuration_id_selector.ids` | [cloudfront.aws_configuration_id_selector.ids](data-sources--protected_application--reference--group-002.md#canonical-0212130313002203-2330022232023102-2113332210113110-2120102010120031-2022110322333023-1021122301102322-1302013113313201-2320021211213200) |
| `cloudfront.aws_configuration_tag_selector` | [cloudfront.aws_configuration_tag_selector](data-sources--protected_application--reference--group-002.md#canonical-3121212122301010-0021023210313022-2012123131120310-0133230001312122-1302310132011011-1010303032323323-1120020221213303-2031113212020230) |
| `cloudfront.aws_configuration_tag_selector.tags` | [cloudfront.aws_configuration_tag_selector.tags](data-sources--protected_application--reference--group-002.md#canonical-0310120031133223-3113330201332303-1331023132032323-2023221313012120-1231033102120101-2311223001203122-1130223123310202-1201312301121000) |
| `cloudfront.continue_mitigation_action_hdr` | [cloudfront.continue_mitigation_action_hdr](data-sources--protected_application--reference--group-002.md#canonical-1301031001003013-2221130322231013-3100001332203031-2221023330132113-2112310033313122-3203101010111211-0202200123001312-0212211011233021) |
| `cloudfront.data_sample` | [cloudfront.data_sample](data-sources--protected_application--reference--group-002.md#canonical-1132102323231201-2202010203312032-2222333222221131-3222300022031011-0001232023301231-2101011201233201-0123331210120113-3112213113101110) |
| `cloudfront.disable_aws_configuration` | [cloudfront.disable_aws_configuration](data-sources--protected_application--reference--group-002.md#canonical-2213221110001032-1000233003211022-1110113313110233-3301010111320201-1211101032222320-0223020010132233-2133333002221101-3000103213100111) |
| `cloudfront.disable_js_insert` | [cloudfront.disable_js_insert](data-sources--protected_application--reference--group-002.md#canonical-3110121231323013-2010102333202133-1223003323100001-1011201130020222-2221313203100310-2010330102021023-3122023000010030-1300121212112323) |
| `cloudfront.disable_mobile_sdk` | [cloudfront.disable_mobile_sdk](data-sources--protected_application--reference--group-002.md#canonical-1202133233323323-0222330321230102-0200201210331330-1220310232323100-0031000031010012-3331001130303213-3022323231212210-1101113200133302) |
| `cloudfront.js_insertion_rules` | [cloudfront.js_insertion_rules](data-sources--protected_application--reference--group-002.md#canonical-2222221321010301-0332022320311011-2331212321332101-1031010113312012-3210233121102310-0021232023021320-0311120231013130-1213111111103113) |
| `cloudfront.js_insertion_rules.exclude_list` | [cloudfront.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-002.md#canonical-1022110102232113-1322331322332002-1001302110311301-0032121222031212-0012233112203111-0032133031121332-1120111131301233-0031313201313003) |
| `cloudfront.js_insertion_rules.exclude_list.any_domain` | [cloudfront.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-002.md#canonical-0220230100013000-1333131101012111-2332130100013303-3231310230201322-0001030012302031-0020302112313120-3020212103103213-3302221102202232) |
| `cloudfront.js_insertion_rules.exclude_list.domain` | [cloudfront.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-002.md#canonical-2310011202103220-2120031201231120-2031111002322120-2101210213321012-0203031233210110-0023310222200011-3330121231230110-0320113000012130) |
| `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` | [cloudfront.js_insertion_rules.exclude_list.domain.exact_value](data-sources--protected_application--reference--group-002.md#canonical-1312003010021121-2131132233112203-1301133311330313-1302011000123203-1231102021200223-1211113223313022-0110311233132032-0320230312121003) |
| `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` | [cloudfront.js_insertion_rules.exclude_list.domain.regex_value](data-sources--protected_application--reference--group-002.md#canonical-2112021322332330-2131110002111133-2131203200303212-2212223020000312-3023023112112233-3220122222213010-0332201231321221-2122212103300133) |
| `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudfront.js_insertion_rules.exclude_list.domain.suffix_value](data-sources--protected_application--reference--group-002.md#canonical-3202103101210013-1322223301031303-2111001232303123-0130100201110211-3111130131332133-2111220210321210-1030001312213131-0300323120333213) |
| `cloudfront.js_insertion_rules.exclude_list.metadata` | [cloudfront.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-002.md#canonical-0111102101213003-1230113111031231-0202211021223213-3122322320000032-2111133303122031-0023010133301211-0230332122011120-2020303230020002) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudfront.js_insertion_rules.exclude_list.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-1212322113323201-3333131200112221-2300302222213123-2231021010230002-1033001230100033-0322330313130030-1202001210233000-1311002130033210) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.name` | [cloudfront.js_insertion_rules.exclude_list.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-0211112121011120-0112010221003202-1131131111110021-2330021223000302-2323011320020313-3100123003313300-2223321130323223-2123131011223222) |
| `cloudfront.js_insertion_rules.exclude_list.path` | [cloudfront.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-002.md#canonical-3121200221120213-1212323031312312-0120132303113322-3220232131333203-0320113033300333-3330203120033131-2230000132210033-1021131132131330) |
| `cloudfront.js_insertion_rules.exclude_list.path.path` | [cloudfront.js_insertion_rules.exclude_list.path.path](data-sources--protected_application--reference--group-002.md#canonical-0233000020202120-0013302133203213-0330231022023002-1011032021223201-3321321330021110-0122300030001123-3032313010202231-0111020030333131) |
| `cloudfront.js_insertion_rules.exclude_list.path.prefix` | [cloudfront.js_insertion_rules.exclude_list.path.prefix](data-sources--protected_application--reference--group-002.md#canonical-2002121032033202-1211222023232312-1223012331322120-3120101303122121-1133110300002201-2132121333312111-3000022310123101-3313302000300003) |
| `cloudfront.js_insertion_rules.exclude_list.path.regex` | [cloudfront.js_insertion_rules.exclude_list.path.regex](data-sources--protected_application--reference--group-002.md#canonical-2010031312303002-3211110123333130-0010300013013233-0013113130331021-0230330333102033-1312101213130333-1323333222201002-2003221002131312) |
| `cloudfront.js_insertion_rules.javascript_location` | [cloudfront.js_insertion_rules.javascript_location](data-sources--protected_application--reference--group-002.md#canonical-2213213031231212-0003203030110002-2220313131211120-2233100001023033-3012103100102332-0001230301020121-3000021330102332-2002210313023222) |
| `cloudfront.js_insertion_rules.javascript_mode` | [cloudfront.js_insertion_rules.javascript_mode](data-sources--protected_application--reference--group-002.md#canonical-0000203020212301-2133120200110232-2100323301212223-0032330210113322-0013322013032122-3110110130013100-0011110111120223-3131221231321320) |
| `cloudfront.js_insertion_rules.js_download_path` | [cloudfront.js_insertion_rules.js_download_path](data-sources--protected_application--reference--group-002.md#canonical-3211210313332010-2210332102000030-0010223111203102-2301100231103132-2201002031100021-1131200212211023-0301231111013133-1233023023232010) |
| `cloudfront.js_insertion_rules.rules` | [cloudfront.js_insertion_rules.rules](data-sources--protected_application--reference--group-002.md#canonical-3112211323200012-1223322323223121-1221102013102200-2231133101331021-2300001032233333-0231120330103210-1212210332300220-1012311313022321) |
| `cloudfront.js_insertion_rules.rules.any_domain` | [cloudfront.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-002.md#canonical-0222003222210323-3202111102203220-1102211320120213-0333111312123211-0013112101330113-2223033210103231-3112222122223112-2010012010101013) |
| `cloudfront.js_insertion_rules.rules.domain` | [cloudfront.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-002.md#canonical-3331221212300133-1223211303033303-2010130322021122-0102122012131310-2011202003022123-0013311223221032-0102102233320230-2132212220311310) |
| `cloudfront.js_insertion_rules.rules.domain.exact_value` | [cloudfront.js_insertion_rules.rules.domain.exact_value](data-sources--protected_application--reference--group-002.md#canonical-2123230330211300-0002023101232120-2012122123131320-0213333223012112-2331100330233130-0213213130033010-0122310001122012-0232300002112131) |
| `cloudfront.js_insertion_rules.rules.domain.regex_value` | [cloudfront.js_insertion_rules.rules.domain.regex_value](data-sources--protected_application--reference--group-002.md#canonical-1311132312311103-1133212022210313-2102113332101133-1122331100203012-3131133212210323-0212122022220222-2303200310303020-3221003033000202) |
| `cloudfront.js_insertion_rules.rules.domain.suffix_value` | [cloudfront.js_insertion_rules.rules.domain.suffix_value](data-sources--protected_application--reference--group-002.md#canonical-3213310122031123-3320313130031030-1133132220220232-1023201321003121-0311032300010220-2232031033203332-1212300320100230-2123332311302032) |
| `cloudfront.js_insertion_rules.rules.exact_path` | [cloudfront.js_insertion_rules.rules.exact_path](data-sources--protected_application--reference--group-002.md#canonical-2330022230132211-3221002200113200-1100131123202310-0210100330031012-0221112021331332-2011101230213322-2121232033121311-0332111332222022) |
| `cloudfront.js_insertion_rules.rules.glob` | [cloudfront.js_insertion_rules.rules.glob](data-sources--protected_application--reference--group-002.md#canonical-0213003102103301-0231133303321310-0010223000220100-0031033033231303-0310101301133333-0111300200213321-3331312121120012-3231012303221121) |
| `cloudfront.js_insertion_rules.rules.metadata` | [cloudfront.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-002.md#canonical-0022310302103031-0331231102030333-3120321222311312-0020131003331111-2212033323301313-1302223032031212-2331130022200020-1322331013311101) |
| `cloudfront.js_insertion_rules.rules.metadata.description_spec` | [cloudfront.js_insertion_rules.rules.metadata.description_spec](data-sources--protected_application--reference--group-002.md#canonical-3001231312311121-1312231001320021-0003313103123212-0330312112110303-1231003321101131-3301230303111333-1231030120122132-2131032220323020) |
| `cloudfront.js_insertion_rules.rules.metadata.name` | [cloudfront.js_insertion_rules.rules.metadata.name](data-sources--protected_application--reference--group-002.md#canonical-2123133012212133-2310202220311233-0213230023302321-3201300321300013-3210203312211110-3112113210112021-1303030212113023-2030110203023211) |
| `cloudfront.js_insertion_rules.rules.prefix` | [cloudfront.js_insertion_rules.rules.prefix](data-sources--protected_application--reference--group-002.md#canonical-1102112322211030-1110013213130020-1203320202302331-0111311100132202-2011023302003030-2110102210033222-3022310112030302-0120230202210230) |
| `cloudfront.loglevel` | [cloudfront.loglevel](data-sources--protected_application--reference--group-002.md#canonical-2201031311110222-3000013232133121-2021212222103210-3331001320232200-0322120223111212-0312110033300130-1232331330103030-3112102232233011) |
| `cloudfront.manual_js_insert` | [cloudfront.manual_js_insert](data-sources--protected_application--reference--group-002.md#canonical-1231011133010223-3122230322022221-3021222003233331-2101311111113001-2313033033311320-2101020223111333-2213123133202022-0212231202212132) |
| `cloudfront.manual_js_insert.javascript_mode` | [cloudfront.manual_js_insert.javascript_mode](data-sources--protected_application--reference--group-002.md#canonical-2031223130012320-1002023220033121-1123303221123200-2103321031032013-3223103101130120-3230211232313203-2030202220003312-2011020133131232) |
| `cloudfront.manual_js_insert.js_download_path` | [cloudfront.manual_js_insert.js_download_path](data-sources--protected_application--reference--group-002.md#canonical-0313331030133212-1032111323032303-0100021032110220-3232111112031033-0131011203020002-0330132231012123-3230212022110120-0201303001002023) |
| `cloudfront.mobile_sdk_config` | [cloudfront.mobile_sdk_config](data-sources--protected_application--reference--group-002.md#canonical-1312021210122012-0331300021020123-2230203331133203-1221133112213030-2301102122213221-0222131110123110-3321211320210302-3312002113320132) |
| `cloudfront.mobile_sdk_config.mobile_identifier` | [cloudfront.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-002.md#canonical-1330032323232100-1320130301232230-3123220222202301-2303203231131111-1230023201202201-3120001300002310-3121012232233113-1220110213011102) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers` | [cloudfront.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-002.md#canonical-1322312110312033-2320011333131330-2123130100300003-0333021023212202-2333312330312103-3120323300202121-2023102131330130-2300112230302233) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.exact](data-sources--protected_application--reference--group-002.md#canonical-2322223030021221-0033131321320231-0331223011022122-3102122102020112-2120010302122123-1021200210122323-3032003103331012-0113323213102132) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.name](data-sources--protected_application--reference--group-002.md#canonical-0102231110012330-0103310313211102-2013233022022133-0120310111110032-2330013120001022-3132122312130133-1210211300111203-3121103301110112) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.regex](data-sources--protected_application--reference--group-002.md#canonical-0332331133103302-3012133233310233-0332121101132122-1201001030222030-0012031130123333-2030011232113130-3131232201021223-3220310003032102) |
| `cloudfront.protected_endpoints` | [cloudfront.protected_endpoints](data-sources--protected_application--reference--group-003.md#canonical-0232213020212323-1020030122103132-2231022031131020-2201113331300303-0111301321330201-2122020311332330-3303032103301223-2001300220202110) |
| `cloudfront.protected_endpoints.any_domain` | [cloudfront.protected_endpoints.any_domain](data-sources--protected_application--reference--group-003.md#canonical-0302122110010130-2031103103101000-0302002022210020-0022230202103223-3313033133303101-1313013020203010-1202213023111123-1323003011003211) |
| `cloudfront.protected_endpoints.domain` | [cloudfront.protected_endpoints.domain](data-sources--protected_application--reference--group-003.md#canonical-2302231201023310-2313302100000222-3123132201202211-0313113312232311-3333220322013203-0011200020212100-0210213201332220-0202230123332103) |
| `cloudfront.protected_endpoints.domain.exact_value` | [cloudfront.protected_endpoints.domain.exact_value](data-sources--protected_application--reference--group-003.md#canonical-2232133333212102-1110113103100322-3223321213132000-2110230230210100-1113302133010222-1223303130133001-2132011213223011-3011122111003330) |
| `cloudfront.protected_endpoints.domain.regex_value` | [cloudfront.protected_endpoints.domain.regex_value](data-sources--protected_application--reference--group-003.md#canonical-2102020000130020-1230311222231003-2103223202013111-2021311330313121-3033021101322130-0230023011311030-2111113312202332-2000203312323320) |
| `cloudfront.protected_endpoints.domain.suffix_value` | [cloudfront.protected_endpoints.domain.suffix_value](data-sources--protected_application--reference--group-003.md#canonical-3100331203013213-2111233123322013-1323121230320310-3130113203100222-3032203021313231-2030330201023023-3120231113012002-0213033130221103) |
| `cloudfront.protected_endpoints.flow_label` | [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--reference--group-003.md#canonical-2020033321111033-2202222133123020-2131133212213323-0222100102031133-2033032012121202-2313202322130212-3310021331311022-3013222210000203) |
| `cloudfront.protected_endpoints.flow_label.account_management` | [cloudfront.protected_endpoints.flow_label.account_management](data-sources--protected_application--reference--group-003.md#canonical-3200011001300121-1012011300020310-1220233013302021-3201222303301312-3211000322133331-3201312001031220-1222300101012023-2220323221201110) |
| `cloudfront.protected_endpoints.flow_label.account_management.create` | [cloudfront.protected_endpoints.flow_label.account_management.create](data-sources--protected_application--reference--group-003.md#canonical-1201211312121211-1211310230220101-2220221303212022-2331111113122211-1220302231220100-0023210303121113-2212103013222033-3202101121310120) |
| `cloudfront.protected_endpoints.flow_label.account_management.password_reset` | [cloudfront.protected_endpoints.flow_label.account_management.password_reset](data-sources--protected_application--reference--group-003.md#canonical-3021002030312000-1311121111111201-2210211023013021-3133033102312302-3113131211312223-0323023003100120-2202320032213223-1000020231130232) |
| `cloudfront.protected_endpoints.flow_label.authentication` | [cloudfront.protected_endpoints.flow_label.authentication](data-sources--protected_application--reference--group-003.md#canonical-0133021313100221-0203311121003231-0123310101323110-3033333332123233-2021031323112332-0132232031002103-2103131331030100-0020021033001330) |
| `cloudfront.protected_endpoints.flow_label.authentication.login` | [cloudfront.protected_endpoints.flow_label.authentication.login](data-sources--protected_application--reference--group-003.md#canonical-1012103133031001-2312020033233231-0020320003313002-3020021322310220-0202223030120333-0201212223222101-0212211012133222-2022113330310122) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--protected_application--reference--group-003.md#canonical-0122200211003002-0130212322323301-1102122331111333-1022123311312002-3333100312023032-2121102301313122-2110203331120233-0200020230232231) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](data-sources--protected_application--reference--group-003.md#canonical-2212111022132332-2202111122203111-3101231323033320-1020002223321303-3111111021122133-3303310301331013-3013112101123012-2310000021322132) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--protected_application--reference--group-003.md#canonical-3310220233022301-2120232333033111-2223213113122231-0133320130220230-3333333122031123-1222300203213213-2022013001223211-0213100330200111) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name](data-sources--protected_application--reference--group-003.md#canonical-0320230031002303-2222132230212120-2320030213131023-1320022222222023-2301202223300311-2220210231110221-3010023030120322-3112112210131202) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values](data-sources--protected_application--reference--group-003.md#canonical-1203203012103232-1030222322300120-3200233103020223-2120203231233323-3123202221003323-0232312200130221-2210201122230110-0313313001100030) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status](data-sources--protected_application--reference--group-003.md#canonical-2223312011002310-1002321110003230-3313303301231320-1003102202313112-3313130310322121-3303223313303303-0331302230233012-2311011331222120) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--protected_application--reference--group-003.md#canonical-3012221033231200-0121302301101011-1033122332123023-0011011020011121-0030030132332003-0330201132202312-2002330030231131-2313312232120302) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name](data-sources--protected_application--reference--group-003.md#canonical-3003122003321110-2031233012202203-3021001023022032-0300220330000103-2022201123303110-1221232102123010-2331213333021013-3201122013303311) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values](data-sources--protected_application--reference--group-003.md#canonical-0011213333313120-2230221211101311-0130111211213103-3331022332001203-2322021122111301-2131221013112032-1211210223132221-2321033333331320) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status](data-sources--protected_application--reference--group-003.md#canonical-2003212000301303-2113020321331212-1211311320301102-2332303030121022-2001302112233300-2121020303201000-2201033012102021-0323133011332232) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` | [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](data-sources--protected_application--reference--group-003.md#canonical-3220033113120103-2101233210303221-1002232301123312-0321023123232323-0121012110130333-0223130032323223-2230023112202232-3222033332211331) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_partner` | [cloudfront.protected_endpoints.flow_label.authentication.login_partner](data-sources--protected_application--reference--group-003.md#canonical-0302023012233013-3221000023200010-2031010330020232-1002320220210023-1003200130132023-1131132111021120-2113330032021110-0320331232321001) |
| `cloudfront.protected_endpoints.flow_label.authentication.logout` | [cloudfront.protected_endpoints.flow_label.authentication.logout](data-sources--protected_application--reference--group-003.md#canonical-1002313002311101-0331131111301323-0133232313003312-2222310222101123-0023101303203211-0133103013010102-1330011002323230-1102111211303301) |
| `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` | [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](data-sources--protected_application--reference--group-003.md#canonical-2301231112210332-3333321122233312-0221330323331200-3233330003013000-2132333300331222-3320331231013100-2121203323120331-0220220300031002) |
| `cloudfront.protected_endpoints.flow_label.financial_services` | [cloudfront.protected_endpoints.flow_label.financial_services](data-sources--protected_application--reference--group-003.md#canonical-2222023211012313-1012110111322331-1201302320020311-1020001302312203-0230132020032331-0231021211110203-2133200311112303-1100233031130313) |
| `cloudfront.protected_endpoints.flow_label.financial_services.apply` | [cloudfront.protected_endpoints.flow_label.financial_services.apply](data-sources--protected_application--reference--group-003.md#canonical-1132102021302213-0122321030023122-1330021103020030-3302210322132330-0333200312300333-0012302301021030-3303312321330010-1123320001302102) |
| `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` | [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](data-sources--protected_application--reference--group-003.md#canonical-3302201311301221-1330002120222203-1313200313230013-3030023120220001-0002203001221011-0100100021200030-1211100222111002-2111232321310021) |
| `cloudfront.protected_endpoints.flow_label.flight` | [cloudfront.protected_endpoints.flow_label.flight](data-sources--protected_application--reference--group-003.md#canonical-1232122302233330-0213332221113131-2320101130100000-1112203011031213-1222121212102130-0330113001231220-3132010210130132-1101310120300233) |
| `cloudfront.protected_endpoints.flow_label.flight.checkin` | [cloudfront.protected_endpoints.flow_label.flight.checkin](data-sources--protected_application--reference--group-003.md#canonical-3202003000010110-1333100201101313-0201201131310222-0033123002200113-2223002211020331-2010103312103302-0111312221103301-3102220011320322) |
| `cloudfront.protected_endpoints.flow_label.profile_management` | [cloudfront.protected_endpoints.flow_label.profile_management](data-sources--protected_application--reference--group-003.md#canonical-2303110110213010-0232001133101302-3333101000112220-0120202302203331-2300031022332211-2333332101033002-0131203101001320-2320311002011232) |
| `cloudfront.protected_endpoints.flow_label.profile_management.create` | [cloudfront.protected_endpoints.flow_label.profile_management.create](data-sources--protected_application--reference--group-003.md#canonical-3033232101003130-0130121022102313-3003233210333122-3322323230200121-3022303311130021-3223022113223210-3332313020132301-1121121011203332) |
| `cloudfront.protected_endpoints.flow_label.profile_management.update` | [cloudfront.protected_endpoints.flow_label.profile_management.update](data-sources--protected_application--reference--group-003.md#canonical-2023002023120122-0321323102011210-0021310330010132-1321001130212333-1323111030222233-0232213310111123-0320323213331010-2113210223003102) |
| `cloudfront.protected_endpoints.flow_label.profile_management.view` | [cloudfront.protected_endpoints.flow_label.profile_management.view](data-sources--protected_application--reference--group-003.md#canonical-0011211333313021-3300023230323311-1323231023201311-3101010312130033-1030231332211103-1212120320103111-1212033200232001-3003232213002103) |
| `cloudfront.protected_endpoints.flow_label.search` | [cloudfront.protected_endpoints.flow_label.search](data-sources--protected_application--reference--group-003.md#canonical-2230122110213122-1110120131032001-3233210112103300-2013030213210220-2010002313212221-2311233132223310-2030032121230020-3001013202212023) |
| `cloudfront.protected_endpoints.flow_label.search.flight_search` | [cloudfront.protected_endpoints.flow_label.search.flight_search](data-sources--protected_application--reference--group-003.md#canonical-3320032311010300-1112022210110230-0231102330123022-0020200011121122-1221023121303013-3010213213002222-2002123103100203-0033302323101213) |
| `cloudfront.protected_endpoints.flow_label.search.product_search` | [cloudfront.protected_endpoints.flow_label.search.product_search](data-sources--protected_application--reference--group-003.md#canonical-2012312231223230-3221012223033223-1101331103003103-1022211110012232-1203233220331321-1131231202322112-3000232313303130-3223013313332001) |
| `cloudfront.protected_endpoints.flow_label.search.reservation_search` | [cloudfront.protected_endpoints.flow_label.search.reservation_search](data-sources--protected_application--reference--group-003.md#canonical-3111021213123323-1331313023032322-3222322122210221-0320113013231302-3111323000211210-2230112110233120-0312030013331113-1100302232301312) |
| `cloudfront.protected_endpoints.flow_label.search.room_search` | [cloudfront.protected_endpoints.flow_label.search.room_search](data-sources--protected_application--reference--group-003.md#canonical-2022111222122031-1023322310111013-3202131313221030-3212000313112132-0300000320332311-0201122112123303-2033113121213203-1313110212131303) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](data-sources--protected_application--reference--group-003.md#canonical-1201220110210032-0222222220330220-2320333230100312-1221212001303310-1133121001233200-3233203133331003-3032131032201103-2101212313111010) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](data-sources--protected_application--reference--group-003.md#canonical-3102312231001223-2231331203211103-3123132322130202-3111223003101031-3020020131203211-2133002122112123-3003202320000332-3301121130133012) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](data-sources--protected_application--reference--group-003.md#canonical-2130201102102003-1312021230333312-0000121200230133-3301323131200211-0301101101220311-2300031213313001-3032330101012120-3110302222030300) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](data-sources--protected_application--reference--group-003.md#canonical-0103011120333033-0231032003232011-0233320210232300-1210032121223310-1210022233310113-2112130220202213-1122330302030201-1313100132301302) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](data-sources--protected_application--reference--group-003.md#canonical-0131112222223321-3220121211321201-2101011010012331-0113020213121030-1012320031200133-3110013002212021-3122030112102233-2001000032011221) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](data-sources--protected_application--reference--group-003.md#canonical-1030022113131102-2020321122310200-0313133330201222-3330011100012132-0023122000313033-0013220331301123-1033012302111222-0130210213330111) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](data-sources--protected_application--reference--group-003.md#canonical-0230101322032220-1313022012323120-1300112112012300-1212020303123320-2330031231033031-2223113013110210-1102221233301221-2330130130322332) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](data-sources--protected_application--reference--group-003.md#canonical-1212321202210332-1113220023331001-3323212222103212-2100231331302022-3001000310112001-0302012003333020-0110211232301122-1302220201002220) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](data-sources--protected_application--reference--group-003.md#canonical-0322202003100020-0213321032201120-2333001311001320-0120213020330012-1032030212001210-0001020012311012-1111321230300222-1103123002022320) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](data-sources--protected_application--reference--group-003.md#canonical-3132132203113113-2223113301110110-2120313233113330-0002031321133200-1101333131010320-3212311103011320-0222223330023220-1022331002233210) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](data-sources--protected_application--reference--group-003.md#canonical-1210312200102232-1233331312321120-0011131202100003-2122203303333111-2320003012003301-0003131010031022-0101213132300213-2310202321232302) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](data-sources--protected_application--reference--group-003.md#canonical-3022012322113332-2120133310310220-3023003111203110-2130202131230333-2002201112020122-0132013211033331-1310221213320301-2112023103201312) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](data-sources--protected_application--reference--group-003.md#canonical-0213020101000331-0113230001002102-0131200111013323-0230003223312211-2021310232121020-0022331330111110-1113110312202103-3003312220103331) |
| `cloudfront.protected_endpoints.http_methods` | [cloudfront.protected_endpoints.http_methods](data-sources--protected_application--reference--group-003.md#canonical-3303203312203031-0120303221223312-0331030213320330-1330310122213111-0232213031112330-2102000122301330-1321101100210332-3000011322133232) |
| `cloudfront.protected_endpoints.metadata` | [cloudfront.protected_endpoints.metadata](data-sources--protected_application--reference--group-003.md#canonical-1110001100020130-3103033011210222-3100333313123103-3322333123021233-0321332013212031-0101103202330130-1121233111030123-2330001011002322) |
| `cloudfront.protected_endpoints.metadata.description_spec` | [cloudfront.protected_endpoints.metadata.description_spec](data-sources--protected_application--reference--group-003.md#canonical-0030231303213012-1312233012332211-0102320012120013-2133211022311301-3230233311032030-3013230220011202-3331000111332210-2210212233222322) |
| `cloudfront.protected_endpoints.metadata.name` | [cloudfront.protected_endpoints.metadata.name](data-sources--protected_application--reference--group-003.md#canonical-3021313132000132-1121210301232100-1222221101003000-2132010321031130-1010321113321113-0321000121003011-3033120003001303-0321112013013222) |
| `cloudfront.protected_endpoints.mobile_client` | [cloudfront.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-003.md#canonical-0012112132311220-3113112100222332-2312111022211122-0100131000310023-1030012122303200-1201102321123003-3310000310201122-3222301102100232) |
| `cloudfront.protected_endpoints.mobile_client.block` | [cloudfront.protected_endpoints.mobile_client.block](data-sources--protected_application--reference--group-003.md#canonical-0310113123330330-1013111210332022-1031311202011212-1101221020222013-0212301000102101-1232323021132002-1322000231302002-3033201233120033) |
| `cloudfront.protected_endpoints.mobile_client.block.body` | [cloudfront.protected_endpoints.mobile_client.block.body](data-sources--protected_application--reference--group-003.md#canonical-0212111030330113-0001323233021023-2303223113021031-1232303222011333-3010202021012012-2302221210013300-1220212003103301-3113012212121301) |
| `cloudfront.protected_endpoints.mobile_client.block.content_type` | [cloudfront.protected_endpoints.mobile_client.block.content_type](data-sources--protected_application--reference--group-003.md#canonical-2101332323131011-3010000130112303-0022230311012213-0321233100300020-0222002120102112-0020333212330212-2303200312030102-3311133323223123) |
| `cloudfront.protected_endpoints.mobile_client.block.status` | [cloudfront.protected_endpoints.mobile_client.block.status](data-sources--protected_application--reference--group-003.md#canonical-3023210001121311-0321023300121333-2233320232333332-3321220333213110-0121032030033012-0333112010123201-0312001211101011-0003111332201032) |
| `cloudfront.protected_endpoints.mobile_client.continue` | [cloudfront.protected_endpoints.mobile_client.continue](data-sources--protected_application--reference--group-003.md#canonical-2220132120301131-3031010223230013-0010201221201202-1013111230331312-1021013100210102-1023332233133023-3013210133133110-2023133121233213) |
| `cloudfront.protected_endpoints.mobile_client.continue.add_header` | [cloudfront.protected_endpoints.mobile_client.continue.add_header](data-sources--protected_application--reference--group-004.md#canonical-2310323321213111-1230121301201001-0323102301200112-3111003003013101-3121303212033003-0230212230332011-0100213230022330-3222032313003221) |
| `cloudfront.protected_endpoints.mobile_client.continue.no_header` | [cloudfront.protected_endpoints.mobile_client.continue.no_header](data-sources--protected_application--reference--group-004.md#canonical-3203232102103023-3003022312320232-0321030031012131-3323302230013213-2000330311023320-2133123321300331-2311012200112310-2200103210201102) |
| `cloudfront.protected_endpoints.path` | [cloudfront.protected_endpoints.path](data-sources--protected_application--reference--group-003.md#canonical-2122012101213113-2221201200303102-2122011020233202-0221332133220331-0210231203321001-2123103300221131-3123322013331210-0213111111132332) |
| `cloudfront.protected_endpoints.query` | [cloudfront.protected_endpoints.query](data-sources--protected_application--reference--group-003.md#canonical-3032233123200112-1222002132032020-0331330222131300-1331113200212013-3133301233321120-0030323330310210-1112032002302112-0231003232201101) |
| `cloudfront.protected_endpoints.undefined_flow_label` | [cloudfront.protected_endpoints.undefined_flow_label](data-sources--protected_application--reference--group-004.md#canonical-1131023123131210-3311232222313021-2103013003211112-1012331100003000-1032222020333332-2200301020001233-3311202020112330-3033101100313113) |
| `cloudfront.protected_endpoints.web_client` | [cloudfront.protected_endpoints.web_client](data-sources--protected_application--reference--group-004.md#canonical-0132013010120303-2103320113312212-2010213313032232-0113232232231121-3212013332200203-1010021321211110-1310000021233222-0100003121223331) |
| `cloudfront.protected_endpoints.web_client.block` | [cloudfront.protected_endpoints.web_client.block](data-sources--protected_application--reference--group-004.md#canonical-0021132222022303-1012311111331313-2102313210011220-2101313032131213-2202303003120320-3012203332330320-1003123011022202-3133213321112012) |
| `cloudfront.protected_endpoints.web_client.block.body` | [cloudfront.protected_endpoints.web_client.block.body](data-sources--protected_application--reference--group-004.md#canonical-0313000223230310-3031123201230300-1012113001332011-2122032322001113-2022222232332322-0021322111101111-0233303030122113-2330223321230203) |
| `cloudfront.protected_endpoints.web_client.block.content_type` | [cloudfront.protected_endpoints.web_client.block.content_type](data-sources--protected_application--reference--group-004.md#canonical-0312123330300313-3131313010103012-0020021223103230-0300030001300310-0031013121313122-3230310003112030-0021101231333210-1203312213010202) |
| `cloudfront.protected_endpoints.web_client.block.status` | [cloudfront.protected_endpoints.web_client.block.status](data-sources--protected_application--reference--group-004.md#canonical-1121313033323211-2213210223022100-2031110323010223-0133202122203010-0310111221022010-3111223222001332-1203012033011211-1233101210132202) |
| `cloudfront.protected_endpoints.web_client.continue` | [cloudfront.protected_endpoints.web_client.continue](data-sources--protected_application--reference--group-004.md#canonical-2232102001222001-0312102033201032-2223022331022231-2131011030113003-3333211011012301-1020100301300310-2022212333112011-3101210111133032) |
| `cloudfront.protected_endpoints.web_client.continue.add_header` | [cloudfront.protected_endpoints.web_client.continue.add_header](data-sources--protected_application--reference--group-004.md#canonical-3312233101012121-2121220331001112-3001121011023123-0023123111103020-3020111221013033-1203000120213213-1302202203202230-2222210331231022) |
| `cloudfront.protected_endpoints.web_client.continue.no_header` | [cloudfront.protected_endpoints.web_client.continue.no_header](data-sources--protected_application--reference--group-004.md#canonical-3002302311202203-1203322011330130-1131330012230002-0232021001233131-2203011132122331-1310003001131212-0332113121033130-3201311110002203) |
| `cloudfront.protected_endpoints.web_client.redirect` | [cloudfront.protected_endpoints.web_client.redirect](data-sources--protected_application--reference--group-004.md#canonical-2313323133022323-3231102321202003-2220233132302023-1221030230320331-3033012022200011-3212302102113110-1213021331222002-3303203203202002) |
| `cloudfront.protected_endpoints.web_client.redirect.location` | [cloudfront.protected_endpoints.web_client.redirect.location](data-sources--protected_application--reference--group-004.md#canonical-3332330033100032-2200021110023011-0023300023200010-0020233033002001-0233003203301233-3020012111102211-2013030321023220-0013220323012302) |
| `cloudfront.protected_endpoints.web_client.redirect.status` | [cloudfront.protected_endpoints.web_client.redirect.status](data-sources--protected_application--reference--group-004.md#canonical-1002013023012123-2102001302201200-0023200221301312-3213110200233133-1202221103300022-0131013300121211-1310103133300113-2233201322212212) |
| `cloudfront.protected_endpoints.web_mobile_client` | [cloudfront.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-004.md#canonical-0210021313132131-3110323200113102-2221311032333000-2022321133212032-0122002022333031-3301313032222201-3330320133102013-3022103003010021) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile](data-sources--protected_application--reference--group-004.md#canonical-3001003000031003-0001311103012310-0210320012020020-2320302020003030-3110232022200331-0222132031010132-1321233302231221-2210033301000333) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.body](data-sources--protected_application--reference--group-004.md#canonical-2121010012310103-3223321003023123-2230332201012112-1113111012121122-2232013332232331-0111323130013303-0021210013203110-2310010023320131) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type](data-sources--protected_application--reference--group-004.md#canonical-3132320012103031-1110022323023033-0023200100201322-3112100211101131-2220030301012102-2202222300121003-3311223032030120-0130002001122121) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.status](data-sources--protected_application--reference--group-004.md#canonical-3101001332322003-1320313123102003-3021133003100130-0310311030002303-0031202132213201-1121011332110133-2321200301201231-3011223020300020) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web` | [cloudfront.protected_endpoints.web_mobile_client.block_web](data-sources--protected_application--reference--group-004.md#canonical-0130303031210023-1132330021311213-3311320011321123-0200321030223001-1331333133301310-3013021233233332-3120212111112311-0021311132002202) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.body` | [cloudfront.protected_endpoints.web_mobile_client.block_web.body](data-sources--protected_application--reference--group-004.md#canonical-1311102033130230-3123232321123003-2113201012023300-1232110002232100-3031203023310300-2230002030103232-2110203133331312-0011023303002323) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_web.content_type](data-sources--protected_application--reference--group-004.md#canonical-2223110020031101-1311232002012322-0030022300311013-1202330122302133-1210130003211023-3131110131330202-2302213100322033-1132300310333032) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.status` | [cloudfront.protected_endpoints.web_mobile_client.block_web.status](data-sources--protected_application--reference--group-004.md#canonical-3203300301211333-3022333203310023-3323330013121321-0331111231010032-3100133112032302-2300203322220110-0310111121223132-1131210300310113) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](data-sources--protected_application--reference--group-004.md#canonical-3311310102322302-3233230333232311-0330321132131311-0302000310012230-1023032000222203-0020002130203302-2111110333302002-1023211120230311) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](data-sources--protected_application--reference--group-004.md#canonical-0312201031012102-2303130210331203-2110031221303111-2210102231332101-1323203102230321-0032202312300321-3032122103222303-2130203012331120) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](data-sources--protected_application--reference--group-004.md#canonical-0101310113020002-3230121001023120-0022000011312131-0110102102121332-0322221312030123-3223002021302011-0102321322201223-1203313231100112) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web` | [cloudfront.protected_endpoints.web_mobile_client.continue_web](data-sources--protected_application--reference--group-004.md#canonical-2333330003333331-3110032200301103-2323000123030110-3322210001001222-1302033032202323-1322002333233221-0112311301003203-0303103203130303) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](data-sources--protected_application--reference--group-004.md#canonical-1210012112012112-1202232330203203-0033333301132102-2211320123333202-2130030201213103-0122210202012132-2331032113032003-2233022120023000) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](data-sources--protected_application--reference--group-004.md#canonical-3321220300202330-0030120001103233-2201210110100212-0131200212133102-0230222130023122-0022221013323110-2122311302120121-0312231302310011) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web](data-sources--protected_application--reference--group-004.md#canonical-3021222322320121-2212232201111231-2123323230201332-0213031021330322-3030210102323331-2233200222202301-2012313000102011-1011121002212102) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.location](data-sources--protected_application--reference--group-004.md#canonical-3120300020323012-1032020012221300-2212022332223101-3110013221132021-0020012210203121-0123320103311223-2232023303011311-1030220031320013) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.status](data-sources--protected_application--reference--group-004.md#canonical-1301113110133302-0302110000110301-2133012333030301-0130200002333233-2022130201301112-2102201321212030-3130112123031202-3232232331120001) |
| `cloudfront.timeout` | [cloudfront.timeout](data-sources--protected_application--reference--group-002.md#canonical-1301230011221011-0212111310230231-2110331232201203-1133011000012311-0320213032210100-1132332330023202-2121330001021322-0003102311110013) |
| `cloudfront.trusted_clients` | [cloudfront.trusted_clients](data-sources--protected_application--reference--group-004.md#canonical-1130010221332230-2323323230311312-1221202002200233-3033122032331300-2111202310333321-1102113333001003-3331102220333303-0310331100113302) |
| `cloudfront.trusted_clients.http_header` | [cloudfront.trusted_clients.http_header](data-sources--protected_application--reference--group-004.md#canonical-0010102220332203-2011120013322001-2202012021331323-0211333201321102-3003122130322022-3310123212213322-0202023331323321-1230300030301321) |
| `cloudfront.trusted_clients.http_header.headers` | [cloudfront.trusted_clients.http_header.headers](data-sources--protected_application--reference--group-004.md#canonical-0232103231001031-3332300221030130-3023023032020223-3113200102001002-1313333201123230-0122333123202031-3320033232300320-1012200121131323) |
| `cloudfront.trusted_clients.http_header.headers.exact` | [cloudfront.trusted_clients.http_header.headers.exact](data-sources--protected_application--reference--group-004.md#canonical-0110332230310333-0111021301202112-0013220212012110-2120003030013010-2301212312133213-3302233201203022-2220033230311301-1321332232331132) |
| `cloudfront.trusted_clients.http_header.headers.name` | [cloudfront.trusted_clients.http_header.headers.name](data-sources--protected_application--reference--group-004.md#canonical-2032000102003201-0221300100122002-0022222101100000-2211331210102201-0113212202333320-3323303113322011-1201331032032311-2201232022110230) |
| `cloudfront.trusted_clients.http_header.headers.regex` | [cloudfront.trusted_clients.http_header.headers.regex](data-sources--protected_application--reference--group-004.md#canonical-2230012103033322-3222301200010112-1010132322123310-3321033033320112-1132331112221201-1131032332300323-1030033130110101-0113121320031103) |
| `cloudfront.trusted_clients.ip_prefix` | [cloudfront.trusted_clients.ip_prefix](data-sources--protected_application--reference--group-004.md#canonical-0223302211003300-0112322130232222-2203120311021230-0111230301011132-2223130123320301-1100010212201010-2022333013121002-3012113003203302) |
| `cloudfront.trusted_clients.metadata` | [cloudfront.trusted_clients.metadata](data-sources--protected_application--reference--group-004.md#canonical-0010202132112020-1001320000300321-0222213011012123-1322323033122102-2022032133030001-1110022020232002-1321232221130233-1233033222233322) |
| `cloudfront.trusted_clients.metadata.description_spec` | [cloudfront.trusted_clients.metadata.description_spec](data-sources--protected_application--reference--group-004.md#canonical-3231302311120302-3032100030203131-1000120302130010-1310311313311322-1032033330020332-1332131103303112-0033223331101321-2012121302021301) |
| `cloudfront.trusted_clients.metadata.name` | [cloudfront.trusted_clients.metadata.name](data-sources--protected_application--reference--group-004.md#canonical-0213120313022013-1333031332010001-3300321031103312-1010223211230233-2010111221002311-3312330313212202-1130313101011302-3220222020202323) |
| `custom_connector` | [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-2102330322032211-0103123332032130-2302031003011130-3102032101232210-2031020330020132-1000303002313201-0220212130010110-1223321031220121) |
| `description` | [description](data-sources--protected_application--reference--group-001.md#canonical-3312333012213211-0323230320123103-0200121311213311-3200112123103321-2201223203322313-0231111121203232-2313210313102233-2111301333220333) |
| `f5_big_ip` | [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-0332302203002303-1323230123130022-0213100021022021-2230320230121123-0321322312323030-3013332201212200-1223333231331113-1331213022012310) |
| `id` | [ID](data-sources--protected_application--reference--group-001.md#canonical-2111001110230031-2101333220323113-3211002122022103-3132331302221002-1211231113133022-3331202333110120-3220321233330201-3211331232110233) |
| `labels` | [labels](data-sources--protected_application--reference--group-001.md#canonical-2132230300201311-3111310330002222-3120230100322021-1022303202131103-0123301030112002-2323220030013321-1123233103333313-0003121203110330) |
| `name` | [name](data-sources--protected_application--reference--group-001.md#canonical-0023123123230212-1103211322111123-0010012032112213-2323002311113300-2120302132200122-3133233233313222-1022002111200102-1001233211312221) |
| `namespace` | [namespace](data-sources--protected_application--reference--group-001.md#canonical-1310103102101223-0230021323212321-3133212230121222-1323131012002322-1230013200010023-1231231003032033-0123021003101203-3100001210023333) |
| `region` | [region](data-sources--protected_application--reference--group-001.md#canonical-0010102001012002-1323203111213100-2032220022322220-3320220112200213-3132313211121332-0031001312231200-1031032001313322-3333310111130122) |
| `salesforce_commerce_connector` | [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-2112311010211032-1110123211222023-1011322203103220-2110303202120102-0020230102200300-2312033213021200-0313232030302100-1031231203200131) |

<a id="canonical-2331011110101100-1200301022022301-0131021310030333-2110222213002020-3100020210112311-0202211320022203-1201212303223212-1233112031130021"></a>

## Next pages — Property reference / 300210112021 / 12

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-2131122002322200-1211102013312210-0332313001322211-3203310002112222-1102230310021123-1302303130013203-0210310220203232-1201203310213203)
- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-3121210303101201-1231303202003123-1231320121022310-3001201012333211-2010131131103323-0230330313113000-3233023131132211-3332203002023000)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1123020123100122-0120011030211112-1131223130120230-0022231222113022-1121113231021321-3323303132213221-0213122003003202-2311323222001310)
- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-2131022120013230-3323322131002320-3112130302320203-0132230022331312-3202130201002000-0310132300310330-1103022101033312-1203333103123122)
- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-2221321233201022-2312300032001020-0333110031033331-0133213100320220-0313313301321000-2131300212321000-1211302300200001-2300130230311301)
- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-3300103021210020-1101231333023233-1202122312030020-1322030202123101-2122211031213030-0203233223330223-3331032100031322-3101033002131222)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2131122002322200-1211102013312210-0332313001322211-3203310002112222-1102230310021123-1302303130013203-0210310220203232-1201203310213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133222131233030-3033121010233213-3112030222220310-3122132321033331-1231132111020132-1332020331232021-1020001220032102-3211330302232301"></a>

## adobe_commerce_connector — adobe_commerce_connector / 213311331222 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- adobe_commerce_connector

<a id="canonical-0011311010322211-0231123221220032-2221230102020200-1320103102222213-3102313202333231-2102220032210002-1112232123103010-3010331333301311"></a>

Type: `["object", {}]`. Computed.

\[OneOf: adobe\_commerce\_connector, big\_ip\_iapp, Cloudflare, CloudFront, custom\_connector,
f5\_big\_ip, Salesforce\_commerce\_connector\] Configuration parameter for adobe commerce connector.

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

- [adobe_commerce_connector](data-sources--protected_application--reference--group-001.md#canonical-0011311010322211-0231123221220032-2221230102020200-1320103102222213-3102313202333231-2102220032210002-1112232123103010-3010331333301311)
- [big_ip_iapp](data-sources--protected_application--reference--group-001.md#canonical-2310010302100121-2100201222033113-0121010133333222-3331301002332312-2003311113020202-2201231133202112-3300310301033320-0102232230321033)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-3023111033212211-1111122330202301-0012113103330120-3112300011213122-1111332310321102-2101102231002202-3320333023323111-0023110123110001)
- [CloudFront](data-sources--protected_application--reference--group-002.md#canonical-1001230210132003-1312122030130213-3323113320202103-0212010010230110-2113003310332323-2200230323320021-0021000333232102-1313302122221002)
- [custom_connector](data-sources--protected_application--reference--group-004.md#canonical-2102330322032211-0103123332032130-2302031003011130-3102032101232210-2031020330020132-1000303002313201-0220212130010110-1223321031220121)
- [f5_big_ip](data-sources--protected_application--reference--group-004.md#canonical-0332302203002303-1323230123130022-0213100021022021-2230320230121123-0321322312323030-3013332201212200-1223333231331113-1331213022012310)
- [salesforce_commerce_connector](data-sources--protected_application--reference--group-004.md#canonical-2112311010211032-1110123211222023-1011322203103220-2110303202120102-0020230102200300-2312033213021200-0313232030302100-1031231203200131)

Select alternatives according to the provider validators above.

<a id="canonical-2321013031033132-1013211310032230-3230001001223330-2213221010130101-1012011331302202-0223232220023320-3211200103232100-0132223303332130"></a>

## Direct properties — adobe_commerce_connector / 213311331222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200122010000233-2332000113301310-0231113310133333-0012300333200231-1123300321323003-3013001321013331-0012202132221102-0303112121133013"></a>

## Next pages — adobe_commerce_connector / 213311331222 / 4

- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3121210303101201-1231303202003123-1231320121022310-3001201012333211-2010131131103323-0230330313113000-3233023131132211-3332203002023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023032121221321-1313220320201002-2131122221333121-1331301123101231-2213230313232302-1100002012220133-2222202113030023-3103211013213232"></a>

## big_ip_iapp — big_ip_iapp / 121231133313 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- big_ip_iapp

<a id="canonical-2310010302100121-2100201222033113-0121010133333222-3331301002332312-2003311113020202-2201231133202112-3300310301033320-0102232230321033"></a>

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

<a id="canonical-0300312300313122-3022303013302012-2313231220010111-0013021331211130-2001210030222210-3021031120203232-3131100333001133-3023203201321300"></a>

## Direct properties — big_ip_iapp / 121231133313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001012230031203-1301012310123023-3303130232300101-3331102223010333-1332210323221132-0012031110032202-3202103211033100-2121322311211023"></a>

## Next pages — big_ip_iapp / 121231133313 / 4

- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300331233201320-3020233122132301-1120003030211211-1112103020221011-1032001130030122-2001013333232210-1022013300212131-3212331202310113"></a>

## Cloudflare — Cloudflare / 210023001213 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- Cloudflare

<a id="canonical-3023111033212211-1111122330202301-0012113103330120-3112300011213122-1111332310321102-2101102231002202-3320333023323111-0023110123110001"></a>

Type: `"single"`. Computed.

Bot Defense policy configuration for Cloudflare.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

<a id="canonical-1303200300022220-1122300220332002-3331233012123203-0022302131003331-2111113220232030-1201012111222322-1310002131113010-3222101211030210"></a>

## Direct properties — Cloudflare / 210023001213 / 3

<a id="canonical-2202200332333113-0321020131310023-2022311121012120-0012102313103032-3210000001332123-3331303330102300-1020222230221302-3113201132313323"></a>

<a id="canonical-1103112012223223-2202030112230211-2120213210221030-0132222203311033-1302123210303231-0131002133230230-2300100201123323-1220003210031222"></a>

## continue_mitigation_action_hdr property — Cloudflare / 210023001213 / 4

Type: `"string"`. Computed.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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

- [disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-0220000022211302-3130322210001221-1001020120122002-2120002012101120-3221002020301330-0013330113022100-0310201001112030-0113321321031101): complete subsection reference.

- [disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-1013222032123130-3112101311023032-3302221203000132-2303332110011200-0111211211013323-0212303321030311-3111312301203213-3113213232003102): complete subsection reference.

- [js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312): complete subsection reference.

<a id="canonical-2232322200300300-2012300233103122-1133113222312110-2132000310233200-3233220032223202-3322100303013000-3130332312330322-1002033213311112"></a>

<a id="canonical-2221200303023030-0101111102010111-2110233301232302-2310312300033022-1103311033002323-3222020023122003-3332321002020300-3103103323013312"></a>

## loglevel property — Cloudflare / 210023001213 / 5

Type: `"string"`. Computed.

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

- [manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-0113203221313211-3112313212212333-2320310302212230-2213100321323013-1030120213323013-0302103123230023-2123012113202120-3302231101113232): complete subsection reference.

- [mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231): complete subsection reference.

- [protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122): complete subsection reference.

<a id="canonical-1222221300223120-3113133302202013-1230032212131120-1032233100202033-2220002130020120-2310202303111132-3132023112133010-0303231322120302"></a>

<a id="canonical-0300103102110321-1302323132123120-0322030110000220-3212100221020122-2323132322030220-2232211321012331-2331100201132033-3233120131132130"></a>

## timeout property — Cloudflare / 210023001213 / 6

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

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

- [trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220): complete subsection reference.

<a id="canonical-0213220013021020-1332221130230023-3100311132330213-3202130101301332-2002201121122012-1300031201000013-0122231232111321-3113021011221211"></a>

## Next pages — Cloudflare / 210023001213 / 7

- [cloudflare.disable_js_insert](data-sources--protected_application--reference--group-001.md#canonical-0220000022211302-3130322210001221-1001020120122002-2120002012101120-3221002020301330-0013330113022100-0310201001112030-0113321321031101)
- [cloudflare.disable_mobile_sdk](data-sources--protected_application--reference--group-001.md#canonical-1013222032123130-3112101311023032-3302221203000132-2303332110011200-0111211211013323-0212303321030311-3111312301203213-3113213232003102)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.manual_js_insert](data-sources--protected_application--reference--group-001.md#canonical-0113203221313211-3112313212212333-2320310302212230-2213100321323013-1030120213323013-0302103123230023-2123012113202120-3302231101113232)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [cloudflare.trusted_clients](data-sources--protected_application--reference--group-002.md#canonical-1323233210131003-0221130101233310-2201231231321300-2031303032000332-3020310033013032-2022001001230003-0021202123302120-1001231012222220)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0220000022211302-3130322210001221-1001020120122002-2120002012101120-3221002020301330-0013330113022100-0310201001112030-0113321321031101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221320202212202-0212200202022232-1310123100021033-0100032022210033-2232122102021111-0133303203231103-3200030133323323-0323311113030333"></a>

## Cloudflare.disable_js_insert — disable_js_insert / 332110300230 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.disable_js_insert

<a id="canonical-0123210122333123-1103213201222301-2032023233121210-2020032023011333-0023010021202111-1020213120212110-3131323313113032-3100031113211000"></a>

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

<a id="canonical-1211310102033311-0300112122203313-0302200003000022-1111211133300202-1323102331133011-3200212323130022-0032200222131112-0030132201300020"></a>

## Direct properties — disable_js_insert / 332110300230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032212203223121-1000310210330332-0233112202203322-1133302120330000-0111223331230303-0000212113101023-2322112212201101-1321111000303203"></a>

## Next pages — disable_js_insert / 332110300230 / 4

- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1013222032123130-3112101311023032-3302221203000132-2303332110011200-0111211211013323-0212303321030311-3111312301203213-3113213232003102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000123201212300-3233301322330222-0030013133320232-1110031100022121-3220333030030003-3311132322112211-1001022002222300-2131311231220233"></a>

## Cloudflare.disable_mobile_sdk — disable_mobile_sdk / 331301133110 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.disable_mobile_sdk

<a id="canonical-0120230020032110-1233130200333211-1223221003313221-0311031310222013-1100121003332310-1022000232021103-2010331030323100-2133302322132300"></a>

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

<a id="canonical-0101330123300302-3230031300123320-2012202203333100-1333233333321130-0310031203221110-1211122211301101-0303003330223120-0322113111132323"></a>

## Direct properties — disable_mobile_sdk / 331301133110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233333121313111-2312210120131320-3013110232202131-1010023120133020-3130201313023003-0121022123033320-2330232320310233-0212201321221211"></a>

## Next pages — disable_mobile_sdk / 331301133110 / 4

- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112001120022030-2303112313200221-2203123102113332-0212101010030000-2013121331220011-0211031010133111-1223233332311032-3211313201321120"></a>

## Cloudflare.js_insertion_rules — js_insertion_rules / 002012130011 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.js_insertion_rules

<a id="canonical-3033130212321203-1220212110233133-3213232032233210-0002012100323023-0330312131223120-3101213120202031-3101313333232311-1033000221111203"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0211332120321121-1113111013111330-1033222111003310-2131303133313132-1330131022233020-1022123033320021-3000331013120301-3231203331211232"></a>

## Direct properties — js_insertion_rules / 002012130011 / 3

- [exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333): complete subsection reference.

<a id="canonical-2302323021011231-2012230020020330-0201021233011121-2031003123321222-3232000033131230-3211201131330222-1230102002220012-0302312203113103"></a>

<a id="canonical-3311022202110312-3222032201013002-2011120002001112-3022122123001021-3212101022120232-2220112302312332-2333333022030102-0302332313333131"></a>

## javascript_location property — js_insertion_rules / 002012130011 / 4

Type: `"string"`. Computed.

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

<a id="canonical-1010320122311211-3113322333100310-2321023213022310-2323120232320000-3200032020312012-2103321300322303-1313321123101320-2212200033332223"></a>

<a id="canonical-2100132011312123-3031200132113200-1213220331230212-0211131320103221-3103103222111101-1003322013201302-0201230300132113-3222121010330220"></a>

## js_download_path property — js_insertion_rules / 002012130011 / 5

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/CommonJS’.

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

- [rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321): complete subsection reference.

<a id="canonical-0233113311111300-2003110320130020-3031130021021012-1022010300102003-3013303022123333-0000301102022113-0100123211301001-2233213220330200"></a>

## Next pages — js_insertion_rules / 002012130011 / 6

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133313202032113-3111320301020333-0110123230301000-2131132123032210-0131311313102110-3023002202033310-3123001203300110-0120112213013122"></a>

## Cloudflare.js_insertion_rules.exclude_list — exclude_list / 203311310000 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- Cloudflare.js_insertion_rules.exclude_list

<a id="canonical-2101020020021303-0213230231003300-3110030130303203-2222211322302222-1102311213023321-3203303022110111-1311331201130023-3030222231021203"></a>

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

<a id="canonical-0000231112102021-3010331112200203-0000231000313223-3323122222200233-2110102011120030-2030321201020320-0130220320120300-1333220033330211"></a>

## Direct properties — exclude_list / 203311310000 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-2222000011122321-2122001312332313-3011333231310313-1033300011033300-2323332213312111-3300033203221302-3020131231231212-2133210203120000): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-2310220121300213-1200311010030000-1231323012323013-2012212030113333-2022213331130321-1030121200212223-0022302310233001-1233033223202321): complete subsection reference.

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-2013111323113301-1132302111011323-3301321132131000-2303302211100322-2323311203333111-3323012100130013-2332110222001211-1110022020022023): complete subsection reference.

- [path](data-sources--protected_application--reference--group-001.md#canonical-2122312210233200-3133121131320021-0031122222221203-3020030122321302-3301003222112333-1103132130230311-3333332301020320-2103010310111110): complete subsection reference.

<a id="canonical-0013320331213020-0211020310211011-2010103003131122-2031311200022031-3331211030200302-3201103021210120-0132331130131200-1020213101333323"></a>

## Next pages — exclude_list / 203311310000 / 4

- [cloudflare.js_insertion_rules.exclude_list.any_domain](data-sources--protected_application--reference--group-001.md#canonical-2222000011122321-2122001312332313-3011333231310313-1033300011033300-2323332213312111-3300033203221302-3020131231231212-2133210203120000)
- [cloudflare.js_insertion_rules.exclude_list.domain](data-sources--protected_application--reference--group-001.md#canonical-2310220121300213-1200311010030000-1231323012323013-2012212030113333-2022213331130321-1030121200212223-0022302310233001-1233033223202321)
- [cloudflare.js_insertion_rules.exclude_list.metadata](data-sources--protected_application--reference--group-001.md#canonical-2013111323113301-1132302111011323-3301321132131000-2303302211100322-2323311203333111-3323012100130013-2332110222001211-1110022020022023)
- [cloudflare.js_insertion_rules.exclude_list.path](data-sources--protected_application--reference--group-001.md#canonical-2122312210233200-3133121131320021-0031122222221203-3020030122321302-3301003222112333-1103132130230311-3333332301020320-2103010310111110)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2222000011122321-2122001312332313-3011333231310313-1033300011033300-2323332213312111-3300033203221302-3020131231231212-2133210203120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213132023001303-0033233201212322-0302303032332131-2203033123330102-2313323130021101-3210022113101133-2213110002312100-1210102323032320"></a>

## Cloudflare.js_insertion_rules.exclude_list.any_domain — any_domain / 110221322020 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- Cloudflare.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2111031200022311-3011302133102211-1103311222311203-2213133002212023-0221031312223013-0023101033203100-2022112311121022-0201220223300000"></a>

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

<a id="canonical-0210113013012222-3302212033010321-2131132020000231-2312112001000231-2112323002023220-1121302303310103-0102133322320212-0221232222010100"></a>

## Direct properties — any_domain / 110221322020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321312323303011-3023130113002111-1230112103330012-2221021200233221-0331220100301132-3303220001010310-3021331231112122-0001013332320203"></a>

## Next pages — any_domain / 110221322020 / 4

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2310220121300213-1200311010030000-1231323012323013-2012212030113333-2022213331130321-1030121200212223-0022302310233001-1233033223202321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331113003033120-2310230331030212-2012133313301022-0131332031123112-0023233031330301-0233200031202022-2222120023220302-3000220120320032"></a>

## Cloudflare.js_insertion_rules.exclude_list.domain — domain / 131320302123 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- Cloudflare.js_insertion_rules.exclude_list.domain

<a id="canonical-2123011120102211-2200020120333020-0121310030011112-2212113000103311-0323012311202222-2231323200130000-2221021132201203-2001001103302213"></a>

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

<a id="canonical-3300130010100103-1133220113030100-3023012113233123-2113201122203322-1113210220102320-0021210110223112-1102300030020311-0330332213320331"></a>

## Direct properties — domain / 131320302123 / 3

<a id="canonical-2030111102032020-2022312100012201-0222120121330300-2203022031122303-2300030200320210-3120100112021211-0033110112020031-1233210112123021"></a>

<a id="canonical-1233123023031131-0101113313221203-1231030103321201-1103020010233223-3211110101001233-2231001123111210-0120032101222203-3233000231023303"></a>

## exact_value property — domain / 131320302123 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-2333203101223103-2220000202122133-0303103102111022-3232032232332033-1333222010221230-0030111110020102-0300023210002232-3003323000220121"></a>

<a id="canonical-1321323103301203-1103300133210032-0130233102213133-2312211103330323-2111323223113333-0310202220102122-2213012031232132-1123230313111200"></a>

## regex_value property — domain / 131320302123 / 5

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

<a id="canonical-1310223000000223-3330001200311103-2201232222113322-2310102010222233-1301231303120110-0031232131301132-0332202231232322-2022333333133100"></a>

<a id="canonical-3033101023032220-2131323032101211-3211213132132010-1220100103222230-2023031201011131-2031100023030331-1122321213201200-1203201231131201"></a>

## suffix_value property — domain / 131320302123 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-1020031113200230-2210102203023333-3203220332112011-0203333111303132-0320020123210330-1031121130213303-3102112102121221-1222213312013031"></a>

## Next pages — domain / 131320302123 / 7

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2013111323113301-1132302111011323-3301321132131000-2303302211100322-2323311203333111-3323012100130013-2332110222001211-1110022020022023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112303230110101-3001010112333303-2011030220323120-2303323323201221-0021210132331212-1333112313321003-0203001032232023-0320201220120002"></a>

## Cloudflare.js_insertion_rules.exclude_list.metadata — metadata / 302133310210 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- Cloudflare.js_insertion_rules.exclude_list.metadata

<a id="canonical-1003023133202221-0220011220231210-1202333023311312-1131321102311030-1222012101232001-1320302213022012-1001013200003301-3223231003113301"></a>

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

<a id="canonical-1332103230001312-3231102003112110-2303303230203323-2133211003312002-1230220031002213-0230221221323310-2002020102302002-3212301012110330"></a>

## Direct properties — metadata / 302133310210 / 3

<a id="canonical-1311002112210132-1110231111100031-2102100001033010-0031132032210131-3101310332002033-0123220300033301-3312023033102020-1033233202120322"></a>

<a id="canonical-0332221123032212-0322012311231103-2331132111022301-1022333323012322-1002023103333313-2031200112133020-3012300321132020-2223132032212220"></a>

## description_spec property — metadata / 302133310210 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3213110102122312-1201012132221011-0131110102232300-1203323321001033-1302103213233021-0302310301310311-3321213011133000-3212102000130220"></a>

<a id="canonical-3211022023210013-2222132302102010-3102313032333311-2010131122201210-3331000213301221-2112102011001330-0300300201100311-1330233102211302"></a>

## name property — metadata / 302133310210 / 5

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

<a id="canonical-1002132122012033-2212302121132203-2003312200010133-0102201232130330-3222321232100101-2010013023201022-2000233011113302-3312023031203201"></a>

## Next pages — metadata / 302133310210 / 6

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2122312210233200-3133121131320021-0031122222221203-3020030122321302-3301003222112333-1103132130230311-3333332301020320-2103010310111110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230210110201100-2322310003112313-3321330001102322-1333031122033302-2331111100220332-0312012233302232-3033332213212330-3112101231023022"></a>

## Cloudflare.js_insertion_rules.exclude_list.path — path / 021230212123 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- Cloudflare.js_insertion_rules.exclude_list.path

<a id="canonical-1103223013202110-0202202100021130-2031002123120321-1011211030033312-2321023232322132-2112033233130210-3100303211223102-1012033111210100"></a>

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

<a id="canonical-1301202303120210-2102332320033313-3033010011202132-3312010103123120-0003113310332300-0333103032220331-3020113323302023-2132120220323122"></a>

## Direct properties — path / 021230212123 / 3

<a id="canonical-0100220122023033-0010223103130022-1232121011101121-0310033231333213-2003103110233112-0013200100203222-2303203312000032-0002231213203130"></a>

<a id="canonical-0132233012320230-1322103200121200-0000220201201220-2131122032110101-2321133203211123-3023321300322112-0230022002031310-2211021100121011"></a>

## path property — path / 021230212123 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

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

<a id="canonical-0313102301032213-0120132111300300-1231100131033331-0002311212001111-2312322120013020-0211303203001222-1312220213300233-0100121223110021"></a>

<a id="canonical-2303323232200302-2020132210102132-0301002320121012-2302311101321313-3102101223003102-0310300031000312-0201100330103221-2113113110023300"></a>

## prefix property — path / 021230212123 / 5

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

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

<a id="canonical-1300311323310003-0010113101101132-2303001312220100-3322000323231002-2223232231123001-0030032033301322-0023100302031303-2002331101232322"></a>

<a id="canonical-0132111232131110-2231302131121333-3113302320222312-3330300301021313-2130212112200220-1001033322021112-2230330102022230-0101001031030323"></a>

## regular expression property — path / 021230212123 / 6

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

<a id="canonical-3200111221002031-3121221331102002-3312230122213330-1323301231022231-0023011321011002-2123102103130201-0222022010312201-0121000000033022"></a>

## Next pages — path / 021230212123 / 7

- [cloudflare.js_insertion_rules.exclude_list](data-sources--protected_application--reference--group-001.md#canonical-2001333113133333-0320013031030321-1112023330212213-0213333230230002-0311001020210222-2020121302003102-2223232311010133-3100201020232333)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021222002222001-1203212211100230-2023013231122333-3113022212200211-3013033001010310-3110202331133120-2321003020113311-3022110233321130"></a>

## Cloudflare.js_insertion_rules.rules — rules / 122210130100 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- Cloudflare.js_insertion_rules.rules

<a id="canonical-0301100233133322-0323212333103212-3023210333331133-1132310322210320-3022123232111223-0333301320000032-2333231033011132-2013232220020013"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

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

<a id="canonical-2312302102332132-0211221120123101-3331130112313113-1313211002233121-1231023102310211-3011313302100110-1123311103200220-0302031322300222"></a>

## Direct properties — rules / 122210130100 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-3000120220231021-3203313003010010-0010001122320210-1310111102301303-1122213020332310-0211200230221100-0330003331001223-3131030012103330): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-1332112120332003-3233312030032232-0131121203222333-3312020123010303-2130233130103031-3020331232131002-3313123113210003-0010200131013010): complete subsection reference.

<a id="canonical-2223310001121233-2012221213002202-0302000021201111-2203023132232010-0020012331131033-3223211310222023-3330131320213022-2130121123332231"></a>

<a id="canonical-2001313102333230-1223022101212000-1030301200222303-1331332023233121-0203100010010130-3033311003233303-3010100033213310-0003300222003333"></a>

## exact_path property — rules / 122210130100 / 4

Type: `"string"`. Computed.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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

<a id="canonical-1013211320002131-3022210331330311-2210201011203032-0121021102331011-2021230332322103-0001021013230111-1212313310310232-0023111013301311"></a>

<a id="canonical-3111311303001001-1023130033323031-0222321333232230-2132201333200221-2231131211101100-2331012233022022-1213223231310200-0322223300021333"></a>

## glob property — rules / 122210130100 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

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

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-2231320222323321-1022223122333132-3002331131121100-3001201320322221-3100002320102333-2321322012102121-3331100012212121-3331123012120301): complete subsection reference.

<a id="canonical-3123010210031302-3313331033313013-2122222112100033-2320213211122331-2001011003033020-2100102120133322-2232010232111002-2101202121013102"></a>

<a id="canonical-3002111333002122-2330220012310010-0322303022030220-1323001110213112-3331331323212222-0030221332130320-2223303011323020-1012132211203201"></a>

## prefix property — rules / 122210130100 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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

<a id="canonical-3201220020113303-3121002311131130-1033002322120303-2221302331120310-2312331101220300-2031011002002331-1133113000222023-0332010221020032"></a>

## Next pages — rules / 122210130100 / 7

- [cloudflare.js_insertion_rules.rules.any_domain](data-sources--protected_application--reference--group-001.md#canonical-3000120220231021-3203313003010010-0010001122320210-1310111102301303-1122213020332310-0211200230221100-0330003331001223-3131030012103330)
- [cloudflare.js_insertion_rules.rules.domain](data-sources--protected_application--reference--group-001.md#canonical-1332112120332003-3233312030032232-0131121203222333-3312020123010303-2130233130103031-3020331232131002-3313123113210003-0010200131013010)
- [cloudflare.js_insertion_rules.rules.metadata](data-sources--protected_application--reference--group-001.md#canonical-2231320222323321-1022223122333132-3002331131121100-3001201320322221-3100002320102333-2321322012102121-3331100012212121-3331123012120301)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3000120220231021-3203313003010010-0010001122320210-1310111102301303-1122213020332310-0211200230221100-0330003331001223-3131030012103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011020023022221-3102133020323213-1003133020100223-3302032131030212-1202021232213022-2110122031202123-2110031232111230-1011133010010231"></a>

## Cloudflare.js_insertion_rules.rules.any_domain — any_domain / 312012213113 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- Cloudflare.js_insertion_rules.rules.any_domain

<a id="canonical-2033023231011110-1332213221203103-3313330330113202-3200010313201210-3022133300121100-2210232302323233-2232101013201031-2013131030213320"></a>

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

<a id="canonical-3332112300303210-0002323122101103-2330320310331323-3211112321332021-0133121332000122-3210323201230312-2322023102110011-3130312201313013"></a>

## Direct properties — any_domain / 312012213113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032311022120320-1130123021232031-0302330030201221-2013030123232330-1122000303001221-3312203023133111-1310310310301003-3333212122321001"></a>

## Next pages — any_domain / 312012213113 / 4

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1332112120332003-3233312030032232-0131121203222333-3312020123010303-2130233130103031-3020331232131002-3313123113210003-0010200131013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222332221223310-2102313222100221-1011220312101000-1111233100201022-3230103013331300-0010231210002303-2132101203213012-1202323300012130"></a>

## Cloudflare.js_insertion_rules.rules.domain — domain / 102132333001 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- Cloudflare.js_insertion_rules.rules.domain

<a id="canonical-0112311321131311-0201101101131033-0331302222031032-2020303230320011-3330130012033310-3030000030231001-1333233330200202-2223203032101321"></a>

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

<a id="canonical-1211112102011201-2003000332101011-0203121202333321-3222312100001322-3021110300210210-2330001231303000-3312312001220331-0321131202201210"></a>

## Direct properties — domain / 102132333001 / 3

<a id="canonical-2020221223131012-2131033123232102-3230223232033302-3310321013100203-2302202221312310-0201333022110210-3332000012110130-3120201010300301"></a>

<a id="canonical-1331310222232123-1130220231011330-0102332020112103-3320303113132311-1203000303320031-0323322211032322-2013300111103220-1120323122001013"></a>

## exact_value property — domain / 102132333001 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-2011002330003321-3031033001111333-2100301023331330-2322032121301120-0200013200112131-0313010230223102-0030032321202002-0220201120122100"></a>

<a id="canonical-2003111202111330-0001110032120321-0223320302223211-1032212111311002-0021032110013011-3232212102303310-2110110201123222-1030100020010101"></a>

## regex_value property — domain / 102132333001 / 5

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

<a id="canonical-2313212302023333-3322000231031333-3111003322131103-2332131131033110-0131120133213031-1210211102310200-2032002123013103-0332033102012001"></a>

<a id="canonical-2230022311222212-3302313112201021-2033320100233220-2221322020331303-2113010312101113-2331331230033131-0131132332201022-3213001202111002"></a>

## suffix_value property — domain / 102132333001 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-3330133302002332-0212021100031032-3013110313022231-2231021321010021-2123323210332100-1321021210021303-1200231211312032-3022322311011000"></a>

## Next pages — domain / 102132333001 / 7

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2231320222323321-1022223122333132-3002331131121100-3001201320322221-3100002320102333-2321322012102121-3331100012212121-3331123012120301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322100231003032-2113123232301222-1231201330111302-2232020311211321-1311202332233102-2220231110112001-0312211310010330-0132320032203122"></a>

## Cloudflare.js_insertion_rules.rules.metadata — metadata / 310131032203 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.js_insertion_rules](data-sources--protected_application--reference--group-001.md#canonical-1231311222220033-0111003303023300-1233020100031131-2212023302233222-3023233331031212-1332321201301030-1122123222031303-2312333321323312)
- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- Cloudflare.js_insertion_rules.rules.metadata

<a id="canonical-1231322113013122-1221103321130330-1003223213221302-3002103232310111-3111232331311123-2120210121021210-3131212120100300-0201301320033120"></a>

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

<a id="canonical-2102233100213103-1002231130102103-1031213303131200-3333303130213013-0003022111313211-0020302320303001-3013200103012003-3130211201203231"></a>

## Direct properties — metadata / 310131032203 / 3

<a id="canonical-1232201330233211-0122102010311001-0111122021231333-1101133020302002-1303023230321223-1300030333210003-3120130023221113-3000200101131322"></a>

<a id="canonical-1313330011201303-0033113102133010-3332210232330200-3011211133313320-2002312022303033-3200031032231322-0010321122010020-0303233122103201"></a>

## description_spec property — metadata / 310131032203 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1230322312130000-3122000310211211-3312022332111110-0230031221330303-2321121221111031-2223121133303021-0132133313100103-3232121011131231"></a>

<a id="canonical-3222002233303200-3111110201000113-3220303222012210-0303120312123120-3333222113101301-3113133321101022-0221203102223100-2131233303121233"></a>

## name property — metadata / 310131032203 / 5

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

<a id="canonical-1032020131331002-3000302100012121-2101032203033321-3121230021231010-3022331310100303-3101120301023301-1022233233222013-3003113213313131"></a>

## Next pages — metadata / 310131032203 / 6

- [cloudflare.js_insertion_rules.rules](data-sources--protected_application--reference--group-001.md#canonical-1100302020232020-1033003010233212-2232230131213103-3313031222210202-0112312323121002-0021221200113311-0221111033100321-0130232033320321)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0113203221313211-3112313212212333-2320310302212230-2213100321323013-1030120213323013-0302103123230023-2123012113202120-3302231101113232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013110332323122-1213212212233330-2311310323303231-1301131321033031-1310303021013001-1303031220312023-2120232322101331-0110103102123203"></a>

## Cloudflare.manual_js_insert — manual_js_insert / 220002111323 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.manual_js_insert

<a id="canonical-3010110321202310-3102122010021121-1212203013132010-2011313220101131-2102122332220000-2330201223031323-2112031322331300-3133301313321010"></a>

Type: `"single"`. Computed.

Insert JavaScript Manually. Insert JavaScript manually.

Upstream description:

Insert JavaScript manually.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203303032030022-1113313121013230-1121302020010211-2313220312121203-3301111231310231-2321110231131231-0001232100330001-3210210202211031"></a>

## Direct properties — manual_js_insert / 220002111323 / 3

<a id="canonical-3233311301103321-1310323121001010-2011230033213021-3010101003311232-1320122212032103-0212201332103000-1321002302222020-2223130123311331"></a>

<a id="canonical-0012302133102300-2312110320030023-1300102132201110-2213002313112133-0233211011131202-1313110302032222-0112311033301033-3110301100000033"></a>

## js_download_path property — manual_js_insert / 220002111323 / 4

Type: `"string"`. Computed.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/CommonJS’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/CommonJS’.

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

<a id="canonical-3331321113110111-0210211231202030-2331311132230202-0131210331232323-0221130232300233-3031132122211202-2133023033311322-1011332100023113"></a>

## Next pages — manual_js_insert / 220002111323 / 5

- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201012000002321-0233102221113322-3003210113113133-0323320130032121-1332213213020110-0332103021322321-2211230231322132-0333230321032020"></a>

## Cloudflare.mobile_sdk_config — mobile_sdk_config / 113113112320 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.mobile_sdk_config

<a id="canonical-1312323003331123-2221012232232111-1020102330032132-0011220123020302-0030122103020233-2310331102131012-3303123321301231-3332300103022002"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010103301313311-0303030212003312-0330320133231213-3132022013010130-1001110001122203-1032223111321001-0312001230221222-0311030322233102"></a>

## Direct properties — mobile_sdk_config / 113113112320 / 3

- [mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-2221122320202302-0022131332112213-3111230031112001-2033031013133132-2021023300122320-1310310323331002-3021330103100310-0102013130221221): complete subsection reference.

<a id="canonical-2120212303303030-1030023233002003-1303210231230001-1122210112333020-2030011011021131-3321011013203330-0003222312003011-0010112120210121"></a>

## Next pages — mobile_sdk_config / 113113112320 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-2221122320202302-0022131332112213-3111230031112001-2033031013133132-2021023300122320-1310310323331002-3021330103100310-0102013130221221)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2221122320202302-0022131332112213-3111230031112001-2033031013133132-2021023300122320-1310310323331002-3021330103100310-0102013130221221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101122111222031-1201221102133221-2132132032031233-0333103303111211-0103203013132112-0010300010030100-1120100233123333-0210313202131220"></a>

## Cloudflare.mobile_sdk_config.mobile_identifier — mobile_identifier / 010233232313 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231)
- Cloudflare.mobile_sdk_config.mobile_identifier

<a id="canonical-0030003221330113-0312012032323011-1003200130310031-2233331123331232-1023121232312320-1231022101310303-2110303113131203-1322221313032212"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132332200203100-2220233112112111-2203101123033023-3123202322203330-3103001132301000-0123323021201032-3111031110010321-3002332101323002"></a>

## Direct properties — mobile_identifier / 010233232313 / 3

- [headers](data-sources--protected_application--reference--group-001.md#canonical-3332133303310002-0322121020113212-3120112003121132-2222220121133133-3131120322130313-2030333212100110-2000002213022121-1220101202122220): complete subsection reference.

<a id="canonical-0102313202120210-3320223023023203-3300003212213122-1332022321222321-1321213132110102-1013201120113331-1202130132300310-3211330301230030"></a>

## Next pages — mobile_identifier / 010233232313 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier.headers](data-sources--protected_application--reference--group-001.md#canonical-3332133303310002-0322121020113212-3120112003121132-2222220121133133-3131120322130313-2030333212100110-2000002213022121-1220101202122220)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-3332133303310002-0322121020113212-3120112003121132-2222220121133133-3131120322130313-2030333212100110-2000002213022121-1220101202122220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312003010022010-2321103223201223-2223231111121223-2112301210223022-0312330311231131-1322313302221310-2002213013110100-3300131130121320"></a>

## Cloudflare.mobile_sdk_config.mobile_identifier.headers — headers / 212111103330 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.mobile_sdk_config](data-sources--protected_application--reference--group-001.md#canonical-1012330233031112-2222321301033031-3302233131022100-2332032300003202-2320003111223303-0031323032112130-2020330111011333-2103332111101231)
- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-2221122320202302-0022131332112213-3111230031112001-2033031013133132-2021023300122320-1310310323331002-3021330103100310-0102013130221221)
- Cloudflare.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1203330130313101-3012123102323230-3133223322013232-2021310000333211-0203200233202033-0121312111332313-2012101200020001-0322203010013033"></a>

Type: `"list"`. Computed.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

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

<a id="canonical-0112030110101033-0333323232133313-3200220310231133-2200230312233323-2013122201232032-1302221321002322-0200133322301003-3033323333120122"></a>

## Direct properties — headers / 212111103330 / 3

<a id="canonical-1000032003130210-3013031231031133-3131113020220301-1030103222011102-2212310300023130-0012013231211002-1210332310200012-1022103332320312"></a>

<a id="canonical-1012322331120212-1023010131132221-2001033131212021-2332131000001113-1133321010101010-3301320312311002-3030200311213231-0112030200003103"></a>

## exact property — headers / 212111103330 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\] Header value to match exactly.

Upstream description:

Exclusive with \[regular expression\] Header value to match exactly.

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

<a id="canonical-0220002132031232-0022233220202113-2201301310010020-1312311233022323-2032000222322101-2312020300103320-3202013233032211-1310023032100232"></a>

<a id="canonical-2212213132302133-1020230122222011-1132212300121131-2020223203331222-3203023133003010-2101001210132011-0011123331013200-3123011222233300"></a>

## name property — headers / 212111103330 / 5

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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

<a id="canonical-0201113000303101-2311121022222213-0311120101300232-2032110212213212-3032301022032312-2120212132000333-3111030302103321-1212232031122303"></a>

<a id="canonical-2310121301010103-3021100223003200-1101333321320210-3320201332000021-1130113111303033-1221022121001023-2211323000301303-2332113032022313"></a>

## regular expression property — headers / 212111103330 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\] regular expression match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] regular expression match of the header value in re2 format.

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

<a id="canonical-3030022110330223-1303132131333023-0011002020301233-3233211330332131-3123311011210100-2232322322023011-3202231031233133-2330103002130313"></a>

## Next pages — headers / 212111103330 / 7

- [cloudflare.mobile_sdk_config.mobile_identifier](data-sources--protected_application--reference--group-001.md#canonical-2221122320202302-0022131332112213-3111230031112001-2033031013133132-2021023300122320-1310310323331002-3021330103100310-0102013130221221)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232301233300330-2120210232213320-0310232023222103-0030013201203003-2023101033111333-2320212211022201-3112210320222132-0010333300102322"></a>

## Cloudflare.protected_endpoints — protected_endpoints / 032211001110 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- Cloudflare.protected_endpoints

<a id="canonical-2300121303010201-0332332210221112-2333212312230100-3011131213021130-2010333220313020-1111221332112102-1021331120331300-0123033012021000"></a>

Type: `"list"`. Computed.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

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

<a id="canonical-0212300121233330-0322000121020211-1000201312202123-3023012212322033-0322222331303231-2020033003100131-2331012212201120-1002200221121212"></a>

## Direct properties — protected_endpoints / 032211001110 / 3

- [any_domain](data-sources--protected_application--reference--group-001.md#canonical-2000001131223120-0013112331203310-2110320202001101-0212131123213002-1012213233001130-3012220233001032-2030222303230110-2233132012021331): complete subsection reference.

- [domain](data-sources--protected_application--reference--group-001.md#canonical-0033332310100122-0322200200230013-2030323113223301-0021302301302103-3020223102303100-2213213210232323-0212111121130312-2123020012012311): complete subsection reference.

<a id="canonical-1300001221100003-3322211312030311-2220311233330010-1200011012013221-0113230030113232-1222312033313101-3301232011033223-0011123313022120"></a>

<a id="canonical-2232001132331020-3231320231101331-2110023112301322-0030302112131103-3303331310331021-2031013330210002-3210312301102110-1310303321311311"></a>

## http_methods property — protected_endpoints / 032211001110 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--protected_application--reference--group-001.md#canonical-2131322001210311-0202030113003201-2222112121200211-3230202002032221-0023301332002232-1133303113201130-0113011221002032-3030101301131033): complete subsection reference.

- [mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223): complete subsection reference.

- [path](data-sources--protected_application--reference--group-002.md#canonical-3232201300123231-0223201311223113-1102002220010310-1220002021030023-0202033222102033-0032122023100220-2313102223310321-3310323301001221): complete subsection reference.

<a id="canonical-0001132201102300-1200112033001221-3000020233021113-0332220233202032-2000330103330031-0212020121031130-2010323102322323-3203012301101131"></a>

<a id="canonical-0323110003121030-2223022113232022-0003033001231133-1301112012000000-3023120111013222-1301013001220200-3130211323313233-0123200233310223"></a>

## query property — protected_endpoints / 032211001110 / 5

Type: `"string"`. Computed.

Enter a regular expression to match your query parameters of interest.

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
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311): complete subsection reference.

- [web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100): complete subsection reference.

<a id="canonical-2011312111113303-1033131032003320-0320210111222122-1122301303232200-2102331222131123-2102133031223221-3232333332302211-2103221000022133"></a>

## Next pages — protected_endpoints / 032211001110 / 6

- [cloudflare.protected_endpoints.any_domain](data-sources--protected_application--reference--group-001.md#canonical-2000001131223120-0013112331203310-2110320202001101-0212131123213002-1012213233001130-3012220233001032-2030222303230110-2233132012021331)
- [cloudflare.protected_endpoints.domain](data-sources--protected_application--reference--group-001.md#canonical-0033332310100122-0322200200230013-2030323113223301-0021302301302103-3020223102303100-2213213210232323-0212111121130312-2123020012012311)
- [cloudflare.protected_endpoints.metadata](data-sources--protected_application--reference--group-001.md#canonical-2131322001210311-0202030113003201-2222112121200211-3230202002032221-0023301332002232-1133303113201130-0113011221002032-3030101301131033)
- [cloudflare.protected_endpoints.mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2030023030101010-1310113333033102-1013023003332100-3233211203312330-0033001100311212-2030010200302302-1023220213200122-0223130313132223)
- [cloudflare.protected_endpoints.path](data-sources--protected_application--reference--group-002.md#canonical-3232201300123231-0223201311223113-1102002220010310-1220002021030023-0202033222102033-0032122023100220-2313102223310321-3310323301001221)
- [cloudflare.protected_endpoints.web_client](data-sources--protected_application--reference--group-002.md#canonical-3230221231323110-2000013213332023-3011112223321020-1023203303123102-3000330332112310-3132023301231112-0121133032331033-1331200011200311)
- [cloudflare.protected_endpoints.web_mobile_client](data-sources--protected_application--reference--group-002.md#canonical-2200323013101112-0010032321332011-2012100323232032-2200330010131301-2312033323311132-1010333032003110-0321301301131112-0212220022020100)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2000001131223120-0013112331203310-2110320202001101-0212131123213002-1012213233001130-3012220233001032-2030222303230110-2233132012021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300232001320012-3323121321333122-1302120003302012-1221300221020013-3201020120313303-2010023111220112-1200100323120132-3231323121032122"></a>

## Cloudflare.protected_endpoints.any_domain — any_domain / 312020110323 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.any_domain

<a id="canonical-3223223133330120-3332331232213112-1101313310221321-1301313100233103-1203332212123212-3033200232101331-0203303332003331-1133233203201003"></a>

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

<a id="canonical-3132103003103231-1010221112003113-1101203031233133-3230103330011333-3320021010302311-0032220023221202-3332101231132110-2202020332110011"></a>

## Direct properties — any_domain / 312020110323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103202220103010-1103301003312123-0213213212130231-2020101302213021-0002111222222012-1230312313201311-0210001002132030-2233112331232120"></a>

## Next pages — any_domain / 312020110323 / 4

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-0033332310100122-0322200200230013-2030323113223301-0021302301302103-3020223102303100-2213213210232323-0212111121130312-2123020012012311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212123100221220-3302202321011033-2133001011322230-3002213123202221-2303023112233123-2123100032203130-2123302230233020-3113333100320220"></a>

## Cloudflare.protected_endpoints.domain — domain / 022120233313 / 2

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)
- [Property reference](data-sources--protected_application--reference--group-001.md#canonical-0001333200233223-0322212201111301-1102320330030121-1213300312333233-0100033112300113-2331020111321030-1332233011232311-1112113101312102)
- [Cloudflare](data-sources--protected_application--reference--group-001.md#canonical-2103030102303220-3032000233032113-1311312033011021-0203031102122023-0312101000122132-2100301000121311-2223211222233310-0201001111112331)
- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- Cloudflare.protected_endpoints.domain

<a id="canonical-0223212010233220-0012100021130123-1213312331203303-1022332002213030-0103310322022021-3030102113333200-3213133303210031-3232231320000330"></a>

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

<a id="canonical-3130112323301201-0300123032031321-3330320323213301-0301211232100202-0103100032002321-0332021312100220-2221133302223213-3002231131033123"></a>

## Direct properties — domain / 022120233313 / 3

<a id="canonical-2011011323331210-0032113323110030-1110332222020011-0330032201211310-2111111033103001-1232232303112120-3001200022311230-1200122110212132"></a>

<a id="canonical-1012313300003331-2333111202322201-3101303110133230-2002102203033121-1102310003131332-3223100120122103-3030003121312200-2223302221131103"></a>

## exact_value property — domain / 022120233313 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

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

<a id="canonical-2112333002330001-0233013001223330-3013331113012032-2122322013312031-0321131021311032-0203123321233230-3310330231231202-3222300010312221"></a>

<a id="canonical-0001312003321230-2202310222211122-1001012121332012-1001210323130231-2031121202332231-2011100032100132-0010322023211321-2230132100101003"></a>

## regex_value property — domain / 022120233313 / 5

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

<a id="canonical-0033133302321322-3133320021003021-3012132123202123-0003100120300010-1110021121121323-0203033103013113-0232001221113231-0210103132120000"></a>

<a id="canonical-3232102203223132-2010003233003212-3323001032200111-0332333230202123-3212131031011000-3010111120121013-3131131020223301-3213002023311102"></a>

## suffix_value property — domain / 022120233313 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

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

<a id="canonical-1303310013302232-2221132302313330-3333223101211312-0032233303010100-0031120023303213-2022312330122310-3321213113002131-2310001321003332"></a>

## Next pages — domain / 022120233313 / 7

- [cloudflare.protected_endpoints](data-sources--protected_application--reference--group-001.md#canonical-1030220202111001-1103110111301100-1322102231233131-0123103313030303-2233000102020131-3103322313321020-0122132330233112-1300111122001122)
- [xcsh_protected_application](../data-sources/protected_application.md#canonical-0321101133212112-2213200303020213-3012101103330221-0200320313210231-2310313032122133-1332031112223203-2302102031333233-2230021122320022)

<a id="canonical-2131322001210311-0202030113003201-2222112121200211-3230202002032221-0023301332002232-1133303113201130-0113011221002032-3030101301131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
