---
page_title: "xcsh_network_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_network_policy reference."
---

# xcsh_network_policy reference

<a id="canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301122133002021-1221033030112230-0013220021333210-3210020313022012-0002120023213210-2102323010323231-2020010102000232-0233030201133303"></a>

## Property reference — Property reference / 322213232030 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- Property reference

<a id="canonical-0110013220103331-2231012223130020-0203201100300103-1231100302332100-1121231132112013-3130120013201311-0133121110302111-3231103222111312"></a>

## Direct properties — Property reference / 322213232030 / 3

<a id="canonical-1131321201212221-0132232210101123-1201300303330003-2003212103300103-0032101330112021-2121100100130123-0332332231210322-3232020131323001"></a>

<a id="canonical-1132023121021220-3331123213120010-3223201101020222-3321311313223230-2322223230333321-3220311131112123-0211020001221322-1022211021202110"></a>

## annotations property — Property reference / 322213232030 / 4

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

<a id="canonical-0013230330321010-2030121212311132-3130302102200021-2212223112330200-1320032001201022-1232230321200122-2320001130233012-1132322112111332"></a>

<a id="canonical-3023332023130210-0313323301133223-2102012130001312-0012001011013130-2310100330200233-0220013322031213-2013333101100302-2002112033303011"></a>

## description property — Property reference / 322213232030 / 5

Type: `"string"`. Computed.

Description of the NetworkPolicy.

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

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332): complete subsection reference.

<a id="canonical-3010313121111023-3010002322102120-3012121023023002-1001102123303001-3313323113033123-0113133303110012-3021022123221021-0201010212012302"></a>

<a id="canonical-2310131100001323-2033121322131130-0331103110202233-3021121333133201-2203102120213101-3303200000202122-2220200203320013-2322232203031320"></a>

## ID property — Property reference / 322213232030 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0031302010302201-2111203012303001-1323000230123132-2122333100313000-0302231011122113-3130121300100023-1132013122000223-2123101133231203"></a>

<a id="canonical-1133203213002320-1301123301030020-2313022200002230-1001200131000311-1102101213232001-0222130021122213-2301023032030232-1010101031312033"></a>

## labels property — Property reference / 322213232030 / 7

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

<a id="canonical-0100030011230122-0112322010322030-2022303010110331-2000310321211103-3010032100030122-3311000333021332-2032131301023203-1230022110012211"></a>

<a id="canonical-1223001300302020-0121130122203312-2011123300112302-1101102122123230-0110010100112020-1013210011131112-0331213331321121-3012131333331020"></a>

## name property — Property reference / 322213232030 / 8

Type: `"string"`. Required.

Name of the NetworkPolicy.

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

<a id="canonical-0131130223013111-0302321322112311-0232123033313302-0100103132123033-1221311210210232-2200230100100201-3323122123220221-3302032332211002"></a>

<a id="canonical-0121310302033323-2220301001110232-3322020032203002-3103322222301221-0110202121002022-1201300212031002-3332312230120100-0120323011222023"></a>

## namespace property — Property reference / 322213232030 / 9

Type: `"string"`. Required.

Namespace where the NetworkPolicy exists.

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

- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212): complete subsection reference.

<a id="canonical-3123103333001321-1132010022312133-3120320022111023-3332103332323102-1220212210212330-3322213311111010-1310332230310112-2303333221213101"></a>

## All schema paths — Property reference / 322213232030 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_policy--reference--group-001.md#canonical-1131321201212221-0132232210101123-1201300303330003-2003212103300103-0032101330112021-2121100100130123-0332332231210322-3232020131323001) |
| `description` | [description](data-sources--network_policy--reference--group-001.md#canonical-0013230330321010-2030121212311132-3130302102200021-2212223112330200-1320032001201022-1232230321200122-2320001130233012-1132322112111332) |
| `endpoint` | [endpoint](data-sources--network_policy--reference--group-001.md#canonical-3232220003100230-2233132333122323-2111102320210222-3303122302023023-2202110223002013-2020231223010123-2130103112031123-1200232100031113) |
| `endpoint.any` | [endpoint.any](data-sources--network_policy--reference--group-001.md#canonical-3232013012203212-0012301022233020-0331311100033233-0213031121330010-2323111012122203-1033111101213223-2220030101112020-2023331321202101) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-2001020302322200-1010103321113112-3112121031023132-3231203301323213-2000120003001001-1202031323212322-0110211113003230-1010123033200101) |
| `endpoint.label_selector` | [endpoint.label_selector](data-sources--network_policy--reference--group-001.md#canonical-1301132102300103-2123120313033300-1003332102202222-3113323221000302-1210000300303021-2021230321331101-1233211223002331-1103322001233021) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-1223230313032311-2302132021301133-2322312203200213-3002223133123102-2123302000221303-3032133122333301-2121311232020230-3023310200311020) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3133113330023020-2100031101132202-3010021100033300-3111020301102021-0200031233111032-3100021301123322-2223032303120220-3300200011131001) |
| `endpoint.prefix_list` | [endpoint.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2102211120313010-3120011321130331-1331213103131220-0222321002230321-0010331133101212-3001003123012003-3121332001200022-3002111133000001) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-3101221120321001-0122330302022033-1233222211101312-0201303111202001-0222001111322002-2111232221012203-3221020112111310-2112231333331233) |
| `id` | [id](data-sources--network_policy--reference--group-001.md#canonical-3010313121111023-3010002322102120-3012121023023002-1001102123303001-3313323113033123-0113133303110012-3021022123221021-0201010212012302) |
| `labels` | [labels](data-sources--network_policy--reference--group-001.md#canonical-0031302010302201-2111203012303001-1323000230123132-2122333100313000-0302231011122113-3130121300100023-1132013122000223-2123101133231203) |
| `name` | [name](data-sources--network_policy--reference--group-001.md#canonical-0100030011230122-0112322010322030-2022303010110331-2000310321211103-3010032100030122-3311000333021332-2032131301023203-1230022110012211) |
| `namespace` | [namespace](data-sources--network_policy--reference--group-001.md#canonical-0131130223013111-0302321322112311-0232123033313302-0100103132123033-1221311210210232-2200230100100201-3323122123220221-3302032332211002) |
| `rules` | [rules](data-sources--network_policy--reference--group-001.md#canonical-2303303121303110-0101310211032322-1222023320000121-1210231212020330-2230232020101023-1202010211201231-0032331221020233-3133331232203133) |
| `rules.egress_rules` | [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-3020133313300320-1120132210333030-0323010230203201-3301331023201100-2130021001232221-2132033022120220-2123123202120231-1222031030133030) |
| `rules.egress_rules.action` | [rules.egress_rules.action](data-sources--network_policy--reference--group-001.md#canonical-2211000200203321-1120311103312323-2322320130330230-1030000112101332-0123120303220330-1131210213310112-0133230031010200-2021221111101313) |
| `rules.egress_rules.adv_action` | [rules.egress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-2103033332130222-1210133102022212-3313100021201032-3112303310333230-0110120203220110-3032303123302013-3111223022232112-3230021030222203) |
| `rules.egress_rules.adv_action.action` | [rules.egress_rules.adv_action.action](data-sources--network_policy--reference--group-001.md#canonical-3110202203323310-3320032121301113-2031010130200021-0020332230220023-0120301321122320-2101020311210022-2210121203001313-1333333312003320) |
| `rules.egress_rules.all_tcp_traffic` | [rules.egress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-3121113101332301-0003133030122220-0303013222331213-1323331030020110-3011130012232030-2002120111120330-0203320331022303-2012022303223312) |
| `rules.egress_rules.all_traffic` | [rules.egress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3111123103100023-3320012323123233-0033223310010021-3100223223333232-2221322322211200-2233013312311003-3013001212131301-2021210200123312) |
| `rules.egress_rules.all_udp_traffic` | [rules.egress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-2023230001111100-1302123101200011-0123023113211303-1333021321231323-2130122221023223-2202131312321331-0201230233332023-0301320320131133) |
| `rules.egress_rules.any` | [rules.egress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-3320323202210033-2113023021232103-1300103320121100-2011332122023000-3232201233312320-3323100031311310-2220110332330100-0130023023002311) |
| `rules.egress_rules.applications` | [rules.egress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-2213000122211303-1020102221220011-1300013010103300-2322133010022221-1123321231012330-2211021220023101-3212203223211330-2211201323000010) |
| `rules.egress_rules.applications.applications` | [rules.egress_rules.applications.applications](data-sources--network_policy--reference--group-001.md#canonical-0122221213201100-0221122333120100-3011100011011022-0012232132011223-3031013023300223-1322222210221100-2003122022121100-0330322231113223) |
| `rules.egress_rules.inside_endpoints` | [rules.egress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-0121132120313202-0212213021111332-1132303110230102-3212021102222310-2330210010323323-2132020033132301-3010111122023013-1230300132210211) |
| `rules.egress_rules.ip_prefix_set` | [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1232200103312122-0322211310033332-1110311312122232-1233131332313311-1131112002330131-1220200221132031-3123222012201130-1303131300120331) |
| `rules.egress_rules.ip_prefix_set.ref` | [rules.egress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-0110233223001103-3222301020211203-1123110121102332-1000121002131322-0312102113201310-3020300331121220-2013323111213022-0323011001220223) |
| `rules.egress_rules.ip_prefix_set.ref.kind` | [rules.egress_rules.ip_prefix_set.ref.kind](data-sources--network_policy--reference--group-001.md#canonical-1130210203103110-3203033112000110-1320311331310003-3000102211300022-2011100313030022-3032031300033122-3023312022012232-2311310130030113) |
| `rules.egress_rules.ip_prefix_set.ref.name` | [rules.egress_rules.ip_prefix_set.ref.name](data-sources--network_policy--reference--group-001.md#canonical-3110112012000223-3032022302120001-0010033230121121-1133021203020310-0222012120010132-1133130033002121-1000220232022020-3031122003332013) |
| `rules.egress_rules.ip_prefix_set.ref.namespace` | [rules.egress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy--reference--group-001.md#canonical-0120222311332133-1132301300132211-2100202202002010-0313220312120130-0123232113323100-1101021333210030-2111203033120313-3112132310032201) |
| `rules.egress_rules.ip_prefix_set.ref.tenant` | [rules.egress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy--reference--group-001.md#canonical-2211031302121233-0322022112102112-1123022112212212-1131221112122311-2231312211210223-1032022230022211-1302233220223122-2211102103223213) |
| `rules.egress_rules.ip_prefix_set.ref.uid` | [rules.egress_rules.ip_prefix_set.ref.uid](data-sources--network_policy--reference--group-001.md#canonical-1012301202030122-2012133301220033-1331010010230121-2201031110211131-2312201330123233-3103300022331100-3022013131030010-1303110112033301) |
| `rules.egress_rules.label_matcher` | [rules.egress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-3133021322311233-3111032130300022-3101312320113102-3211030033133032-0332301320301222-2202111012333032-2220010300131123-2322100312111132) |
| `rules.egress_rules.label_matcher.keys` | [rules.egress_rules.label_matcher.keys](data-sources--network_policy--reference--group-001.md#canonical-0000123302320331-0332330300210003-0221303133032233-1103011220121202-0322332331010130-1220002203101321-3002233211003111-2312110311021032) |
| `rules.egress_rules.label_selector` | [rules.egress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-1113212230023132-3131020111023100-3132101230121312-2021232133303031-1330303321033130-0111022333102213-1021302100221301-2332322122103330) |
| `rules.egress_rules.label_selector.expressions` | [rules.egress_rules.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-3003100200200331-1123122003023112-2021222012100121-2202200013320030-1333323120111020-2312301003111323-0033133201101123-0102111212132132) |
| `rules.egress_rules.metadata` | [rules.egress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-3233012220112130-2132322302333113-1211310102020232-0211111331000310-2110110103311303-3031013301013102-1333101221120133-3202100110122132) |
| `rules.egress_rules.metadata.description_spec` | [rules.egress_rules.metadata.description_spec](data-sources--network_policy--reference--group-001.md#canonical-1102120313333230-3030320310322201-0021200033233120-3213200302332202-3332023220232323-0000203321303221-1112210101321322-1022332023301011) |
| `rules.egress_rules.metadata.name` | [rules.egress_rules.metadata.name](data-sources--network_policy--reference--group-001.md#canonical-0120001032103320-3111301212111002-2003113302322323-0232121232202112-0211123101213101-1110212300322221-2100210112212233-2001213032012112) |
| `rules.egress_rules.outside_endpoints` | [rules.egress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-0011311033130012-2030102120321312-1000210002100320-0300211301232333-3213002023212101-3021200003322112-0321010332212001-0011310003230100) |
| `rules.egress_rules.prefix_list` | [rules.egress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-0312301331233210-3233131200000103-0321122233110320-0313003233310202-3100000010000103-0001222023221212-3001200212022323-1011031312011013) |
| `rules.egress_rules.prefix_list.prefixes` | [rules.egress_rules.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-1333021121022222-1100323223230113-2302210021201302-0131110221100003-0011331123311113-0210003112010301-2033320213231110-2213031122223212) |
| `rules.egress_rules.protocol_port_range` | [rules.egress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-1110032130122230-2220300320323132-0203303113021330-1301322031212301-2032230101030322-0012313001122032-2112123000203203-1111030011202220) |
| `rules.egress_rules.protocol_port_range.port_ranges` | [rules.egress_rules.protocol_port_range.port_ranges](data-sources--network_policy--reference--group-001.md#canonical-2331210022220002-3332230120222130-2002121121210010-3123100121133020-2202133032331000-2231300013031032-0202203001120322-3110201232221212) |
| `rules.egress_rules.protocol_port_range.protocol` | [rules.egress_rules.protocol_port_range.protocol](data-sources--network_policy--reference--group-001.md#canonical-0000021313033202-2303003202331021-1201133011013023-2130111233230211-3333133212011100-0122300100032121-1003011322212233-2020320012201211) |
| `rules.ingress_rules` | [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0222111101121101-0021210133123033-0220210221001200-2130100211120323-0133130300200002-2312021023203113-0001222213233201-1333203222222212) |
| `rules.ingress_rules.action` | [rules.ingress_rules.action](data-sources--network_policy--reference--group-001.md#canonical-2112323133112111-0233132032013323-3031330303013011-3222131102331323-2113210231131111-2031000112331212-2013310133103012-2101032322322131) |
| `rules.ingress_rules.adv_action` | [rules.ingress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-3211031101300331-3332123102120111-2211303223330131-0323121332110020-2002113001132212-0213100223300031-2323003011303310-3331120303132000) |
| `rules.ingress_rules.adv_action.action` | [rules.ingress_rules.adv_action.action](data-sources--network_policy--reference--group-001.md#canonical-2333310112210211-2313110322222331-2333330013111033-1131310232203332-2212111131002233-1101331213003023-2212230320003312-2331110333101010) |
| `rules.ingress_rules.all_tcp_traffic` | [rules.ingress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-2113232303110312-1311000103312012-3031013230302313-0230100321301321-2201103131100222-0221322133223002-2303021210022333-1332300011112101) |
| `rules.ingress_rules.all_traffic` | [rules.ingress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-0323331202002230-3002002202210123-1032121122320103-1013101102333032-3312020333002212-3132130121231200-1303103230012323-0010333233011322) |
| `rules.ingress_rules.all_udp_traffic` | [rules.ingress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-2130332111332100-0313302310002203-0202230233102331-0032102202212332-2111201201201323-2320103220232120-3212120301021232-2231231133213202) |
| `rules.ingress_rules.any` | [rules.ingress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-0023103301121211-2001213000212311-2303002320213213-1023210310211203-2331222232111321-1230102030002201-2033331010100210-0330200330031122) |
| `rules.ingress_rules.applications` | [rules.ingress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-0231333211022302-2300311103303130-1010102031233121-3112100310321222-2230103002322212-3301231101332310-2013203313003000-2002103131122002) |
| `rules.ingress_rules.applications.applications` | [rules.ingress_rules.applications.applications](data-sources--network_policy--reference--group-001.md#canonical-3032003223231231-2012212103333012-1302322302230310-2111021032311321-1100302101232101-3012011023323220-0201223312103302-3212233033201200) |
| `rules.ingress_rules.inside_endpoints` | [rules.ingress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3032031302223030-0330220320221012-0130231203121023-3331203200132223-1131300033311030-3003002310331003-3121232132233110-0322322100010100) |
| `rules.ingress_rules.ip_prefix_set` | [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1031131221033202-2002131303013320-2000211003321323-3020123202220020-3113333002110123-1100012322320332-0132120332330213-1223322223220212) |
| `rules.ingress_rules.ip_prefix_set.ref` | [rules.ingress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-0103103231202110-2123112220130331-3121112030310332-3211233133111233-3121101112101130-1112022011200123-0103322101113200-1201011330300111) |
| `rules.ingress_rules.ip_prefix_set.ref.kind` | [rules.ingress_rules.ip_prefix_set.ref.kind](data-sources--network_policy--reference--group-001.md#canonical-0211003221312310-0130102212332011-0130121330130020-0122113223220203-1100202332221311-2321011133321112-3021221001101233-0030032023113203) |
| `rules.ingress_rules.ip_prefix_set.ref.name` | [rules.ingress_rules.ip_prefix_set.ref.name](data-sources--network_policy--reference--group-001.md#canonical-0233003103101221-2321020023221100-1030101203001230-2220201200000310-3321102022210033-1113300210300110-1231232133300203-1223313322210230) |
| `rules.ingress_rules.ip_prefix_set.ref.namespace` | [rules.ingress_rules.ip_prefix_set.ref.namespace](data-sources--network_policy--reference--group-001.md#canonical-2010103211001311-3220211233102312-3321310102131010-3231322003033331-1330103021233332-1012011220223230-1110003123331100-3100310310212222) |
| `rules.ingress_rules.ip_prefix_set.ref.tenant` | [rules.ingress_rules.ip_prefix_set.ref.tenant](data-sources--network_policy--reference--group-001.md#canonical-2301220333302331-0133000121122130-1300023203312313-0120121201013033-3212022331323312-1033311111010103-2133032031213331-2201001013333133) |
| `rules.ingress_rules.ip_prefix_set.ref.uid` | [rules.ingress_rules.ip_prefix_set.ref.uid](data-sources--network_policy--reference--group-001.md#canonical-2121333232030000-3022333312303311-3330000013111323-2332113222331231-3222230333131213-1122222310202220-0012032213013021-2233221313113002) |
| `rules.ingress_rules.label_matcher` | [rules.ingress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-1013313033123130-3110010312111122-0021020103031231-3223100031233210-2000122222103320-1012110011323213-0310010302021020-0301031122102210) |
| `rules.ingress_rules.label_matcher.keys` | [rules.ingress_rules.label_matcher.keys](data-sources--network_policy--reference--group-001.md#canonical-3230121322321222-0301112333301211-1130011210132332-3232103220102332-1110102002111331-3102001321123021-1003212112230023-3220202311000331) |
| `rules.ingress_rules.label_selector` | [rules.ingress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-1132021230211323-3020113122133100-0220132001102020-3230230002131012-2122232103132321-1322210213320002-3203233222223230-0231123023011202) |
| `rules.ingress_rules.label_selector.expressions` | [rules.ingress_rules.label_selector.expressions](data-sources--network_policy--reference--group-001.md#canonical-3101201201333233-1223222033102023-1132202113310303-0232112302020030-1122313331001011-0232111323223023-1231223131113023-0323330300213333) |
| `rules.ingress_rules.metadata` | [rules.ingress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-0203122132002323-1033321023130211-2012310300213210-2203320312312101-2110310020333121-2110123303233311-1213210233120202-0233011113121010) |
| `rules.ingress_rules.metadata.description_spec` | [rules.ingress_rules.metadata.description_spec](data-sources--network_policy--reference--group-001.md#canonical-1011021321313130-1312303012000202-2002131131023111-1312223002022100-3102223021331131-0121003110122130-3013333010013100-1210332133130031) |
| `rules.ingress_rules.metadata.name` | [rules.ingress_rules.metadata.name](data-sources--network_policy--reference--group-001.md#canonical-0002123310220030-2110112001133231-2330003021330201-0110231220131012-1133000111022011-1310113133032310-3022213111023032-1023330213210330) |
| `rules.ingress_rules.outside_endpoints` | [rules.ingress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-1022223220201130-2101020202101210-2123211313300312-1233311022313030-0221222231223321-3303330013211012-1321113122103023-1103221023220322) |
| `rules.ingress_rules.prefix_list` | [rules.ingress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2221020203121002-0000123112313031-0320220002130321-2213030002100321-0202123003222112-2102232132033333-1112001032103013-0111333300222012) |
| `rules.ingress_rules.prefix_list.prefixes` | [rules.ingress_rules.prefix_list.prefixes](data-sources--network_policy--reference--group-001.md#canonical-2322100301021103-2211122033201233-2221113230102222-0102110202132231-1031110211331121-1031001232203113-2032110033331023-3330103321320223) |
| `rules.ingress_rules.protocol_port_range` | [rules.ingress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-1301231020022111-3101121231313233-1230130301021331-3313223032103110-1301023133112331-3103231000200020-0110032333031112-2223120320211201) |
| `rules.ingress_rules.protocol_port_range.port_ranges` | [rules.ingress_rules.protocol_port_range.port_ranges](data-sources--network_policy--reference--group-001.md#canonical-2010201233323233-1011213222103213-1233001101232321-2312311302312203-1213310022230301-0032012232311130-1012113323112010-3201213222012012) |
| `rules.ingress_rules.protocol_port_range.protocol` | [rules.ingress_rules.protocol_port_range.protocol](data-sources--network_policy--reference--group-001.md#canonical-2003131001231131-0200002323202133-0211223120022020-1013030010011121-3310123100221020-1233032210333030-3032112323301023-2113330223313010) |

<a id="canonical-0033310333232110-1310300111233223-1110121113313311-0131322112020310-2003011313001112-1323122201333303-3233202303000123-3030321203023110"></a>

## Next pages — Property reference / 322213232030 / 11

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033213311230303-0130001101323113-1013030230202103-2330002213122123-0110311230330333-1222221120132002-3122120130323222-3203100300112221"></a>

## endpoint — endpoint / 323023130331 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- endpoint

<a id="canonical-3232220003100230-2233132333122323-2111102320210222-3303122302023023-2202110223002013-2020231223010123-2130103112031123-1200232100031113"></a>

Type: `"single"`. Computed.

Shape of the endpoint choices for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

<a id="canonical-2201112313333120-0012101321203113-0130323113203331-2101222211121323-2103330131003133-2323203203311322-2022333313100332-2021323112321232"></a>

## Direct properties — endpoint / 323023130331 / 3

- [any](data-sources--network_policy--reference--group-001.md#canonical-1100220321311001-3323101312111330-3202110012132120-2212023133302233-1120133211023030-2030301332013000-0103001303102132-2101222022130121): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3010321223011300-2003110322110112-2220322030312232-0122121302033332-0302031021030102-3301011100331001-0230130322023320-2233321220111231): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-2030203133002310-0331033121120000-0010033231200120-3321223011112123-0212231112310300-2233332001031021-1302001122201002-2032021302300031): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-2002001313202100-1000232220202113-1213103333113330-3113001310320312-0102232203011001-2103020323202311-1331200220122203-2331301201032301): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2011001230022323-0321200100010203-1331302202002133-2202131121002113-0131313323202121-3303032303001000-2223331323303231-3121132332300331): complete subsection reference.

<a id="canonical-3333331113002003-1300123133223301-2322001133101030-2200322131333333-3323301023231222-0032301201221000-1022111310232331-1221030000121212"></a>

## Next pages — endpoint / 323023130331 / 4

- [endpoint.any](data-sources--network_policy--reference--group-001.md#canonical-1100220321311001-3323101312111330-3202110012132120-2212023133302233-1120133211023030-2030301332013000-0103001303102132-2101222022130121)
- [endpoint.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3010321223011300-2003110322110112-2220322030312232-0122121302033332-0302031021030102-3301011100331001-0230130322023320-2233321220111231)
- [endpoint.label_selector](data-sources--network_policy--reference--group-001.md#canonical-2030203133002310-0331033121120000-0010033231200120-3321223011112123-0212231112310300-2233332001031021-1302001122201002-2032021302300031)
- [endpoint.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-2002001313202100-1000232220202113-1213103333113330-3113001310320312-0102232203011001-2103020323202311-1331200220122203-2331301201032301)
- [endpoint.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2011001230022323-0321200100010203-1331302202002133-2202131121002113-0131313323202121-3303032303001000-2223331323303231-3121132332300331)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1100220321311001-3323101312111330-3202110012132120-2212023133302233-1120133211023030-2030301332013000-0103001303102132-2101222022130121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311000333001300-1230010011002001-0001201323103221-2002102100023223-1131301123132303-3203312100020111-1333303301023233-2132131001130111"></a>

## endpoint.any — any / 200131021131 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- endpoint.any

<a id="canonical-3232013012203212-0012301022233020-0331311100033233-0213031121330010-2323111012122203-1033111101213223-2220030101112020-2023331321202101"></a>

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

<a id="canonical-1300010303310102-3102011132021103-2002311121131133-3003232330033032-2301123110232031-2333130100012331-3312112120201331-0220220232220030"></a>

## Direct properties — any / 200131021131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121330031201112-3321323001121031-1333201201103132-2332132020112321-3002212201231113-0003031333031003-1200101333231203-3001101300120031"></a>

## Next pages — any / 200131021131 / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3010321223011300-2003110322110112-2220322030312232-0122121302033332-0302031021030102-3301011100331001-0230130322023320-2233321220111231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013312221133013-0102113333122203-2210013313310130-1310000103001212-0303133312301311-3003110203122233-0203023030130201-3233331011011302"></a>

## endpoint.inside_endpoints — inside_endpoints / 333032300200 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- endpoint.inside_endpoints

<a id="canonical-2001020302322200-1010103321113112-3112121031023132-3231203301323213-2000120003001001-1202031323212322-0110211113003230-1010123033200101"></a>

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

<a id="canonical-2010321001133003-0130231310233000-2212212310132031-2230223333121212-2120133010000020-0121311121320133-1301302131103211-2321021333031120"></a>

## Direct properties — inside_endpoints / 333032300200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312020320020313-2112032020313322-0002123203100313-0320103122301102-3121201022003112-1020232310001101-1330312202213021-0322111223010011"></a>

## Next pages — inside_endpoints / 333032300200 / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2030203133002310-0331033121120000-0010033231200120-3321223011112123-0212231112310300-2233332001031021-1302001122201002-2032021302300031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113203032210320-0001012131213122-1113111123311231-3121123312132310-1102312313030310-2323031212101103-2331303121022233-3321212120231312"></a>

## endpoint.label_selector — label_selector / 203301032211 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- endpoint.label_selector

<a id="canonical-1301132102300103-2123120313033300-1003332102202222-3113323221000302-1210000300303021-2021230321331101-1233211223002331-1103322001233021"></a>

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

<a id="canonical-3000202123130311-1323220020020100-1332031310002032-3230200323031020-2211303230220112-0312303123002202-2023102302012130-1110203302010002"></a>

## Direct properties — label_selector / 203301032211 / 3

<a id="canonical-1223230313032311-2302132021301133-2322312203200213-3002223133123102-2123302000221303-3032133122333301-2121311232020230-3023310200311020"></a>

<a id="canonical-1223312212113120-3112201002310010-0012222300133111-3213231333023310-0102013111123310-0012020123020123-0023132203022300-0130310022223131"></a>

## expressions property — label_selector / 203301032211 / 4

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

<a id="canonical-2320112310203232-0212220032003000-0312132322201010-2131311330323020-2013121130223332-2332031222022321-2000311101202323-1212100321023113"></a>

## Next pages — label_selector / 203301032211 / 5

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2002001313202100-1000232220202113-1213103333113330-3113001310320312-0102232203011001-2103020323202311-1331200220122203-2331301201032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310310122220013-1010332330310032-1223200122303302-2302221222301322-3223031130033013-2300130031321313-0223111313013303-1311213201011113"></a>

## endpoint.outside_endpoints — outside_endpoints / 002232030001 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- endpoint.outside_endpoints

<a id="canonical-3133113330023020-2100031101132202-3010021100033300-3111020301102021-0200031233111032-3100021301123322-2223032303120220-3300200011131001"></a>

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

<a id="canonical-3210203133000131-3103101132032220-3302300120011001-3210320210122202-1210220321211120-0003003032322132-3010021010233303-0020222002310213"></a>

## Direct properties — outside_endpoints / 002232030001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302030112020232-0212231031011022-1201130300111120-0321023223013312-0322003033313310-3210100010310133-1322200031232101-1022331111231220"></a>

## Next pages — outside_endpoints / 002232030001 / 4

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2011001230022323-0321200100010203-1331302202002133-2202131121002113-0131313323202121-3303032303001000-2223331323303231-3121132332300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121213332112000-0131302020233132-2222200133322233-2231303112020131-3020211333121022-3121301323212322-2301231012103221-1021100320023031"></a>

## endpoint.prefix_list — prefix_list / 332212321110 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- endpoint.prefix_list

<a id="canonical-2102211120313010-3120011321130331-1331213103131220-0222321002230321-0010331133101212-3001003123012003-3121332001200022-3002111133000001"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0113022032002121-2021202101120033-0102233022212122-0002300123320121-3320013121200020-1030122231022210-2311203320003022-2233223131120213"></a>

## Direct properties — prefix_list / 332212321110 / 3

<a id="canonical-3101221120321001-0122330302022033-1233222211101312-0201303111202001-0222001111322002-2111232221012203-3221020112111310-2112231333331233"></a>

<a id="canonical-2111321130230110-0033010223200023-0101322320113211-0211333122223010-2330122322303310-0111103223111011-1031330011221101-2300331213313311"></a>

## prefixes property — prefix_list / 332212321110 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1203323320222320-3021221311220100-2130303132030203-0132301130230111-1033012003220022-0212313213201012-0212011301031033-1332233312032331"></a>

## Next pages — prefix_list / 332212321110 / 5

- [endpoint](data-sources--network_policy--reference--group-001.md#canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310020231123221-3020203231322021-3112213320013133-2110322221103113-1332020102012231-2212023321213120-2131111131210021-0122333311223302"></a>

## rules — rules / 230110122301 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- rules

<a id="canonical-2303303121303110-0101310211032322-1222023320000121-1210231212020330-2230232020101023-1202010211201231-0032331221020233-3133331232203133"></a>

Type: `"single"`. Computed.

Rule Choice. Shape of Rule Choice.

Upstream description:

Shape of Rule Choice.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2223232001013132-3203313332102001-1123220233233212-1010132233313100-0232301230230033-1312302102323211-1330203022333313-1323331331311101"></a>

## Direct properties — rules / 230110122301 / 3

- [egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323): complete subsection reference.

- [ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202): complete subsection reference.

<a id="canonical-2132311222223202-2203221331202001-0121102110212132-0002123001320013-3301310112000100-2312123120232011-3323202103220310-0133211303230200"></a>

## Next pages — rules / 230110122301 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222221333223101-0030021212210003-1120111323030331-3120011132033022-1313233022001313-2112300102322220-1232330112233102-2320012123033111"></a>

## rules.egress_rules — egress_rules / 203022012121 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- rules.egress_rules

<a id="canonical-3020133313300320-1120132210333030-0323010230203201-3301331023201100-2130021001232221-2132033022120220-2123123202120231-1222031030133030"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections from policy endpoints.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0013100320212001-1032123121011000-2203133233311013-2201100300031321-1201012333333222-3013301313322012-3330123100303303-1313323232011010"></a>

## Direct properties — egress_rules / 203022012121 / 3

<a id="canonical-2211000200203321-1120311103312323-2322320130330230-1030000112101332-0123120303220330-1131210213310112-0133230031010200-2021221111101313"></a>

<a id="canonical-0003200203333123-3000302120123122-3333010003310002-1001303213033033-1033302202020002-0100020330122322-3213321221310133-0232033301222000"></a>

## action property — egress_rules / 203022012121 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy--reference--group-001.md#canonical-0132021322103113-2023200012110000-3032122311021022-0201233103322113-2012022000321220-0213110231030322-0133321230100323-3121031221310012): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-3030203003100013-2211231131101323-3011222033200131-0112310230320233-1233202321011010-3030332023102310-1012311233210310-3101210103013231): complete subsection reference.

- [all_traffic](data-sources--network_policy--reference--group-001.md#canonical-0331231232133311-2322313021102030-2030021003323113-1023022110300002-3003320320030333-0110113112001203-0023003220003212-0331210111021331): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-0022112302302331-1133222210133011-0221030101201323-3300022220202023-2011313121211202-1202031113131011-1332302122212131-0303311132301312): complete subsection reference.

- [any](data-sources--network_policy--reference--group-001.md#canonical-1200130113330200-1130022300301120-3300200320122220-0113201121131133-1200021130011311-1220012100333322-3332020332301313-0221032013030301): complete subsection reference.

- [applications](data-sources--network_policy--reference--group-001.md#canonical-3102123131233031-2120101113230113-0020021220223203-0301330100002131-3222120133321221-0220010023032013-0331332323310202-3320002100220031): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3210322030011222-0213323032133313-0023203331202331-2112021230220110-2120000223202200-2223213212231220-2000203303112101-1030200103333323): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220): complete subsection reference.

- [label_matcher](data-sources--network_policy--reference--group-001.md#canonical-3103120202331220-3020222320021103-3231231302132201-0221111120010303-3201333021311002-2122002331032200-0321021001230303-3230301111333210): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-0133213113103123-1131333122110211-3113132022322331-1022023230120130-1101302321213313-2003013202322202-0322202232202123-1123220011323330): complete subsection reference.

- [metadata](data-sources--network_policy--reference--group-001.md#canonical-0310100221210010-1101220102211012-3033202320032101-0012113013303332-1230203310113230-1123303111103211-3112030102012300-3331110003130322): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-0020033132321010-1030031023032103-3022000321110302-2210112133130031-2122122030011310-1221110312031121-2201122121213000-3331203302131030): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2100022021230120-0133323031302332-0300232221000131-2221233300230323-1033312203101331-1302000011220012-0221202203122033-0030001212302100): complete subsection reference.

- [protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-0031302222203212-0111321312131102-3210003232110220-0100003032332011-1022301113211323-2223201223212203-3111111001002303-2222323133311003): complete subsection reference.

<a id="canonical-2300020232321122-2131330020211203-0021103203131030-0010103030302023-3303210010320003-1101011201021123-2013210212311011-3313311111233001"></a>

## Next pages — egress_rules / 203022012121 / 5

- [rules.egress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-0132021322103113-2023200012110000-3032122311021022-0201233103322113-2012022000321220-0213110231030322-0133321230100323-3121031221310012)
- [rules.egress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-3030203003100013-2211231131101323-3011222033200131-0112310230320233-1233202321011010-3030332023102310-1012311233210310-3101210103013231)
- [rules.egress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-0331231232133311-2322313021102030-2030021003323113-1023022110300002-3003320320030333-0110113112001203-0023003220003212-0331210111021331)
- [rules.egress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-0022112302302331-1133222210133011-0221030101201323-3300022220202023-2011313121211202-1202031113131011-1332302122212131-0303311132301312)
- [rules.egress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-1200130113330200-1130022300301120-3300200320122220-0113201121131133-1200021130011311-1220012100333322-3332020332301313-0221032013030301)
- [rules.egress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-3102123131233031-2120101113230113-0020021220223203-0301330100002131-3222120133321221-0220010023032013-0331332323310202-3320002100220031)
- [rules.egress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-3210322030011222-0213323032133313-0023203331202331-2112021230220110-2120000223202200-2223213212231220-2000203303112101-1030200103333323)
- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220)
- [rules.egress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-3103120202331220-3020222320021103-3231231302132201-0221111120010303-3201333021311002-2122002331032200-0321021001230303-3230301111333210)
- [rules.egress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-0133213113103123-1131333122110211-3113132022322331-1022023230120130-1101302321213313-2003013202322202-0322202232202123-1123220011323330)
- [rules.egress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-0310100221210010-1101220102211012-3033202320032101-0012113013303332-1230203310113230-1123303111103211-3112030102012300-3331110003130322)
- [rules.egress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-0020033132321010-1030031023032103-3022000321110302-2210112133130031-2122122030011310-1221110312031121-2201122121213000-3331203302131030)
- [rules.egress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-2100022021230120-0133323031302332-0300232221000131-2221233300230323-1033312203101331-1302000011220012-0221202203122033-0030001212302100)
- [rules.egress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-0031302222203212-0111321312131102-3210003232110220-0100003032332011-1022301113211323-2223201223212203-3111111001002303-2222323133311003)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0132021322103113-2023200012110000-3032122311021022-0201233103322113-2012022000321220-0213110231030322-0133321230100323-3121031221310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122100330330111-1120211331320330-0110300202033301-3111112102122120-2022220312232023-3303121202210003-3121112123002133-1011131211123323"></a>

## rules.egress_rules.adv_action — adv_action / 201133033321 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.adv_action

<a id="canonical-2103033332130222-1210133102022212-3313100021201032-3112303310333230-0110120203220110-3032303123302013-3111223022232112-3230021030222203"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301233223303220-0220203110230031-0030030130320211-0303111222000312-0031312302000031-0312232101110003-1023003211122133-2011020102033131"></a>

## Direct properties — adv_action / 201133033321 / 3

<a id="canonical-3110202203323310-3320032121301113-2031010130200021-0020332230220023-0120301321122320-2101020311210022-2210121203001313-1333333312003320"></a>

<a id="canonical-2011132233031133-2200011201103031-1201122231110301-0323003212331101-0301213200222100-2132130220111221-2311321323102300-3202130211022232"></a>

## action property — adv_action / 201133033321 / 4

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1022032013232302-2121122333232322-0033212223003233-0221011002321202-1310323320231211-0201313211222323-2323221030300331-3330010030303301"></a>

## Next pages — adv_action / 201133033321 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3030203003100013-2211231131101323-3011222033200131-0112310230320233-1233202321011010-3030332023102310-1012311233210310-3101210103013231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320333001302301-0203113100320312-1311002123011121-0120323320221233-3133001210033313-1333032011311121-2332201103123102-1120332011032321"></a>

## rules.egress_rules.all_tcp_traffic — all_tcp_traffic / 023230300103 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.all_tcp_traffic

<a id="canonical-3121113101332301-0003133030122220-0303013222331213-1323331030020110-3011130012232030-2002120111120330-0203320331022303-2012022303223312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-2302020103321302-1102010002310201-1010123232130133-3232213303110223-3313111102122122-0130120313233330-0302131021001100-2003133013210020"></a>

## Direct properties — all_tcp_traffic / 023230300103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122130233133203-1023203121300001-2210221120212222-3321031210012103-0202231233232212-0300000211320312-3232132111200221-0303032102101002"></a>

## Next pages — all_tcp_traffic / 023230300103 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0331231232133311-2322313021102030-2030021003323113-1023022110300002-3003320320030333-0110113112001203-0023003220003212-0331210111021331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220121320312001-1033112300110330-1021130332332221-2120220321021322-2213312300321323-1130123130012200-1131200111303021-3023210331231213"></a>

## rules.egress_rules.all_traffic — all_traffic / 133000303320 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.all_traffic

<a id="canonical-3111123103100023-3320012323123233-0033223310010021-3100223223333232-2221322322211200-2233013312311003-3013001212131301-2021210200123312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-0000201103111303-1213201300201102-2002303221200122-3120012020131220-2321230121202322-3122131222330021-1232300302130223-1101332001113220"></a>

## Direct properties — all_traffic / 133000303320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033123120122312-0123130030213323-0211033211000211-3323012231130123-3033132323133011-0113010232121012-3321310233102221-3112301302030300"></a>

## Next pages — all_traffic / 133000303320 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0022112302302331-1133222210133011-0221030101201323-3300022220202023-2011313121211202-1202031113131011-1332302122212131-0303311132301312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301323222331133-1301032211111032-3010011302322201-3231332203130000-1121130221221101-3310112321122001-2300301301303211-1013233032301101"></a>

## rules.egress_rules.all_udp_traffic — all_udp_traffic / 221233202100 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.all_udp_traffic

<a id="canonical-2023230001111100-1302123101200011-0123023113211303-1333021321231323-2130122221023223-2202131312321331-0201230233332023-0301320320131133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-2332203302323021-2001121131211102-0301301233213130-2310102201131012-2331031120310311-0132211032332300-1302011003233232-3300210121123300"></a>

## Direct properties — all_udp_traffic / 221233202100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323020110210023-1030032001112130-3202301032031111-3202110213223030-3031201032233123-2111230213303313-1311333311110012-0021110202121302"></a>

## Next pages — all_udp_traffic / 221233202100 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1200130113330200-1130022300301120-3300200320122220-0113201121131133-1200021130011311-1220012100333322-3332020332301313-0221032013030301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113031221211220-3200102000200311-2033033110122210-1322101212230321-0302313123210312-1123202331312102-2301221101330130-3032312212330012"></a>

## rules.egress_rules.any — any / 223001110332 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.any

<a id="canonical-3320323202210033-2113023021232103-1300103320121100-2011332122023000-3232201233312320-3323100031311310-2220110332330100-0130023023002311"></a>

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

<a id="canonical-1013013123112232-2101031203233323-1220223000201130-1020011100012101-1133203002230120-2110110301303312-2021212131103032-0112101202233333"></a>

## Direct properties — any / 223001110332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021331312212130-1123301302321133-3301211101313303-1000211211130202-1322210111122311-3333302301030130-3231232003203203-0212013130202230"></a>

## Next pages — any / 223001110332 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3102123131233031-2120101113230113-0020021220223203-0301330100002131-3222120133321221-0220010023032013-0331332323310202-3320002100220031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001203222032300-0232201230120312-1023202210231301-0120103313102323-2100221200323232-0201000023312103-1131010012132132-2101020203303011"></a>

## rules.egress_rules.applications — applications / 001212231223 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.applications

<a id="canonical-2213000122211303-1020102221220011-1300013010103300-2322133010022221-1123321231012330-2211021220023101-3212203223211330-2211201323000010"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233320333101230-2312122033023330-0020221013233202-3132132130213103-3210012312200030-2101103113021002-0033230213233120-2002311203000222"></a>

## Direct properties — applications / 001212231223 / 3

<a id="canonical-0122221213201100-0221122333120100-3011100011011022-0012232132011223-3031013023300223-1322222210221100-2003122022121100-0330322231113223"></a>

<a id="canonical-1222121203120302-2102100132310100-1202102232130200-0122003312230302-1033221321323100-2232201110031001-0232020203020222-2111220012333122"></a>

## applications property — applications / 001212231223 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103203302010222-1133232320020331-0323212112033213-3122133102201202-3031003330311123-2321101133120200-2031112111331333-0223211030333120"></a>

## Next pages — applications / 001212231223 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3210322030011222-0213323032133313-0023203331202331-2112021230220110-2120000223202200-2223213212231220-2000203303112101-1030200103333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103201013303101-0033013121211102-0321331122300212-1103202213303022-1031332030013110-2120002233320131-1211003201203103-2012121122122201"></a>

## rules.egress_rules.inside_endpoints — inside_endpoints / 330203322233 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.inside_endpoints

<a id="canonical-0121132120313202-0212213021111332-1132303110230102-3212021102222310-2330210010323323-2132020033132301-3010111122023013-1230300132210211"></a>

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

<a id="canonical-2133230313131000-3320032123101133-0200233133130003-3201233100010100-2111121323333011-2321300003020101-0223211020313312-0103121011300000"></a>

## Direct properties — inside_endpoints / 330203322233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330031300001122-3203303131300011-1033211033012102-3222021112302202-0222110102120321-3323233023220131-2033003211023200-3120231011302002"></a>

## Next pages — inside_endpoints / 330203322233 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123131022003213-3100221122302110-3330301232121332-3133233330031322-3023333210212331-1100210111130121-3032330120222300-3133002331230211"></a>

## rules.egress_rules.ip_prefix_set — ip_prefix_set / 100301322113 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.ip_prefix_set

<a id="canonical-1232200103312122-0322211310033332-1110311312122232-1233131332313311-1131112002330131-1220200221132031-3123222012201130-1303131300120331"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011003112100131-2300131312001210-1120002302221022-0101210233333110-0113201013202021-2300112232003100-2023130202203120-3301312301311201"></a>

## Direct properties — ip_prefix_set / 100301322113 / 3

- [ref](data-sources--network_policy--reference--group-001.md#canonical-2222031330202300-3201211330220222-2311300313233132-3132301333003221-3223030010133303-3000023112120301-0300001303120212-0011000033111301): complete subsection reference.

<a id="canonical-3311203233002101-3200300232103200-0032012033122101-1033033222332231-2201123202123033-1201310232013023-2230233320331121-0212121331032130"></a>

## Next pages — ip_prefix_set / 100301322113 / 4

- [rules.egress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-2222031330202300-3201211330220222-2311300313233132-3132301333003221-3223030010133303-3000023112120301-0300001303120212-0011000033111301)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2222031330202300-3201211330220222-2311300313233132-3132301333003221-3223030010133303-3000023112120301-0300001303120212-0011000033111301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011103100323023-1312130033312222-1221312021101310-0023010333220111-0010323202202131-3320002021210130-2012330012001202-0010101032101030"></a>

## rules.egress_rules.ip_prefix_set.ref — ref / 022222100120 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220)
- rules.egress_rules.ip_prefix_set.ref

<a id="canonical-0110233223001103-3222301020211203-1123110121102332-1000121002131322-0312102113201310-3020300331121220-2013323111213022-0323011001220223"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0131300332012303-1320130233302031-2330130010220022-0222102301330311-2032310022032111-3323030233303033-3322231323212200-3313001211333302"></a>

## Direct properties — ref / 022222100120 / 3

<a id="canonical-1130210203103110-3203033112000110-1320311331310003-3000102211300022-2011100313030022-3032031300033122-3023312022012232-2311310130030113"></a>

<a id="canonical-2130003222210232-2111233133223312-2033323331231122-2222212120020110-0121321211023201-3222333231220000-0201102020120320-0021110331330123"></a>

## kind property — ref / 022222100120 / 4

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

<a id="canonical-3110112012000223-3032022302120001-0010033230121121-1133021203020310-0222012120010132-1133130033002121-1000220232022020-3031122003332013"></a>

<a id="canonical-3203111001002132-2113121303313320-0301220310331030-3220213203222200-1323300103123230-0120233002333013-0110130230133330-1201021101221202"></a>

## name property — ref / 022222100120 / 5

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

<a id="canonical-0120222311332133-1132301300132211-2100202202002010-0313220312120130-0123232113323100-1101021333210030-2111203033120313-3112132310032201"></a>

<a id="canonical-1121332331133122-1302113301222030-3010223310122022-3301303213122220-2131023202012003-0303021303121010-2202212331023123-2210322111211020"></a>

## namespace property — ref / 022222100120 / 6

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

<a id="canonical-2211031302121233-0322022112102112-1123022112212212-1131221112122311-2231312211210223-1032022230022211-1302233220223122-2211102103223213"></a>

<a id="canonical-3220231102112221-0300303332312132-0330132203132103-1332011112100102-1032123121330200-2022102000222121-3000212322001301-2003123033322032"></a>

## tenant property — ref / 022222100120 / 7

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

<a id="canonical-1012301202030122-2012133301220033-1331010010230121-2201031110211131-2312201330123233-3103300022331100-3022013131030010-1303110112033301"></a>

<a id="canonical-3033020033203230-1332010232131223-0312031110122210-3012221321130021-0322132333232322-1010110112310331-0303023112112302-0313002333130030"></a>

## uid property — ref / 022222100120 / 8

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

<a id="canonical-1101132031210201-3312300331113023-2031222210303303-3312221000033023-1301221131211003-0300020223131000-2110022121102000-0303231101132032"></a>

## Next pages — ref / 022222100120 / 9

- [rules.egress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-1312312110232031-1203301113033300-2032231212221031-2013232311020211-0320030010121223-3332230221103331-1111333303323003-3320220200210220)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3103120202331220-3020222320021103-3231231302132201-0221111120010303-3201333021311002-2122002331032200-0321021001230303-3230301111333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230120113312030-1313202302330213-3013232111302123-3211320010011201-3333130323103102-0132331321132300-0100003332330313-0310132100210112"></a>

## rules.egress_rules.label_matcher — label_matcher / 220000232130 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.label_matcher

<a id="canonical-3133021322311233-3111032130300022-3101312320113102-3211030033133032-0332301320301222-2202111012333032-2220010300131123-2322100312111132"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221331313231232-0310122232111203-1231123003123301-2033013300133220-3121123231111211-3103200303320021-0213323133232330-3301331222113322"></a>

## Direct properties — label_matcher / 220000232130 / 3

<a id="canonical-0000123302320331-0332330300210003-0221303133032233-1103011220121202-0322332331010130-1220002203101321-3002233211003111-2312110311021032"></a>

<a id="canonical-2221313131322100-0100012102201021-3333302233020131-3213210222010023-0221130110102011-2000223310300332-1022212220320123-1011121331133232"></a>

## keys property — label_matcher / 220000232130 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310010112120111-1023012122012332-1210320102121220-2111121211011011-3112320033103013-2230230301210011-3210323110010013-1322212013122133"></a>

## Next pages — label_matcher / 220000232130 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0133213113103123-1131333122110211-3113132022322331-1022023230120130-1101302321213313-2003013202322202-0322202232202123-1123220011323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131002031122313-2020213101012133-0313010021230300-3312032003003232-1101130212110013-3220131312122000-0301103030332022-2233203100322030"></a>

## rules.egress_rules.label_selector — label_selector / 102221012121 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.label_selector

<a id="canonical-1113212230023132-3131020111023100-3132101230121312-2021232133303031-1330303321033130-0111022333102213-1021302100221301-2332322122103330"></a>

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

<a id="canonical-3312121220020200-0020121300030310-3131210122232130-0310112212132200-0230210310023302-0003200131013013-2333032202131013-0133322203302212"></a>

## Direct properties — label_selector / 102221012121 / 3

<a id="canonical-3003100200200331-1123122003023112-2021222012100121-2202200013320030-1333323120111020-2312301003111323-0033133201101123-0102111212132132"></a>

<a id="canonical-3231222121220212-2302002232122013-3101121021131221-1231032321210323-1010110222303032-2023321332133013-0232330210213111-2132010013213100"></a>

## expressions property — label_selector / 102221012121 / 4

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

<a id="canonical-1321131030103312-0330333012211230-2301003110120301-3030120033310213-0320022133001220-2312023233003310-3200213031113213-0012133013020223"></a>

## Next pages — label_selector / 102221012121 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0310100221210010-1101220102211012-3033202320032101-0012113013303332-1230203310113230-1123303111103211-3112030102012300-3331110003130322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223310133123101-0201311232331011-1302120031200132-1322220002332221-0101133222120103-1001210301322102-3331332200233112-0033013303133321"></a>

## rules.egress_rules.metadata — metadata / 121123001021 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.metadata

<a id="canonical-3233012220112130-2132322302333113-1211310102020232-0211111331000310-2110110103311303-3031013301013102-1333101221120133-3202100110122132"></a>

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

<a id="canonical-2203011323230310-1000332300023023-1020231203223230-3311233221100112-2111310002303000-0131331111300013-1221121332213211-1311110220222102"></a>

## Direct properties — metadata / 121123001021 / 3

<a id="canonical-1102120313333230-3030320310322201-0021200033233120-3213200302332202-3332023220232323-0000203321303221-1112210101321322-1022332023301011"></a>

<a id="canonical-2013202100210000-2020000232130030-1332223130122313-1020203001133300-1022330211133232-0300333200211013-3300312010020232-1000001111303021"></a>

## description_spec property — metadata / 121123001021 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0120001032103320-3111301212111002-2003113302322323-0232121232202112-0211123101213101-1110212300322221-2100210112212233-2001213032012112"></a>

<a id="canonical-1232220202032121-3301212133221302-1300203020110001-3223223200302200-0120312210233322-0112220230230122-0232311222021231-0312120132223030"></a>

## name property — metadata / 121123001021 / 5

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

<a id="canonical-0011132020220202-3131301100233000-3112313033030012-1301231033233101-1223102202332002-0233013210201020-3333222130333322-2322031310031232"></a>

## Next pages — metadata / 121123001021 / 6

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0020033132321010-1030031023032103-3022000321110302-2210112133130031-2122122030011310-1221110312031121-2201122121213000-3331203302131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202121001023302-0023022031200330-1310210321213321-3322200333030312-2220020220011201-3201310111200212-0130003020020323-0120300313032201"></a>

## rules.egress_rules.outside_endpoints — outside_endpoints / 120010021310 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.outside_endpoints

<a id="canonical-0011311033130012-2030102120321312-1000210002100320-0300211301232333-3213002023212101-3021200003322112-0321010332212001-0011310003230100"></a>

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

<a id="canonical-1120120313230200-0032032303320321-3230101330103223-2223331120211000-1111301123130132-2011312231011313-3230203023112213-3231311301200330"></a>

## Direct properties — outside_endpoints / 120010021310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211113011321232-3003102011011211-1230332130020121-3011213122100110-3131222032313023-3203201113302221-1301310221320012-3320100330302203"></a>

## Next pages — outside_endpoints / 120010021310 / 4

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2100022021230120-0133323031302332-0300232221000131-2221233300230323-1033312203101331-1302000011220012-0221202203122033-0030001212302100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020121113002102-3230222231230003-1010333112331210-2113233301023322-2020020111312023-0023131310130201-2033013211130100-2321311101102102"></a>

## rules.egress_rules.prefix_list — prefix_list / 100213013002 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.prefix_list

<a id="canonical-0312301331233210-3233131200000103-0321122233110320-0313003233310202-3100000010000103-0001222023221212-3001200212022323-1011031312011013"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2202300003113133-2012323123022200-1333100103221133-1030122233100302-1001120333221133-0010020333320112-0332333232000000-3300020200113031"></a>

## Direct properties — prefix_list / 100213013002 / 3

<a id="canonical-1333021121022222-1100323223230113-2302210021201302-0131110221100003-0011331123311113-0210003112010301-2033320213231110-2213031122223212"></a>

<a id="canonical-3230120332013103-2232123111212321-0212112021100322-1131221303302332-1300230133223313-1032301102213332-1131332010123113-0200330231203132"></a>

## prefixes property — prefix_list / 100213013002 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3103123133301002-0133210323221301-2112323223020030-3133312201111330-0222023001231311-1122000221022000-2203213213302120-2323331312102322"></a>

## Next pages — prefix_list / 100213013002 / 5

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0031302222203212-0111321312131102-3210003232110220-0100003032332011-1022301113211323-2223201223212203-3111111001002303-2222323133311003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003010023112032-0202330103011313-3003102200012110-1330013021100100-0331030031222313-1000211120331002-2310132311223123-1212201202231031"></a>

## rules.egress_rules.protocol_port_range — protocol_port_range / 211301030220 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- rules.egress_rules.protocol_port_range

<a id="canonical-1110032130122230-2220300320323132-0203303113021330-1301322031212301-2032230101030322-0012313001122032-2112123000203203-1111030011202220"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3110212023121102-2132110121210000-3312213123202231-2023321013023133-1222220101313320-1222321321312132-0131103100002303-1021210100020033"></a>

## Direct properties — protocol_port_range / 211301030220 / 3

<a id="canonical-2331210022220002-3332230120222130-2002121121210010-3123100121133020-2202133032331000-2231300013031032-0202203001120322-3110201232221212"></a>

<a id="canonical-0101233013122133-0023211122201130-1012101002231101-3212201130232211-0102022003332210-1031103323010022-1101322233001210-3112101222211113"></a>

## port_ranges property — protocol_port_range / 211301030220 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-0000021313033202-2303003202331021-1201133011013023-2130111233230211-3333133212011100-0122300100032121-1003011322212233-2020320012201211"></a>

<a id="canonical-0313300331022023-0133323100100302-2031101030220001-2111120000123322-0323222102121312-1110222301230232-2123202230130201-0213233302331003"></a>

## protocol property — protocol_port_range / 211301030220 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-2122132120000002-3220232103122221-1130013011133022-0213233300212232-2210030303220213-1102210021012130-2201011210132313-2122023231321232"></a>

## Next pages — protocol_port_range / 211301030220 / 6

- [rules.egress_rules](data-sources--network_policy--reference--group-001.md#canonical-0022023111312301-0010121303310200-2112103033110321-0303113001310213-1201311310021300-3313331233021133-3010211323032111-1102332013000323)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210220322233333-2111200121201011-0103002020102123-2233010032132221-0312102010020112-0220313102120113-1112223320212002-0123210111212131"></a>

## rules.ingress_rules — ingress_rules / 023220010331 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- rules.ingress_rules

<a id="canonical-0222111101121101-0021210133123033-0220210221001200-2130100211120323-0133130300200002-2312021023203113-0001222213233201-1333203222222212"></a>

Type: `"list"`. Computed.

Ordered list of rules applied to connections to policy endpoints.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0132311303023123-2020022112000303-0132331223310030-2310033222223030-0030000120301200-3123301120101232-2313310203312313-0213020230133131"></a>

## Direct properties — ingress_rules / 023220010331 / 3

<a id="canonical-2112323133112111-0233132032013323-3031330303013011-3222131102331323-2113210231131111-2031000112331212-2013310133103012-2101032322322131"></a>

<a id="canonical-2113301000323023-3021221021303013-0222033220312202-2132310011003013-2011010021331320-0001222133213200-1212021103111111-3110230033302200"></a>

## action property — ingress_rules / 023220010331 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](data-sources--network_policy--reference--group-001.md#canonical-0332230313200212-0113130102200222-1202230220131032-2111013333220003-3123032001003210-0021333013131103-1332002031332111-0131300203330301): complete subsection reference.

- [all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-1001230320213003-2312310222000232-0221201001132000-0111213102321112-1300013312122130-0011003132302113-3001132302330031-1123021210330022): complete subsection reference.

- [all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3313101201130012-1303232111123120-1302200320120003-2120100213323213-0233221210033233-2120202002103330-1311203232103210-3310202301021210): complete subsection reference.

- [all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-2210002020212100-1002311211213012-3030313022123111-2201002001220030-3001000302322332-2121221310000201-1113010312213230-0330102121221313): complete subsection reference.

- [any](data-sources--network_policy--reference--group-001.md#canonical-2120300011201210-3020321010122120-1133003020121233-0023332112122121-0030121121212010-0230101113103302-0211111103313233-3010131100201222): complete subsection reference.

- [applications](data-sources--network_policy--reference--group-001.md#canonical-1310332200120013-2312203301321200-0310211101310311-3323211312100023-0222323031210130-1212201220111231-1313331332222101-3120202221330100): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-2203313213121122-1312113030330330-2122003330300231-2211232110321223-2210222133313013-0020100322110332-3011231030133011-1203101122223100): complete subsection reference.

- [ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-0320211313302223-3003013000310222-0002312302333230-0013230003321110-0301013033333230-0013022133323213-2110330310302103-3022113220221212): complete subsection reference.

- [label_matcher](data-sources--network_policy--reference--group-001.md#canonical-1301233302131132-3011120031132132-1230313022130201-0313223213202121-0003222332112212-1132112102011323-1132311011113033-0011202312330310): complete subsection reference.

- [label_selector](data-sources--network_policy--reference--group-001.md#canonical-3113220210330013-3131103230111213-2121223312200331-0000001323013313-0313132000323212-2303020201023203-3320232000330030-0221010310300210): complete subsection reference.

- [metadata](data-sources--network_policy--reference--group-001.md#canonical-0032310312311012-1032033002032303-2033232103311200-2130321310212112-1301222013211300-0032121123333010-0311023112201012-2303102132222332): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-1312022311333022-3122122032022231-3223032322333033-2221000222102033-3002101023021210-2003301131012011-1133110301303101-2100130210320113): complete subsection reference.

- [prefix_list](data-sources--network_policy--reference--group-001.md#canonical-3201322311211300-3310203230212101-1012320202301031-2302110003200200-1221110303110022-0302111232213133-3333111002202130-0003000333312300): complete subsection reference.

- [protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-3101101020222321-2010332301012111-1221300330332300-3203120020312213-1102112301122113-2133313123220023-1122330033013220-2213022100120111): complete subsection reference.

<a id="canonical-3113223303302122-2133001330110021-1232313000003302-1022001000112132-0330002312031013-3021133102020123-2000002001031001-3010311212310203"></a>

## Next pages — ingress_rules / 023220010331 / 5

- [rules.ingress_rules.adv_action](data-sources--network_policy--reference--group-001.md#canonical-0332230313200212-0113130102200222-1202230220131032-2111013333220003-3123032001003210-0021333013131103-1332002031332111-0131300203330301)
- [rules.ingress_rules.all_tcp_traffic](data-sources--network_policy--reference--group-001.md#canonical-1001230320213003-2312310222000232-0221201001132000-0111213102321112-1300013312122130-0011003132302113-3001132302330031-1123021210330022)
- [rules.ingress_rules.all_traffic](data-sources--network_policy--reference--group-001.md#canonical-3313101201130012-1303232111123120-1302200320120003-2120100213323213-0233221210033233-2120202002103330-1311203232103210-3310202301021210)
- [rules.ingress_rules.all_udp_traffic](data-sources--network_policy--reference--group-001.md#canonical-2210002020212100-1002311211213012-3030313022123111-2201002001220030-3001000302322332-2121221310000201-1113010312213230-0330102121221313)
- [rules.ingress_rules.any](data-sources--network_policy--reference--group-001.md#canonical-2120300011201210-3020321010122120-1133003020121233-0023332112122121-0030121121212010-0230101113103302-0211111103313233-3010131100201222)
- [rules.ingress_rules.applications](data-sources--network_policy--reference--group-001.md#canonical-1310332200120013-2312203301321200-0310211101310311-3323211312100023-0222323031210130-1212201220111231-1313331332222101-3120202221330100)
- [rules.ingress_rules.inside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-2203313213121122-1312113030330330-2122003330300231-2211232110321223-2210222133313013-0020100322110332-3011231030133011-1203101122223100)
- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-0320211313302223-3003013000310222-0002312302333230-0013230003321110-0301013033333230-0013022133323213-2110330310302103-3022113220221212)
- [rules.ingress_rules.label_matcher](data-sources--network_policy--reference--group-001.md#canonical-1301233302131132-3011120031132132-1230313022130201-0313223213202121-0003222332112212-1132112102011323-1132311011113033-0011202312330310)
- [rules.ingress_rules.label_selector](data-sources--network_policy--reference--group-001.md#canonical-3113220210330013-3131103230111213-2121223312200331-0000001323013313-0313132000323212-2303020201023203-3320232000330030-0221010310300210)
- [rules.ingress_rules.metadata](data-sources--network_policy--reference--group-001.md#canonical-0032310312311012-1032033002032303-2033232103311200-2130321310212112-1301222013211300-0032121123333010-0311023112201012-2303102132222332)
- [rules.ingress_rules.outside_endpoints](data-sources--network_policy--reference--group-001.md#canonical-1312022311333022-3122122032022231-3223032322333033-2221000222102033-3002101023021210-2003301131012011-1133110301303101-2100130210320113)
- [rules.ingress_rules.prefix_list](data-sources--network_policy--reference--group-001.md#canonical-3201322311211300-3310203230212101-1012320202301031-2302110003200200-1221110303110022-0302111232213133-3333111002202130-0003000333312300)
- [rules.ingress_rules.protocol_port_range](data-sources--network_policy--reference--group-001.md#canonical-3101101020222321-2010332301012111-1221300330332300-3203120020312213-1102112301122113-2133313123220023-1122330033013220-2213022100120111)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0332230313200212-0113130102200222-1202230220131032-2111013333220003-3123032001003210-0021333013131103-1332002031332111-0131300203330301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321113230333021-0323322023000021-0100110223202301-2131313131101233-3120323031230103-0322001000332322-1211133313332113-2002031133122011"></a>

## rules.ingress_rules.adv_action — adv_action / 213000011202 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.adv_action

<a id="canonical-3211031101300331-3332123102120111-2211303223330131-0323121332110020-2002113001132212-0213100223300031-2323003011303310-3331120303132000"></a>

Type: `"single"`. Computed.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0322121123003101-3031033120102223-3322003131011332-2000000200133132-3103112300302330-2130231123300310-2112230120022223-0311202210223203"></a>

## Direct properties — adv_action / 213000011202 / 3

<a id="canonical-2333310112210211-2313110322222331-2333330013111033-1131310232203332-2212111131002233-1101331213003023-2212230320003312-2331110333101010"></a>

<a id="canonical-2130331203300332-0211122201213322-2221203033123223-1202103101213320-3011000301122133-2112131203120113-1211101233222022-0202300021021011"></a>

## action property — adv_action / 213000011202 / 4

Type: `"string"`. Computed.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212312202033212-2120122031111030-2321211300112100-3032122032303030-1003023333221331-1321103030031222-3103331212310032-3001122100030332"></a>

## Next pages — adv_action / 213000011202 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1001230320213003-2312310222000232-0221201001132000-0111213102321112-1300013312122130-0011003132302113-3001132302330031-1123021210330022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020221030303322-3020113003122101-0133313332333023-2033110311031010-1031230013111132-0113133220310210-3301133002232020-2012102110021213"></a>

## rules.ingress_rules.all_tcp_traffic — all_tcp_traffic / 102332101233 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.all_tcp_traffic

<a id="canonical-2113232303110312-1311000103312012-3031013230302313-0230100321301321-2201103131100222-0221322133223002-2303021210022333-1332300011112101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all tcp traffic.

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

<a id="canonical-0213021133233321-0101331012220212-3200022023130030-0333011203020101-3310223001020112-3121031231122231-2002010112332100-0212103011231101"></a>

## Direct properties — all_tcp_traffic / 102332101233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311022000113122-1202133122301022-2021102131201103-2030031120321103-0011203011322001-2013111023300112-0213220210132212-3203102202300302"></a>

## Next pages — all_tcp_traffic / 102332101233 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3313101201130012-1303232111123120-1302200320120003-2120100213323213-0233221210033233-2120202002103330-1311203232103210-3310202301021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301122320323022-0001200030233232-0230130320131023-3132111020132223-3113030031221201-2231333310221230-3322311003303302-0202001111232231"></a>

## rules.ingress_rules.all_traffic — all_traffic / 033323133003 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.all_traffic

<a id="canonical-0323331202002230-3002002202210123-1032121122320103-1013101102333032-3312020333002212-3132130121231200-1303103230012323-0010333233011322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all traffic.

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

<a id="canonical-0320321212221211-2001211010033331-2021303123323313-3300002030233303-1130312322013221-1002132011030133-2212023231113001-0222100200320303"></a>

## Direct properties — all_traffic / 033323133003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101333100323323-3100200303303222-1022302022231012-0113122011322023-2012333311120011-1123031021212301-3123222100111212-2023101100212112"></a>

## Next pages — all_traffic / 033323133003 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2210002020212100-1002311211213012-3030313022123111-2201002001220030-3001000302322332-2121221310000201-1113010312213230-0330102121221313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102020000011333-1100020231010203-0130100112211233-0033211123301200-2221303202211311-1212102320121021-2212223001132130-2233111030031313"></a>

## rules.ingress_rules.all_udp_traffic — all_udp_traffic / 121002213001 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.all_udp_traffic

<a id="canonical-2130332111332100-0313302310002203-0202230233102331-0032102202212332-2111201201201323-2320103220232120-3212120301021232-2231231133213202"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all udp traffic.

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

<a id="canonical-1231033300310210-1301030110031230-1122202111100312-0322310311230212-1021322311000003-2132323013211123-1012320002201321-2030333202200310"></a>

## Direct properties — all_udp_traffic / 121002213001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3120301222232121-3330212101231221-2000012010231022-2320310103013030-0212103321331121-3102103133023121-0220001312331230-0320010322200030"></a>

## Next pages — all_udp_traffic / 121002213001 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2120300011201210-3020321010122120-1133003020121233-0023332112122121-0030121121212010-0230101113103302-0211111103313233-3010131100201222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002101302203312-2002120301201021-1321013333023011-1223010233103102-3021220113212233-0213302122012213-2212311021113331-3031213110131330"></a>

## rules.ingress_rules.any — any / 200311020013 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.any

<a id="canonical-0023103301121211-2001213000212311-2303002320213213-1023210310211203-2331222232111321-1230102030002201-2033331010100210-0330200330031122"></a>

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

<a id="canonical-1001321122013202-1022302033023222-0023232130320001-1113222320332233-2211011011113100-2012130123030320-0102211100023233-2310301103303200"></a>

## Direct properties — any / 200311020013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333311030010222-1323003012221223-0330333311203112-0102010212102300-0013132020203333-0020100320000121-0020121330013333-0033322130133303"></a>

## Next pages — any / 200311020013 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1310332200120013-2312203301321200-0310211101310311-3323211312100023-0222323031210130-1212201220111231-1313331332222101-3120202221330100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120121211210320-3020220301103322-3013303301102332-3033302301303003-3331011022032033-1320021232011021-1210312111020102-1123000231322113"></a>

## rules.ingress_rules.applications — applications / 322223110221 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.applications

<a id="canonical-0231333211022302-2300311103303130-1010102031233121-3112100310321222-2230103002322212-3301231101332310-2013203313003000-2002103131122002"></a>

Type: `"single"`. Computed.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1302202312120222-1101322103010323-2232123313313022-2200030101223101-2311020021310002-2202333233223002-3302312002030033-1013133122311303"></a>

## Direct properties — applications / 322223110221 / 3

<a id="canonical-3032003223231231-2012212103333012-1302322302230310-2111021032311321-1100302101232101-3012011023323220-0201223312103302-3212233033201200"></a>

<a id="canonical-3213332300230103-1011121213000233-1100100322120302-2002003112112331-0010012320013221-1231300301000123-2011313322201222-1030130021023212"></a>

## applications property — applications / 322223110221 / 4

Type: `["list", "string"]`. Computed.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0033221322122110-2212323303032232-0200030223103101-3322132301311322-3110233310211002-1021221010003002-2102220203032131-1303133112030132"></a>

## Next pages — applications / 322223110221 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2203313213121122-1312113030330330-2122003330300231-2211232110321223-2210222133313013-0020100322110332-3011231030133011-1203101122223100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300022132113010-0131321330011012-0301201202312311-1010321021100203-1321223123230001-1031233230032030-0213303121220313-3300013021310103"></a>

## rules.ingress_rules.inside_endpoints — inside_endpoints / 012333103111 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.inside_endpoints

<a id="canonical-3032031302223030-0330220320221012-0130231203121023-3331203200132223-1131300033311030-3003002310331003-3121232132233110-0322322100010100"></a>

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

<a id="canonical-3010202011230023-2332100322210011-3302110213021101-2131133310002333-2333210331322302-3002131002133131-2322130012100320-1211022011133231"></a>

## Direct properties — inside_endpoints / 012333103111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331332001032213-3212311202311313-3323031203323100-2010031222333320-1122111030130212-3230333130223121-2333103102221302-2202211302000313"></a>

## Next pages — inside_endpoints / 012333103111 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0320211313302223-3003013000310222-0002312302333230-0013230003321110-0301013033333230-0013022133323213-2110330310302103-3022113220221212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321113300132232-3023232303210331-2221022332201122-3122123120110213-1222313000130231-0101212233322033-2220302203220002-2033123132300133"></a>

## rules.ingress_rules.ip_prefix_set — ip_prefix_set / 323332223211 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.ip_prefix_set

<a id="canonical-1031131221033202-2002131303013320-2000211003321323-3020123202220020-3113333002110123-1100012322320332-0132120332330213-1223322223220212"></a>

Type: `"single"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1221323333022030-0303023030202112-2303212013210130-2313221210123230-3211101202120301-0323211332101012-3122133230322130-3312112202232113"></a>

## Direct properties — ip_prefix_set / 323332223211 / 3

- [ref](data-sources--network_policy--reference--group-001.md#canonical-2311322310332112-2302012023321333-1113100103231122-2013010122130302-2220312210131032-0320102000233233-2300232320112231-0203322023322223): complete subsection reference.

<a id="canonical-0000200312222033-1101021222010330-1122201203222213-1322102220010032-0031333311331212-2332231202230111-1222000001320310-3033033332332313"></a>

## Next pages — ip_prefix_set / 323332223211 / 4

- [rules.ingress_rules.ip_prefix_set.ref](data-sources--network_policy--reference--group-001.md#canonical-2311322310332112-2302012023321333-1113100103231122-2013010122130302-2220312210131032-0320102000233233-2300232320112231-0203322023322223)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-2311322310332112-2302012023321333-1113100103231122-2013010122130302-2220312210131032-0320102000233233-2300232320112231-0203322023322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321023132013213-2303013022103212-3102111022223130-3110313112023133-2010302332101330-2102233111321111-0113200110313030-2302113022230221"></a>

## rules.ingress_rules.ip_prefix_set.ref — ref / 202132301230 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-0320211313302223-3003013000310222-0002312302333230-0013230003321110-0301013033333230-0013022133323213-2110330310302103-3022113220221212)
- rules.ingress_rules.ip_prefix_set.ref

<a id="canonical-0103103231202110-2123112220130331-3121112030310332-3211233133111233-3121101112101130-1112022011200123-0103322101113200-1201011330300111"></a>

Type: `"list"`. Computed.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1333202002011320-2333132330223200-0112212122033303-0131213220131021-0032110220013220-3301013201320123-1110211132010111-0032123322020300"></a>

## Direct properties — ref / 202132301230 / 3

<a id="canonical-0211003221312310-0130102212332011-0130121330130020-0122113223220203-1100202332221311-2321011133321112-3021221001101233-0030032023113203"></a>

<a id="canonical-2330211012031210-1333030010122020-1121211123133030-1312230122111211-1331310310201210-2232201122101133-2310012220022302-3011330301003230"></a>

## kind property — ref / 202132301230 / 4

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

<a id="canonical-0233003103101221-2321020023221100-1030101203001230-2220201200000310-3321102022210033-1113300210300110-1231232133300203-1223313322210230"></a>

<a id="canonical-0332221311110222-3032030233100122-1202012022323232-0320100213322230-0202301121230131-3120223030331313-0323010021123310-0132221110120311"></a>

## name property — ref / 202132301230 / 5

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

<a id="canonical-2010103211001311-3220211233102312-3321310102131010-3231322003033331-1330103021233332-1012011220223230-1110003123331100-3100310310212222"></a>

<a id="canonical-0013313312212303-2130111333233103-0102333212332033-2201322100332311-2311022020100213-2011330310312321-1200332203130301-0011231112211322"></a>

## namespace property — ref / 202132301230 / 6

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

<a id="canonical-2301220333302331-0133000121122130-1300023203312313-0120121201013033-3212022331323312-1033311111010103-2133032031213331-2201001013333133"></a>

<a id="canonical-1031102113222023-3132022233033003-0120022003132220-0022112113012302-0212222021300310-1321111312100103-0120300012122101-1101122103001111"></a>

## tenant property — ref / 202132301230 / 7

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

<a id="canonical-2121333232030000-3022333312303311-3330000013111323-2332113222331231-3222230333131213-1122222310202220-0012032213013021-2233221313113002"></a>

<a id="canonical-2033112220000030-0113001310012300-1101232332200322-0003313133301213-0213002032313302-1112132313032323-1023330022121320-1231323022113220"></a>

## uid property — ref / 202132301230 / 8

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

<a id="canonical-1311301220201312-1230310001113300-0210021023101333-0303113211210200-1030230312100221-1120310012010232-1211202032001111-0003330223212123"></a>

## Next pages — ref / 202132301230 / 9

- [rules.ingress_rules.ip_prefix_set](data-sources--network_policy--reference--group-001.md#canonical-0320211313302223-3003013000310222-0002312302333230-0013230003321110-0301013033333230-0013022133323213-2110330310302103-3022113220221212)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1301233302131132-3011120031132132-1230313022130201-0313223213202121-0003222332112212-1132112102011323-1132311011113033-0011202312330310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123013303223200-3300101123133023-1321300020220322-3010311302202212-2203331230112202-3013102123223220-2200021211220220-0110211330111331"></a>

## rules.ingress_rules.label_matcher — label_matcher / 313232002333 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.label_matcher

<a id="canonical-1013313033123130-3110010312111122-0021020103031231-3223100031233210-2000122222103320-1012110011323213-0310010302021020-0301031122102210"></a>

Type: `"single"`. Computed.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3211333003032122-3203003111112131-3323101000031202-3100032212310132-0111111102223012-2103212303231312-0202031233211303-2111101131110111"></a>

## Direct properties — label_matcher / 313232002333 / 3

<a id="canonical-3230121322321222-0301112333301211-1130011210132332-3232103220102332-1110102002111331-3102001321123021-1003212112230023-3220202311000331"></a>

<a id="canonical-2112332303103312-1003121010131032-0023202110112011-0333301231123130-3202031013013003-2313132002213211-0033023103212101-2102322201313130"></a>

## keys property — label_matcher / 313232002333 / 4

Type: `["list", "string"]`. Computed.

The list of label key names that have to match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0031220322130230-3321103230221131-1301123220132212-1010211010133233-3101020100010030-1200003331301102-3303131212223220-2112002131223021"></a>

## Next pages — label_matcher / 313232002333 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3113220210330013-3131103230111213-2121223312200331-0000001323013313-0313132000323212-2303020201023203-3320232000330030-0221010310300210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121000003220332-1131223002010310-3023110012310302-2310102030103203-1300100302302330-2210133200021010-1231311301301333-1103301023301031"></a>

## rules.ingress_rules.label_selector — label_selector / 033330000112 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.label_selector

<a id="canonical-1132021230211323-3020113122133100-0220132001102020-3230230002131012-2122232103132321-1322210213320002-3203233222223230-0231123023011202"></a>

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

<a id="canonical-2202303211311020-3331120213010123-0332111300123311-0302212001033212-3022031020322323-2202302003130200-0002001202330031-2010002000322011"></a>

## Direct properties — label_selector / 033330000112 / 3

<a id="canonical-3101201201333233-1223222033102023-1132202113310303-0232112302020030-1122313331001011-0232111323223023-1231223131113023-0323330300213333"></a>

<a id="canonical-1133311032002102-3203022313000233-0002132322322332-1321331321320203-3333112132113102-0232000031120110-2121032110221112-2132113220233131"></a>

## expressions property — label_selector / 033330000112 / 4

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

<a id="canonical-1320322100210030-1221123111112122-3201111133330203-3331333012201101-2202023201312211-3101212202011332-0110332323331231-3310020210032021"></a>

## Next pages — label_selector / 033330000112 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-0032310312311012-1032033002032303-2033232103311200-2130321310212112-1301222013211300-0032121123333010-0311023112201012-2303102132222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230110013200220-1231123231303223-0010321220120312-0212230020222323-2231323201010110-3311312321203222-0011210120332313-2111010211132032"></a>

## rules.ingress_rules.metadata — metadata / 022033112321 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.metadata

<a id="canonical-0203122132002323-1033321023130211-2012310300213210-2203320312312101-2110310020333121-2110123303233311-1213210233120202-0233011113121010"></a>

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

<a id="canonical-0302233220101122-2032303122112133-0011223013113130-3103103231122220-1313301132021303-0020221112312102-3311223113001002-3000221310213123"></a>

## Direct properties — metadata / 022033112321 / 3

<a id="canonical-1011021321313130-1312303012000202-2002131131023111-1312223002022100-3102223021331131-0121003110122130-3013333010013100-1210332133130031"></a>

<a id="canonical-3121330131123230-2310130201220330-1013311123212230-0000230012103323-2313121131330201-3333233000230010-3111112330133112-1230233031331000"></a>

## description_spec property — metadata / 022033112321 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0002123310220030-2110112001133231-2330003021330201-0110231220131012-1133000111022011-1310113133032310-3022213111023032-1023330213210330"></a>

<a id="canonical-3312303012300132-3030303121221333-2210132331133310-2122121302103110-2222120223311031-1030203301102120-2310023113213331-2102120121122211"></a>

## name property — metadata / 022033112321 / 5

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

<a id="canonical-2333211210223101-2120230012201031-2000320213302132-1023213033311020-3121232031023013-1030313003310202-1010033333321212-1010022231132111"></a>

## Next pages — metadata / 022033112321 / 6

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-1312022311333022-3122122032022231-3223032322333033-2221000222102033-3002101023021210-2003301131012011-1133110301303101-2100130210320113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102102020120321-0001033232320022-0320333110130321-2013130330201000-3102331203001312-3232122110313310-2102030003113211-3122212223321100"></a>

## rules.ingress_rules.outside_endpoints — outside_endpoints / 200101102013 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.outside_endpoints

<a id="canonical-1022223220201130-2101020202101210-2123211313300312-1233311022313030-0221222231223321-3303330013211012-1321113122103023-1103221023220322"></a>

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

<a id="canonical-1323020030010233-3012032003023100-3312000132323321-1311231300113302-2000031010220313-2212100031020012-2120203012313101-2311233332320132"></a>

## Direct properties — outside_endpoints / 200101102013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031222023013312-3310332320230102-0201031001201333-1211000111332012-2333310011121211-3032000010011202-2120332323123022-0223212201000112"></a>

## Next pages — outside_endpoints / 200101102013 / 4

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3201322311211300-3310203230212101-1012320202301031-2302110003200200-1221110303110022-0302111232213133-3333111002202130-0003000333312300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122231213211332-1000121203123203-0032331110121023-2033223002010303-0001232120121011-3322111102312023-3330002110321101-0131133000200033"></a>

## rules.ingress_rules.prefix_list — prefix_list / 311121103012 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.prefix_list

<a id="canonical-2221020203121002-0000123112313031-0320220002130321-2213030002100321-0202123003222112-2102232132033333-1112001032103013-0111333300222012"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010311133231211-1031223103003313-3001223331111320-0311002003121023-1132021033212002-2321203223213111-3313231133112231-3212313232203221"></a>

## Direct properties — prefix_list / 311121103012 / 3

<a id="canonical-2322100301021103-2211122033201233-2221113230102222-0102110202132231-1031110211331121-1031001232203113-2032110033331023-3330103321320223"></a>

<a id="canonical-2202301003330000-0130001302301023-0100220123302113-3231320332211223-3213200023013101-2122003201332100-3131133323300020-0312022021200202"></a>

## prefixes property — prefix_list / 311121103012 / 4

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1232312102131230-2130213002023000-3302130202323223-2001110213021132-1211322230101303-2223220110223121-2232021022020121-1030312323332211"></a>

## Next pages — prefix_list / 311121103012 / 5

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)

<a id="canonical-3101101020222321-2010332301012111-1221300330332300-3203120020312213-1102112301122113-2133313123220023-1122330033013220-2213022100120111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020331321102320-2030122021213010-0330232233123220-2100213311100002-3030220010110311-3233101211002111-2022012300203032-1020031111212020"></a>

## rules.ingress_rules.protocol_port_range — protocol_port_range / 012111301222 / 2

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
- [Property reference](data-sources--network_policy--reference--group-001.md#canonical-1311103202221102-1230113110322222-0330120323322123-0223212230313111-1113112233013330-2101223313322112-3011012123300123-2111301323113333)
- [rules](data-sources--network_policy--reference--group-001.md#canonical-0330200303122200-2303303233120202-0133310221030110-3023123232222121-0221333023110201-3023220333211320-3001222113111213-0310132131330212)
- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- rules.ingress_rules.protocol_port_range

<a id="canonical-1301231020022111-3101121231313233-1230130301021331-3313223032103110-1301023133112331-3103231000200020-0110032333031112-2223120320211201"></a>

Type: `"single"`. Computed.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0032313022103122-2110222203333332-0011133212002200-1103012133213130-1120230002232130-2011002303022112-2013003021031023-1100210010001002"></a>

## Direct properties — protocol_port_range / 012111301222 / 3

<a id="canonical-2010201233323233-1011213222103213-1233001101232321-2312311302312203-1213310022230301-0032012232311130-1012113323112010-3201213222012012"></a>

<a id="canonical-1123023210123103-3303130032101313-3111201230321030-2113102322203013-0030000331232022-1130110030210321-1112302111020323-0123030101203231"></a>

## port_ranges property — protocol_port_range / 012111301222 / 4

Type: `["list", "string"]`. Computed.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-2003131001231131-0200002323202133-0211223120022020-1013030010011121-3310123100221020-1233032210333030-3032112323301023-2113330223313010"></a>

<a id="canonical-0000021132013120-0213303230012221-2130033000332131-1230332311032033-2323123211313013-1023022030321133-1032100222321020-2021121002002022"></a>

## protocol property — protocol_port_range / 012111301222 / 5

Type: `"string"`. Computed.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-1230311103331022-2323213222220313-0031023020312131-2213333103311021-0231302203322132-3301312001130013-0003210000222010-0303322212222220"></a>

## Next pages — protocol_port_range / 012111301222 / 6

- [rules.ingress_rules](data-sources--network_policy--reference--group-001.md#canonical-0123023032311111-1202222220101231-2331122200222311-1300323303221200-0130022103330103-2302201300011233-1210330301202201-1313112200031202)
- [xcsh_network_policy](../data-sources/network_policy.md#canonical-0310211312011330-3122113202012230-1220102020331323-3311020133111313-2012202311123131-3000120301000221-0313212120103323-3000022233312033)
