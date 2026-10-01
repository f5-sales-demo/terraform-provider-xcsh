---
page_title: "xcsh_alert_policy reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy reference."
---

# xcsh_alert_policy reference

<a id="canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021030101222023-2202100223330210-3001132020221112-2101330312022001-3210330323113233-3212231111322100-0303331120132210-1020023000222303"></a>

## Property reference — Property reference / 031021233130 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- Property reference

<a id="canonical-0321221031032022-1311313013022113-2003232311031010-0233023323021303-2133322333011131-2311112221212300-1300130230130031-0003303123202312"></a>

## Direct properties — Property reference / 031021233130 / 3

<a id="canonical-2321232223211231-3000102330301032-1300013101233212-2312132030222310-2113032031022122-3121221311013233-0033132200022232-1102201112130122"></a>

<a id="canonical-0111220123031221-3222110232203223-1021303130330020-2113300103130121-1200002300020200-2210303113102211-3110110321123322-1230313132000123"></a>

## annotations property — Property reference / 031021233130 / 4

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

<a id="canonical-0220122302333010-3103320113200210-3001002003120130-3100232313310002-1012110212131023-1223021233133113-2232132103031222-0331202222133213"></a>

<a id="canonical-3122101002101020-0212201020220223-2022000303311100-3220300211011222-2222122200222333-1332201100121222-1231100222222202-3321013000331133"></a>

## description property — Property reference / 031021233130 / 5

Type: `"string"`. Computed.

Description of the AlertPolicy.

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

<a id="canonical-3132322230013332-3201102122303223-1322030212330122-2001031302013133-0223223120010122-1020212220110012-1331022103232100-1112231130002121"></a>

<a id="canonical-3020003130131022-2001333203300101-0233121111103121-0032102113330000-2330012322321021-3002010301322120-3221102100102110-3133311023033021"></a>

## ID property — Property reference / 031021233130 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3022000010122133-2030232232202020-2301010102112303-1231323223130023-0311310232032013-1212001201100121-0011311330231013-3321112222312323"></a>

<a id="canonical-2200323320313311-1210232303032132-2221332210322130-1110130021123001-0010011132110120-0223130013320120-3301232112121210-2102111013212233"></a>

## labels property — Property reference / 031021233130 / 7

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

<a id="canonical-0220222121111320-2323303132201203-1331003303110300-3333221223312332-0200213133331333-0311031231120000-3132222113111122-2122221020113110"></a>

<a id="canonical-2131031123213133-1331102110033121-2003110213311013-1312112323220031-1021331023311202-3003301311120123-2130310320213332-3233323331233132"></a>

## name property — Property reference / 031021233130 / 8

Type: `"string"`. Required.

Name of the AlertPolicy.

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

<a id="canonical-2023031003203022-0130111220133030-3222311311012101-0112222001132120-1003023302213123-2202030122123133-0200131310223333-3030213023030212"></a>

<a id="canonical-2031203221211321-2301033221311010-2030221212123320-1221200011321333-0021112321332022-3312203320032220-1331132032333200-3021201210021231"></a>

## namespace property — Property reference / 031021233130 / 9

Type: `"string"`. Required.

Namespace where the AlertPolicy exists.

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

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200): complete subsection reference.

- [receivers](data-sources--alert_policy--reference--group-001.md#canonical-3101011111113203-3320322310232130-1320032032033121-1221133020000113-1021233103201030-0221021130100132-1022002300200000-1102333030221331): complete subsection reference.

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202): complete subsection reference.

<a id="canonical-3232131010023103-3323333122223323-1012011122300311-0302100322322302-3312220303331303-0021302000232123-3111322020213200-3321333312210111"></a>

## All schema paths — Property reference / 031021233130 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--alert_policy--reference--group-001.md#canonical-2321232223211231-3000102330301032-1300013101233212-2312132030222310-2113032031022122-3121221311013233-0033132200022232-1102201112130122) |
| `description` | [description](data-sources--alert_policy--reference--group-001.md#canonical-0220122302333010-3103320113200210-3001002003120130-3100232313310002-1012110212131023-1223021233133113-2232132103031222-0331202222133213) |
| `id` | [id](data-sources--alert_policy--reference--group-001.md#canonical-3132322230013332-3201102122303223-1322030212330122-2001031302013133-0223223120010122-1020212220110012-1331022103232100-1112231130002121) |
| `labels` | [labels](data-sources--alert_policy--reference--group-001.md#canonical-3022000010122133-2030232232202020-2301010102112303-1231323223130023-0311310232032013-1212001201100121-0011311330231013-3321112222312323) |
| `name` | [name](data-sources--alert_policy--reference--group-001.md#canonical-0220222121111320-2323303132201203-1331003303110300-3333221223312332-0200213133331333-0311031231120000-3132222113111122-2122221020113110) |
| `namespace` | [namespace](data-sources--alert_policy--reference--group-001.md#canonical-2023031003203022-0130111220133030-3222311311012101-0112222001132120-1003023302213123-2202030122123133-0200131310223333-3030213023030212) |
| `notification_parameters` | [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3311331223111121-3201001311000211-1103202321032032-1112220303220133-0023330222201023-0312132021320321-1311323123020213-2000112032123131) |
| `notification_parameters.custom` | [notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-3230212211103322-2232201330030103-0203220031323203-0100010013022321-1133320101321333-1221010021132221-2030323333322332-1033102132300030) |
| `notification_parameters.custom.labels` | [notification_parameters.custom.labels](data-sources--alert_policy--reference--group-001.md#canonical-0310132202000001-0300301312011033-2021200331232103-2323213022311100-1122111220232022-0002223011030121-2121131230013022-0000113220001123) |
| `notification_parameters.default` | [notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-2023103200222212-3123222100101012-3030020021233310-1323200312001013-0031221031111000-2113231231022313-1313213332331021-3132310012001101) |
| `notification_parameters.group_interval` | [notification_parameters.group_interval](data-sources--alert_policy--reference--group-001.md#canonical-3113231230230322-2123200030030021-3321032002132112-3213202201032320-1310333000313130-1011221213030200-0101232313213302-0022312021332310) |
| `notification_parameters.group_wait` | [notification_parameters.group_wait](data-sources--alert_policy--reference--group-001.md#canonical-0300233112301300-2123020011322210-2003331103331033-3331230113002222-0001212011232100-3010032102030010-3012231002220031-1312103000231020) |
| `notification_parameters.individual` | [notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-0321010112112321-1332223023001311-0301130223232302-3323220333210333-1031212303201222-3230000010233100-2112202110031313-2220302303211233) |
| `notification_parameters.repeat_interval` | [notification_parameters.repeat_interval](data-sources--alert_policy--reference--group-001.md#canonical-3102132323202022-3321313000331200-0011121312202030-3232111320333330-0322233300213101-1303020310301230-1020113201231222-2101310102101233) |
| `notification_parameters.ves_io_group` | [notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-0231012113013021-3100203200012122-2010133021121211-0013222013030300-2213220203232023-0102132020310120-3010013013021210-0300131230333302) |
| `receivers` | [receivers](data-sources--alert_policy--reference--group-001.md#canonical-0131030030013223-1302101022213233-1320310131320031-0332112203223102-3103013311003033-2112323331103130-1303012130321213-2320132322113003) |
| `receivers.kind` | [receivers.kind](data-sources--alert_policy--reference--group-001.md#canonical-0002010300223012-3331102033322000-2000211110223203-2031031232203000-0001013010200033-3310031310012103-3322300202013323-3122110220113023) |
| `receivers.name` | [receivers.name](data-sources--alert_policy--reference--group-001.md#canonical-3022223111321033-2232032120201001-3012013013131002-3101030131000030-1230113101321131-1103120011120201-1303100110102333-0120021021013003) |
| `receivers.namespace` | [receivers.namespace](data-sources--alert_policy--reference--group-001.md#canonical-2300221220032301-2030013103031233-3222123311231022-3010121020330232-2103311022330033-2010011320131010-0101032312301332-3231133110233310) |
| `receivers.tenant` | [receivers.tenant](data-sources--alert_policy--reference--group-001.md#canonical-3311301122320110-2032222010301130-3300133100133130-0112030213111033-0023101023300332-0131130223211110-1013022003303012-1300311013123111) |
| `receivers.uid` | [receivers.uid](data-sources--alert_policy--reference--group-001.md#canonical-3233100333023301-0011201331202122-0330211030121322-2122311323023012-3202031122223230-0333210003001132-1011100232311132-0011020003100032) |
| `routes` | [routes](data-sources--alert_policy--reference--group-001.md#canonical-1102033113313133-3133012120331022-3210303121203103-2323112203132111-0310011032310210-3210101320122323-0320311232330103-2321300301313200) |
| `routes.alertname` | [routes.alertname](data-sources--alert_policy--reference--group-001.md#canonical-0102101012203312-2131012303223211-0323213011202101-2120310303303102-2003021032201121-0312301113300113-2211121312201211-0312131230011012) |
| `routes.alertname_regex` | [routes.alertname_regex](data-sources--alert_policy--reference--group-001.md#canonical-1232002102300130-1113131031212233-0202132101330112-1323213313323213-3302002320103310-2023030301121230-0031023222003011-3111133030011101) |
| `routes.any` | [routes.any](data-sources--alert_policy--reference--group-001.md#canonical-3321203133232201-3101030300021333-1201112331303301-3102312230232222-1321000003330021-3132022032212120-0213013230213133-2310310101232332) |
| `routes.custom` | [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-1100133110101131-3003321221333201-0303103011131003-0212230132100123-3130101110021213-3201010101230001-2310130002221212-1330031112023101) |
| `routes.custom.alertlabel` | [routes.custom.alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-0300320320213020-0201230112033000-3331021103003010-3300201111133231-0222303022022212-0230010311313031-3221320232022131-2100031022031301) |
| `routes.custom.alertname` | [routes.custom.alertname](data-sources--alert_policy--reference--group-001.md#canonical-0211320010210201-1301121311130311-1000010231132322-2000022330021212-2231022101122301-3213112231000232-1330222320222101-2223011201213302) |
| `routes.custom.alertname.exact_match` | [routes.custom.alertname.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-0230003113310122-2333002013113122-0203302223321030-3110332303321132-2100332102020212-1211002022133311-3302131122030111-1030313323120322) |
| `routes.custom.alertname.regex_match` | [routes.custom.alertname.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-2320202230230311-0123222223302222-0323230001103323-1113322202123333-2120111120020112-3110010233323311-0203201103100300-1123033323113231) |
| `routes.custom.group` | [routes.custom.group](data-sources--alert_policy--reference--group-001.md#canonical-1222333231122131-3011102100102133-3021212120101302-3013213303101303-2233231301222312-3320221310123230-1021022121202313-1223133100010033) |
| `routes.custom.group.exact_match` | [routes.custom.group.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-0001333013300020-0310122312201131-2330213022122013-3202120201110222-0312333000200232-0310131033203302-0313202030032011-1132321333303002) |
| `routes.custom.group.regex_match` | [routes.custom.group.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-2232200020132213-1002220112003230-0221300223211213-1113130130021031-3223313000331300-1210202322133230-0231000221232130-0123322112302012) |
| `routes.custom.severity` | [routes.custom.severity](data-sources--alert_policy--reference--group-001.md#canonical-0300013320001122-3323123022101302-1021220302111321-2103233221302030-0310213331020212-3100033011220000-3211112133133102-2102111211113211) |
| `routes.custom.severity.exact_match` | [routes.custom.severity.exact_match](data-sources--alert_policy--reference--group-001.md#canonical-0101123330230002-2300322301133212-3132302100130103-1330113301320311-0313221221110131-2321212210102232-1030131110323303-0101120013221012) |
| `routes.custom.severity.regex_match` | [routes.custom.severity.regex_match](data-sources--alert_policy--reference--group-001.md#canonical-3311010123011222-0331233032021011-1123003301220133-1313213011222330-1211101021002213-2330123122223123-1100021300011223-3032131120020200) |
| `routes.dont_send` | [routes.dont_send](data-sources--alert_policy--reference--group-001.md#canonical-0132223031320311-0110122222210231-1231002203330213-0222312111123120-2033200333103022-3031310012301033-0032021130010031-1331321210020102) |
| `routes.group` | [routes.group](data-sources--alert_policy--reference--group-001.md#canonical-3002103111102220-1012333023113210-3300121023010111-0222122323001132-0032300202231221-1211331111113223-0203221000211132-2000302302123101) |
| `routes.group.groups` | [routes.group.groups](data-sources--alert_policy--reference--group-001.md#canonical-1000022221320103-1222000202021100-3122111223112020-1113023001301201-3130001110011122-1220300132223311-1323032000001310-2133111301220330) |
| `routes.notification_parameters` | [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-0103333310332002-0201030011110130-2120020133223320-2231013303031222-0331211013312220-3321000102132322-2011320332311223-2112131031211002) |
| `routes.notification_parameters.custom` | [routes.notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-0101011321000010-3333111112023110-3211331011302320-1303311131122212-0321201233000333-0012223012033210-3131322001002322-3130101213120100) |
| `routes.notification_parameters.custom.labels` | [routes.notification_parameters.custom.labels](data-sources--alert_policy--reference--group-001.md#canonical-0111203001231020-0210120122022221-1111233220011120-3321133313001113-1223203223313310-3112312033211211-2212321302111000-2322003111001122) |
| `routes.notification_parameters.default` | [routes.notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-3031002220332011-2010231333312330-0230101120133333-0001211210030301-3031321310323230-3300102031333302-3103313102223330-0211222113202213) |
| `routes.notification_parameters.group_interval` | [routes.notification_parameters.group_interval](data-sources--alert_policy--reference--group-001.md#canonical-1012210210331202-2221210310320011-1120203002202133-1122032003111011-1331323102202223-2302222323310312-1312022030223331-1230232332021211) |
| `routes.notification_parameters.group_wait` | [routes.notification_parameters.group_wait](data-sources--alert_policy--reference--group-001.md#canonical-3003203122231102-3132213122212201-3331122203210002-0000200231110332-1312132012001113-2233002021010311-3310323110303003-2111331223212113) |
| `routes.notification_parameters.individual` | [routes.notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-2200101031311233-1133100000221312-2313122032111302-0322132223310021-2303102120331223-3030233312132313-3203001123302233-2322100231021001) |
| `routes.notification_parameters.repeat_interval` | [routes.notification_parameters.repeat_interval](data-sources--alert_policy--reference--group-001.md#canonical-2000223001003011-3302231232030221-3030233001231023-1203222131203120-3132032210221133-2333000010231023-1311103203030310-1023020020033001) |
| `routes.notification_parameters.ves_io_group` | [routes.notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-2311122002223123-2031012112311122-3200001202033201-2100223230232301-2310101330121123-3131222302213100-1322320033130230-0300023233111321) |
| `routes.send` | [routes.send](data-sources--alert_policy--reference--group-001.md#canonical-0330200220023020-0100022021313312-2123332120122200-0112222003300113-1233123213113211-1213223312333323-2231333221230033-1312233121013021) |
| `routes.severity` | [routes.severity](data-sources--alert_policy--reference--group-001.md#canonical-3213211203012230-1320230122202221-2202232123231210-2122013320122133-3333330111202110-2020203302002300-3222201000020101-1030001202211002) |
| `routes.severity.severities` | [routes.severity.severities](data-sources--alert_policy--reference--group-001.md#canonical-0111213102232313-2011123113000302-2220121202132303-0012023332131221-3330010131301333-2313103332023012-3021130130330211-2213112032101303) |

<a id="canonical-2101211202102221-1110123323203321-3122021311100211-2031133020123103-0201103203120303-0302301003110223-0121020010033221-2032312103221013"></a>

## Next pages — Property reference / 031021233130 / 11

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- [receivers](data-sources--alert_policy--reference--group-001.md#canonical-3101011111113203-3320322310232130-1320032032033121-1221133020000113-1021233103201030-0221021130100132-1022002300200000-1102333030221331)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023003020212012-3032220330310003-1312113233022211-3120222132022121-0212013110211000-3012101312132202-2302112212113013-3310212302233100"></a>

## notification_parameters — notification_parameters / 220332130331 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- notification_parameters

<a id="canonical-3311331223111121-3201001311000211-1103202321032032-1112220303220133-0023330222201023-0312132021320321-1311323123020213-2000112032123131"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

<a id="canonical-0103202313022203-3003210301112200-3201333300321133-0300022200211221-1133112201332331-1302023101133002-0112122230002112-1102301203302301"></a>

## Direct properties — notification_parameters / 220332130331 / 3

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-0233321223231111-0120003030221221-3003223000130200-1310312013101223-1101010232232020-0321120122230010-2310002333110102-2321221321201310): complete subsection reference.

- [default](data-sources--alert_policy--reference--group-001.md#canonical-0103132313333320-1033203232332303-1201112210221213-0332233222331120-2121101230232101-0330310111232000-0332311221212221-0103030131303100): complete subsection reference.

<a id="canonical-3113231230230322-2123200030030021-3321032002132112-3213202201032320-1310333000313130-1011221213030200-0101232313213302-0022312021332310"></a>

<a id="canonical-1133200331200313-3130100230023011-0033323130223033-0132112120233120-3213112022322303-3200011232211223-2101211021233303-1020131013303201"></a>

## group_interval property — notification_parameters / 220332130331 / 4

Type: `"string"`. Computed.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-0300233112301300-2123020011322210-2003331103331033-3331230113002222-0001212011232100-3010032102030010-3012231002220031-1312103000231020"></a>

<a id="canonical-0023333113332300-3310331333101000-2032223123101102-3310203322013212-3211201100013212-2210101000321230-2211212313321212-3102021032301010"></a>

## group_wait property — notification_parameters / 220332130331 / 5

Type: `"string"`. Computed.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](data-sources--alert_policy--reference--group-001.md#canonical-2332033320102111-2322210023211030-3120032111320232-3011232021023103-1120103111111313-0010202331231133-0120110303012211-0313223232023032): complete subsection reference.

<a id="canonical-3102132323202022-3321313000331200-0011121312202030-3232111320333330-0322233300213101-1303020310301230-1020113201231222-2101310102101233"></a>

<a id="canonical-3020033002303300-3013333203003111-3023123302123332-0231220112022103-3303130033320313-1113321110333301-2131131032222132-2222111200012323"></a>

## repeat_interval property — notification_parameters / 220332130331 / 6

Type: `"string"`. Computed.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-2212200313101322-2121230020303002-1022102223230022-2220220201232210-3123211133202122-3003011011032011-0032023210030201-0101210300313020): complete subsection reference.

<a id="canonical-1013023221001111-3212221002003113-3023030210122031-3031221332100001-0230013222023113-0332101200230002-1001303112101312-1332110023121131"></a>

## Next pages — notification_parameters / 220332130331 / 7

- [notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-0233321223231111-0120003030221221-3003223000130200-1310312013101223-1101010232232020-0321120122230010-2310002333110102-2321221321201310)
- [notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-0103132313333320-1033203232332303-1201112210221213-0332233222331120-2121101230232101-0330310111232000-0332311221212221-0103030131303100)
- [notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-2332033320102111-2322210023211030-3120032111320232-3011232021023103-1120103111111313-0010202331231133-0120110303012211-0313223232023032)
- [notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-2212200313101322-2121230020303002-1022102223230022-2220220201232210-3123211133202122-3003011011032011-0032023210030201-0101210300313020)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-0233321223231111-0120003030221221-3003223000130200-1310312013101223-1101010232232020-0321120122230010-2310002333110102-2321221321201310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203311003102123-3131201210211333-0333310003232223-2310221111211223-3113121230323201-1032121232031210-3310323301022112-1020201331213132"></a>

## notification_parameters.custom — custom / 133233122111 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- notification_parameters.custom

<a id="canonical-3230212211103322-2232201330030103-0203220031323203-0100010013022321-1133320101321333-1221010021132221-2030323333322332-1033102132300030"></a>

Type: `"single"`. Computed.

Specify list of custom labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233001021101123-1310102332312201-0101001300130113-3021312110212011-0021302302133231-0010022121301103-2131101310032020-3312133230232232"></a>

## Direct properties — custom / 133233122111 / 3

<a id="canonical-0310132202000001-0300301312011033-2021200331232103-2323213022311100-1122111220232022-0002223011030121-2121131230013022-0000113220001123"></a>

<a id="canonical-2201222201012220-2001321311013013-3020011112211311-2320322210210301-0232110300112301-3022311211301112-0010230010011133-3000231033320023"></a>

## labels property — custom / 133233122111 / 4

Type: `["list", "string"]`. Computed.

Name of labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0321200023220321-3122301332013012-1222001321302022-1331211101233222-1201322002231031-3131111013003132-2113030322010223-1111122333322211"></a>

## Next pages — custom / 133233122111 / 5

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-0103132313333320-1033203232332303-1201112210221213-0332233222331120-2121101230232101-0330310111232000-0332311221212221-0103030131303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313010123323222-2313332123232202-0120013202331022-0213100011223000-1030311233230030-1023201121101231-3031310323022312-2233131312223100"></a>

## notification_parameters.default — default / 101221100001 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- notification_parameters.default

<a id="canonical-2023103200222212-3123222100101012-3030020021233310-1323200312001013-0031221031111000-2113231231022313-1313213332331021-3132310012001101"></a>

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

<a id="canonical-2120211102113322-3201023332201203-1033013231112001-0322232011103111-0303222201220310-1232112023102132-1332212312233030-0013320000010212"></a>

## Direct properties — default / 101221100001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033010031320012-2020113233121013-1120332311321132-0000030300021031-2102201122231032-2223131233113323-0313101132212233-0301201012001220"></a>

## Next pages — default / 101221100001 / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2332033320102111-2322210023211030-3120032111320232-3011232021023103-1120103111111313-0010202331231133-0120110303012211-0313223232023032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203210301300021-1212103102201203-1000013313211003-0022203032012111-1332020113320000-0131001102212301-1120121100210001-0100312013021222"></a>

## notification_parameters.individual — individual / 033321212122 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- notification_parameters.individual

<a id="canonical-0321010112112321-1332223023001311-0301130223232302-3323220333210333-1031212303201222-3230000010233100-2112202110031313-2220302303211233"></a>

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

<a id="canonical-3313310213122032-2131323322303031-3002330331320222-0012130202100233-1212103201202220-2322023301030200-2120123210211000-3222223110020000"></a>

## Direct properties — individual / 033321212122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203102303312221-1330232121112212-0113320303331030-1333101323222320-2022030012130103-0000113030211220-2030303113310333-2000201210133301"></a>

## Next pages — individual / 033321212122 / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2212200313101322-2121230020303002-1022102223230022-2220220201232210-3123211133202122-3003011011032011-0032023210030201-0101210300313020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012220103120310-3013130101100321-1112100302331020-1113001122020323-3032033202010330-3333030102220112-1310013211231010-2113130233223110"></a>

## notification_parameters.ves_io_group — ves_io_group / 311101020022 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- notification_parameters.ves_io_group

<a id="canonical-0231012113013021-3100203200012122-2010133021121211-0013222013030300-2213220203232023-0102132020310120-3010013013021210-0300131230333302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ves io group.

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

<a id="canonical-2032213312010111-3021002201032101-2020221000020023-3223122131331021-2311123001100312-2210023332130001-2111302103020223-2103321302002200"></a>

## Direct properties — ves_io_group / 311101020022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223301213311003-3233301221202022-1200000033113230-1301223131011103-3303220002110313-0133301112222200-0131100021120031-1312220112132023"></a>

## Next pages — ves_io_group / 311101020022 / 4

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-2012132313221311-3323230320132032-3021211202211100-2311103030201203-3201200032113120-3202102012130301-3022210210213311-0331112303130200)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-3101011111113203-3320322310232130-1320032032033121-1221133020000113-1021233103201030-0221021130100132-1022002300200000-1102333030221331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303111011013021-3010333123021321-3223030103211031-2232022212102211-2031011010320012-3332131013110231-0010011012002132-2111323331000232"></a>

## receivers — receivers / 120223231012 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- receivers

<a id="canonical-0131030030013223-1302101022213233-1320310131320031-0332112203223102-3103013311003033-2112323331103130-1303012130321213-2320132322113003"></a>

Type: `"list"`. Computed.

List of Alert Receivers where the alerts will be sent.

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

<a id="canonical-0130022322013133-1312023022130020-1202001013031323-2222003201031301-0100203122113100-3011000311332133-3113312302023200-3021100200332030"></a>

## Direct properties — receivers / 120223231012 / 3

<a id="canonical-0002010300223012-3331102033322000-2000211110223203-2031031232203000-0001013010200033-3310031310012103-3322300202013323-3122110220113023"></a>

<a id="canonical-2022102001000211-2033113110113303-0220202112310100-3231211303022002-3233211313301030-3222102311312303-3022102332132330-2313313203110111"></a>

## kind property — receivers / 120223231012 / 4

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

<a id="canonical-3022223111321033-2232032120201001-3012013013131002-3101030131000030-1230113101321131-1103120011120201-1303100110102333-0120021021013003"></a>

<a id="canonical-2120312231332300-3203322021113332-1310030323010011-3212001131302111-1021323121201021-0330213201231010-0230120312203013-1022121311003321"></a>

## name property — receivers / 120223231012 / 5

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

<a id="canonical-2300221220032301-2030013103031233-3222123311231022-3010121020330232-2103311022330033-2010011320131010-0101032312301332-3231133110233310"></a>

<a id="canonical-2200213303300112-1203203220000311-2011322023032330-3201012113321311-1030310130210203-0012221013132232-3003330123221231-3112332312222110"></a>

## namespace property — receivers / 120223231012 / 6

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

<a id="canonical-3311301122320110-2032222010301130-3300133100133130-0112030213111033-0023101023300332-0131130223211110-1013022003303012-1300311013123111"></a>

<a id="canonical-0301310202002230-1232332031231132-0221102300200312-0101300233133001-3201333022033321-0233110222130011-0033030321303311-0202303201131313"></a>

## tenant property — receivers / 120223231012 / 7

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

<a id="canonical-3233100333023301-0011201331202122-0330211030121322-2122311323023012-3202031122223230-0333210003001132-1011100232311132-0011020003100032"></a>

<a id="canonical-2212203121011301-0321330201223230-3032030301313333-2111022001002112-0001221302002232-2233211110210010-3031111232320320-0213132210121010"></a>

## uid property — receivers / 120223231012 / 8

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

<a id="canonical-0000113322113233-0100310221031312-2012122101013111-2330112122002202-3330111010333101-0012330332321023-3020220000003220-3312223130212300"></a>

## Next pages — receivers / 120223231012 / 9

- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232031322212102-1322000000301000-2231200021201131-3131211011320023-0312023212133201-3322112101012102-3022032220032031-0202221013011033"></a>

## routes — routes / 100023333102 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- routes

<a id="canonical-1102033113313133-3133012120331022-3210303121203103-2323112203132111-0310011032310210-3210101320122323-0320311232330103-2321300301313200"></a>

Type: `"list"`. Computed.

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

Upstream description:

Set of routes to match the incoming alert. The routes are evaluated in the specified order and
terminates on the first match.

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

<a id="canonical-3300013201113022-0122201130013120-1221102320020222-1122303132122020-1031000210102321-1023123222030310-1023103112013301-2233013333033222"></a>

## Direct properties — routes / 100023333102 / 3

<a id="canonical-0102101012203312-2131012303223211-0323213011202101-2120310303303102-2003021032201121-0312301113300113-2211121312201211-0312131230011012"></a>

<a id="canonical-1201323213232131-1112233131320330-3222020011203102-2003101321220332-2002023031133310-1223030223202001-3133231120201102-2210210100230211"></a>

## alertname property — routes / 100023333102 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN|SITE\_PHYSICAL\_INTERFACE\_DOWN|TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN|SERVICE\_SERVER\_ERROR|SERVICE\_CLIENT\_ERROR|SERVICE\_HEALTH\_LOW|SERVICE\_UNAVAILABLE|SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE|SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE|SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL|MALICIOUS\_USER\_DETECTED|WAF\_TOO\_MANY\_ATTACKS|API\_SECURITY\_TOO\_MANY\_ATTACKS|SERVICE\_POLICY\_TOO\_MANY\_ATTACKS|WAF\_TOO\_MANY\_MALICIOUS\_BOTS|BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS|THREAT\_CAMPAIGN|VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN|VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE|TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING|TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING|TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON|TLS\_CUSTOM\_CERTIFICATE\_EXPIRED|L7DDOS|DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD|API\_SECURITY\_UNUSED\_API\_DETECTED|API\_SECURITY\_SHADOW\_API\_DETECTED|API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED|API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED|ROUTED\_DDOS\_ALERT\_NOTIFICATION|ROUTED\_DDOS\_MITIGATION\_NOTIFICATION|ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION|L7\_DDOS\_AUTO\_MITIGATION\]
List of Alert Names Customer tunnel interface down Physical Interface down Tunnel Interfaces to
Customer Site Down Virtual Host server error Virtual Host client error Service Health Low Service
Unavailable Virtual Host server error Virtual Host client error Endpoint Healthcheck failure
Synthetic.. Possible values are \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`,
\`SITE\_PHYSICAL\_INTERFACE\_DOWN\`, \`TUNNELS\_TO\_CUSTOMER\_SITE\_DOWN\`,
\`SERVICE\_SERVER\_ERROR\`, \`SERVICE\_CLIENT\_ERROR\`, \`SERVICE\_HEALTH\_LOW\`,
\`SERVICE\_UNAVAILABLE\`, \`SERVICE\_SERVER\_ERROR\_PER\_SOURCE\_SITE\`,
\`SERVICE\_CLIENT\_ERROR\_PER\_SOURCE\_SITE\`, \`SERVICE\_ENDPOINT\_HEALTHCHECK\_FAILURE\`,
\`SYNTHETIC\_MONITOR\_HEALTH\_CRITICAL\`, \`MALICIOUS\_USER\_DETECTED\`,
\`WAF\_TOO\_MANY\_ATTACKS\`, \`API\_SECURITY\_TOO\_MANY\_ATTACKS\`,
\`SERVICE\_POLICY\_TOO\_MANY\_ATTACKS\`, \`WAF\_TOO\_MANY\_MALICIOUS\_BOTS\`,
\`BOT\_DEFENSE\_TOO\_MANY\_SECURITY\_EVENTS\`, \`THREAT\_CAMPAIGN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SUSPICIOUS\_DOMAIN\`,
\`VES\_CLIENT\_SIDE\_DEFENSE\_SENSITIVE\_FIELD\_READ\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_FAILURE\`,
\`TLS\_AUTOMATIC\_CERTIFICATE\_RENEWAL\_STILL\_FAILING\`, \`TLS\_AUTOMATIC\_CERTIFICATE\_EXPIRED\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\`, \`TLS\_CUSTOM\_CERTIFICATE\_EXPIRING\_SOON\`,
\`TLS\_CUSTOM\_CERTIFICATE\_EXPIRED\`, \`L7DDOS\`, \`DNS\_ZONE\_IGNORED\_DUPLICATE\_RECORD\`,
\`API\_SECURITY\_UNUSED\_API\_DETECTED\`, \`API\_SECURITY\_SHADOW\_API\_DETECTED\`,
\`API\_SECURITY\_SENSITIVE\_DATA\_IN\_RESPONSE\_DETECTED\`,
\`API\_SECURITY\_RISK\_SCORE\_HIGH\_DETECTED\`, \`ROUTED\_DDOS\_ALERT\_NOTIFICATION\`,
\`ROUTED\_DDOS\_MITIGATION\_NOTIFICATION\`, \`ROUTED\_DDOS\_TUNNEL\_STATUS\_UPDATE\_NOTIFICATION\`,
\`L7\_DDOS\_AUTO\_MITIGATION\`. Defaults to \`SITE\_CUSTOMER\_TUNNEL\_INTERFACE\_DOWN\`.

Upstream description:

List of Alert Names

Customer tunnel interface down Physical Interface down Tunnel Interfaces to Customer Site Down
Virtual Host server error Virtual Host client error Service Health Low Service Unavailable Virtual
Host server error Virtual Host client error Endpoint Healthcheck failure Synthetic monitor health
critical Malicious user detected Virtual Host WAF security events detected Virtual Host API security
events detected Virtual Host Service Policy security events detected Virtual Host Many Malicious
Bots based WAF security events detected Virtual Host Many Malicious Bots based Bot Defense security
events detected Virtual Host Many Threat campaign based WAF security events detected Suspicious
domain identified by Client-Side Defense service Client-Side Defense has identified a suspicious
script that is reading sensitive form field TLS Automatic Certificate renewal is failing TLS
Automatic Certificate renewal is still failing after multiple retries TLS Automatic Certificate has
expired TLS Custom Certificate will expire in less than 28 days TLS Custom Certificate will expire
in less than 15 days TLS Custom Certificate has expired DDoS security event detected DNS Zone
Ignored a Duplicate Record Create Request Unused APIs Detected Shadow APIs Detected Endpoints With
Sensitive Data In Response Detected High Risk Score Endpoints Detected A routed DDoS traffic anomaly
has been detected A routed DDoS mitigation has been implemented to block malicious traffic A routed
DDoS tunnel status has been changed L7 DDoS attack was detected, automatic mitigation is taking
place.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
  "enum": [
    "SITE_CUSTOMER_TUNNEL_INTERFACE_DOWN",
    "SITE_PHYSICAL_INTERFACE_DOWN",
    "TUNNELS_TO_CUSTOMER_SITE_DOWN",
    "SERVICE_SERVER_ERROR",
    "SERVICE_CLIENT_ERROR",
    "SERVICE_HEALTH_LOW",
    "SERVICE_UNAVAILABLE",
    "SERVICE_SERVER_ERROR_PER_SOURCE_SITE",
    "SERVICE_CLIENT_ERROR_PER_SOURCE_SITE",
    "SERVICE_ENDPOINT_HEALTHCHECK_FAILURE",
    "SYNTHETIC_MONITOR_HEALTH_CRITICAL",
    "MALICIOUS_USER_DETECTED",
    "WAF_TOO_MANY_ATTACKS",
    "API_SECURITY_TOO_MANY_ATTACKS",
    "SERVICE_POLICY_TOO_MANY_ATTACKS",
    "WAF_TOO_MANY_MALICIOUS_BOTS",
    "BOT_DEFENSE_TOO_MANY_SECURITY_EVENTS",
    "THREAT_CAMPAIGN",
    "VES_CLIENT_SIDE_DEFENSE_SUSPICIOUS_DOMAIN",
    "VES_CLIENT_SIDE_DEFENSE_SENSITIVE_FIELD_READ",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_FAILURE",
    "TLS_AUTOMATIC_CERTIFICATE_RENEWAL_STILL_FAILING",
    "TLS_AUTOMATIC_CERTIFICATE_EXPIRED",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING",
    "TLS_CUSTOM_CERTIFICATE_EXPIRING_SOON",
    "TLS_CUSTOM_CERTIFICATE_EXPIRED",
    "L7DDOS",
    "DNS_ZONE_IGNORED_DUPLICATE_RECORD",
    "API_SECURITY_UNUSED_API_DETECTED",
    "API_SECURITY_SHADOW_API_DETECTED",
    "API_SECURITY_SENSITIVE_DATA_IN_RESPONSE_DETECTED",
    "API_SECURITY_RISK_SCORE_HIGH_DETECTED",
    "ROUTED_DDOS_ALERT_NOTIFICATION",
    "ROUTED_DDOS_MITIGATION_NOTIFICATION",
    "ROUTED_DDOS_TUNNEL_STATUS_UPDATE_NOTIFICATION",
    "L7_DDOS_AUTO_MITIGATION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1232002102300130-1113131031212233-0202132101330112-1323213313323213-3302002320103310-2023030301121230-0031023222003011-3111133030011101"></a>

<a id="canonical-3212020003112212-3103132312303011-0310122220130022-2001110122312231-1201202013233131-0123032031202100-0121110032312023-3023211002231222"></a>

## alertname_regex property — routes / 100023333102 / 5

Type: `"string"`. Computed.

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

Upstream description:

Exclusive with \[alertname any custom group severity\] Regular Expression match for the alertname.

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

- [any](data-sources--alert_policy--reference--group-001.md#canonical-1111130100211122-3222210232030202-2221333323112311-1121201113020300-2321311012312100-2232233102302103-1332323302320030-2333211031110032): complete subsection reference.

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000): complete subsection reference.

- [dont_send](data-sources--alert_policy--reference--group-001.md#canonical-3213231322233200-3122333111133131-1331323303311222-1202333313320223-3301303133202133-0033131233332030-3002203323131311-2323310302033111): complete subsection reference.

- [group](data-sources--alert_policy--reference--group-001.md#canonical-2233223112013011-1321003322230102-0301001203213333-0122000222202110-3121223221010132-2301332032303312-0110021331332102-1030231113112203): complete subsection reference.

- [notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322): complete subsection reference.

- [send](data-sources--alert_policy--reference--group-001.md#canonical-2301303200003012-0222210012300032-1222303311120232-0131120201030010-2220103330210030-3033332302303230-2102202012311123-0103231122212002): complete subsection reference.

- [severity](data-sources--alert_policy--reference--group-001.md#canonical-0023021201013110-2032203003113120-0330000033303313-1323320123013030-1130013333022020-1220223110321132-0232223012001321-2122302330130311): complete subsection reference.

<a id="canonical-1300230310220313-3223221131022113-2231221010022222-1101102123310303-0322122310002030-1000232120133303-2200333202313001-3301332230211022"></a>

## Next pages — routes / 100023333102 / 6

- [routes.any](data-sources--alert_policy--reference--group-001.md#canonical-1111130100211122-3222210232030202-2221333323112311-1121201113020300-2321311012312100-2232233102302103-1332323302320030-2333211031110032)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- [routes.dont_send](data-sources--alert_policy--reference--group-001.md#canonical-3213231322233200-3122333111133131-1331323303311222-1202333313320223-3301303133202133-0033131233332030-3002203323131311-2323310302033111)
- [routes.group](data-sources--alert_policy--reference--group-001.md#canonical-2233223112013011-1321003322230102-0301001203213333-0122000222202110-3121223221010132-2301332032303312-0110021331332102-1030231113112203)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- [routes.send](data-sources--alert_policy--reference--group-001.md#canonical-2301303200003012-0222210012300032-1222303311120232-0131120201030010-2220103330210030-3033332302303230-2102202012311123-0103231122212002)
- [routes.severity](data-sources--alert_policy--reference--group-001.md#canonical-0023021201013110-2032203003113120-0330000033303313-1323320123013030-1130013333022020-1220223110321132-0232223012001321-2122302330130311)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-1111130100211122-3222210232030202-2221333323112311-1121201113020300-2321311012312100-2232233102302103-1332323302320030-2333211031110032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023232101223313-1002023130032321-2032123200221322-3020300210213230-2002321233201202-0303312201123132-3313133213021003-0130200230213221"></a>

## routes.any — any / 211320201303 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.any

<a id="canonical-3321203133232201-3101030300021333-1201112331303301-3102312230232222-1321000003330021-3132022032212120-0213013230213133-2310310101232332"></a>

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

<a id="canonical-2311301222121310-0220132010102122-3320133320120021-3300023101200031-1312031102331122-3102000032201213-0130212200323131-2222303202202313"></a>

## Direct properties — any / 211320201303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002122121000300-3333133013020210-3122232300332330-1321111200200003-3210323122301033-1132110332133311-3110202210131002-0200210010210010"></a>

## Next pages — any / 211320201303 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312301132231310-1320230300130100-1130320023210233-2021222213221211-0123331120230333-3103232022131202-3100303331132201-3221321101201130"></a>

## routes.custom — custom / 311021023122 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.custom

<a id="canonical-1100133110101131-3003321221333201-0303103011131003-0212230132100123-3130101110021213-3201010101230001-2310130002221212-1330031112023101"></a>

Type: `"single"`. Computed.

Set of matchers an alert has to fulfill to match the route.

Upstream description:

A set of matchers an alert has to fulfill to match the route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2002201231031201-1120213110303322-2113233010003233-2212012101200013-0022213120020311-0000221310001033-2203112110232131-0100310123111222"></a>

## Direct properties — custom / 311021023122 / 3

- [alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-0123003123222100-3123103230023332-2130113213210332-1212230331013201-1011220121233331-0123130201200320-1313121332220101-3013113311311000): complete subsection reference.

- [alertname](data-sources--alert_policy--reference--group-001.md#canonical-1010331031221220-1030103211013011-2303202332032322-1321212012131321-2121310312130001-0212333021201322-2322201020003030-1321202200021010): complete subsection reference.

- [group](data-sources--alert_policy--reference--group-001.md#canonical-3123330332320212-3011100103333130-1213121013213331-0200203320222032-0200230303123002-2131111031321222-2211333211033313-1121112231020111): complete subsection reference.

- [severity](data-sources--alert_policy--reference--group-001.md#canonical-2223332111322120-2212302211303331-1121033012003020-3012032003021332-3001232020323333-3302321132331233-0220230103100022-3233130331330020): complete subsection reference.

<a id="canonical-2230110333101011-2001231201313211-0322231113211302-1100110133001023-2312210320300320-3112313111013330-0233301321200322-2031301203301322"></a>

## Next pages — custom / 311021023122 / 4

- [routes.custom.alertlabel](data-sources--alert_policy--reference--group-001.md#canonical-0123003123222100-3123103230023332-2130113213210332-1212230331013201-1011220121233331-0123130201200320-1313121332220101-3013113311311000)
- [routes.custom.alertname](data-sources--alert_policy--reference--group-001.md#canonical-1010331031221220-1030103211013011-2303202332032322-1321212012131321-2121310312130001-0212333021201322-2322201020003030-1321202200021010)
- [routes.custom.group](data-sources--alert_policy--reference--group-001.md#canonical-3123330332320212-3011100103333130-1213121013213331-0200203320222032-0200230303123002-2131111031321222-2211333211033313-1121112231020111)
- [routes.custom.severity](data-sources--alert_policy--reference--group-001.md#canonical-2223332111322120-2212302211303331-1121033012003020-3012032003021332-3001232020323333-3302321132331233-0220230103100022-3233130331330020)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-0123003123222100-3123103230023332-2130113213210332-1212230331013201-1011220121233331-0123130201200320-1313121332220101-3013113311311000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112101012003123-0312320211130310-1330230010211323-1201220321030202-0333211322013123-3111213202121201-1233011100323033-3303001320221232"></a>

## routes.custom.alertlabel — alertlabel / 022111123111 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- routes.custom.alertlabel

<a id="canonical-0300320320213020-0201230112033000-3331021103003010-3300201111133231-0222303022022212-0230010311313031-3221320232022131-2100031022031301"></a>

Type: `"single"`. Computed.

AlertLabel to configure the alert policy rule.

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
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.keys.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.map.max_pairs": "3"
  }
}
```

<a id="canonical-0001213100333321-0212233132300033-1120211001132200-3110313320220230-3031313333013303-1001210201233011-2231122300210202-1311011313332202"></a>

## Direct properties — alertlabel / 022111123111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023122020103133-1232132023201110-3132120330113033-2020332210002222-3201232123001223-2002000231202213-1300321331302331-3100220013233113"></a>

## Next pages — alertlabel / 022111123111 / 4

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-1010331031221220-1030103211013011-2303202332032322-1321212012131321-2121310312130001-0212333021201322-2322201020003030-1321202200021010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010232211132232-1130311121021203-3123223111331301-1312222311102012-3332002123200002-0303022000110102-3320231102103002-3100003330220011"></a>

## routes.custom.alertname — alertname / 320030003012 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- routes.custom.alertname

<a id="canonical-0211320010210201-1301121311130311-1000010231132322-2000022330021212-2231022101122301-3213112231000232-1330222320222101-2223011201213302"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-0333113203322310-3020130322321120-2221011030001333-0130121230312233-0013110211312312-2100032030112131-1101222111202002-2120312021321312"></a>

## Direct properties — alertname / 320030003012 / 3

<a id="canonical-0230003113310122-2333002013113122-0203302223321030-3110332303321132-2100332102020212-1211002022133311-3302131122030111-1030313323120322"></a>

<a id="canonical-2220312020012010-0212003031232202-1220312330101322-1220020321221012-0133013010313212-0210221320120131-2003233013100210-1220302131332200"></a>

## exact_match property — alertname / 320030003012 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-2320202230230311-0123222223302222-0323230001103323-1113322202123333-2120111120020112-3110010233323311-0203201103100300-1123033323113231"></a>

<a id="canonical-2213322012110101-0302330110232333-3120111320110232-0202332212320100-2232221301213020-3133300311230202-3320220203030200-3323123113100202"></a>

## regex_match property — alertname / 320030003012 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-2212000030033031-0103202023130102-0113002232032022-0032211031123232-1122233113120330-0332302203132333-0103020133223011-3221122201231133"></a>

## Next pages — alertname / 320030003012 / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-3123330332320212-3011100103333130-1213121013213331-0200203320222032-0200230303123002-2131111031321222-2211333211033313-1121112231020111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011333001210233-3231232303230320-1032130013223132-3010103022210001-3201120221113311-3120013010321231-0001201321220023-0113111013130301"></a>

## routes.custom.group — group / 011110312110 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- routes.custom.group

<a id="canonical-1222333231122131-3011102100102133-3021212120101302-3013213303101303-2233231301222312-3320221310123230-1021022121202313-1223133100010033"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-1303001200032301-3322001010020331-0330222321133031-1213331101201201-1303312321300020-3120232320020320-3201001223030110-3122122320131311"></a>

## Direct properties — group / 011110312110 / 3

<a id="canonical-0001333013300020-0310122312201131-2330213022122013-3202120201110222-0312333000200232-0310131033203302-0313202030032011-1132321333303002"></a>

<a id="canonical-0221101010233200-1321033002222313-0311301212120013-3202201030311013-0023222232102000-3113112333001201-1233123223303121-2222120301012323"></a>

## exact_match property — group / 011110312110 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-2232200020132213-1002220112003230-0221300223211213-1113130130021031-3223313000331300-1210202322133230-0231000221232130-0123322112302012"></a>

<a id="canonical-2131303212231103-3310133101131131-1312333300232232-2312222323311203-0122002233021020-1202213111302332-3011230212110332-1131013220213122"></a>

## regex_match property — group / 011110312110 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-0111021200333010-3310211003110023-3011001121022200-3001120332101010-0222020023031031-0201203301030122-0102220100021303-2303322322021310"></a>

## Next pages — group / 011110312110 / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2223332111322120-2212302211303331-1121033012003020-3012032003021332-3001232020323333-3302321132331233-0220230103100022-3233130331330020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020311222301111-2332000231130102-2333203201031310-3323212312032201-3201133130022030-2232332331200212-2220103023003303-0303033301012030"></a>

## routes.custom.severity — severity / 013020232322 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- routes.custom.severity

<a id="canonical-0300013320001122-3323123022101302-1021220302111321-2103233221302030-0310213331020212-3100033011220000-3211112133133102-2102111211113211"></a>

Type: `"single"`. Computed.

Label Matcher.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-matcher_type": "[\"exact_match\",\"regex_match\"]"
}
```

<a id="canonical-2132120112121202-3221202223113323-0221011121202121-3232122132033032-1111012022013323-1322233002121020-2213220200003103-1103131120100213"></a>

## Direct properties — severity / 013020232322 / 3

<a id="canonical-0101123330230002-2300322301133212-3132302100130103-1330113301320311-0313221221110131-2321212210102232-1030131110323303-0101120013221012"></a>

<a id="canonical-3001131021223213-3211312121121233-3031101111130323-1312100003122011-2112202020001303-1012012311333000-3331202221032233-3323030322203111"></a>

## exact_match property — severity / 013020232322 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_match\] Equality match value for the label.

Upstream description:

Exclusive with \[regex\_match\] Equality match value for the label.

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

<a id="canonical-3311010123011222-0331233032021011-1123003301220133-1313213011222330-1211101021002213-2330123122223123-1100021300011223-3032131120020200"></a>

<a id="canonical-0233313233322212-0022231031303232-3203023331120101-2212030122223003-1121313330230121-3011230300033203-2000331021023122-2002332122020011"></a>

## regex_match property — severity / 013020232322 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_match\] Regular expression match value for the label.

Upstream description:

Exclusive with \[exact\_match\] Regular expression match value for the label.

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

<a id="canonical-2233313220032232-0231223201211201-0210303011202323-2100333320002121-3220132103011230-3232113003021101-1032311122320031-0110212023333032"></a>

## Next pages — severity / 013020232322 / 6

- [routes.custom](data-sources--alert_policy--reference--group-001.md#canonical-2111032222302223-0233011322210313-1123202120313233-2330130321123330-0231130212120033-2213022223222302-1301321103330130-3002100333101000)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-3213231322233200-3122333111133131-1331323303311222-1202333313320223-3301303133202133-0033131233332030-3002203323131311-2323310302033111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032332001103310-2220202323001103-1213321210013001-2231012313223101-2303012330203120-1332210331221132-3331022301302011-0120311212000222"></a>

## routes.dont_send — dont_send / 311113120211 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.dont_send

<a id="canonical-0132223031320311-0110122222210231-1231002203330213-0222312111123120-2033200333103022-3031310012301033-0032021130010031-1331321210020102"></a>

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

<a id="canonical-0212211002300132-3310200223203311-2320200132111003-2131030320123100-1330232011110331-1113223113321020-2000110322132133-1121211011131002"></a>

## Direct properties — dont_send / 311113120211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201311111201223-1011012032223313-1300020311020231-3133023310303211-3033303001313202-1131133212300013-0032301113101332-0132221031103331"></a>

## Next pages — dont_send / 311113120211 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2233223112013011-1321003322230102-0301001203213333-0122000222202110-3121223221010132-2301332032303312-0110021331332102-1030231113112203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200200323212232-3213212012321220-3032013203220033-3133033031330010-3021333213200300-1123031320111030-3223100013000203-1300202213133313"></a>

## routes.group — group / 003111310301 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.group

<a id="canonical-3002103111102220-1012333023113210-3300121023010111-0222122323001132-0032300202231221-1211331111113223-0203221000211132-2000302302123101"></a>

Type: `"single"`. Computed.

Select one or more known group names to match the incoming alert.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1211002323000303-3123322220230103-2203010331003323-0111131130321032-1120112021000321-3231033333022200-1032223312002203-2303312002330131"></a>

## Direct properties — group / 003111310301 / 3

<a id="canonical-1000022221320103-1222000202021100-3122111223112020-1113023001301201-3130001110011122-1220300132223311-1323032000001310-2133111301220330"></a>

<a id="canonical-1003231231012312-1302121230120023-0301310230110031-2010211103012311-1120003323100120-3111310023320310-2100011031020032-2202002003321031"></a>

## groups property — group / 003111310301 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
INFRASTRUCTURE|IAAS\_CAAS|VIRTUAL\_HOST|VOLT\_SHARE|UAM|SECURITY|TIMESERIES\_ANOMALY|SHAPE\_SECURITY|SECURITY\_CSD|CDN|SYNTHETIC\_MONITORS|TLS|SECURITY\_BOT\_DEFENSE|CLOUD\_LINK|DNS|ROUTED\_DDOS\]
Groups. Name of groups to match the alert. Possible values are \`INFRASTRUCTURE\`, \`IAAS\_CAAS\`,
\`VIRTUAL\_HOST\`, \`VOLT\_SHARE\`, \`UAM\`, \`SECURITY\`, \`TIMESERIES\_ANOMALY\`,
\`SHAPE\_SECURITY\`, \`SECURITY\_CSD\`, \`CDN\`, \`SYNTHETIC\_MONITORS\`, \`TLS\`,
\`SECURITY\_BOT\_DEFENSE\`, \`CLOUD\_LINK\`, \`DNS\`, \`ROUTED\_DDOS\`. Defaults to
\`INFRASTRUCTURE\`.

Upstream description:

Name of groups to match the alert.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1221300312023303-2131130232120030-0113031112232310-3131221222020122-2210302123223303-3120123001202030-2011200303303333-3313020111220333"></a>

## Next pages — group / 003111310301 / 5

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212022103311102-2202130223103223-1033120021312000-3202131002010231-0013133302120223-0302122302310300-3022102220110000-0211322102110022"></a>

## routes.notification_parameters — notification_parameters / 302130012011 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.notification_parameters

<a id="canonical-0103333310332002-0201030011110130-2120020133223320-2231013303031222-0331211013312220-3321000102132322-2011320332311223-2112131031211002"></a>

Type: `"single"`. Computed.

Set of notification parameters to decide how and when the alert notifications should be sent to the
receivers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-group_by": "[\"custom\",\"default\",\"individual\",\"ves_io_group\"]"
}
```

<a id="canonical-1301131001322101-2020213102230202-3301233300322100-2221332313212120-3103230231303233-1323332303323210-1203302333100113-0023210011130101"></a>

## Direct properties — notification_parameters / 302130012011 / 3

- [custom](data-sources--alert_policy--reference--group-001.md#canonical-1010231102310202-0032222133032113-2000321113021313-0331323220033212-3202133213233002-0322312220221230-1032322231022032-2201203232220130): complete subsection reference.

- [default](data-sources--alert_policy--reference--group-001.md#canonical-2030222302312212-1321303222330102-2033130010200120-3021213101000022-3000231031103013-3023330011021113-2220211222120233-2332311021010112): complete subsection reference.

<a id="canonical-1012210210331202-2221210310320011-1120203002202133-1122032003111011-1331323102202223-2302222323310312-1312022030223331-1230232332021211"></a>

<a id="canonical-2310201330333203-1220031322010130-1101221320300022-2001230102332022-1231020230010310-1222321001131010-2230223222023323-1000213323332320"></a>

## group_interval property — notification_parameters / 302130012011 / 4

Type: `"string"`. Computed.

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval..

Upstream description:

Group Interval is used to specify how long to wait before sending a notification about new alerts
that are added to the group for which an initial notification has already been sent. Format:
\[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours, d - days If not specified,
group\_interval defaults to "1m"

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "30s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3003203122231102-3132213122212201-3331122203210002-0000200231110332-1312132012001113-2233002021010311-3310323110303003-2111331223212113"></a>

<a id="canonical-0131021320102221-2120312223111230-1110101212123311-3113322100200232-2010230323320213-2201021332121110-1112220302032011-1032322303330032"></a>

## group_wait property — notification_parameters / 302130012011 / 5

Type: `"string"`. Computed.

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to '30s'.

Upstream description:

Time value used to specify how long to initially wait for an inhibiting alert to arrive or collect
more alerts for the same group. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_wait defaults to "30s"

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
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "5m",
    "ves.io.schema.rules.string.min_time_interval": "0s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [individual](data-sources--alert_policy--reference--group-001.md#canonical-3112321112233210-3311210130223311-3123321123133300-2322322132122111-3311213223221121-2033100202112101-3313111231031300-3022331012003100): complete subsection reference.

<a id="canonical-2000223001003011-3302231232030221-3030233001231023-1203222131203120-3132032210221133-2333000010231023-1311103203030310-1023020020033001"></a>

<a id="canonical-3321302130321010-1013132122213102-0103301013222301-0110023312133331-1213220101332031-2332130232132110-2101132130112032-0320010312302111"></a>

## repeat_interval property — notification_parameters / 302130012011 / 6

Type: `"string"`. Computed.

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to '4h'.

Upstream description:

Repeat Interval is used to specify how long to wait before sending a notification again if it has
already been sent successfully. Format: \[0-9\]\[smhd\], where s - seconds, m - minutes, h - hours,
d - days If not specified, group\_interval defaults to "4h"

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
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10d",
    "ves.io.schema.rules.string.min_time_interval": "30m",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-1030030013101223-2100310011303130-0111031203222000-1211131233011002-1212223121310201-0023321313011322-1012333032321003-3300230111200001): complete subsection reference.

<a id="canonical-2111100102220122-3300321330021120-2032133210213030-1132112121031003-3032030202032311-1120121122130230-0212202011012010-1131200121200113"></a>

## Next pages — notification_parameters / 302130012011 / 7

- [routes.notification_parameters.custom](data-sources--alert_policy--reference--group-001.md#canonical-1010231102310202-0032222133032113-2000321113021313-0331323220033212-3202133213233002-0322312220221230-1032322231022032-2201203232220130)
- [routes.notification_parameters.default](data-sources--alert_policy--reference--group-001.md#canonical-2030222302312212-1321303222330102-2033130010200120-3021213101000022-3000231031103013-3023330011021113-2220211222120233-2332311021010112)
- [routes.notification_parameters.individual](data-sources--alert_policy--reference--group-001.md#canonical-3112321112233210-3311210130223311-3123321123133300-2322322132122111-3311213223221121-2033100202112101-3313111231031300-3022331012003100)
- [routes.notification_parameters.ves_io_group](data-sources--alert_policy--reference--group-001.md#canonical-1030030013101223-2100310011303130-0111031203222000-1211131233011002-1212223121310201-0023321313011322-1012333032321003-3300230111200001)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-1010231102310202-0032222133032113-2000321113021313-0331323220033212-3202133213233002-0322312220221230-1032322231022032-2201203232220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133220113020013-3022311321100133-3330000033311120-1122233303310330-2033023131320201-3232130303020002-0123022102311112-2000230322023123"></a>

## routes.notification_parameters.custom — custom / 033100321132 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- routes.notification_parameters.custom

<a id="canonical-0101011321000010-3333111112023110-3211331011302320-1303311131122212-0321201233000333-0012223012033210-3131322001002322-3130101213120100"></a>

Type: `"single"`. Computed.

Specify list of custom labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1213332303300123-1331102310001311-3213203132032110-0320232310020033-1233103230121122-0300123121211131-1230233333222012-3000202220232201"></a>

## Direct properties — custom / 033100321132 / 3

<a id="canonical-0111203001231020-0210120122022221-1111233220011120-3321133313001113-1223203223313310-3112312033211211-2212321302111000-2322003111001122"></a>

<a id="canonical-0220012222322012-1231013120200223-1033230222032120-0103332121130010-1302033211012013-0133210310321020-0012231120101003-1031020021133213"></a>

## labels property — custom / 033100321132 / 4

Type: `["list", "string"]`. Computed.

Name of labels to group/aggregate the alerts.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[a-zA-Z_][a-zA-Z0-9_]*$",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3011232120013323-0031002111213121-1222031010332112-0230031101121303-3111303232000132-2103312203233132-3102321230010322-1321120112223131"></a>

## Next pages — custom / 033100321132 / 5

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2030222302312212-1321303222330102-2033130010200120-3021213101000022-3000231031103013-3023330011021113-2220211222120233-2332311021010112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100322002202120-0310211012321211-1121023121102112-1210103121101302-0202131132030223-3010302022211101-0232223000130013-1223232312331100"></a>

## routes.notification_parameters.default — default / 300100202222 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- routes.notification_parameters.default

<a id="canonical-3031002220332011-2010231333312330-0230101120133333-0001211210030301-3031321310323230-3300102031333302-3103313102223330-0211222113202213"></a>

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

<a id="canonical-1000102201122303-2310033223103002-1132312113200130-3313022232131223-3020003301103300-2010101300022123-0213022122020012-1111312111223103"></a>

## Direct properties — default / 300100202222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112212333123021-3331202233031131-2212101221221303-1303002223233222-3010232203320322-3213023301223102-0030112231200110-0213110010123031"></a>

## Next pages — default / 300100202222 / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-3112321112233210-3311210130223311-3123321123133300-2322322132122111-3311213223221121-2033100202112101-3313111231031300-3022331012003100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323111003101302-2131022223022122-0232313121331102-3201200031322013-3223210301300011-1210200012002221-0012103332131113-2220323221313201"></a>

## routes.notification_parameters.individual — individual / 011112322201 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- routes.notification_parameters.individual

<a id="canonical-2200101031311233-1133100000221312-2313122032111302-0322132223310021-2303102120331223-3030233312132313-3203001123302233-2322100231021001"></a>

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

<a id="canonical-3031322320330332-3132202231112323-3113233212303020-3300023120322201-0201220022303203-3000223001122231-0023031130201131-2103220020312301"></a>

## Direct properties — individual / 011112322201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210311203220303-2303103202011012-1233203122333313-0302302221101233-3012122011333223-2001332203231323-0101030301301223-0300203110310211"></a>

## Next pages — individual / 011112322201 / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-1030030013101223-2100310011303130-0111031203222000-1211131233011002-1212223121310201-0023321313011322-1012333032321003-3300230111200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000113302231200-2310121132233232-3333200223232232-0103213231200022-3021313333203313-0200312031032023-3321201300213302-2303023001220310"></a>

## routes.notification_parameters.ves_io_group — ves_io_group / 122000032121 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- routes.notification_parameters.ves_io_group

<a id="canonical-2311122002223123-2031012112311122-3200001202033201-2100223230232301-2310101330121123-3131222302213100-1322320033130230-0300023233111321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ves io group.

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

<a id="canonical-3010023212300320-3131100302123233-3230232212322202-3133113000110330-0033212133103232-1130231012102130-2311230001103311-3232133213032021"></a>

## Direct properties — ves_io_group / 122000032121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030222001003131-2030330030022201-2133233103321303-2213302200103230-0120031300312321-2103132020303201-1110111313033111-3123110222201030"></a>

## Next pages — ves_io_group / 122000032121 / 4

- [routes.notification_parameters](data-sources--alert_policy--reference--group-001.md#canonical-3222233311320102-0030322313100000-1023230020323111-3023210212302030-3002100220323100-0111221013302013-0223102120002031-0233120202310322)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-2301303200003012-0222210012300032-1222303311120232-0131120201030010-2220103330210030-3033332302303230-2102202012311123-0103231122212002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203333133211130-2211323301103122-3300320320121211-0211123301230202-0122333023203120-3202301000022002-2022102120102212-2031230133012102"></a>

## routes.send — send / 221113001011 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.send

<a id="canonical-0330200220023020-0100022021313312-2123332120122200-0112222003300113-1233123213113211-1213223312333323-2231333221230033-1312233121013021"></a>

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

<a id="canonical-2231112310013131-3132312310133302-1130321033102010-0000102010312000-1303200031233210-3320133033010031-0312301200211023-1323011022323311"></a>

## Direct properties — send / 221113001011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111022032221112-0131023210123212-3011012300323221-3110210330222100-1221011112020010-3001012223212013-1332333013010032-2111303101131102"></a>

## Next pages — send / 221113001011 / 4

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)

<a id="canonical-0023021201013110-2032203003113120-0330000033303313-1323320123013030-1130013333022020-1220223110321132-0232223012001321-2122302330130311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313222001113011-3111112311211332-3130301012232033-1301301001023113-0003103103010312-0121010033311232-2212120113131301-0133213131233303"></a>

## routes.severity — severity / 232003033111 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
- [Property reference](data-sources--alert_policy--reference--group-001.md#canonical-2112221201120012-2232133313303133-2202012021221221-3312133021112123-2200311330232132-2302002223011013-0001113332303200-1210002023212123)
- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- routes.severity

<a id="canonical-3213211203012230-1320230122202221-2202232123231210-2122013320122133-3333330111202110-2020203302002300-3222201000020101-1030001202211002"></a>

Type: `"single"`. Computed.

Select one or more severity levels to match the incoming alert.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2030112321232322-0232132313102131-0010220033233132-3130323002330011-3000001332301113-0323223222203110-3023120123100212-3303312013021231"></a>

## Direct properties — severity / 232003033111 / 3

<a id="canonical-0111213102232313-2011123113000302-2220121202132303-0012023332131221-3330010131301333-2313103332023012-3021130130330211-2213112032101303"></a>

<a id="canonical-1101130021000331-1111312110110120-3220222121222331-0113331032212323-3322131130031100-1020312231333221-2022202101312301-1031013221332313"></a>

## severities property — severity / 232003033111 / 4

Type: `["list", "string"]`. Computed.

\[Enum: MINOR|MAJOR|CRITICAL\] Severities. List of severity levels. Possible values are \`MINOR\`,
\`MAJOR\`, \`CRITICAL\`. Defaults to \`MINOR\`.

Upstream description:

List of severity levels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1202231113111213-1302111030300221-1201120212133101-3310111331202312-2130101220221303-2333232200112121-1001311233220231-0200032123202133"></a>

## Next pages — severity / 232003033111 / 5

- [routes](data-sources--alert_policy--reference--group-001.md#canonical-1021230332031130-2213013032020203-1113313122013232-3230212020122200-2010101322212202-2331133113011102-2113130202310120-2003213023103202)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-0320032130000223-0013312111113203-3221033200310232-1301022120210012-3300321311100201-0013103020333103-2030012213311032-1033111200111130)
