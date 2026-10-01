---
page_title: "xcsh_fast_acl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl reference."
---

# xcsh_fast_acl reference

<a id="canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201110221200021-1220132032130130-3301203303201023-0033022122103131-2130013201002133-1133002121302122-1032321212220231-0101020200100020"></a>

## Property reference — Property reference / 333203021003 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- Property reference

<a id="canonical-2213302033212103-3113223201013231-1002313110012330-2321200233121330-0132113301223301-2203310213131133-0120231211022102-3001331011133210"></a>

## Direct properties — Property reference / 333203021003 / 3

<a id="canonical-3012312132321121-1121301331020330-3201113132010100-2121022313032212-3200122221233332-0312030300202330-3223203110100021-1331313120210113"></a>

<a id="canonical-3320111313032203-2123022120113322-1121021100211001-2021303321313010-3111300023102110-3000132021332102-2330102312121023-2331112033321223"></a>

## annotations property — Property reference / 333203021003 / 4

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

<a id="canonical-3102003103101132-2111231100122131-3233031013301103-3231313300301123-2323213301101021-2313111022203330-1330213211333121-1301101130023011"></a>

<a id="canonical-1202200123032302-3220003311130013-1013032103103103-3023311302231021-1133131333302303-0313231231123203-3323321211012002-2000210213012033"></a>

## description property — Property reference / 333203021003 / 5

Type: `"string"`. Computed.

Description of the FastACL.

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

<a id="canonical-3022321032210321-1233210012333332-1130323300113232-3222120012022132-2303023031313211-0232031031102303-1303321230011130-2310332103112223"></a>

<a id="canonical-1003322230220111-3003030022332211-1011113310103122-2223233012312331-2322200131222321-1011123132003103-2302313132033223-3010111121312112"></a>

## ID property — Property reference / 333203021003 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1130230031110323-1302213323323303-3132302203221202-2322032220102222-1330121031001110-1302322301330230-1023213233230003-1321331101023021"></a>

<a id="canonical-3332123330213123-0121213210121001-1031332213313322-2201002200020333-2333021301232220-1133100200120132-3203310311323221-1233133110010300"></a>

## labels property — Property reference / 333203021003 / 7

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

<a id="canonical-1200232223222301-1103130322011113-3101101323132320-2330121123332112-1230010331333332-2233312120003202-2222001221100113-1313321130202001"></a>

<a id="canonical-2222121121231310-2013033021213120-0130120113300303-3320232012320022-3231230001333213-0100203332000111-0030022202022301-3122102222121132"></a>

## name property — Property reference / 333203021003 / 8

Type: `"string"`. Required.

Name of the FastACL.

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

<a id="canonical-0320301312323211-3030130320120323-2302111213202303-2223203033002123-1303112013312103-2203010003121133-3313320031103003-3321003212123122"></a>

<a id="canonical-3122222113312222-2111000003022000-1320212220031033-1100312322033313-2331210000332012-1203301312033121-0133221210130133-2232000321200130"></a>

## namespace property — Property reference / 333203021003 / 9

Type: `"string"`. Optional, Computed.

Namespace where the FastACL exists.

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

- [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-0311301022103313-0032311323131023-2000131002233033-3213031223300331-0130121332023020-0012320112122022-1031332013011200-0003312303033221): complete subsection reference.

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221): complete subsection reference.

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320): complete subsection reference.

<a id="canonical-1211012002233001-1001123020300130-0233003232022032-2300122322213111-3133202132201322-0110130033103202-2322300032013330-0012311323012311"></a>

## All schema paths — Property reference / 333203021003 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--fast_acl--reference--group-001.md#canonical-3012312132321121-1121301331020330-3201113132010100-2121022313032212-3200122221233332-0312030300202330-3223203110100021-1331313120210113) |
| `description` | [description](data-sources--fast_acl--reference--group-001.md#canonical-3102003103101132-2111231100122131-3233031013301103-3231313300301123-2323213301101021-2313111022203330-1330213211333121-1301101130023011) |
| `id` | [ID](data-sources--fast_acl--reference--group-001.md#canonical-3022321032210321-1233210012333332-1130323300113232-3222120012022132-2303023031313211-0232031031102303-1303321230011130-2310332103112223) |
| `labels` | [labels](data-sources--fast_acl--reference--group-001.md#canonical-1130230031110323-1302213323323303-3132302203221202-2322032220102222-1330121031001110-1302322301330230-1023213233230003-1321331101023021) |
| `name` | [name](data-sources--fast_acl--reference--group-001.md#canonical-1200232223222301-1103130322011113-3101101323132320-2330121123332112-1230010331333332-2233312120003202-2222001221100113-1313321130202001) |
| `namespace` | [namespace](data-sources--fast_acl--reference--group-001.md#canonical-0320301312323211-3030130320120323-2302111213202303-2223203033002123-1303112013312103-2203010003121133-3313320031103003-3321003212123122) |
| `protocol_policer` | [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-2111010220230013-2112302311313331-3313220222313022-2312111203021001-0021220132130003-1120101033023222-0012013033030030-3202000123112213) |
| `protocol_policer.name` | [protocol_policer.name](data-sources--fast_acl--reference--group-001.md#canonical-2012123132032332-3013301003032120-3302203313222003-2031003030202110-1210023011220120-0221201213011210-2113233300313221-2231202001223001) |
| `protocol_policer.namespace` | [protocol_policer.namespace](data-sources--fast_acl--reference--group-001.md#canonical-3121031031313232-1011121022021012-2213123310213222-0320310133311232-3221320002202120-3302333300131312-2300111033203302-0100020322111300) |
| `protocol_policer.tenant` | [protocol_policer.tenant](data-sources--fast_acl--reference--group-001.md#canonical-2021102211231313-1130113020113331-0313103203330030-2133332220003001-3221302001130032-1002110023023100-0331113130303202-3133111031231133) |
| `re_acl` | [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-0111200112123112-1201123020211310-2332023333003303-0133311221133011-1310202220232303-3232110330313223-0123233201033320-0123313122323121) |
| `re_acl.all_public_vips` | [re_acl.all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-2133302111302000-1212132332231033-1000122031320321-0103133231201122-2101301313331321-1312320032202001-3203222132011011-2131010220022001) |
| `re_acl.default_tenant_vip` | [re_acl.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-2131131000332002-0122101211312223-3302323002100023-2002331320102230-1212322123110332-0012101100301301-2001223131212133-2000231212301023) |
| `re_acl.fast_acl_rules` | [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0100000011300320-0022100303023020-2103221102110111-3003003221001300-3133033322323323-2001213002220332-1011333111110312-3312023303101223) |
| `re_acl.fast_acl_rules.action` | [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-1030031112100211-2330222201312100-1110010223032221-2110110002121012-3301121222230212-1222021102320202-2000232001220120-2203233102333302) |
| `re_acl.fast_acl_rules.action.policer_action` | [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3130113210321131-0113103232103133-3323001321102031-1222000210130000-3213101122232120-2212033131030001-3003023301011322-2002121211012323) |
| `re_acl.fast_acl_rules.action.policer_action.ref` | [re_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-1020222223103000-0031232003210320-3312033031212213-3111013313333002-1221023312020002-0033113233300211-1123113011323013-1111322001112332) |
| `re_acl.fast_acl_rules.action.policer_action.ref.kind` | [re_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-3103321210100232-1310112131102201-0213323233113231-2330001332332120-2011023113213111-3223022201300203-0303120320321213-0223201322122230) |
| `re_acl.fast_acl_rules.action.policer_action.ref.name` | [re_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-2333320030111303-2310312233222030-2120102001021020-0011230321321131-3302102131312310-3331332311001232-3321303311313021-3121231122011023) |
| `re_acl.fast_acl_rules.action.policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-2202131013302013-2220020022030332-1203133103133101-2311022030001222-2201221030331332-3332030302000231-2313220122230211-0033132131113321) |
| `re_acl.fast_acl_rules.action.policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-2131011213030021-2022112210103002-2130000230221131-2131131201132302-0312122312111311-2220211021131320-1002213230313220-3003000330032321) |
| `re_acl.fast_acl_rules.action.policer_action.ref.uid` | [re_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-3122220021321010-0130130120122122-3023311300100331-0032323300300302-0133301100011030-1230223020132121-3302320331303030-2100302130320321) |
| `re_acl.fast_acl_rules.action.protocol_policer_action` | [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3312031033002100-1213031012312212-2122001011200013-2232100022130012-2230032332032321-0300120312013121-0020232120123221-2302002122121300) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-2010012301322011-1330221132003110-1212213103233101-2332002030003233-2322113322110313-2211133221111330-0223223220121203-3300011131233011) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-0100330020203131-0011021122330312-3213203123332332-0132330201111233-3100012132200201-0202020203321103-1022213311113131-0233220102102103) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-0301000230213202-2323230123323023-1100133101321002-1011223201131221-1310020311010213-1300033131001212-1322100312300031-1101233111113232) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-2213333312301110-1330302002022322-3231321310333113-2132230023220100-0330301223321333-1111203333110300-0132212331330023-0013233021120001) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-1313011203132020-2003210201232001-1110330320033001-2010121023222323-3302321031123031-1300002010003000-1211322120101233-2000103310132213) |
| `re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [re_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-2213212333313330-3302113221312200-1132311031321012-2000002201023120-3023031300023300-1312202331000123-3320002101301301-0011223131313212) |
| `re_acl.fast_acl_rules.action.simple_action` | [re_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--reference--group-001.md#canonical-1020010331203032-1122010001210212-3032030002012210-1203112232120202-0311330331021030-0002201122122132-3013023031023000-1020032211031303) |
| `re_acl.fast_acl_rules.ip_prefix_set` | [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-1021233113320302-1322102330203021-1101123102211210-1021303012113332-1310303011302302-3132200320211112-3033203032030222-2102332312000123) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref` | [re_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-3203203100001000-1331131203020332-3331323330113201-1001310200211023-0000331331021110-2323123031033033-1220312300021233-2230321010303122) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [re_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-1333233321120131-1200020113013320-3323321122302131-3301232110110201-1101211233310011-1301102001032111-2200330131210133-3211330013121131) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.name` | [re_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-0031123033300210-0300013003102200-0122301333302212-2112121211121333-0311132221132303-2033121301201221-3313030031320213-1132202001021211) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [re_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-1331121313302020-3132210311121230-2313122030303131-1301323011331320-0110033202100100-0000220200303222-0332001211023031-3210031233131222) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [re_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-1121022332122201-1230223201232212-2033231322231010-1030030002323333-2013222112112222-1201113200131101-3210120132032331-0120022121120023) |
| `re_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [re_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-1002303112233123-3211333121223210-3032000223022030-3103201312202000-2002003100220332-3311213030102012-2303312322313310-0223120320220122) |
| `re_acl.fast_acl_rules.metadata` | [re_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-1313023111221011-2032322300212120-2321311110222010-0223212323100133-3113100021021112-2233222330323110-0133213211311200-0001333133200200) |
| `re_acl.fast_acl_rules.metadata.description_spec` | [re_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--reference--group-001.md#canonical-0220003320332013-1221003012002120-1231021321021003-2221102030301100-1230121330220330-1322012330201333-0130011133023110-1323220021213210) |
| `re_acl.fast_acl_rules.metadata.name` | [re_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--reference--group-001.md#canonical-1023132100211110-2320130001133133-2322010111230132-3231121021202322-2131111132112122-0203121101320110-2313101233320131-3221310122132222) |
| `re_acl.fast_acl_rules.port` | [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-0020001103331101-2022032023031123-3220311030000222-3313123330311230-3201312123301001-0131201232221122-1211321210223202-3202000103122203) |
| `re_acl.fast_acl_rules.port.all` | [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-3330211213020021-3212230330323012-0333131022120333-2012300222322313-1333220301021022-3031003132131302-2302131012102102-1131020221203101) |
| `re_acl.fast_acl_rules.port.dns` | [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-1201021113320311-2300133030221022-1303311112213012-1031311202013302-2102002003002320-1331312123020332-0100223020002303-3322213032212023) |
| `re_acl.fast_acl_rules.port.user_defined` | [re_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--reference--group-001.md#canonical-1311302333103132-2211022330031210-0033300203212301-3031131222000221-3222220123321333-2010001003332320-0231231311010332-2121231303232033) |
| `re_acl.fast_acl_rules.prefix` | [re_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-0210123112030200-0021210101232331-0233232311220031-1032312101220320-0102033132020032-1232100330320333-1122302303201111-3221110200112203) |
| `re_acl.fast_acl_rules.prefix.prefix` | [re_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--reference--group-001.md#canonical-2111030300311310-1023202123010212-3020133021123230-2132001010011101-0022012030011232-2312200133021222-1103202032112132-0330130322311201) |
| `re_acl.selected_tenant_vip` | [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-3013301000123133-0310210332330321-0033313030213213-3101023122020222-0101301011013233-3303000122212200-0210011112113303-0000010201213300) |
| `re_acl.selected_tenant_vip.default_tenant_vip` | [re_acl.selected_tenant_vip.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-3121020320211021-3020131013220213-0300331311310130-2231221030021010-0103220301310300-0222032010030130-1110303213233213-0333133321103010) |
| `re_acl.selected_tenant_vip.public_ip_refs` | [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-2030223122120233-3322233002100201-2300103030032201-2022233201231132-0123330230123023-1113200122012131-0221233213031332-1021110111000302) |
| `re_acl.selected_tenant_vip.public_ip_refs.name` | [re_acl.selected_tenant_vip.public_ip_refs.name](data-sources--fast_acl--reference--group-001.md#canonical-3123030023110112-2003230222003031-0123022023000303-1033330311310322-1133221013131320-1330300303010023-0101330112222213-3321001120323200) |
| `re_acl.selected_tenant_vip.public_ip_refs.namespace` | [re_acl.selected_tenant_vip.public_ip_refs.namespace](data-sources--fast_acl--reference--group-001.md#canonical-3220012021200103-2122102201022130-3220023122122302-1023300321220100-2133003331022313-0110003333222301-1312033021232321-2020223221202021) |
| `re_acl.selected_tenant_vip.public_ip_refs.tenant` | [re_acl.selected_tenant_vip.public_ip_refs.tenant](data-sources--fast_acl--reference--group-001.md#canonical-2132001133323113-1332233331310012-2332013222313003-3131103212230003-0000232020103123-2320213302211131-3110030330101102-2202111300230112) |
| `site_acl` | [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-0331211011301320-1133001113102311-0001132302121130-0020011101030233-3223233222131221-3013131100322113-2221100330031103-0221021322223110) |
| `site_acl.all_services` | [site_acl.all_services](data-sources--fast_acl--reference--group-001.md#canonical-2130333322212111-1221003031120023-2233211022333332-3333213310311121-2321100213103131-0002210311030230-0112200131323010-3202213113110312) |
| `site_acl.fast_acl_rules` | [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-3112130121101333-3220210230032311-1023311220100200-2331222302211111-0223200123221321-3223013022223120-1003030130201003-3301323333131013) |
| `site_acl.fast_acl_rules.action` | [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-3222000113123220-3030201030001021-2203031201103320-0132120332022313-2131303200210100-2122220111100302-1131222323120123-1210223131201321) |
| `site_acl.fast_acl_rules.action.policer_action` | [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2202232320122231-0132110223012132-2022202102211211-0010332211231023-0221220320011300-1121002302103311-2203031003020201-1000121230321300) |
| `site_acl.fast_acl_rules.action.policer_action.ref` | [site_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-0322133030132010-0231300112030001-1000201220213212-2133321301331301-0002103103203301-0022120212300322-3032302001210130-0310022232202200) |
| `site_acl.fast_acl_rules.action.policer_action.ref.kind` | [site_acl.fast_acl_rules.action.policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-2022120301101132-0323033012113203-1313221203333212-3133103131031020-2302023011111320-0033220200333122-2310232030011210-2310101000222102) |
| `site_acl.fast_acl_rules.action.policer_action.ref.name` | [site_acl.fast_acl_rules.action.policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-2303313331102102-1202111023121000-1220313211220121-3001010222033020-1020221011120233-3201003222001332-2231123302320202-0321210023121233) |
| `site_acl.fast_acl_rules.action.policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-1121023133100300-3020323320030210-2111312323013221-1121303231220003-2321110130101321-3012312123131002-3320312122102003-1122213121231221) |
| `site_acl.fast_acl_rules.action.policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-1111111101322223-2212301011310232-0023211011102122-1033321120121223-0230231120033013-0000201323221332-0010022320020321-0002320300200102) |
| `site_acl.fast_acl_rules.action.policer_action.ref.uid` | [site_acl.fast_acl_rules.action.policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-1031302120000230-0033121020220022-1011311033221300-2232021123130312-0120131122011021-1221211322331331-3031223202321213-1311311300232013) |
| `site_acl.fast_acl_rules.action.protocol_policer_action` | [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2333110203212222-3111321200133102-3223111221110021-2233221221303132-1312321231012121-0033101012030030-2210102213230210-3230321213332331) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-0120101223313111-2330003231132132-0121113101322331-1022330201312022-1010323133112121-1123302100020310-1003003031033023-2033313101301011) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-0112233320310103-0032023011123323-3330102212313101-3211032311033223-1032200123032321-3333011030033033-3333113100322103-3031320312122121) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.name` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-2022103232031202-3003233231120320-1201303230031222-1200111332002101-2121322023120212-2310113212213203-1221220013112203-0023300332213010) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-1021300113012122-1001121333012211-0133212022012021-3211033103020031-1302012311002312-3022002103211031-0021323211213003-1301321122013220) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-3310131330020231-0212000321301301-1100011133213121-2320210223221022-0230113120203121-0303311133030211-0231131011300031-2110101101122013) |
| `site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid` | [site_acl.fast_acl_rules.action.protocol_policer_action.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-0110021232033133-1332301110100221-3032221312322021-2323131230203100-3302103323201213-3000202020222033-0130123023002301-3211221101210130) |
| `site_acl.fast_acl_rules.action.simple_action` | [site_acl.fast_acl_rules.action.simple_action](data-sources--fast_acl--reference--group-001.md#canonical-3013211213212132-1220232102322202-3223211103112232-2021013012203133-3302223302133201-2220023333222210-3110120220312102-0333112020111002) |
| `site_acl.fast_acl_rules.ip_prefix_set` | [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-2311003303222221-2012332033100033-1320020301331012-0001120203212123-1223003133213022-1323012120100311-3012332010002231-1113233312202200) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref` | [site_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-3031132210123230-0331033010211312-1233131031032222-1231121030313321-1100310131003120-2222301311111000-2322302101211000-3311011130022031) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.kind` | [site_acl.fast_acl_rules.ip_prefix_set.ref.kind](data-sources--fast_acl--reference--group-001.md#canonical-2022213201300111-3031033301303031-0120000112313001-3230231323332113-1320023130020013-2000130210212320-3300133220301312-3320010332033321) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.name` | [site_acl.fast_acl_rules.ip_prefix_set.ref.name](data-sources--fast_acl--reference--group-001.md#canonical-3221012333020303-2001033010232320-3330320323200303-3200132021222222-1022200200313113-0301220232113121-2203213011330123-3112102331220230) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.namespace` | [site_acl.fast_acl_rules.ip_prefix_set.ref.namespace](data-sources--fast_acl--reference--group-001.md#canonical-0210302101311332-3102131013130201-3013301023022122-1121002022231033-3213123212012023-3100233221220123-1112030133311231-3311111203102013) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.tenant` | [site_acl.fast_acl_rules.ip_prefix_set.ref.tenant](data-sources--fast_acl--reference--group-001.md#canonical-3330112133313010-1332000203033302-2121100200031103-1331132033331232-2113023202001322-3022012013122300-3210012023020223-2001232313102023) |
| `site_acl.fast_acl_rules.ip_prefix_set.ref.uid` | [site_acl.fast_acl_rules.ip_prefix_set.ref.uid](data-sources--fast_acl--reference--group-001.md#canonical-3003320101133013-2011320230321201-1211021030230001-3032331231033210-1301120112120230-0310230030012111-1211031011211102-1002320103133032) |
| `site_acl.fast_acl_rules.metadata` | [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-2013312111111221-2302233123230220-0102311101231200-3000033110331112-2030222001120022-2121223013232302-1301311200121130-1132122202232203) |
| `site_acl.fast_acl_rules.metadata.description_spec` | [site_acl.fast_acl_rules.metadata.description_spec](data-sources--fast_acl--reference--group-001.md#canonical-3021300300101112-2301123012122122-1001200132100123-2132100023303101-3212103331310100-1032132322030121-3101000221310201-0332230003330122) |
| `site_acl.fast_acl_rules.metadata.name` | [site_acl.fast_acl_rules.metadata.name](data-sources--fast_acl--reference--group-001.md#canonical-3031201100303330-1313033313132033-1221012222001132-3223220320000001-3312132131132032-1102311033002002-3111313121013211-3020023003010002) |
| `site_acl.fast_acl_rules.port` | [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-3010233133203020-2020303223311301-2023121133310201-0213200201221311-3101020121033203-2332223133223111-1003013002010112-2312102103030120) |
| `site_acl.fast_acl_rules.port.all` | [site_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-1211123000302023-1310300310123332-2303222023300301-3111122001110312-0122022101233230-2012303203102111-2113221320121021-3113021233233122) |
| `site_acl.fast_acl_rules.port.dns` | [site_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-3032231211331131-0010322203311121-1110021121013213-0321331130220222-0322003003210133-2213000013222332-1032312300102001-3002102030121330) |
| `site_acl.fast_acl_rules.port.user_defined` | [site_acl.fast_acl_rules.port.user_defined](data-sources--fast_acl--reference--group-001.md#canonical-3111213233323003-0332302013030333-2301010101032313-0110312030131010-3330023331210310-0210010011102010-0222333332200110-2231123202013110) |
| `site_acl.fast_acl_rules.prefix` | [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-0223300323001312-2103122200321122-3022200310300130-0313132202132333-2023113230032120-2221033001022301-3100002320033133-2112001311223003) |
| `site_acl.fast_acl_rules.prefix.prefix` | [site_acl.fast_acl_rules.prefix.prefix](data-sources--fast_acl--reference--group-001.md#canonical-1311201310001202-1132332022012113-2121023123330322-2112122310112122-0232000301013110-2222131222210210-3223320003010211-1122210203301121) |
| `site_acl.inside_network` | [site_acl.inside_network](data-sources--fast_acl--reference--group-001.md#canonical-0031222033203333-0100330213230111-1310010003212100-1102131313202112-1113120133332330-0220021231330030-0312213120111333-3023332122000032) |
| `site_acl.interface_services` | [site_acl.interface_services](data-sources--fast_acl--reference--group-001.md#canonical-1021131310320323-1101133233222213-2310013232033020-3302122102101231-1011313113333303-1000320002101301-3323321230330012-1213103133322332) |
| `site_acl.outside_network` | [site_acl.outside_network](data-sources--fast_acl--reference--group-001.md#canonical-2323312133323311-3313202201231031-2200112321021131-1021332133101230-0313133031130122-2222033122333130-2121332231300321-2301332222010303) |
| `site_acl.vip_services` | [site_acl.vip_services](data-sources--fast_acl--reference--group-001.md#canonical-3233130210000133-2033220130031221-3333222302212021-1012003033031320-3130231012120302-0233100100333001-0012200311320310-0210311030001302) |

<a id="canonical-0331123212012013-1232102300321230-0130321201023032-0212032013200032-3332312123332111-2210021313021112-0323330113300201-3332211123133023"></a>

## Next pages — Property reference / 333203021003 / 11

- [protocol_policer](data-sources--fast_acl--reference--group-001.md#canonical-0311301022103313-0032311323131023-2000131002233033-3213031223300331-0130121332023020-0012320112122022-1031332013011200-0003312303033221)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0311301022103313-0032311323131023-2000131002233033-3213031223300331-0130121332023020-0012320112122022-1031332013011200-0003312303033221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323320203231311-0311013020132133-0003132000112111-1313233131331132-0212333313100121-1223111332323302-0132032032320121-3100100002333231"></a>

## protocol_policer — protocol_policer / 113023010321 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- protocol_policer

<a id="canonical-2111010220230013-2112302311313331-3313220222313022-2312111203021001-0021220132130003-1120101033023222-0012013033030030-3202000123112213"></a>

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

<a id="canonical-1223230023303301-2331332213211013-3222203221100013-0213222231021100-0102222032133220-0313020333022000-1212222111332132-3012033321101301"></a>

## Direct properties — protocol_policer / 113023010321 / 3

<a id="canonical-2012123132032332-3013301003032120-3302203313222003-2031003030202110-1210023011220120-0221201213011210-2113233300313221-2231202001223001"></a>

<a id="canonical-2313002212112200-3002212202320031-0333322003122322-1210310130011302-3201330203033113-1301210230233331-1302103000200021-3203100202333002"></a>

## name property — protocol_policer / 113023010321 / 4

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

<a id="canonical-3121031031313232-1011121022021012-2213123310213222-0320310133311232-3221320002202120-3302333300131312-2300111033203302-0100020322111300"></a>

<a id="canonical-2331311300203122-0002200323230033-0000220312121100-2232012131210110-1321001101111200-3302120013323022-2321000203331132-3221102300320130"></a>

## namespace property — protocol_policer / 113023010321 / 5

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

<a id="canonical-2021102211231313-1130113020113331-0313103203330030-2133332220003001-3221302001130032-1002110023023100-0331113130303202-3133111031231133"></a>

<a id="canonical-0020100211001123-2300200130320302-2111131211032310-0123303011232012-0302322203300120-1221313120332332-2102121201222211-2030323222031200"></a>

## tenant property — protocol_policer / 113023010321 / 6

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

<a id="canonical-0313132323302222-3000022320002023-2011032232212102-2010023001310131-0233112022301220-3332313320101323-0221102302202023-2132313003321210"></a>

## Next pages — protocol_policer / 113023010321 / 7

- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111012212313122-3103330332200131-1300021320130331-1101301222030313-0013320020113230-3233202021121231-1130032032213032-1213012122222332"></a>

## re_acl — re_acl / 302130101001 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- re_acl

<a id="canonical-0111200112123112-1201123020211310-2332023333003303-0133311221133011-1310202220232303-3232110330313223-0123233201033320-0123313122323121"></a>

Type: `"single"`. Computed.

\[OneOf: re\_acl, site\_acl\] Fast ACL for RE. Fast ACL definition for RE.

Upstream description:

Fast ACL definition for RE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-vip_choice": "[\"all_public_vips\",\"default_tenant_vip\",\"selected_tenant_vip\"]"
}
```

OneOf alternatives in this subsection:

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-0111200112123112-1201123020211310-2332023333003303-0133311221133011-1310202220232303-3232110330313223-0123233201033320-0123313122323121)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-0331211011301320-1133001113102311-0001132302121130-0020011101030233-3223233222131221-3013131100322113-2221100330031103-0221021322223110)

Select alternatives according to the provider validators above.

<a id="canonical-1203231013311233-0230212112222323-0100020001210132-2112122312113231-3301212311213320-1221010321100331-3333333001301012-1031120132321232"></a>

## Direct properties — re_acl / 302130101001 / 3

- [all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-2203011201211022-2131111331023122-2102332333121231-0222323002032201-0030130331102230-2323301212102313-3003212000213200-1233103002311222): complete subsection reference.

- [default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-2020031213211322-1211230012130223-3331321000321122-2102021213223112-1213133231021310-2120333101223030-1022002003002332-3102210321131331): complete subsection reference.

- [fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332): complete subsection reference.

- [selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201): complete subsection reference.

<a id="canonical-1021020001223111-2011031301130121-0130000000312130-1111202301122001-2302133231321201-3122132003133123-0022102100102031-3323332211122223"></a>

## Next pages — re_acl / 302130101001 / 4

- [re_acl.all_public_vips](data-sources--fast_acl--reference--group-001.md#canonical-2203011201211022-2131111331023122-2102332333121231-0222323002032201-0030130331102230-2323301212102313-3003212000213200-1233103002311222)
- [re_acl.default_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-2020031213211322-1211230012130223-3331321000321122-2102021213223112-1213133231021310-2120333101223030-1022002003002332-3102210321131331)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2203011201211022-2131111331023122-2102332333121231-0222323002032201-0030130331102230-2323301212102313-3003212000213200-1233103002311222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201021023220002-1223031203011132-0101222332330000-3320013212131032-3220001211201302-1231002301333301-3311212010332202-3033212321232321"></a>

## re_acl.all_public_vips — all_public_vips / 011200323211 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- re_acl.all_public_vips

<a id="canonical-2133302111302000-1212132332231033-1000122031320321-0103133231201122-2101301313331321-1312320032202001-3203222132011011-2131010220022001"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3023122331310010-3023301130033211-1013133003022102-3223030021220211-3001213000030312-2330300222211301-0131321300330022-3233320103021010"></a>

## Direct properties — all_public_vips / 011200323211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312231230332323-2323012233300213-1020102022123130-3030301332011233-1031303232310323-2121321130323303-2113332023231002-1211100010102100"></a>

## Next pages — all_public_vips / 011200323211 / 4

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2020031213211322-1211230012130223-3331321000321122-2102021213223112-1213133231021310-2120333101223030-1022002003002332-3102210321131331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231311103100233-2113022113322012-0110330001021101-1000322000133330-3011230020333031-3203201312211310-3100121100222321-0213202303211010"></a>

## re_acl.default_tenant_vip — default_tenant_vip / 133311230232 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- re_acl.default_tenant_vip

<a id="canonical-2131131000332002-0122101211312223-3302323002100023-2002331320102230-1212322123110332-0012101100301301-2001223131212133-2000231212301023"></a>

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

<a id="canonical-3211030013131022-3301023303110330-2011031012123003-3201312211311313-1222133032021331-3122202113332333-0112002230303133-0323120020323302"></a>

## Direct properties — default_tenant_vip / 133311230232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103323212130021-2323321020003322-0122100103122200-3323111032312301-3002311003303310-3132131133112203-2020332102310010-2320203330120202"></a>

## Next pages — default_tenant_vip / 133311230232 / 4

- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321111101030331-2222300131011210-1002021211123313-1331111033201003-1131312133033300-0103211330223102-3100020310112120-1203230333332011"></a>

## re_acl.fast_acl_rules — fast_acl_rules / 123030133331 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- re_acl.fast_acl_rules

<a id="canonical-0100000011300320-0022100303023020-2103221102110111-3003003221001300-3133033322323323-2001213002220332-1011333111110312-3312023303101223"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match. Defaults to \`\[\]\`. Server applies default when omitted.

Upstream description:

Fast ACL rules to match.

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

<a id="canonical-3002333232133301-2212221333020023-2032211300101030-0123112002020310-3302323113232132-2103312212130032-0200103310220102-1221112311233212"></a>

## Direct properties — fast_acl_rules / 123030133331 / 3

- [action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003): complete subsection reference.

- [ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-2100123010320303-2302121023220133-3212132233230323-2130122330100331-2032311323003020-1001110322331302-1010112112320330-3010332220001012): complete subsection reference.

- [metadata](data-sources--fast_acl--reference--group-001.md#canonical-2303100011321101-1111011200021211-3322011203231333-2201112300022021-2133310033131213-0102210100202111-3320031013011311-0230030110223312): complete subsection reference.

- [port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222): complete subsection reference.

- [prefix](data-sources--fast_acl--reference--group-001.md#canonical-2333121103020211-1132230120300010-3220033202123021-3220113120112312-0301230122021222-3220313031300330-3300303313202000-3130310002001320): complete subsection reference.

<a id="canonical-3322331001302021-3113032213213201-2020301221013101-3320332200223110-3113221310132121-2213133130033212-3002232001030203-2010122323300123"></a>

## Next pages — fast_acl_rules / 123030133331 / 4

- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-2100123010320303-2302121023220133-3212132233230323-2130122330100331-2032311323003020-1001110322331302-1010112112320330-3010332220001012)
- [re_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-2303100011321101-1111011200021211-3322011203231333-2201112300022021-2133310033131213-0102210100202111-3320031013011311-0230030110223312)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222)
- [re_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-2333121103020211-1132230120300010-3220033202123021-3220113120112312-0301230122021222-3220313031300330-3300303313202000-3130310002001320)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120112133313230-0331100010111202-3332202210101001-3211120203232003-2013131202220110-1210331231122101-1001312211321112-2203220021123110"></a>

## re_acl.fast_acl_rules.action — action / 222213300002 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- re_acl.fast_acl_rules.action

<a id="canonical-1030031112100211-2330222201312100-1110010223032221-2110110002121012-3301121222230212-1222021102320202-2000232001220120-2203233102333302"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-0003031230001111-0210001001123101-0011232322121123-3122210300311023-3101231312003213-2220321213222321-0101231220021303-1232131211200022"></a>

## Direct properties — action / 222213300002 / 3

- [policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3200302111031232-2232112022202311-3010232033011023-0111121312201003-2033112331003023-0001323203110023-0123111130111113-2101003132023023): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322): complete subsection reference.

<a id="canonical-1020010331203032-1122010001210212-3032030002012210-1203112232120202-0311330331021030-0002201122122132-3013023031023000-1020032211031303"></a>

<a id="canonical-3001310113131113-2122220000302003-1231012012120002-1332022220222231-1022212323131010-2031220000031022-0123113020203321-1213200001300111"></a>

## simple_action property — action / 222213300002 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

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

<a id="canonical-2303311230322210-1323302102321110-3000123022020212-3323120230330230-2200002103303301-1211011200100031-0313303320310133-2020032203203112"></a>

## Next pages — action / 222213300002 / 5

- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3200302111031232-2232112022202311-3010232033011023-0111121312201003-2033112331003023-0001323203110023-0123111130111113-2101003132023023)
- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3200302111031232-2232112022202311-3010232033011023-0111121312201003-2033112331003023-0001323203110023-0123111130111113-2101003132023023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012321303221300-2201332323333320-1223023311011220-3203313223330333-3331200300000310-2102032222030313-0011211112033120-2322102032220131"></a>

## re_acl.fast_acl_rules.action.policer_action — policer_action / 321312100223 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- re_acl.fast_acl_rules.action.policer_action

<a id="canonical-3130113210321131-0113103232103133-3323001321102031-1222000210130000-3213101122232120-2212033131030001-3003023301011322-2002121211012323"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1012030021133323-0333020101022323-2233211030323203-2312120120110331-0311301021211102-1330312211301302-1000011302011203-2000313322101222"></a>

## Direct properties — policer_action / 321312100223 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-1330123121201312-3233013123321023-0002222321023002-0200202230333013-1111103322110203-3231112031330301-2112133112202033-1330223211333210): complete subsection reference.

<a id="canonical-0033000213000312-3120110210013220-0133132221211111-0010103330123012-0023123311212303-3313313301013030-0131220113121202-3200012013113330"></a>

## Next pages — policer_action / 321312100223 / 4

- [re_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-1330123121201312-3233013123321023-0002222321023002-0200202230333013-1111103322110203-3231112031330301-2112133112202033-1330223211333210)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1330123121201312-3233013123321023-0002222321023002-0200202230333013-1111103322110203-3231112031330301-2112133112202033-1330223211333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031021133110013-2221121113230013-0200211203320331-2123121302310003-3211231223230210-0113313101133231-3302113213323103-3221323103233201"></a>

## re_acl.fast_acl_rules.action.policer_action.ref — ref / 132232300021 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3200302111031232-2232112022202311-3010232033011023-0111121312201003-2033112331003023-0001323203110023-0123111130111113-2101003132023023)
- re_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-1020222223103000-0031232003210320-3312033031212213-3111013313333002-1221023312020002-0033113233300211-1123113011323013-1111322001112332"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-3302013133103330-0100003311201121-2230322333232122-1121022122122313-3311031122220321-3232210332030112-2300011320131021-3131013322123203"></a>

## Direct properties — ref / 132232300021 / 3

<a id="canonical-3103321210100232-1310112131102201-0213323233113231-2330001332332120-2011023113213111-3223022201300203-0303120320321213-0223201322122230"></a>

<a id="canonical-2200012132112213-3002001313312011-3312210102233301-1301200220032132-1001301211303310-1130120122233120-2200321112321333-1113032203313020"></a>

## kind property — ref / 132232300021 / 4

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

<a id="canonical-2333320030111303-2310312233222030-2120102001021020-0011230321321131-3302102131312310-3331332311001232-3321303311313021-3121231122011023"></a>

<a id="canonical-1303112013133301-0030120312303202-0221123133112132-3323333210011332-3111022301332121-3130001222303031-2322203023130230-1130202330133233"></a>

## name property — ref / 132232300021 / 5

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

<a id="canonical-2202131013302013-2220020022030332-1203133103133101-2311022030001222-2201221030331332-3332030302000231-2313220122230211-0033132131113321"></a>

<a id="canonical-2212031213123032-0113233132103121-3302000202122321-2333232213331222-1233213221130203-0020303032000000-0231301310320203-1323303030000130"></a>

## namespace property — ref / 132232300021 / 6

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

<a id="canonical-2131011213030021-2022112210103002-2130000230221131-2131131201132302-0312122312111311-2220211021131320-1002213230313220-3003000330032321"></a>

<a id="canonical-1133002301103300-2032311203030121-0133221223112221-2012012212112203-0301023223232232-3321321320202200-3130122121130002-3232112121331301"></a>

## tenant property — ref / 132232300021 / 7

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

<a id="canonical-3122220021321010-0130130120122122-3023311300100331-0032323300300302-0133301100011030-1230223020132121-3302320331303030-2100302130320321"></a>

<a id="canonical-2213000323032323-2103302230110000-1022233000203313-3311032132103232-3300103203300103-3003221331103331-0331202012202233-3021120112202031"></a>

## uid property — ref / 132232300021 / 8

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

<a id="canonical-3012101031323300-0302201212132331-3023033313021032-0333210003033002-3002033202020030-1211122220110320-2202100320310230-0122030020000013"></a>

## Next pages — ref / 132232300021 / 9

- [re_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3200302111031232-2232112022202311-3010232033011023-0111121312201003-2033112331003023-0001323203110023-0123111130111113-2101003132023023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011020301300231-0010130013130011-0201323321321210-0000231233120301-3303223030213033-2232101010222100-1011120212311121-1320131111113022"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action — protocol_policer_action / 123203310001 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- re_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-3312031033002100-1213031012312212-2122001011200013-2232100022130012-2230032332032321-0300120312013121-0020232120123221-2302002122121300"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1331222130203131-2210312110123131-1123302323022111-1020331210121330-2201231211301002-0012322222322121-1131332002011320-3223103220032223"></a>

## Direct properties — protocol_policer_action / 123203310001 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-3030223213131103-0012231221313333-3120120020312210-1133023200331120-3131012123001312-2212202112203333-1231120200011333-1211333320011033): complete subsection reference.

<a id="canonical-0100333222210002-3022010212221300-2323330302133023-0311023313133233-3323023311113213-1320331101000212-3132320131212220-2101123213211331"></a>

## Next pages — protocol_policer_action / 123203310001 / 4

- [re_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-3030223213131103-0012231221313333-3120120020312210-1133023200331120-3131012123001312-2212202112203333-1231120200011333-1211333320011033)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3030223213131103-0012231221313333-3120120020312210-1133023200331120-3131012123001312-2212202112203333-1231120200011333-1211333320011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010110121001102-3033211310131003-0213213310012031-0103302023223132-1132223021031321-1200230023103332-3102313010020330-3203103130201331"></a>

## re_acl.fast_acl_rules.action.protocol_policer_action.ref — ref / 202132202313 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-2210130221002133-3220300231000003-2330202130010302-1302122313003201-2100313211001131-2020033311210323-3332133332030001-2330001100131003)
- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322)
- re_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-2010012301322011-1330221132003110-1212213103233101-2332002030003233-2322113322110313-2211133221111330-0223223220121203-3300011131233011"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-2212111201001203-3320223121330101-3010231031000130-2310300103100230-0220213201230032-2111313010211001-3133003221000002-0301103213212310"></a>

## Direct properties — ref / 202132202313 / 3

<a id="canonical-0100330020203131-0011021122330312-3213203123332332-0132330201111233-3100012132200201-0202020203321103-1022213311113131-0233220102102103"></a>

<a id="canonical-1122232021203120-0210223022131132-1222131002012031-0032113011020102-2312020021003321-0320121032023303-3113022021211230-3012012330131212"></a>

## kind property — ref / 202132202313 / 4

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

<a id="canonical-0301000230213202-2323230123323023-1100133101321002-1011223201131221-1310020311010213-1300033131001212-1322100312300031-1101233111113232"></a>

<a id="canonical-0223100133322020-3010222221222301-0301113101003102-0111233220202221-0323332222013213-3300233123011031-0330120130222323-3332010032323033"></a>

## name property — ref / 202132202313 / 5

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

<a id="canonical-2213333312301110-1330302002022322-3231321310333113-2132230023220100-0330301223321333-1111203333110300-0132212331330023-0013233021120001"></a>

<a id="canonical-1032200013100333-3120122132210130-3120212313003202-1023310232202301-1013032122011002-3220122323202212-3021010001032310-1201331203121222"></a>

## namespace property — ref / 202132202313 / 6

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

<a id="canonical-1313011203132020-2003210201232001-1110330320033001-2010121023222323-3302321031123031-1300002010003000-1211322120101233-2000103310132213"></a>

<a id="canonical-0320020330222220-3222311230033013-2103101303111232-1113012213000031-2011011332002011-2110333302122122-3213113121220033-1301002301213011"></a>

## tenant property — ref / 202132202313 / 7

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

<a id="canonical-2213212333313330-3302113221312200-1132311031321012-2000002201023120-3023031300023300-1312202331000123-3320002101301301-0011223131313212"></a>

<a id="canonical-2001101001203300-2303120311100323-0131110030123312-2203223233123220-1230333101100001-2202023121201113-0302003002123131-2101001212212312"></a>

## uid property — ref / 202132202313 / 8

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

<a id="canonical-2120131123133031-2223210023331322-0122023120103132-3331231213322102-0030202033131120-0201130233120023-3130013003000023-3322323210113101"></a>

## Next pages — ref / 202132202313 / 9

- [re_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-3022303333123333-2130222103330000-1021102233030130-2311330100031233-0021100322313032-2103231320000001-3213101301221100-0110320211321322)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2100123010320303-2302121023220133-3212132233230323-2130122330100331-2032311323003020-1001110322331302-1010112112320330-3010332220001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031100111220010-0201202322020113-2113200030301310-2111333222311003-1032003000302301-2130301032311022-2313301101001021-1201033333323001"></a>

## re_acl.fast_acl_rules.ip_prefix_set — ip_prefix_set / 221123313202 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- re_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-1021233113320302-1322102330203021-1101123102211210-1021303012113332-1310303011302302-3132200320211112-3033203032030222-2102332312000123"></a>

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

<a id="canonical-0113310233123203-2211121232313021-3122022113201122-1302110132101322-3212111100300332-2303210322331210-2011223232120020-2202230110330302"></a>

## Direct properties — ip_prefix_set / 221123313202 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-1000302321100322-3313203003220321-2222101122101003-0032120002002331-1113203122200112-2232200213320332-1213322223223202-2203311022332110): complete subsection reference.

<a id="canonical-0233313001000331-3203302130111101-2331213030013123-0201202303130111-3113013111111121-2213021111303122-2022123120111320-3232122112321112"></a>

## Next pages — ip_prefix_set / 221123313202 / 4

- [re_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-1000302321100322-3313203003220321-2222101122101003-0032120002002331-1113203122200112-2232200213320332-1213322223223202-2203311022332110)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1000302321100322-3313203003220321-2222101122101003-0032120002002331-1113203122200112-2232200213320332-1213322223223202-2203311022332110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221322220132332-2103201310201303-2230230121131312-1101023000223133-1010111031321312-1030020312320021-3123323003022211-2003120032331111"></a>

## re_acl.fast_acl_rules.ip_prefix_set.ref — ref / 303000331101 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-2100123010320303-2302121023220133-3212132233230323-2130122330100331-2032311323003020-1001110322331302-1010112112320330-3010332220001012)
- re_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-3203203100001000-1331131203020332-3331323330113201-1001310200211023-0000331331021110-2323123031033033-1220312300021233-2230321010303122"></a>

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

<a id="canonical-3112230311110321-0001221020311220-2002011332000030-0011101233101133-2111302032003111-2020313203103232-0211300020033002-3331211001230020"></a>

## Direct properties — ref / 303000331101 / 3

<a id="canonical-1333233321120131-1200020113013320-3323321122302131-3301232110110201-1101211233310011-1301102001032111-2200330131210133-3211330013121131"></a>

<a id="canonical-2121121023202001-2022233322000102-3021232022321022-3311130113023231-1202130121003110-2202023000320130-3001001003111133-0322210330212212"></a>

## kind property — ref / 303000331101 / 4

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

<a id="canonical-0031123033300210-0300013003102200-0122301333302212-2112121211121333-0311132221132303-2033121301201221-3313030031320213-1132202001021211"></a>

<a id="canonical-1033203133033232-0110332123212011-0032003003112110-0330112133212313-2102131231133302-2100123110030103-3023330212221023-2013112230313331"></a>

## name property — ref / 303000331101 / 5

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

<a id="canonical-1331121313302020-3132210311121230-2313122030303131-1301323011331320-0110033202100100-0000220200303222-0332001211023031-3210031233131222"></a>

<a id="canonical-1212003100320103-0132100031013130-1312301111021212-3122310123331323-2002203023110303-1201230213100332-3022023132313122-1001312203210101"></a>

## namespace property — ref / 303000331101 / 6

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

<a id="canonical-1121022332122201-1230223201232212-2033231322231010-1030030002323333-2013222112112222-1201113200131101-3210120132032331-0120022121120023"></a>

<a id="canonical-0322003223211300-1203232102322222-1100220121120322-0300320133103030-2033220120101112-2222120123003032-0203232222323331-0332011122110020"></a>

## tenant property — ref / 303000331101 / 7

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

<a id="canonical-1002303112233123-3211333121223210-3032000223022030-3103201312202000-2002003100220332-3311213030102012-2303312322313310-0223120320220122"></a>

<a id="canonical-0031311002001321-1233101000222101-3021110321321212-2130103013011000-1233111220311333-2320300003333322-1100320112321112-1212102010021232"></a>

## uid property — ref / 303000331101 / 8

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

<a id="canonical-2031331123123313-3303031012113221-3111022322110232-3022201223130201-0331313100123133-2032332120002002-2221201130202101-2022033103322032"></a>

## Next pages — ref / 303000331101 / 9

- [re_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-2100123010320303-2302121023220133-3212132233230323-2130122330100331-2032311323003020-1001110322331302-1010112112320330-3010332220001012)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2303100011321101-1111011200021211-3322011203231333-2201112300022021-2133310033131213-0102210100202111-3320031013011311-0230030110223312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013233023201000-1203232232013202-2132111201122002-2110200232230012-1332121132212002-0210320331202120-0231312113110332-3203203120313102"></a>

## re_acl.fast_acl_rules.metadata — metadata / 123212000112 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- re_acl.fast_acl_rules.metadata

<a id="canonical-1313023111221011-2032322300212120-2321311110222010-0223212323100133-3113100021021112-2233222330323110-0133213211311200-0001333133200200"></a>

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

<a id="canonical-1011220321213210-2300021131212120-0323310311010030-0213112300231110-0011132110033232-1202301100330312-2201210030203213-2202302103231331"></a>

## Direct properties — metadata / 123212000112 / 3

<a id="canonical-0220003320332013-1221003012002120-1231021321021003-2221102030301100-1230121330220330-1322012330201333-0130011133023110-1323220021213210"></a>

<a id="canonical-0230102200303110-3133311011232330-2311321330211001-2202221031022212-1330131230232033-3320022211123013-3133022130121121-1333032321032001"></a>

## description_spec property — metadata / 123212000112 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1023132100211110-2320130001133133-2322010111230132-3231121021202322-2131111132112122-0203121101320110-2313101233320131-3221310122132222"></a>

<a id="canonical-2211321030223130-2111310331332202-1000320213010021-3102321110002230-3302333100201131-0231212221222123-2332133211302312-0312221133000131"></a>

## name property — metadata / 123212000112 / 5

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

<a id="canonical-2212022300131110-1222222121310313-1112100003200002-3113211022311320-0101203222211300-3303131200301020-3201132301310221-3303011022300231"></a>

## Next pages — metadata / 123212000112 / 6

- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321122313312230-0103001311332030-3313033020112103-1313333101311320-3210002301312231-1131113101102023-1033212301313321-0230130231012131"></a>

## re_acl.fast_acl_rules.port — port / 000200331001 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- re_acl.fast_acl_rules.port

<a id="canonical-0020001103331101-2022032023031123-3220311030000222-3313123330311230-3201312123301001-0131201232221122-1211321210223202-3202000103122203"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1322131221030233-2100120320133022-2303012302231322-1013123332123311-0101032213330223-1102011102120100-0322130323132010-2323310300113103"></a>

## Direct properties — port / 000200331001 / 3

- [all](data-sources--fast_acl--reference--group-001.md#canonical-1100000332123101-3111102121200032-1123113310303002-1101311030103202-3000101221133230-2233231301021013-3013330112221101-3020321133031331): complete subsection reference.

- [DNS](data-sources--fast_acl--reference--group-001.md#canonical-3103110301323101-3033312231330233-2312121332203113-2331331210230221-3232131323133130-1321022331331202-1130100212311120-1102232122330302): complete subsection reference.

<a id="canonical-1311302333103132-2211022330031210-0033300203212301-3031131222000221-3222220123321333-2010001003332320-0231231311010332-2121231303232033"></a>

<a id="canonical-0221210012012021-3113120023331113-1011011202130133-1022000220233303-1121112100221021-0012131323033131-1020332313131330-2222303302333323"></a>

## user_defined property — port / 000200331001 / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1303200222122322-0313110221002320-1103221213211230-1233101302320100-0103122003033313-3133322311003102-0221101013012201-1000201323213213"></a>

## Next pages — port / 000200331001 / 5

- [re_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-1100000332123101-3111102121200032-1123113310303002-1101311030103202-3000101221133230-2233231301021013-3013330112221101-3020321133031331)
- [re_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-3103110301323101-3033312231330233-2312121332203113-2331331210230221-3232131323133130-1321022331331202-1130100212311120-1102232122330302)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1100000332123101-3111102121200032-1123113310303002-1101311030103202-3000101221133230-2233231301021013-3013330112221101-3020321133031331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021213303102020-2230203330033232-3221323302013223-1300122031000231-1313010333032132-3200210012102331-2011320200002003-1231212232121320"></a>

## re_acl.fast_acl_rules.port.all — all / 133102131013 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222)
- re_acl.fast_acl_rules.port.all

<a id="canonical-3330211213020021-3212230330323012-0333131022120333-2012300222322313-1333220301021022-3031003132131302-2302131012102102-1131020221203101"></a>

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

<a id="canonical-0110023122212220-1312201332000301-2310101133333101-1120102010123100-0321003222012102-0100122211202112-2300013010100313-2201223111011022"></a>

## Direct properties — all / 133102131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313013231103132-0000011021323320-2233031022111313-0020123332102133-3022313133302331-2311230131301301-2010232013110011-1331310021221211"></a>

## Next pages — all / 133102131013 / 4

- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3103110301323101-3033312231330233-2312121332203113-2331331210230221-3232131323133130-1321022331331202-1130100212311120-1102232122330302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131233300023222-2233230113312222-1313130302123313-0222320311202113-0321113322313222-2331103122031301-2223311222111303-0013311331120111"></a>

## re_acl.fast_acl_rules.port.DNS — DNS / 120132332110 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222)
- re_acl.fast_acl_rules.port.DNS

<a id="canonical-1201021113320311-2300133030221022-1303311112213012-1031311202013302-2102002003002320-1331312123020332-0100223020002303-3322213032212023"></a>

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

<a id="canonical-0021312333122120-0210132120202321-1111233010312322-3102222202101213-0333332130020302-3033003123303322-3233132131010302-1011113013211221"></a>

## Direct properties — DNS / 120132332110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020130110133130-1030200133123012-2223013112033331-2302012012033023-3220011323131201-3102212221323101-0101000123003101-2032321011212221"></a>

## Next pages — DNS / 120132332110 / 4

- [re_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1320003012203321-1303300000203312-3100201332131023-3230031110320012-0330323320013231-0332313231230000-0222212201212201-1000032323003222)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2333121103020211-1132230120300010-3220033202123021-3220113120112312-0301230122021222-3220313031300330-3300303313202000-3130310002001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110002123121211-2322001112331330-0321202233112311-3020203012202030-1201301133321230-2311323122000132-1213231322213011-0321130300112122"></a>

## re_acl.fast_acl_rules.prefix — prefix / 300331102201 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- re_acl.fast_acl_rules.prefix

<a id="canonical-0210123112030200-0021210101232331-0233232311220031-1032312101220320-0102033132020032-1232100330320333-1122302303201111-3221110200112203"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132111020110000-3221101221001103-3110000103330020-1123030111301003-3113212131323301-3322021321101011-3303110301023033-3200301200102313"></a>

## Direct properties — prefix / 300331102201 / 3

<a id="canonical-2111030300311310-1023202123010212-3020133021123230-2132001010011101-0022012030011232-2312200133021222-1103202032112132-0330130322311201"></a>

<a id="canonical-1121030110200102-3302220200113331-3022023010232320-2330030231213301-0203021103200333-1120300302100330-3003222212232322-3113200001200223"></a>

## prefix property — prefix / 300331102201 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-2222220213213030-2021200211111313-1323323011030221-2223311203300003-0131330231020302-0112302321223201-2223213100300111-0132101023101130"></a>

## Next pages — prefix / 300331102201 / 5

- [re_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-1011011302130320-3000222303031132-0001030031032212-1123011021121022-3032231332001101-3020002232321020-3011020322001031-3313212221221332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233300230023302-0230302033033003-2030120203003211-1211102111031200-3223213313333311-3312112210100310-2000110030210122-3120301023211100"></a>

## re_acl.selected_tenant_vip — selected_tenant_vip / 331313032013 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- re_acl.selected_tenant_vip

<a id="canonical-3013301000123133-0310210332330321-0033313030213213-3101023122020222-0101301011013233-3303000122212200-0210011112113303-0000010201213300"></a>

Type: `"single"`. Computed.

Specific Tenant VIP. Select various tenant public VIP(s)

Upstream description:

Select various tenant public VIP(s)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132212202011223-3022132000220130-2213331111001303-0330110321020332-1112303023033000-3330000231001100-2230022323301323-1023212333131330"></a>

## Direct properties — selected_tenant_vip / 331313032013 / 3

<a id="canonical-3121020320211021-3020131013220213-0300331311310130-2231221030021010-0103220301310300-0222032010030130-1110303213233213-0333133321103010"></a>

<a id="canonical-2010333003130211-2103211333321323-3231003312312230-2322322020102331-3030223220212322-2212120332000213-1113131030112321-1330232133333201"></a>

## default_tenant_vip property — selected_tenant_vip / 331313032013 / 4

Type: `"bool"`. Computed.

Include tenant VIP in list of specific VIP(s).

Upstream description:

Include tenant VIP in list of specific VIP(s)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-3313201002022202-3120031112113113-1000330311210231-2033331202302013-0303031100012122-3212213121110102-2100113322103010-3123013221230221): complete subsection reference.

<a id="canonical-1131010221333020-0330301223022100-3313131332211302-2011310132100333-1232120231312313-3311130101212013-1232002321233233-2321310311333332"></a>

## Next pages — selected_tenant_vip / 331313032013 / 5

- [re_acl.selected_tenant_vip.public_ip_refs](data-sources--fast_acl--reference--group-001.md#canonical-3313201002022202-3120031112113113-1000330311210231-2033331202302013-0303031100012122-3212213121110102-2100113322103010-3123013221230221)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3313201002022202-3120031112113113-1000330311210231-2033331202302013-0303031100012122-3212213121110102-2100113322103010-3123013221230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232213330200120-3223303011113003-2313030121123330-2121231022322121-1121010212212120-1200012122132133-3003232222002020-0010110111013001"></a>

## re_acl.selected_tenant_vip.public_ip_refs — public_ip_refs / 332321002320 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [re_acl](data-sources--fast_acl--reference--group-001.md#canonical-1232213301133030-3012322202103331-0331103312312201-2312222320032020-2212200332220323-0120111331230021-1030000332120133-0311333221132221)
- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201)
- re_acl.selected_tenant_vip.public_ip_refs

<a id="canonical-2030223122120233-3322233002100201-2300103030032201-2022233201231132-0123330230123023-1113200122012131-0221233213031332-1021110111000302"></a>

Type: `"list"`. Computed.

Select Public VIP(s). Select additional public VIP(s)

Upstream description:

Select additional public VIP(s)

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0001311313101320-3100102320310311-0012130330222113-0220301101123301-3210323310113023-3201331021232200-2311210133330022-0331132233313112"></a>

## Direct properties — public_ip_refs / 332321002320 / 3

<a id="canonical-3123030023110112-2003230222003031-0123022023000303-1033330311310322-1133221013131320-1330300303010023-0101330112222213-3321001120323200"></a>

<a id="canonical-0032331303230032-3120000233203021-0023010331333000-1122312021220222-3212131120110322-0213113323000010-2200313001311203-3013001331221012"></a>

## name property — public_ip_refs / 332321002320 / 4

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

<a id="canonical-3220012021200103-2122102201022130-3220023122122302-1023300321220100-2133003331022313-0110003333222301-1312033021232321-2020223221202021"></a>

<a id="canonical-1013110331212133-1101222001100130-3301332031330331-2033310301002022-1112213321330202-2132323230313232-0033121102033300-0020111021210100"></a>

## namespace property — public_ip_refs / 332321002320 / 5

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

<a id="canonical-2132001133323113-1332233331310012-2332013222313003-3131103212230003-0000232020103123-2320213302211131-3110030330101102-2202111300230112"></a>

<a id="canonical-3003230313023333-1312222020102100-0030132202231001-3122330113002201-3213202132302321-1010032203123110-3201033200100123-0321230213312132"></a>

## tenant property — public_ip_refs / 332321002320 / 6

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

<a id="canonical-0210311103333022-3100311110112222-0131300113120310-0031020201203011-0203013120311010-0032332330111320-3110112312131221-3220121313003011"></a>

## Next pages — public_ip_refs / 332321002320 / 7

- [re_acl.selected_tenant_vip](data-sources--fast_acl--reference--group-001.md#canonical-1210100311322300-1002001032202203-0012331030302121-3022301020332021-1322201112200222-3300122330231313-3030120323233302-3001322221233201)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030030210211321-1100103331233112-2000200100210120-3310231212313203-2233100112313033-3023022023212002-3322010213120320-1033031310120201"></a>

## site_acl — site_acl / 323232311321 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- site_acl

<a id="canonical-0331211011301320-1133001113102311-0001132302121130-0020011101030233-3223233222131221-3013131100322113-2221100330031103-0221021322223110"></a>

Type: `"single"`. Computed.

Fast ACL for Site. Fast ACL definition for Site.

Upstream description:

Fast ACL definition for Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]",
  "x-ves-oneof-field-vip_choice": "[\"all_services\",\"interface_services\",\"vip_services\"]"
}
```

<a id="canonical-2011032211123020-1001233011022120-0013103102111331-2311333220310003-0220131131311331-3101212333223333-2333100013211112-3111133113030110"></a>

## Direct properties — site_acl / 323232311321 / 3

- [all_services](data-sources--fast_acl--reference--group-001.md#canonical-0131103113020223-3003011302221023-1032232001320031-2323103011330031-1320000132322010-3331310320013013-0111101003310233-1303231130210213): complete subsection reference.

- [fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023): complete subsection reference.

- [inside_network](data-sources--fast_acl--reference--group-001.md#canonical-1000303121232311-0013131021311012-2201230332233303-2002202010322001-3113023331221202-1222302333002112-2100312033213030-3031221301030013): complete subsection reference.

- [interface_services](data-sources--fast_acl--reference--group-001.md#canonical-1302223030123000-0013022032301302-0120120233110210-0300333210232103-2112220312011310-3030331331300130-3301213312222103-3301321210230113): complete subsection reference.

- [outside_network](data-sources--fast_acl--reference--group-001.md#canonical-1213000220331003-2212012300032022-1301023232032011-0131133003103231-3032133030010100-1112021300302131-3202023310301110-0122202102113020): complete subsection reference.

- [vip_services](data-sources--fast_acl--reference--group-001.md#canonical-2330310120132320-0030113032103000-2100202300223313-3321131323320221-1031013320030201-1200302013101000-2330012121302112-1132210100002332): complete subsection reference.

<a id="canonical-0002223211330330-0303323201100121-2321213021003331-2321023100230323-1231323000132013-1123230102132010-2102022311312133-2113211322022013"></a>

## Next pages — site_acl / 323232311321 / 4

- [site_acl.all_services](data-sources--fast_acl--reference--group-001.md#canonical-0131103113020223-3003011302221023-1032232001320031-2323103011330031-1320000132322010-3331310320013013-0111101003310233-1303231130210213)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.inside_network](data-sources--fast_acl--reference--group-001.md#canonical-1000303121232311-0013131021311012-2201230332233303-2002202010322001-3113023331221202-1222302333002112-2100312033213030-3031221301030013)
- [site_acl.interface_services](data-sources--fast_acl--reference--group-001.md#canonical-1302223030123000-0013022032301302-0120120233110210-0300333210232103-2112220312011310-3030331331300130-3301213312222103-3301321210230113)
- [site_acl.outside_network](data-sources--fast_acl--reference--group-001.md#canonical-1213000220331003-2212012300032022-1301023232032011-0131133003103231-3032133030010100-1112021300302131-3202023310301110-0122202102113020)
- [site_acl.vip_services](data-sources--fast_acl--reference--group-001.md#canonical-2330310120132320-0030113032103000-2100202300223313-3321131323320221-1031013320030201-1200302013101000-2330012121302112-1132210100002332)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0131103113020223-3003011302221023-1032232001320031-2323103011330031-1320000132322010-3331310320013013-0111101003310233-1303231130210213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112111300121133-1003212032320102-1200213002011010-0103332113110320-2031203223223303-1133231210122103-0213002223300031-2032230203223232"></a>

## site_acl.all_services — all_services / 102223213121 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.all_services

<a id="canonical-2130333322212111-1221003031120023-2233211022333332-3333213310311121-2321100213103131-0002210311030230-0112200131323010-3202213113110312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all services.

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

<a id="canonical-1101111223123101-1312013312200230-3212011222011111-2331030301100033-2231130101113103-2303203221133213-3332222101010001-2023002323033301"></a>

## Direct properties — all_services / 102223213121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103302002322100-1102232213001010-2101132131211221-3002300320013000-3033223111300302-0233020102110003-2111331131113200-3222131012303321"></a>

## Next pages — all_services / 102223213121 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110310120233122-0133112320303332-1032012232020310-3222321130020332-3021131100210012-3331123220332213-3333201310003120-0231220112303003"></a>

## site_acl.fast_acl_rules — fast_acl_rules / 012101113133 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.fast_acl_rules

<a id="canonical-3112130121101333-3220210230032311-1023311220100200-2331222302211111-0223200123221321-3223013022223120-1003030130201003-3301323333131013"></a>

Type: `"list"`. Computed.

Rules. Fast ACL rules to match.

Upstream description:

Fast ACL rules to match.

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

<a id="canonical-2123003103332312-3133130201110311-0220200202032213-0322301200112023-0120322001133123-0331312120302303-0300023201031032-2302123033021202"></a>

## Direct properties — fast_acl_rules / 012101113133 / 3

- [action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332): complete subsection reference.

- [ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-0113312003201033-2302331301211213-0013302112313220-2001201332113130-2231223222130322-0230032211032223-3013101130010010-0303100230113213): complete subsection reference.

- [metadata](data-sources--fast_acl--reference--group-001.md#canonical-0130300102303311-0211102101203002-2220131323231320-0200212032310133-1110122120130030-2033011300002001-1031211100312011-3330301301012113): complete subsection reference.

- [port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312): complete subsection reference.

- [prefix](data-sources--fast_acl--reference--group-001.md#canonical-3032222320121131-1200233103132111-1211223013221331-3032033230313000-3302333110003003-2131103313133032-2332113111233312-3100002230232033): complete subsection reference.

<a id="canonical-0220110102032121-1202132333113002-2100231111123200-3233200200121321-2220003012110000-0121310213333313-3000122221131310-3300311102312230"></a>

## Next pages — fast_acl_rules / 012101113133 / 4

- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-0113312003201033-2302331301211213-0013302112313220-2001201332113130-2231223222130322-0230032211032223-3013101130010010-0303100230113213)
- [site_acl.fast_acl_rules.metadata](data-sources--fast_acl--reference--group-001.md#canonical-0130300102303311-0211102101203002-2220131323231320-0200212032310133-1110122120130030-2033011300002001-1031211100312011-3330301301012113)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312)
- [site_acl.fast_acl_rules.prefix](data-sources--fast_acl--reference--group-001.md#canonical-3032222320121131-1200233103132111-1211223013221331-3032033230313000-3302333110003003-2131103313133032-2332113111233312-3100002230232033)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113021303001313-3030033232103213-3002213220030010-1011133032223023-0013331311311023-3111203020232212-0021233033031133-3030321000001112"></a>

## site_acl.fast_acl_rules.action — action / 330031003220 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- site_acl.fast_acl_rules.action

<a id="canonical-3222000113123220-3030201030001021-2203031201103320-0132120332022313-2131303200210100-2122220111100302-1131222323120123-1210223131201321"></a>

Type: `"single"`. Computed.

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Upstream description:

FastAclRuleAction specifies possible action to be applied on traffic, possible action include
dropping, forwarding or ratelimiting the traffic.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action": "[\"policer_action\",\"protocol_policer_action\",\"simple_action\"]"
}
```

<a id="canonical-1103220220323120-2112021001120022-2203033123100200-1132221020112221-0120300202232301-3030131003222302-0023233200311200-1022213110031310"></a>

## Direct properties — action / 330031003220 / 3

- [policer_action](data-sources--fast_acl--reference--group-001.md#canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332): complete subsection reference.

- [protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2203021310031132-3231310303230323-3332231003202203-2003001023211201-0001230332132131-1223131211002320-2022330110022120-0231103213022113): complete subsection reference.

<a id="canonical-3013211213212132-1220232102322202-3223211103112232-2021013012203133-3302223302133201-2220023333222210-3110120220312102-0333112020111002"></a>

<a id="canonical-1111300220311012-1021331031230211-3120202123313111-3020322212221223-1233031220210010-1030302013113310-3203311103201322-3003320310002210"></a>

## simple_action property — action / 330031003220 / 4

Type: `"string"`. Computed.

\[Enum: DENY|ALLOW\] FastAclRuleSimpleAction specifies simple action like PASS or DENY Drop the
traffic Forward the traffic. Possible values are \`DENY\`, \`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

FastAclRuleSimpleAction specifies simple action like PASS or DENY

Drop the traffic Forward the traffic.

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

<a id="canonical-0023311122102012-3101232023032020-1023221303113032-2211001332030303-0013113200133013-3123033113030131-0223323321023200-0213310121310031"></a>

## Next pages — action / 330031003220 / 5

- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332)
- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2203021310031132-3231310303230323-3332231003202203-2003001023211201-0001230332132131-1223131211002320-2022330110022120-0231103213022113)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333333220221000-2010130023311311-1212131310020033-1323022020232202-1220223330123130-0021123300313301-2111313123003222-1130021322131202"></a>

## site_acl.fast_acl_rules.action.policer_action — policer_action / 321001233000 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- site_acl.fast_acl_rules.action.policer_action

<a id="canonical-2202232320122231-0132110223012132-2022202102211211-0010332211231023-0221220320011300-1121002302103311-2203031003020201-1000121230321300"></a>

Type: `"single"`. Computed.

Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0003321332221313-1220032331230323-0213123310012300-1012200302232310-2110221300331021-2323301202102331-1121310120300302-3031112213233023"></a>

## Direct properties — policer_action / 321001233000 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-2033003220013313-3332330100330021-3100020312203020-2311322010210221-1122333331321213-3323020101131333-3101313230113003-2203123210310020): complete subsection reference.

<a id="canonical-3210111231223010-3300223201301030-3211020203212332-2031202003300012-2111110303031212-3212030120121010-3300021113312021-0003021022011122"></a>

## Next pages — policer_action / 321001233000 / 4

- [site_acl.fast_acl_rules.action.policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-2033003220013313-3332330100330021-3100020312203020-2311322010210221-1122333331321213-3323020101131333-3101313230113003-2203123210310020)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2033003220013313-3332330100330021-3100020312203020-2311322010210221-1122333331321213-3323020101131333-3101313230113003-2203123210310020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113023201131122-3132110321111302-1202003202220031-1032100210201322-1123112222201210-0121231013002222-3103200130032332-3003112323211231"></a>

## site_acl.fast_acl_rules.action.policer_action.ref — ref / 321121131000 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332)
- site_acl.fast_acl_rules.action.policer_action.ref

<a id="canonical-0322133030132010-0231300112030001-1000201220213212-2133321301331301-0002103103203301-0022120212300322-3032302001210130-0310022232202200"></a>

Type: `"list"`. Computed.

Reference. A policer direct reference.

Upstream description:

A policer direct reference.

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

<a id="canonical-1233101132312212-2013111233233131-2130100230313202-3022012102132221-0110032202313301-0123200200013203-3312201201111333-3330301010221011"></a>

## Direct properties — ref / 321121131000 / 3

<a id="canonical-2022120301101132-0323033012113203-1313221203333212-3133103131031020-2302023011111320-0033220200333122-2310232030011210-2310101000222102"></a>

<a id="canonical-2200232032013203-2120313203101220-1121213033301310-3131320301020331-3311112333223332-1112023023111231-1201003001212023-0132233333112221"></a>

## kind property — ref / 321121131000 / 4

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

<a id="canonical-2303313331102102-1202111023121000-1220313211220121-3001010222033020-1020221011120233-3201003222001332-2231123302320202-0321210023121233"></a>

<a id="canonical-2133312003012023-1300123000120232-0133000113201022-2330023201030201-3110020113212103-0012022321212232-0130223201222100-0330322013230013"></a>

## name property — ref / 321121131000 / 5

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

<a id="canonical-1121023133100300-3020323320030210-2111312323013221-1121303231220003-2321110130101321-3012312123131002-3320312122102003-1122213121231221"></a>

<a id="canonical-3232112303031123-3110300000323232-1111101322201010-1000233321132210-2222023232131010-0132322233133103-0232133002031231-1221132111012213"></a>

## namespace property — ref / 321121131000 / 6

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

<a id="canonical-1111111101322223-2212301011310232-0023211011102122-1033321120121223-0230231120033013-0000201323221332-0010022320020321-0002320300200102"></a>

<a id="canonical-0213313310301221-3001001233101131-1211312311123332-0311102132002132-2112000122322120-0101130220233231-1211211212233130-0100331331223120"></a>

## tenant property — ref / 321121131000 / 7

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

<a id="canonical-1031302120000230-0033121020220022-1011311033221300-2232021123130312-0120131122011021-1221211322331331-3031223202321213-1311311300232013"></a>

<a id="canonical-0313020031133213-0133333131220031-3012222012210213-2112021032122112-1311302312200223-1203023211332010-2210121210021113-1202021023102101"></a>

## uid property — ref / 321121131000 / 8

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

<a id="canonical-2110200002333102-1110222022011020-1313123333120002-1233012200022313-3111031010230223-3201110201231121-3121312121022313-2230101030303322"></a>

## Next pages — ref / 321121131000 / 9

- [site_acl.fast_acl_rules.action.policer_action](data-sources--fast_acl--reference--group-001.md#canonical-1000110003033231-0123100213212023-2213112112021320-1001123013012000-3323311311022333-2101201032033201-3021301213302113-1001220323323332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2203021310031132-3231310303230323-3332231003202203-2003001023211201-0001230332132131-1223131211002320-2022330110022120-0231103213022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211230320030103-0321133122302233-3113311001302221-3322113031133313-2133312233130103-0301313303033120-3111202221301203-2202233213020330"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action — protocol_policer_action / 203211312033 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- site_acl.fast_acl_rules.action.protocol_policer_action

<a id="canonical-2333110203212222-3111321200133102-3223111221110021-2233221221303132-1312321231012121-0033101012030030-2210102213230210-3230321213332331"></a>

Type: `"single"`. Computed.

Protocol Policer Reference. Reference to policer object.

Upstream description:

Reference to policer object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200331012112001-2312121001021301-0313233232012322-2213300202130222-2202012001223121-1122030033122312-2332133311131033-0322120000023313"></a>

## Direct properties — protocol_policer_action / 203211312033 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-0312312120030130-1220003320310230-3323303013333033-2120013302111031-0131023032210031-0031211331012223-2101000332133001-1120323320012002): complete subsection reference.

<a id="canonical-1223031200313122-2012302323200310-3111221113102031-3023000123320010-1112131213010313-0023210332002122-3332131220012033-3203021010300001"></a>

## Next pages — protocol_policer_action / 203211312033 / 4

- [site_acl.fast_acl_rules.action.protocol_policer_action.ref](data-sources--fast_acl--reference--group-001.md#canonical-0312312120030130-1220003320310230-3323303013333033-2120013302111031-0131023032210031-0031211331012223-2101000332133001-1120323320012002)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0312312120030130-1220003320310230-3323303013333033-2120013302111031-0131023032210031-0031211331012223-2101000332133001-1120323320012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030331003122022-3313333123003003-3310201133020110-0012121101012001-2010320302301031-0232203033202101-3121322021300322-1310003320002222"></a>

## site_acl.fast_acl_rules.action.protocol_policer_action.ref — ref / 313101332021 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.action](data-sources--fast_acl--reference--group-001.md#canonical-0021121121010230-2321333120320031-2221223110112032-3330203231200002-2132020133232032-1030223130311213-3012301213203022-1222210222230332)
- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2203021310031132-3231310303230323-3332231003202203-2003001023211201-0001230332132131-1223131211002320-2022330110022120-0231103213022113)
- site_acl.fast_acl_rules.action.protocol_policer_action.ref

<a id="canonical-0120101223313111-2330003231132132-0121113101322331-1022330201312022-1010323133112121-1123302100020310-1003003031033023-2033313101301011"></a>

Type: `"list"`. Computed.

Protocol policer Reference. Reference to protocol policer object.

Upstream description:

Reference to protocol policer object.

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

<a id="canonical-3103301010230120-2033320111321013-2120211023002230-1021012300011201-1300323020102031-2023002111310320-2220201023330331-1201302301101020"></a>

## Direct properties — ref / 313101332021 / 3

<a id="canonical-0112233320310103-0032023011123323-3330102212313101-3211032311033223-1032200123032321-3333011030033033-3333113100322103-3031320312122121"></a>

<a id="canonical-1330100222332302-1211002203200311-2102103121220232-2010010300320132-1123112222100121-2102220121322210-1220320003321131-3012331113102100"></a>

## kind property — ref / 313101332021 / 4

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

<a id="canonical-2022103232031202-3003233231120320-1201303230031222-1200111332002101-2121322023120212-2310113212213203-1221220013112203-0023300332213010"></a>

<a id="canonical-3213110003110310-0300231230211100-1312033210310100-0320322211022321-0020300223100022-0300223330323122-1211220130201320-2323132123132011"></a>

## name property — ref / 313101332021 / 5

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

<a id="canonical-1021300113012122-1001121333012211-0133212022012021-3211033103020031-1302012311002312-3022002103211031-0021323211213003-1301321122013220"></a>

<a id="canonical-0323321123130321-2100113131133030-1233101211103131-2313100331231012-3120302130200302-0312103130130333-3321111003112303-1101203332302011"></a>

## namespace property — ref / 313101332021 / 6

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

<a id="canonical-3310131330020231-0212000321301301-1100011133213121-2320210223221022-0230113120203121-0303311133030211-0231131011300031-2110101101122013"></a>

<a id="canonical-3111301302033232-2302212021200202-3230300221212033-1222123011300232-1133033333332120-2203113312321321-3130012130222002-3310310101122320"></a>

## tenant property — ref / 313101332021 / 7

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

<a id="canonical-0110021232033133-1332301110100221-3032221312322021-2323131230203100-3302103323201213-3000202020222033-0130123023002301-3211221101210130"></a>

<a id="canonical-3013013221130113-1032320222232121-0213121210002001-3233013330301010-0003210230111033-2031301122202103-2201312001103300-0213013313030201"></a>

## uid property — ref / 313101332021 / 8

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

<a id="canonical-0101131220221301-2322333330333132-1323113313003132-2131031122301000-2322122311332012-3321220220001120-2233301332102320-2030112203332011"></a>

## Next pages — ref / 313101332021 / 9

- [site_acl.fast_acl_rules.action.protocol_policer_action](data-sources--fast_acl--reference--group-001.md#canonical-2203021310031132-3231310303230323-3332231003202203-2003001023211201-0001230332132131-1223131211002320-2022330110022120-0231103213022113)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0113312003201033-2302331301211213-0013302112313220-2001201332113130-2231223222130322-0230032211032223-3013101130010010-0303100230113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202233132111112-1131103202203122-1020310333012001-2331232202011213-0313211121030222-3033130310003212-3210321133100010-0011122022223101"></a>

## site_acl.fast_acl_rules.ip_prefix_set — ip_prefix_set / 131320312020 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- site_acl.fast_acl_rules.ip_prefix_set

<a id="canonical-2311003303222221-2012332033100033-1320020301331012-0001120203212123-1223003133213022-1323012120100311-3012332010002231-1113233312202200"></a>

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

<a id="canonical-1230000201210032-0103301223012030-3223223023302020-3121022330212132-0211001230210002-3113003011310223-2102302311231220-2223030000213231"></a>

## Direct properties — ip_prefix_set / 131320312020 / 3

- [ref](data-sources--fast_acl--reference--group-001.md#canonical-3300222211103001-1323011012123203-1300202213103323-3131323221230100-3321000030011111-3031132000133332-2201003233220010-3013120212121032): complete subsection reference.

<a id="canonical-0332120020101231-2113303221323012-0323113333122133-1311101131311030-2322310313132001-0112331011300003-2013200203003203-3013302211323333"></a>

## Next pages — ip_prefix_set / 131320312020 / 4

- [site_acl.fast_acl_rules.ip_prefix_set.ref](data-sources--fast_acl--reference--group-001.md#canonical-3300222211103001-1323011012123203-1300202213103323-3131323221230100-3321000030011111-3031132000133332-2201003233220010-3013120212121032)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3300222211103001-1323011012123203-1300202213103323-3131323221230100-3321000030011111-3031132000133332-2201003233220010-3013120212121032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132311230220201-2001033210233023-2231011032112221-1303131332321201-2301030013022100-3212100310310200-1100301301310201-1032322332132231"></a>

## site_acl.fast_acl_rules.ip_prefix_set.ref — ref / 121001021033 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-0113312003201033-2302331301211213-0013302112313220-2001201332113130-2231223222130322-0230032211032223-3013101130010010-0303100230113213)
- site_acl.fast_acl_rules.ip_prefix_set.ref

<a id="canonical-3031132210123230-0331033010211312-1233131031032222-1231121030313321-1100310131003120-2222301311111000-2322302101211000-3311011130022031"></a>

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

<a id="canonical-3333110030320133-1130020330211230-0130302110111223-0121323330323103-3302232010132112-1000312301201303-0310223312213202-0313120112003131"></a>

## Direct properties — ref / 121001021033 / 3

<a id="canonical-2022213201300111-3031033301303031-0120000112313001-3230231323332113-1320023130020013-2000130210212320-3300133220301312-3320010332033321"></a>

<a id="canonical-0201012102320303-0301122033032202-0201211012033330-0302310232122020-1312102133210200-0030010222221202-1033113301210013-1211212020013001"></a>

## kind property — ref / 121001021033 / 4

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

<a id="canonical-3221012333020303-2001033010232320-3330320323200303-3200132021222222-1022200200313113-0301220232113121-2203213011330123-3112102331220230"></a>

<a id="canonical-3023130003321232-3032223222100323-1300001311312132-0300320020032322-0321311323002312-0211011202023333-2132311133100123-1312330012220223"></a>

## name property — ref / 121001021033 / 5

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

<a id="canonical-0210302101311332-3102131013130201-3013301023022122-1121002022231033-3213123212012023-3100233221220123-1112030133311231-3311111203102013"></a>

<a id="canonical-2011313331210233-0231211213321313-0322330210301011-2200332321123030-0123003221110133-3000231132100331-0313011300331233-3022232110003002"></a>

## namespace property — ref / 121001021033 / 6

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

<a id="canonical-3330112133313010-1332000203033302-2121100200031103-1331132033331232-2113023202001322-3022012013122300-3210012023020223-2001232313102023"></a>

<a id="canonical-0320200213233010-2113201003323230-1331023122232220-0031220203133101-3322100130302003-0212001012033202-3231221110313133-1220333132010211"></a>

## tenant property — ref / 121001021033 / 7

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

<a id="canonical-3003320101133013-2011320230321201-1211021030230001-3032331231033210-1301120112120230-0310230030012111-1211031011211102-1002320103133032"></a>

<a id="canonical-1002013121000333-2201111320112220-0310321120011210-1230232130020321-2203320223121301-2213120210113220-0300012033111301-1102230231001102"></a>

## uid property — ref / 121001021033 / 8

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

<a id="canonical-0132200113323200-2032032021231013-0321333011302211-2101022002210320-3022311102031333-3233133300222210-0022200032323120-1312022203032301"></a>

## Next pages — ref / 121001021033 / 9

- [site_acl.fast_acl_rules.ip_prefix_set](data-sources--fast_acl--reference--group-001.md#canonical-0113312003201033-2302331301211213-0013302112313220-2001201332113130-2231223222130322-0230032211032223-3013101130010010-0303100230113213)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-0130300102303311-0211102101203002-2220131323231320-0200212032310133-1110122120130030-2033011300002001-1031211100312011-3330301301012113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212013230303220-1033200323032102-2210333222100213-3232132101233023-1200010201001121-3303230202300001-2302220131310120-1123212002222333"></a>

## site_acl.fast_acl_rules.metadata — metadata / 013320201331 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- site_acl.fast_acl_rules.metadata

<a id="canonical-2013312111111221-2302233123230220-0102311101231200-3000033110331112-2030222001120022-2121223013232302-1301311200121130-1132122202232203"></a>

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

<a id="canonical-1110331200131131-2202210223200010-3303303202321200-0121102223333122-0123120221323323-3210131311132132-2102320232000103-2333320030300131"></a>

## Direct properties — metadata / 013320201331 / 3

<a id="canonical-3021300300101112-2301123012122122-1001200132100123-2132100023303101-3212103331310100-1032132322030121-3101000221310201-0332230003330122"></a>

<a id="canonical-2211022202030011-1023011002100030-2223102212123003-1133230312031330-3301003333102210-0022000123331103-2210212322023231-0323211000130013"></a>

## description_spec property — metadata / 013320201331 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3031201100303330-1313033313132033-1221012222001132-3223220320000001-3312132131132032-1102311033002002-3111313121013211-3020023003010002"></a>

<a id="canonical-1133100110030332-1211212113131213-3220021121111012-3101133210101202-2002303023332121-2022202321232100-2101301012123301-2103231000133311"></a>

## name property — metadata / 013320201331 / 5

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

<a id="canonical-0230021022131231-2103012322320131-0223103320203102-0200301100133303-0311030123032332-3232202132112021-3221202223003333-0212301303331132"></a>

## Next pages — metadata / 013320201331 / 6

- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221322103211021-0223131223032021-2013003033231231-3001302030330103-1022211132222310-0000331323133303-0220311301000231-0000332112131103"></a>

## site_acl.fast_acl_rules.port — port / 203010201323 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- site_acl.fast_acl_rules.port

<a id="canonical-3010233133203020-2020303223311301-2023121133310201-0213200201221311-3101020121033203-2332223133223111-1003013002010112-2312102103030120"></a>

Type: `"list"`. Computed.

Source Ports. L4 port numbers to match.

Upstream description:

L4 port numbers to match.

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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1010321222100031-2201233002123321-1301032332130233-3120103131202322-2333111312222120-0310022100012322-2333131112311011-0333231233203001"></a>

## Direct properties — port / 203010201323 / 3

- [all](data-sources--fast_acl--reference--group-001.md#canonical-3022212131120001-2102010032101332-0322002322222211-2301102023313001-2321312112111012-1132200330123030-2300001201123332-2301233201313111): complete subsection reference.

- [DNS](data-sources--fast_acl--reference--group-001.md#canonical-2210221201300130-0210230103110320-2010103131030202-1332212003002231-1003122031231132-3332122032212213-3313200211113312-3303213121321321): complete subsection reference.

<a id="canonical-3111213233323003-0332302013030333-2301010101032313-0110312030131010-3330023331210310-0210010011102010-0222333332200110-2231123202013110"></a>

<a id="canonical-0332033111303230-3032013010211202-1301122132002110-2332231201221103-3311101133212301-3221231212220012-0131122322132033-2333310232110200"></a>

## user_defined property — port / 203010201323 / 4

Type: `"number"`. Computed.

Exclusive with \[all DNS\] Matches the user defined port.

Upstream description:

Exclusive with \[all DNS\] Matches the user defined port.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0220313212230322-0013201011021030-0112132121202120-2132121222100020-2012210223303211-1102020321331123-2320021211321222-0000202223031230"></a>

## Next pages — port / 203010201323 / 5

- [site_acl.fast_acl_rules.port.all](data-sources--fast_acl--reference--group-001.md#canonical-3022212131120001-2102010032101332-0322002322222211-2301102023313001-2321312112111012-1132200330123030-2300001201123332-2301233201313111)
- [site_acl.fast_acl_rules.port.dns](data-sources--fast_acl--reference--group-001.md#canonical-2210221201300130-0210230103110320-2010103131030202-1332212003002231-1003122031231132-3332122032212213-3313200211113312-3303213121321321)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3022212131120001-2102010032101332-0322002322222211-2301102023313001-2321312112111012-1132200330123030-2300001201123332-2301233201313111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202231302030232-3112211311013032-2103332120001002-2330033031130123-0111102223102230-1032223020303310-1330303130101213-2133003111233103"></a>

## site_acl.fast_acl_rules.port.all — all / 220101112221 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312)
- site_acl.fast_acl_rules.port.all

<a id="canonical-1211123000302023-1310300310123332-2303222023300301-3111122001110312-0122022101233230-2012303203102111-2113221320121021-3113021233233122"></a>

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

<a id="canonical-0001131303323022-2222212030332122-2013221302222102-3022303222302031-2002001301120021-2020101303313202-1310213123120120-2130321100001033"></a>

## Direct properties — all / 220101112221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201101221300303-0321212122332332-0213213000323231-1221132112110132-1021023311332203-0123113120003102-3011233210322323-3123112323200310"></a>

## Next pages — all / 220101112221 / 4

- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2210221201300130-0210230103110320-2010103131030202-1332212003002231-1003122031231132-3332122032212213-3313200211113312-3303213121321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111120203130021-3131113110032103-2200002133012221-2320030030300322-3302222211133321-3132112203330030-1012202102023100-3130312322110211"></a>

## site_acl.fast_acl_rules.port.DNS — DNS / 111132303332 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312)
- site_acl.fast_acl_rules.port.DNS

<a id="canonical-3032231211331131-0010322203311121-1110021121013213-0321331130220222-0322003003210133-2213000013222332-1032312300102001-3002102030121330"></a>

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

<a id="canonical-3100301002333300-2110120032000021-3011031302333003-0213022210130120-3321333210232323-3202231033113103-2203101003210002-2132332033011231"></a>

## Direct properties — DNS / 111132303332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023000010322002-0211333210212321-3110330303310322-3212332103100003-3002111212300003-3301223330032133-0120021322101323-3322011010032121"></a>

## Next pages — DNS / 111132303332 / 4

- [site_acl.fast_acl_rules.port](data-sources--fast_acl--reference--group-001.md#canonical-1321111102121011-2233023323232130-1123312032111111-1001111220113123-0113120101222223-3023120230033130-2231320323233002-2031220121112312)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-3032222320121131-1200233103132111-1211223013221331-3032033230313000-3302333110003003-2131103313133032-2332113111233312-3100002230232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312211330120130-1121100233223021-0020123333220311-1312102323223321-0101230210230113-0002123203200313-2333123200230213-3222131223131123"></a>

## site_acl.fast_acl_rules.prefix — prefix / 021310103320 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- site_acl.fast_acl_rules.prefix

<a id="canonical-0223300323001312-2103122200321122-3022200310300130-0313132202132333-2023113230032120-2221033001022301-3100002320033133-2112001311223003"></a>

Type: `"single"`. Computed.

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Upstream description:

List of IP Address prefixes. Prefix must contain both prefix and prefix-length The list can contain
mix of both IPv4 and IPv6 prefixes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3020202211120013-0331211020312020-1321110303231012-3121012131222112-0210220131103231-0230201303213020-2110301030311101-1021121221331113"></a>

## Direct properties — prefix / 021310103320 / 3

<a id="canonical-1311201310001202-1132332022012113-2121023123330322-2112122310112122-0232000301013110-2222131222210210-3223320003010211-1122210203301121"></a>

<a id="canonical-3121101123220030-3220132231212111-1023332300101020-2133033302102313-1312001030022112-3300013233310001-3030212000022220-3323231013322201"></a>

## prefix property — prefix / 021310103320 / 4

Type: `["list", "string"]`. Computed.

IP Address prefix in string format. String must contain both prefix and prefix-length.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-0201301320100111-0030003121112211-3210100131333201-1013233130010231-1312301132322322-0202201131000000-1102203130200002-1111211130230302"></a>

## Next pages — prefix / 021310103320 / 5

- [site_acl.fast_acl_rules](data-sources--fast_acl--reference--group-001.md#canonical-0033210120011020-2212331003113120-0231233101332301-3330033303222300-0232321100222210-2233201323133113-0131020302232121-1202112122300023)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1000303121232311-0013131021311012-2201230332233303-2002202010322001-3113023331221202-1222302333002112-2100312033213030-3031221301030013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001302010132220-2011221321200112-0012321012002302-3033222023000302-2301110332023230-2323203220311103-2011322202200032-3122132122220113"></a>

## site_acl.inside_network — inside_network / 133120002120 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.inside_network

<a id="canonical-0031222033203333-0100330213230111-1310010003212100-1102131313202112-1113120133332330-0220021231330030-0312213120111333-3023332122000032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-3023131133122220-0300120030122133-2201333231203331-1030120332301332-0223312103322021-3312032130110202-3232100110233233-3112010211321013"></a>

## Direct properties — inside_network / 133120002120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122323131102103-0233113231110013-1001332112213233-3022131213310230-3013321012101130-1330231330211210-1033210031300000-3033121330022012"></a>

## Next pages — inside_network / 133120002120 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1302223030123000-0013022032301302-0120120233110210-0300333210232103-2112220312011310-3030331331300130-3301213312222103-3301321210230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002322223101323-3012031222233031-2023021210301031-1313221031030221-1100330230122002-2133230120230123-3010311031102123-1310020232310211"></a>

## site_acl.interface_services — interface_services / 311221301112 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.interface_services

<a id="canonical-1021131310320323-1101133233222213-2310013232033020-3302122102101231-1011313113333303-1000320002101301-3323321230330012-1213103133322332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for interface services.

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

<a id="canonical-2102313010000110-2101032121323003-3312121221020122-3202131011123113-2301033120331012-2201332232210132-2011113002212323-1122001201320120"></a>

## Direct properties — interface_services / 311221301112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101122321303331-0003113202121311-2020212231303203-3233211221333000-0300323233110102-1312133031213231-3331220212122002-2003322312021303"></a>

## Next pages — interface_services / 311221301112 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-1213000220331003-2212012300032022-1301023232032011-0131133003103231-3032133030010100-1112021300302131-3202023310301110-0122202102113020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212010023313211-0323330113122002-3221130312223232-2021231110231131-3223331010301101-2201120003332100-0013130131311210-2031201021110103"></a>

## site_acl.outside_network — outside_network / 102030122133 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.outside_network

<a id="canonical-2323312133323311-3313202201231031-2200112321021131-1021332133101230-0313133031130122-2222033122333130-2121332231300321-2301332222010303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-1201020330332200-1130323222002031-2233013200000310-3002003022023221-0033333313310233-1130333200301110-0121223033233303-3022221301133302"></a>

## Direct properties — outside_network / 102030122133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212201022011200-0201213102110013-3323032211132200-2120221213323221-2203122113013111-3310300012320121-2021121201203220-0332233123213112"></a>

## Next pages — outside_network / 102030122133 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)

<a id="canonical-2330310120132320-0030113032103000-2100202300223313-3321131323320221-1031013320030201-1200302013101000-2330012121302112-1132210100002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323330022000131-1203323312000122-2132123212021303-2132331313132110-1010131030031132-1003030203001030-0220232032100012-3232111132120220"></a>

## site_acl.vip_services — vip_services / 322213001033 / 2

Breadcrumbs:

- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
- [Property reference](data-sources--fast_acl--reference--group-001.md#canonical-3100210133132111-2130133321102022-0301000003213330-0331021223122003-3020013020111312-3302130121121220-1230123210322121-1330323032133311)
- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- site_acl.vip_services

<a id="canonical-3233130210000133-2033220130031221-3333222302212021-1012003033031320-3130231012120302-0233100100333001-0012200311320310-0210311030001302"></a>

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

<a id="canonical-1212102321002100-0222200012213330-0310033211331001-2121003123310333-2130113010102333-0113203111223033-3320103321123003-3023113123312102"></a>

## Direct properties — vip_services / 322213001033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220001312132022-0111120103131211-2111213330020311-2203303021101333-0000101011321000-1102132102022020-1320211100132210-2110312010112310"></a>

## Next pages — vip_services / 322213001033 / 4

- [site_acl](data-sources--fast_acl--reference--group-001.md#canonical-1110011201121121-1001033233300213-2112212131220213-2322300103123200-0110120030023100-0321031322101033-2023110230330231-2122113130222320)
- [xcsh_fast_acl](../data-sources/fast_acl.md#canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210)
