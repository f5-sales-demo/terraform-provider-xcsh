---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-2201011201213310-2123230201311121-3000013211032102-1232301102333202-0000323230332311-1020302323121023-1131232313130121-2002200211130333"></a>

## regex_value property — domain / 133202123322 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-1113210213111333-0100103100333231-1110121023100323-1221011213212333-2111132103220133-2010233323130002-0202022020303100-3102311022122233"></a>

<a id="canonical-3312021323233030-1223212231120333-1021201110332000-3033203201321022-3220322233103301-1023202322013133-3213120232012202-3333123100013200"></a>

## suffix_value property — domain / 133202123322 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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

<a id="canonical-2020001120233201-2121131212113312-3200003002331111-2100230300321100-3123122233121213-0032101221310103-2201131303312330-0102231031233012"></a>

## Next pages — domain / 133202123322 / 7

- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100100201210000-2331210332000112-0001233211303011-2200221102301033-3013203202112011-0101101322211332-2031331300302000-1102200303001002"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — flow_label / 213001023100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-1011312213232312-3032223231220203-3103001202211310-0121011222230301-2301213231303313-0103210211222322-3021300031203132-0332010231330211"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("account_management",
    "authentication"),
  validators.ConflictingObjectAttributes("account_management",
    "financial_services"),
  validators.ConflictingObjectAttributes("account_management",
    "flight"),
  validators.ConflictingObjectAttributes("account_management",
    "profile_management"),
  validators.ConflictingObjectAttributes("account_management",
    "search"),
  validators.ConflictingObjectAttributes("account_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("authentication",
    "financial_services"),
  validators.ConflictingObjectAttributes("authentication",
    "flight"),
  validators.ConflictingObjectAttributes("authentication",
    "profile_management"),
  validators.ConflictingObjectAttributes("authentication",
    "search"),
  validators.ConflictingObjectAttributes("authentication",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("financial_services",
    "flight"),
  validators.ConflictingObjectAttributes("financial_services",
    "profile_management"),
  validators.ConflictingObjectAttributes("financial_services",
    "search"),
  validators.ConflictingObjectAttributes("financial_services",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("flight",
    "profile_management"),
  validators.ConflictingObjectAttributes("flight",
    "search"),
  validators.ConflictingObjectAttributes("flight",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("profile_management",
    "search"),
  validators.ConflictingObjectAttributes("profile_management",
    "shopping_gift_cards"),
  validators.ConflictingObjectAttributes("search",
    "shopping_gift_cards")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

Terraform syntax:

```terraform
flow_label {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231210313031302-0320020101130003-0033032111212231-0303221012311322-3322302233000330-1203312220023332-0232213121200203-2023300111300013"></a>

## Direct properties — flow_label / 213001023100 / 3

- [account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020): complete subsection reference.

- [authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002): complete subsection reference.

- [financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332): complete subsection reference.

- [flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230): complete subsection reference.

- [profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032): complete subsection reference.

- [search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013): complete subsection reference.

- [shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102): complete subsection reference.

<a id="canonical-3122231330221101-1123200333031130-1223023101002002-3231120332001030-3320031123302113-2200111021103321-0010111212211201-2221010120211032"></a>

## Next pages — flow_label / 213001023100 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300111300031102-0331221330022130-1130131023222330-3330111211100132-1121023202023002-2120102133303131-0012130101131311-3100011100031013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — account_management / 312200330220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-2002100110000300-0333320310121211-1311122301202220-2101311102220102-2113203231332301-2223121322121310-3200031000331310-0112110223132010"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Account Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "password_reset")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

Terraform syntax:

```terraform
account_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213333321313310-1200312220212301-3032233112101012-3010130203221130-1322301331333112-0211332003130230-2010023131002133-0113031121200002"></a>

## Direct properties — account_management / 312200330220 / 3

- [create](resources--cdn_loadbalancer--reference--group-008.md#canonical-1332230231313033-0102312302132011-2003030131132321-1121112300332222-3223220031331123-1223323110303332-0300222032310332-3223312120112131): complete subsection reference.

- [password_reset](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102320210330032-0123310121132330-1211333003200101-2003330000101001-1233011310322213-1223031111332323-1301321332202312-3130221202230212): complete subsection reference.

<a id="canonical-2201121213321110-3023211303331130-2302121010213002-2212122000111230-1101223100230301-3220322012203300-2311101303333103-0331321033302331"></a>

## Next pages — account_management / 312200330220 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](resources--cdn_loadbalancer--reference--group-008.md#canonical-1332230231313033-0102312302132011-2003030131132321-1121112300332222-3223220031331123-1223323110303332-0300222032310332-3223312120112131)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102320210330032-0123310121132330-1211333003200101-2003330000101001-1233011310322213-1223031111332323-1301321332202312-3130221202230212)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1332230231313033-0102312302132011-2003030131132321-1121112300332222-3223220031331123-1223323110303332-0300222032310332-3223312120112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122222033030012-1202231203211121-0213131013122302-2310022322031100-0102231000230202-2002203233230032-0100031211303001-2201001010032032"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — create / 023312013220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-3123003210021212-1222202112013203-0102312121103312-0302121103301021-0023000021311333-3000302130322131-2023223233113032-1322211103222003"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
create = {}
```

<a id="canonical-0102132311110033-1032213231011203-2133112201230023-2211311311210220-2012232120221220-3302101301013201-2101103213012110-1101103122113320"></a>

## Direct properties — create / 023312013220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123323210010123-3201220212201132-3203130011013312-3330333330120123-1210122221021123-0102200011301232-2330120313020112-2123311300120310"></a>

## Next pages — create / 023312013220 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2102320210330032-0123310121132330-1211333003200101-2003330000101001-1233011310322213-1223031111332323-1301321332202312-3130221202230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112310112020221-1031310121210110-3002301110023333-3121203132011010-1022101312100323-2211030033030011-1010033322211210-2010212132200331"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — password_reset / 232300333201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-3323230032302103-2311120233232030-2123301022132120-3213030313132113-3001122311002012-2203132221212002-3223313011211300-1201121033310222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for password reset.

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

Terraform syntax:

```terraform
password_reset = {}
```

<a id="canonical-0311303212320020-3121201213302321-3132112000001212-0121011222130231-2320233131322223-1113303303132302-2333103330011220-0102033323313300"></a>

## Direct properties — password_reset / 232300333201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112313102120103-0230222113020333-0101103030011030-3011320232130213-3232213123033223-0212023100112222-1211013133022131-0300310013231132"></a>

## Next pages — password_reset / 232300333201 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-1301130230221023-0033031011202011-1201121010220133-1111211320102213-1011300322200313-1313011010133122-1001222102120030-2331221311233020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203232310021331-3013122102300021-1221130232100110-0031110031013131-0232223200312113-0102022030012220-1230233210113103-2120221231323003"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — authentication / 102020111230 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-1020033010322112-2302112101002333-1011123320001330-0311200111000221-3201000133022101-1003033203203320-0211310323323113-1102030200231221"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Authentication Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("login",
    "login_mfa"),
  validators.ConflictingObjectAttributes("login",
    "login_partner"),
  validators.ConflictingObjectAttributes("login",
    "logout"),
  validators.ConflictingObjectAttributes("login",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_mfa",
    "login_partner"),
  validators.ConflictingObjectAttributes("login_mfa",
    "logout"),
  validators.ConflictingObjectAttributes("login_mfa",
    "token_refresh"),
  validators.ConflictingObjectAttributes("login_partner",
    "logout"),
  validators.ConflictingObjectAttributes("login_partner",
    "token_refresh"),
  validators.ConflictingObjectAttributes("logout",
    "token_refresh")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

Terraform syntax:

```terraform
authentication {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321211010103323-3131112230012332-2121023303231232-3033233010331002-2012231321112102-1232020311222003-3103120203111320-0132022111323231"></a>

## Direct properties — authentication / 102020111230 / 3

- [login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012): complete subsection reference.

- [login_mfa](resources--cdn_loadbalancer--reference--group-008.md#canonical-0032323121300012-1122323330311013-3213111031231312-0300210213201032-0133201211210123-0102013130222231-1311213023012122-3101231001132133): complete subsection reference.

- [login_partner](resources--cdn_loadbalancer--reference--group-008.md#canonical-3123033300213313-1220322323020321-0230030023032231-3003332231333212-3033013013011100-0301100001200122-1123130332202002-3010010210232021): complete subsection reference.

- [logout](resources--cdn_loadbalancer--reference--group-008.md#canonical-1111203002230011-0011321031000231-1012130131030312-2210213101320113-0301211130000233-1332230113333313-2022002032132310-3023132020113233): complete subsection reference.

- [token_refresh](resources--cdn_loadbalancer--reference--group-008.md#canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112): complete subsection reference.

<a id="canonical-2320032313033133-3013012100231331-2002333301033113-1201311310123023-3031120110303211-2110331221113311-2332302130133221-2122001312122201"></a>

## Next pages — authentication / 102020111230 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](resources--cdn_loadbalancer--reference--group-008.md#canonical-0032323121300012-1122323330311013-3213111031231312-0300210213201032-0133201211210123-0102013130222231-1311213023012122-3101231001132133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](resources--cdn_loadbalancer--reference--group-008.md#canonical-3123033300213313-1220322323020321-0230030023032231-3003332231333212-3033013013011100-0301100001200122-1123130332202002-3010010210232021)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](resources--cdn_loadbalancer--reference--group-008.md#canonical-1111203002230011-0011321031000231-1012130131030312-2210213101320113-0301211130000233-1332230113333313-2022002032132310-3023132020113233)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](resources--cdn_loadbalancer--reference--group-008.md#canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011021330222201-2123013002210202-1211303310011333-1100122002302230-0001201030121120-1132010133330120-2333331112223201-2113002320212013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — login / 131131120321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-1221223320103133-0300303230133101-0021221133331021-1213320233233311-3020211101010002-0332301131113020-0233111300001321-0030231202201221"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_transaction_result",
    "transaction_result")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

Terraform syntax:

```terraform
login {
  # Configure direct properties listed below.
}
```

<a id="canonical-1222311002323121-1113112213221112-3010013323311313-3121202002013110-1131203321121311-1003010201311212-1222001102013201-0230223100212011"></a>

## Direct properties — login / 131131120321 / 3

- [disable_transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332003023320000-2122131312310000-2322001111211232-1000121330300013-3123013021031330-3001212011201012-0332320110013332-2323123312101200): complete subsection reference.

- [transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222): complete subsection reference.

<a id="canonical-0321301233201031-2311230130221033-1223202333321100-3212020032223200-1110110223123323-2011203200223333-2211322133301111-2000002123231321"></a>

## Next pages — login / 131131120321 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332003023320000-2122131312310000-2322001111211232-1000121330300013-3123013021031330-3001212011201012-0332320110013332-2323123312101200)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0332003023320000-2122131312310000-2322001111211232-1000121330300013-3123013021031330-3001212011201012-0332320110013332-2323123312101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002003002113323-2113013023200310-3130202011200303-0210223212101132-2223203103232100-2020121121311233-2133111310313323-0121021120111122"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — disable_transaction_result / 101222210333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-3231231120023111-1311332203130311-2013311030333223-3023020330231101-0031200110021220-3031000202010222-2110020133130332-2331033013330100"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_transaction_result = {}
```

<a id="canonical-0122222123330202-2112112231030120-1110011000103213-2110322223131031-2230130213322030-1020012100211020-0232330313230212-1202313210301000"></a>

## Direct properties — disable_transaction_result / 101222210333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100331223012222-2333232330200021-0222020300313232-3302300202200120-2112112223003213-0012132201332022-3121213321230010-1303300223121011"></a>

## Next pages — disable_transaction_result / 101222210333 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131001002300303-0100230032021113-2000033132302330-3011023002123212-0212330021110232-1012231013000230-3032220322002313-0213202132321132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result — transaction_result / 210213230312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-0030011013121023-2331323033110033-3230101202021003-3301212132121331-3232032302022332-3113020113203113-0002222113112002-1112222001121223"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
transaction_result {
  # Configure direct properties listed below.
}
```

<a id="canonical-3323111131302000-3011212202011121-2131133333132332-0331130000011221-2002110223031200-2011113123121011-1200002002310022-1022213222330123"></a>

## Direct properties — transaction_result / 210213230312 / 3

- [failure_conditions](resources--cdn_loadbalancer--reference--group-008.md#canonical-0310201310333111-3101220000322112-1322133033020100-1200202121200110-1223001131233111-1033333303301331-3333220322130001-0310301010021211): complete subsection reference.

- [success_conditions](resources--cdn_loadbalancer--reference--group-008.md#canonical-0310301023302210-0212002010011003-3011221213211023-2321220120023121-0311320333123230-1232003213303203-0030333130300303-2211012331333222): complete subsection reference.

<a id="canonical-1321033331021311-3123301303122230-3301313031132202-0221332300333102-3011121132100022-0010313130021201-0322213221121022-0000210001300033"></a>

## Next pages — transaction_result / 210213230312 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--cdn_loadbalancer--reference--group-008.md#canonical-0310201310333111-3101220000322112-1322133033020100-1200202121200110-1223001131233111-1033333303301331-3333220322130001-0310301010021211)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--cdn_loadbalancer--reference--group-008.md#canonical-0310301023302210-0212002010011003-3011221213211023-2321220120023121-0311320333123230-1232003213303203-0030333130300303-2211012331333222)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0310201310333111-3101220000322112-1322133033020100-1200202121200110-1223001131233111-1033333303301331-3333220322130001-0310301010021211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031212102133312-3013210331330113-2021203233233000-2221300101230112-0012331302210003-1023221031001103-3120220120231313-0221320310202002"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — failure_conditions / 120010202122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-1222213221333321-0302000210223332-0312233013323101-2302230021020110-0123130120231221-3221211232101122-3202332321031212-1113112222211013"></a>

Type: `"object"`. list nested block, Optional.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
failure_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022103121313321-2101033220320013-1012003202001230-0133010231231120-1000220033303201-0001112110123203-0210001122120010-0310022303111013"></a>

## Direct properties — failure_conditions / 120010202122 / 3

<a id="canonical-3322012200000212-2220202232221102-1310332300223223-1030013230112011-2033032000000102-3310000023003202-0302210332332201-2003301103023213"></a>

<a id="canonical-1233311231123201-3130301020023001-3002221000033003-1321003103200131-1310213331030033-1111312223301032-2121332211000103-2032011033213313"></a>

## name property — failure_conditions / 120010202122 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3300310212301113-2123222230002001-2013310220220232-2200103302333223-0333222320032122-0001220100333021-1211211021203212-1313220232211010"></a>

<a id="canonical-2333213100112311-0233331302330320-0331122022110202-0121203013212111-1212133213230010-3230000233120131-0131231023121013-3132031022233201"></a>

## regex_values property — failure_conditions / 120010202122 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200100232100011-0220113010312213-0032102301203223-2201020210221000-2133030211123001-1022332321120321-0122210023333100-3321121303133031"></a>

<a id="canonical-0212332230020323-0220110201113313-1013301321123311-1202103321200200-3333010133232302-1133030222003020-1033301111303120-3120113132132013"></a>

## status property — failure_conditions / 120010202122 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2301201023213022-2100122112003213-2212210000130131-1132021123030333-2013332233102131-3001111222300002-3231002112130032-1001200033302101"></a>

## Next pages — failure_conditions / 120010202122 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0310301023302210-0212002010011003-3011221213211023-2321220120023121-0311320333123230-1232003213303203-0030333130300303-2211012331333222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310001311011220-0312230213010222-0230133323213101-1012113322301310-0013032111200122-0011023121233320-0003111100320312-1020102010301232"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions — success_conditions / 200331000032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102212102023000-0020132013033223-1131220131013003-0002130001222122-1301313212110333-2013100320322231-1211312021222112-2323121333122012)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-0233033321122202-2021113100200120-1202203333300233-3232323032331102-3321211130100201-1321321013123202-1321022333101231-0233220022123300"></a>

Type: `"object"`. list nested block, Optional.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
success_conditions {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321021103030311-2020320223301103-3332223322201103-2132031133012121-1210211301230101-1203200103012123-0103131031130011-3333101013320123"></a>

## Direct properties — success_conditions / 200331000032 / 3

<a id="canonical-3013211113112313-1011133023220230-0300331002122231-1221130312111102-2002003303323032-3200313100003233-2200033203321131-3222202232101030"></a>

<a id="canonical-3100221021331233-2230032233121033-3230321321211211-0002313030120100-1003121330031313-0212312332131212-1021210332013132-3220221220033210"></a>

## name property — success_conditions / 200331000032 / 4

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2132212323211022-1302300321220221-0323130300333112-3111100301320320-3222231100121332-0110120133231130-0223232320313210-1033001221133111"></a>

<a id="canonical-2120230332211233-3303102031321332-0030323011022110-2222033203322232-0103000202313301-1211211320211102-0011013222110020-3320112322222102"></a>

## regex_values property — success_conditions / 200331000032 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310332231223333-1312122031330312-0202302103303212-1030213032110313-1300103333211333-2131210032113001-1112301201300012-0322022133013122"></a>

<a id="canonical-1301111320320201-3012312201232111-3230003113103333-3030012100222201-2330231322012301-0222211220200232-0212332221101323-1220202200002111"></a>

## status property — success_conditions / 200331000032 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3123122303122332-0132031300310313-2212111031321101-2020031033113331-2220331212121233-2101022112201121-3020010022002200-0312332223231300"></a>

## Next pages — success_conditions / 200331000032 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--cdn_loadbalancer--reference--group-008.md#canonical-2022133231331001-1332123030223032-1003322321000121-0103101300111001-1132210031232123-2223311332001221-3222101132230023-0002132232220222)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0032323121300012-1122323330311013-3213111031231312-0300210213201032-0133201211210123-0102013130222231-1311213023012122-3101231001132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320222313131002-3223001020121320-1232122200120223-0031110031103032-1030222201131031-1100302003220310-3022023313110312-0202013331330221"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa — login_mfa / 131032020121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-1211202323231233-0221301312110323-0001012320320103-0222102121312103-0120310303003220-1213213032331132-3323013233010230-3022001113012000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
login_mfa = {}
```

<a id="canonical-2010203223110020-2031122102303130-0002333121002223-0112002133011113-3211131232222133-0101123133112202-3303101331233323-2011331221331130"></a>

## Direct properties — login_mfa / 131032020121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131211232120313-2003231022033223-2320132331030003-3313003130203032-1110130323122203-2103133012222000-0010100022202312-0023323100102133"></a>

## Next pages — login_mfa / 131032020121 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3123033300213313-1220322323020321-0230030023032231-3003332231333212-3033013013011100-0301100001200122-1123130332202002-3010010210232021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103031201222030-3312322331301220-2012102133130211-1101311101032120-1200113111101131-2321202100102213-2122330220232313-2321001223300203"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner — login_partner / 001200203020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-3130232100323301-3201303210311230-1230120013203102-0231113320232011-1111131303133022-2013301310200333-1333120011232202-0112022030103233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for login partner.

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

Terraform syntax:

```terraform
login_partner = {}
```

<a id="canonical-2010323330330210-2120223332130032-0000223112201302-1231102300301121-3023323103320230-0011312330001303-3321203132323011-0210303230201121"></a>

## Direct properties — login_partner / 001200203020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302202333110132-3221121301132002-1303031302222103-0310133322200122-2213313310030010-0301113102111302-3002330123331123-2012210212312322"></a>

## Next pages — login_partner / 001200203020 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1111203002230011-0011321031000231-1012130131030312-2210213101320113-0301211130000233-1332230113333313-2022002032132310-3023132020113233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302012012211303-3301321213112213-3132002112112202-2303312213001130-0322231302222103-3323131212002100-0232133221112121-2132012033022022"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout — logout / 231202001212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-0331200230220202-2302122223031221-2103202101213203-1033213000221231-1322210012202220-3022113013031011-2332312030202332-0123021130222021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
logout = {}
```

<a id="canonical-3223133110012132-3232131103111122-0313122100331312-2223322020002100-1302321233101203-3000220000001320-1130331032332203-2232312211020200"></a>

## Direct properties — logout / 231202001212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212221122222103-0002020302100030-3100001211123212-1300302113221300-1033003310003021-1223203322130133-2000032211230221-1320211111021011"></a>

## Next pages — logout / 231202001212 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0133212203133310-1201032011111112-3320330112303200-1103301310102232-0003031301102113-0233132230001030-0131230002001102-3002331120301112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211330331010212-0202111320313222-3220302132032302-3220213023210301-1312231031022131-2331233232323312-3303033123023231-1122123010203313"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh — token_refresh / 002132211011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-0320233020133300-1122220301130032-1021320030033031-0203002013103013-1023322020231012-0211211131132123-3110312131013003-3030212031120110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for token refresh.

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

Terraform syntax:

```terraform
token_refresh = {}
```

<a id="canonical-1003010131120303-1321322321002003-1312033102122223-2302310321111001-3102233112333302-3022223210213322-1132033110030101-1221201311210020"></a>

## Direct properties — token_refresh / 002132211011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131110131302110-0111020102231113-1021001220311110-1021010101233303-1332130223002132-1322000133113111-2001110213311210-1330010100310210"></a>

## Next pages — token_refresh / 002132211011 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--cdn_loadbalancer--reference--group-008.md#canonical-3203312202013322-0311223333121013-3311102120333110-0000023203100332-0322320111122100-0000213300100221-2120023020023311-2020033310012002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221021202231012-3103211302222011-1001002331222301-0001302113213013-1130031221022121-2303211032012323-0212313132110312-3320101123231021"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services — financial_services / 030130233330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-3213130002330221-0313333132331133-3312303323110112-2121111332212102-2112102131231231-0133212310330013-0322113023302321-0332202302311021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Financial Services Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("apply",
    "money_transfer")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"apply\",\"money_transfer\"]"
}
```

Terraform syntax:

```terraform
financial_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220112232121131-3323231122312013-1022012123213112-3211330101223330-1323333223301232-1202103131112322-0021033003132001-2223211232010030"></a>

## Direct properties — financial_services / 030130233330 / 3

- [apply](resources--cdn_loadbalancer--reference--group-008.md#canonical-3201110323313302-0330112021210310-0022331313100311-0301223313001032-0103232301222313-2331311332030121-1323210130303203-1001012023123012): complete subsection reference.

- [money_transfer](resources--cdn_loadbalancer--reference--group-008.md#canonical-3112132023123100-0321333011000312-3321121030203121-0312303332013131-2020130103321003-2032313321112213-1012212001023201-3013032010213030): complete subsection reference.

<a id="canonical-1232111112311130-2221023002113313-3022332031313002-0311212303330131-3220101132203101-1320213300110023-2013231111323010-3312310020012320"></a>

## Next pages — financial_services / 030130233330 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply](resources--cdn_loadbalancer--reference--group-008.md#canonical-3201110323313302-0330112021210310-0022331313100311-0301223313001032-0103232301222313-2331311332030121-1323210130303203-1001012023123012)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer](resources--cdn_loadbalancer--reference--group-008.md#canonical-3112132023123100-0321333011000312-3321121030203121-0312303332013131-2020130103321003-2032313321112213-1012212001023201-3013032010213030)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3201110323313302-0330112021210310-0022331313100311-0301223313001032-0103232301222313-2331311332030121-1323210130303203-1001012023123012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101103021023122-1112211303032200-3022221023220120-3220222212220020-0002112302310213-1010131200310333-3332320221103312-1231132101200210"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply — apply / 302211213212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-1111132212323020-3313200133331132-0312012312012123-3201122302210123-2021312301121210-0102233121111123-3012123020110202-0010212133001011"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
apply = {}
```

<a id="canonical-0003033112211200-3131030012333012-0111311222332012-2212230231301202-1132310321101033-1102232312010331-1220332131302011-0003130330112201"></a>

## Direct properties — apply / 302211213212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333331332012022-2032023022102303-2230303233233121-1200313210223231-2003022031012120-2002301010302230-1122132323101210-2332300003032202"></a>

## Next pages — apply / 302211213212 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3112132023123100-0321333011000312-3321121030203121-0312303332013131-2020130103321003-2032313321112213-1012212001023201-3013032010213030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231211321001310-0211200133012122-3023020022330122-0100333110210120-2021102301103020-2332112110110012-1101021023102222-3102330310101130"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer — money_transfer / 103012133222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0313213020223323-2133202013012220-0201311200000033-3121212010002312-3021102313313011-1101001113312300-1110232003221301-0000311120212220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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

Terraform syntax:

```terraform
money_transfer = {}
```

<a id="canonical-2221133212103132-2220201132310001-3102302332011113-2012023313203222-2221133122100231-2030120011232102-0230312120022312-1301233232003013"></a>

## Direct properties — money_transfer / 103012133222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303213300313103-2021100103003301-2011312323332322-3123332120113013-3330311321320233-2002301323033110-0011002132321322-1011021301303102"></a>

## Next pages — money_transfer / 103012133222 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321213233222323-2320001331202231-1020211320330323-3221203233311102-2220301031330220-2110321111000220-2013121022213021-2032120302302332)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333100120121220-3223033131301232-1330031211022303-3323000330033133-3102121133201100-1200033201223023-2123021332120130-0212012133033103"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight — flight / 312101310131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-2030223102300103-2011312222001213-2123321100131211-0123022033123320-0232013223221202-3310321202212113-0203123300212231-2032011332000021"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

Upstream description:

Bot Defense Flow Label Flight Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"checkin\"]"
}
```

Terraform syntax:

```terraform
flight {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313221321010021-3031301023133112-1033213112131230-2202001123030303-3033332321232022-1202213002210322-2202120022320032-2311123111002033"></a>

## Direct properties — flight / 312101310131 / 3

- [checkin](resources--cdn_loadbalancer--reference--group-008.md#canonical-0001323123302232-2023131101313303-0000202332322123-0201322031333121-3322122102203111-2020021010111313-0110013000200132-1123300203203232): complete subsection reference.

<a id="canonical-3122231301332211-2001000033331000-3030221021220130-0131130220200022-2322100033223323-1113001021232223-2130202023103000-3313302010322030"></a>

## Next pages — flight / 312101310131 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin](resources--cdn_loadbalancer--reference--group-008.md#canonical-0001323123302232-2023131101313303-0000202332322123-0201322031333121-3322122102203111-2020021010111313-0110013000200132-1123300203203232)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0001323123302232-2023131101313303-0000202332322123-0201322031333121-3322122102203111-2020021010111313-0110013000200132-1123300203203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023213200121212-1001022101231303-0222320300303022-3222001232301010-0203221101310302-0033322010302221-1322220211003131-3322023023201031"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin — checkin / 102112021233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-1313321031200131-1321111222230111-1223203333321332-1302332223320030-1221232030031220-2213000321012302-3131210022000313-0030333311131332"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
checkin {}
```

<a id="canonical-1220101100323011-3020201003112231-2220031022032000-3330022321211210-2003313212030230-2321233132000231-0210002000333310-2330100120332113"></a>

## Direct properties — checkin / 102112021233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312332222030232-0003023332112212-1210121022032330-2233203122312100-0221312030103322-0221132320102023-2000103111122010-3313100100120030"></a>

## Next pages — checkin / 102112021233 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--cdn_loadbalancer--reference--group-008.md#canonical-0332030030201123-3112022110233212-2303113312031010-2131212031312021-1221203201003101-0002221213221131-1011101101203320-0321231311301230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201333103312301-3303120300301101-0021320303130112-0200111222001230-0121130330200202-0221221130110013-2102131233002331-3211213311311301"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management — profile_management / 003123101102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-3213021230003323-3303210330033302-0030221121310001-2312313010320333-3301111022202120-0012112201302023-0223332301213322-0102301303133111"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("create",
    "update"),
  validators.ConflictingObjectAttributes("create",
    "view"),
  validators.ConflictingObjectAttributes("update",
    "view")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"update\",\"view\"]"
}
```

Terraform syntax:

```terraform
profile_management {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212022311032131-2332232003010101-2300333023132212-1123023313332012-1121001000131012-3222031211333312-0031030021233302-0320211200103312"></a>

## Direct properties — profile_management / 003123101102 / 3

- [create](resources--cdn_loadbalancer--reference--group-008.md#canonical-3330202013011331-2020012321322023-1321223331121013-3300023101203020-3030233303131102-2133001001110230-3311231311200013-1203203022133111): complete subsection reference.

- [update](resources--cdn_loadbalancer--reference--group-008.md#canonical-0033223221211332-2310331312211003-1323221010220133-0103002201031121-1323201113321331-1232331201203220-2230132210031120-3203221022010210): complete subsection reference.

- [view](resources--cdn_loadbalancer--reference--group-008.md#canonical-0321130032220302-2322003033030333-2332311132320213-0310332231001123-2321200010302230-1032011103233021-2333102302033221-1330031000110310): complete subsection reference.

<a id="canonical-3100302001220111-2001300223002330-1212111320220013-0223312310311030-0320313321331123-3221203132302233-2132211003111012-0023032230302333"></a>

## Next pages — profile_management / 003123101102 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create](resources--cdn_loadbalancer--reference--group-008.md#canonical-3330202013011331-2020012321322023-1321223331121013-3300023101203020-3030233303131102-2133001001110230-3311231311200013-1203203022133111)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update](resources--cdn_loadbalancer--reference--group-008.md#canonical-0033223221211332-2310331312211003-1323221010220133-0103002201031121-1323201113321331-1232331201203220-2230132210031120-3203221022010210)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view](resources--cdn_loadbalancer--reference--group-008.md#canonical-0321130032220302-2322003033030333-2332311132320213-0310332231001123-2321200010302230-1032011103233021-2333102302033221-1330031000110310)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3330202013011331-2020012321322023-1321223331121013-3300023101203020-3030233303131102-2133001001110230-3311231311200013-1203203022133111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110120230123101-2012122211000010-3213310132232112-1232111013100203-0112111111002312-1022033033123321-0000233121222030-2313212322000003"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create — create / 100320112102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-1131320302002201-0133231020013032-1300311020201302-1022213302021313-2012202210311331-2201223223222102-3100121213121132-2221013312001000"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
create = {}
```

<a id="canonical-0101320012102021-3121231331131330-2012022331213213-0310220330320212-2302101100003230-1133000100021320-1110121323212210-1311010311330312"></a>

## Direct properties — create / 100320112102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102111121232013-1233330210210312-0031020232031230-0202311123222113-3122113210220131-1132201203230021-1003003303203133-2102310111002001"></a>

## Next pages — create / 100320112102 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0033223221211332-2310331312211003-1323221010220133-0103002201031121-1323201113321331-1232331201203220-2230132210031120-3203221022010210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311332132030011-1333201111210130-0223111333101302-1000313122013300-2122203021113230-2301331120330222-3212222301132310-2000202222302131"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update — update / 020002113210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-0212313111122013-0322310022000100-3320102201322120-2222031013213011-2010220111010320-3321211230233113-2330313013112332-2222313003320330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
update = {}
```

<a id="canonical-0210113213301031-1030330101013111-2212221302013221-2330220321221113-2032033101010011-0201012000002022-1333303231110020-1001333100131031"></a>

## Direct properties — update / 020002113210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332211033111331-3111001002133101-3001003303202301-1103000301222332-2203023313333130-0302221131001303-0110031332112020-2320102230023033"></a>

## Next pages — update / 020002113210 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0321130032220302-2322003033030333-2332311132320213-0310332231001123-2321200010302230-1032011103233021-2333102302033221-1330031000110310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101012012000031-2010012022322300-0301333221211013-0312103233202111-3022331302031331-3211312023123212-3021021113231031-1221333333201300"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view — view / 020312223111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-2331332223331101-2331203203210311-0110131031020002-0201013331010310-3022101233212330-2303103032231231-0202030230223013-3031101012202130"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
view = {}
```

<a id="canonical-3120023322312031-0023021111323011-0302031223012221-2310032203302222-0022222331130003-2233120131130030-1223202112323211-2231112211303110"></a>

## Direct properties — view / 020312223111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221120120012120-2133120321230020-2020110222021210-1011122020013021-2002211223011210-3213101021221200-0013310003200331-0323221003313223"></a>

## Next pages — view / 020312223111 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313313310310323-1302211322101301-1122323222111032-1000122130202223-2302330100332033-3122323230003202-1220313212003031-3313001202000210"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search — search / 311213213002 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-1011131100202302-1023003320223223-0130032212110020-3202310212123201-0021032313301213-0310130132320202-1201220302312200-3020210023221200"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("flight_search",
    "product_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "room_search"),
  validators.ConflictingObjectAttributes("product_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("product_search",
    "room_search"),
  validators.ConflictingObjectAttributes("reservation_search",
    "room_search")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012011021321203-0311200123131113-2001130333313020-3332103123313232-2002202030201122-2021303232300211-3121122000303031-1321211033220303"></a>

## Direct properties — search / 311213213002 / 3

- [flight_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-3312300230020231-1230221111102010-3131200230302102-3303302012313030-0212131321310212-3110013300130010-1121121132332333-2133111022221030): complete subsection reference.

- [product_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1331200222232021-1010121202101102-2122330331020202-2232132220113030-3112202223021110-3212131112331320-0130001330020203-2230112130000023): complete subsection reference.

- [reservation_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1300333230313320-3332010101002012-2010230332313111-1003101211113021-2302011300130010-3112002221103332-1202222102023031-1030121110301231): complete subsection reference.

- [room_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-0200000033211031-2110022331222333-0300321133131001-3132220222022020-0001030121103122-0120300102002011-3031322010333130-1320330301102010): complete subsection reference.

<a id="canonical-0301323230100122-2211213130333021-1202200300131201-3010313113013330-2211020201333131-2301012112023101-0131222203131211-3012213101323100"></a>

## Next pages — search / 311213213002 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-3312300230020231-1230221111102010-3131200230302102-3303302012313030-0212131321310212-3110013300130010-1121121132332333-2133111022221030)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1331200222232021-1010121202101102-2122330331020202-2232132220113030-3112202223021110-3212131112331320-0130001330020203-2230112130000023)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1300333230313320-3332010101002012-2010230332313111-1003101211113021-2302011300130010-3112002221103332-1202222102023031-1030121110301231)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-0200000033211031-2110022331222333-0300321133131001-3132220222022020-0001030121103122-0120300102002011-3031322010333130-1320330301102010)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3312300230020231-1230221111102010-3131200230302102-3303302012313030-0212131321310212-3110013300130010-1121121132332333-2133111022221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010131102323030-3130001033213022-1313032333221113-0120323001010232-2100331233103113-3231131023302210-3131230001303123-3200332111011213"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search — flight_search / 022203333131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-2131130313033223-0300323110312122-1121320112133301-3311020122223230-1113013002222200-0001220110233301-1113310323003221-2111330131312202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for flight search.

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

Terraform syntax:

```terraform
flight_search = {}
```

<a id="canonical-0023012033300201-2133121311003200-1330011331210032-2132211133030331-1133311100230310-2333212133033012-0131121221333212-0300300323330312"></a>

## Direct properties — flight_search / 022203333131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122130220222122-1021121303112122-1023132021231300-3012202002000003-0012030111031312-3111231231000223-2203101102032103-2312102021121100"></a>

## Next pages — flight_search / 022203333131 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1331200222232021-1010121202101102-2122330331020202-2232132220113030-3112202223021110-3212131112331320-0130001330020203-2230112130000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021231302330220-1222332023011113-0330123120332001-2211130112010001-3331122032121302-2330220122221033-0231201103022011-0103322220010023"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.product_search — product_search / 132323021310 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-1103030111323210-3110233203020023-3322100230322323-2212333200100230-1200010312031300-2302302021100211-2111120133131211-1323030001030322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for product search.

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

Terraform syntax:

```terraform
product_search = {}
```

<a id="canonical-2033102011032221-3103121013322112-2301123230110111-1130023122021203-3032123301020212-2121213223321032-3002123110020320-1000311221133103"></a>

## Direct properties — product_search / 132323021310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111331033112300-3122220102211031-0032230121300012-3231333021021111-1301310221111001-0030222102201312-3321210012001332-3230031203031002"></a>

## Next pages — product_search / 132323021310 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1300333230313320-3332010101002012-2010230332313111-1003101211113021-2302011300130010-3112002221103332-1202222102023031-1030121110301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132330113011220-2233001211100103-1330213000221302-0310303113223003-1310221313021311-3302322200133020-3032110122202022-3132322031321013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search — reservation_search / 021231321030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-1002133332300013-0011031122322220-1213111031203200-1323011220121001-0122310212230102-3123021331002132-2203000331332301-3031102201203311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reservation search.

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

Terraform syntax:

```terraform
reservation_search = {}
```

<a id="canonical-0310123211230120-0230303310210321-1121012001012020-1001233233203113-3322133131132333-2001131101220100-2102321321033002-2103122133211023"></a>

## Direct properties — reservation_search / 021231321030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300022320300302-2310000033230100-2311011302002302-3011221210200122-0330013322012012-3101210321232031-1121201130112322-0030300132122112"></a>

## Next pages — reservation_search / 021231321030 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0200000033211031-2110022331222333-0300321133131001-3132220222022020-0001030121103122-0120300102002011-3031322010333130-1320330301102010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112012211132020-2003112003023012-0223330112232100-3010101221220311-0011221110002210-2303311101303231-2131133002032331-2010102330302030"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.room_search — room_search / 202133120130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-1131202200031301-2200323232312322-0132002013111023-2030132230321230-1323111202022320-1313123322132333-1221103020213330-1313323012333320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for room search.

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

Terraform syntax:

```terraform
room_search = {}
```

<a id="canonical-3332333112303301-0210001130202302-3331120232033032-2200101221033303-3133020131122202-1213333203112220-3301222303012213-0333101302001010"></a>

## Direct properties — room_search / 202133120130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203323033102333-2112203300031030-3002303031021213-1030110222322333-0220033121332333-3113303330323020-2223020103032002-1211312230011322"></a>

## Next pages — room_search / 202133120130 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132212330301122-1123130333300123-3013230213133133-0323220320323032-0100100120010320-3103221002312023-0203112332033220-0102310031330023"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards — shopping_gift_cards / 011213301320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-1200320023030131-1012230032333313-1232311022030220-3233033003230131-2230200010223200-2032311120003331-1133022220122210-1132323331330002"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "gift_card_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_purchase_gift_card",
    "shop_update_quantity")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122003021120322-3331121123222130-2123230201312011-0013210111222121-2323133301211002-0321102110002000-2120023032012023-0223331131301100"></a>

## Direct properties — shopping_gift_cards / 011213301320 / 3

- [gift_card_make_purchase_with_gift_card](resources--cdn_loadbalancer--reference--group-008.md#canonical-3221302321130012-1021323010330332-0232233321122331-1031002022233001-0132311210322320-2131330300122220-0311231320233312-2012230202321121): complete subsection reference.

- [gift_card_validation](resources--cdn_loadbalancer--reference--group-008.md#canonical-0230222010002321-3122033232111200-0232200333223211-0221113020201100-2300300010133322-1222332012230211-2230233210220222-3212120303113313): complete subsection reference.

- [shop_add_to_cart](resources--cdn_loadbalancer--reference--group-008.md#canonical-0102332233313221-3230201311013231-3332113122213000-0033233131231223-1213202012010322-2300110212332222-3200112212102230-2301201313103221): complete subsection reference.

- [shop_checkout](resources--cdn_loadbalancer--reference--group-008.md#canonical-0312012131201110-1220012100103001-0230233212300222-2231131111020300-1000322200210303-1011301113333132-1332011032100132-1233000123033110): complete subsection reference.

- [shop_choose_seat](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000132111103122-0012000020321120-0022121331002201-0121102003030310-2020223213122023-1321230302000231-1301230120200301-1021210030213100): complete subsection reference.

- [shop_enter_drawing_submission](resources--cdn_loadbalancer--reference--group-008.md#canonical-0103220311012301-3332112213030231-1213330311123213-0133100320032320-2220002000310113-3333123130202223-1130321200202132-0213020210013121): complete subsection reference.

- [shop_make_payment](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000122332103313-2023021003231010-3100122011033233-2200013320233303-2202211000212000-1012013102112223-3020110210001323-3123211311031223): complete subsection reference.

- [shop_order](resources--cdn_loadbalancer--reference--group-008.md#canonical-1112310200102333-2132020132311100-0133310332211100-3122130123002133-1012120131302020-0303030100312111-3212112130111002-1121223323111003): complete subsection reference.

- [shop_price_inquiry](resources--cdn_loadbalancer--reference--group-008.md#canonical-1330011013210322-3321011001101001-1023103132023230-1201110123211103-3232323202322332-2231212200030332-3233320030121220-1321121233312332): complete subsection reference.

- [shop_promo_code_validation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3331130003130133-0113302011022001-3111013010132020-3133323213022300-3002131021232133-3321303301112121-2313123023033000-2023122131203212): complete subsection reference.

- [shop_purchase_gift_card](resources--cdn_loadbalancer--reference--group-008.md#canonical-1311312001112121-2113030033113022-0022203223122233-2131331110033331-3000033210003031-2103302000321130-3313130202300301-2231133103211311): complete subsection reference.

- [shop_update_quantity](resources--cdn_loadbalancer--reference--group-008.md#canonical-1010222323231020-0103010320210313-3221223233333120-0000122131303311-3302311133013100-1132202331333233-1321222300332001-3120032022320030): complete subsection reference.

<a id="canonical-2011203221320320-3111101001230010-0330012022301232-0310200200030030-3113320020313010-0321321201012223-1330322032120130-0300030003220033"></a>

## Next pages — shopping_gift_cards / 011213301320 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--cdn_loadbalancer--reference--group-008.md#canonical-3221302321130012-1021323010330332-0232233321122331-1031002022233001-0132311210322320-2131330300122220-0311231320233312-2012230202321121)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--cdn_loadbalancer--reference--group-008.md#canonical-0230222010002321-3122033232111200-0232200333223211-0221113020201100-2300300010133322-1222332012230211-2230233210220222-3212120303113313)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--cdn_loadbalancer--reference--group-008.md#canonical-0102332233313221-3230201311013231-3332113122213000-0033233131231223-1213202012010322-2300110212332222-3200112212102230-2301201313103221)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--cdn_loadbalancer--reference--group-008.md#canonical-0312012131201110-1220012100103001-0230233212300222-2231131111020300-1000322200210303-1011301113333132-1332011032100132-1233000123033110)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000132111103122-0012000020321120-0022121331002201-0121102003030310-2020223213122023-1321230302000231-1301230120200301-1021210030213100)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--cdn_loadbalancer--reference--group-008.md#canonical-0103220311012301-3332112213030231-1213330311123213-0133100320032320-2220002000310113-3333123130202223-1130321200202132-0213020210013121)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000122332103313-2023021003231010-3100122011033233-2200013320233303-2202211000212000-1012013102112223-3020110210001323-3123211311031223)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order](resources--cdn_loadbalancer--reference--group-008.md#canonical-1112310200102333-2132020132311100-0133310332211100-3122130123002133-1012120131302020-0303030100312111-3212112130111002-1121223323111003)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--cdn_loadbalancer--reference--group-008.md#canonical-1330011013210322-3321011001101001-1023103132023230-1201110123211103-3232323202322332-2231212200030332-3233320030121220-1321121233312332)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3331130003130133-0113302011022001-3111013010132020-3133323213022300-3002131021232133-3321303301112121-2313123023033000-2023122131203212)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--cdn_loadbalancer--reference--group-008.md#canonical-1311312001112121-2113030033113022-0022203223122233-2131331110033331-3000033210003031-2103302000321130-3313130202300301-2231133103211311)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--cdn_loadbalancer--reference--group-008.md#canonical-1010222323231020-0103010320210313-3221223233333120-0000122131303311-3302311133013100-1132202331333233-1321222300332001-3120032022320030)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3221302321130012-1021323010330332-0232233321122331-1031002022233001-0132311210322320-2131330300122220-0311231320233312-2012230202321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011331210030022-0333311302110320-3102033312230110-1113302313023220-3013000332003020-2022130310130000-3222132330031322-0222020112300030"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — gift_card_make_purchase_with_gift_card / 130322031302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0300131133132312-0032110312020221-0200111320233230-1110013102212200-3203323112200110-3302231002030112-1012003323013130-1020210032110321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card make purchase with gift card.

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

Terraform syntax:

```terraform
gift_card_make_purchase_with_gift_card = {}
```

<a id="canonical-0302323231120203-1100020001030021-3131322130010221-2032000101232010-0210233120203222-0102321323132013-3103110231212330-2222012012301213"></a>

## Direct properties — gift_card_make_purchase_with_gift_card / 130322031302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013310003001213-2020322133101323-2022013010030031-0133031122333332-2022222103202212-1112020003322313-3310301133030233-2332133322301003"></a>

## Next pages — gift_card_make_purchase_with_gift_card / 130322031302 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0230222010002321-3122033232111200-0232200333223211-0221113020201100-2300300010133322-1222332012230211-2230233210220222-3212120303113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231030002101211-3021012201223332-3132213000113022-3003331132101013-1231113211220210-1103220033220033-0310320323213022-0311011010300023"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation — gift_card_validation / 101301020222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2233110211300212-0131001300211131-3232201331112201-0030210033200023-2223302100120222-1130303321200021-0012300003023320-2200001121122320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card validation.

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

Terraform syntax:

```terraform
gift_card_validation = {}
```

<a id="canonical-3100012023110201-0330310113102211-1130321131131203-0222121032031302-1100122122032210-3203330322033021-2102310003331220-3130311020303033"></a>

## Direct properties — gift_card_validation / 101301020222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313231113031220-0010032333120020-3232323222200302-3323103102200332-1231123211332311-2311002321213111-3031211021333322-1113013233313233"></a>

## Next pages — gift_card_validation / 101301020222 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0102332233313221-3230201311013231-3332113122213000-0033233131231223-1213202012010322-2300110212332222-3200112212102230-2301201313103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103200332200030-2020222130303212-1010030122111032-0223303211202220-3313132121130131-3030330112032220-0231021002022233-0123021120223333"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — shop_add_to_cart / 333320122331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1201020110002233-2321322330023222-2021023012112332-1001213112332131-1202323020201021-2121202322302013-1310132200010210-0220311313032011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop add to cart.

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

Terraform syntax:

```terraform
shop_add_to_cart = {}
```

<a id="canonical-3203033310323231-3033013302230232-3103100303002312-0320231322003020-3212233300131203-2123302313010122-2113033121022212-2033102213230021"></a>

## Direct properties — shop_add_to_cart / 333320122331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032212211211101-2132031303310213-1302000110333001-0131120302111221-2122333310020333-3010020211330021-2110023311123323-0310222203220321"></a>

## Next pages — shop_add_to_cart / 333320122331 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0312012131201110-1220012100103001-0230233212300222-2231131111020300-1000322200210303-1011301113333132-1332011032100132-1233000123033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101311211011110-3232303201123320-2303030131211231-1023123232123103-1103311032011132-3203100102022320-2110102032221202-0122102101220113"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout — shop_checkout / 011231122321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-3011102211003230-2033012321032232-1030013132113020-1021123230321000-0301303311303010-1210102132200312-0110320300211103-0232222022013223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop checkout.

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

Terraform syntax:

```terraform
shop_checkout = {}
```

<a id="canonical-2330303231013112-2121313013110111-0310020322313233-1211323120111123-3232333023302101-2110121311212011-1003331202302013-3021302322220030"></a>

## Direct properties — shop_checkout / 011231122321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111031031122031-3113203313210210-2032231202001122-0022320333211301-2130002211120212-0220111000232231-1113322201011302-0000031231013232"></a>

## Next pages — shop_checkout / 011231122321 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3000132111103122-0012000020321120-0022121331002201-0121102003030310-2020223213122023-1321230302000231-1301230120200301-1021210030213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313031023023013-2021232103112232-2113301321202331-1100133122320013-2202300121133001-3332322102200332-1033003021000213-1202120110110001"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — shop_choose_seat / 330330301213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-3131313001221131-1132003101033233-0021000211321011-2012230331011302-1333213322013220-3020102023202012-3002203210323001-3210013031220321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop choose seat.

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

Terraform syntax:

```terraform
shop_choose_seat = {}
```

<a id="canonical-2203322231232122-2221032300231201-3323011020302202-1312232102020303-3110323202023233-3100020201021130-2223333133023332-2121021213110211"></a>

## Direct properties — shop_choose_seat / 330330301213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201033022320313-3322213200132302-2212233103232310-3210030012133013-2001223333223202-3122321112032300-3001202322013010-2111230020322001"></a>

## Next pages — shop_choose_seat / 330330301213 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0103220311012301-3332112213030231-1213330311123213-0133100320032320-2220002000310113-3333123130202223-1130321200202132-0213020210013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121132033122232-3013220303100312-3203032320302130-3212322323220333-0022313332233022-2320313111120313-2201120010311102-3221302130212323"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — shop_enter_drawing_submission / 230320201122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-3300100013031323-0212123121330231-0122013133013313-1111203201330231-0020310131003303-2133211031021103-0032030111302330-1220100202203111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop enter drawing submission.

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

Terraform syntax:

```terraform
shop_enter_drawing_submission = {}
```

<a id="canonical-1013301201302023-1303230100011323-2132113030121333-2122011221102221-1223303232011103-3110001311012301-0013223213001101-1100110000222330"></a>

## Direct properties — shop_enter_drawing_submission / 230320201122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030221211001113-0332033030020031-1033121033122232-3300202232120311-3310301121111332-1100002023230101-1021013333010302-2202313333221220"></a>

## Next pages — shop_enter_drawing_submission / 230320201122 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3000122332103313-2023021003231010-3100122011033233-2200013320233303-2202211000212000-1012013102112223-3020110210001323-3123211311031223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301302212120133-2310022301133222-0230133032301320-2332303210230000-2031010300222213-0232301022222332-2002120213113230-3112300330303120"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment — shop_make_payment / 100332303031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-3022311302233011-2101132131033230-1301102223110321-0321300232121320-1331330311223231-1300033123100333-1310132300322100-1220103223013103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop make payment.

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

Terraform syntax:

```terraform
shop_make_payment = {}
```

<a id="canonical-3301233321323010-2000212002202312-1333202103023102-3110123030110122-0023033132212232-1130013122002110-0123301011110231-1003330221332211"></a>

## Direct properties — shop_make_payment / 100332303031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123012122021012-3302221113312113-2212011133013200-3031120213111100-1202323103003211-0312101210133302-3002013032300321-3330230033122212"></a>

## Next pages — shop_make_payment / 100332303031 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1112310200102333-2132020132311100-0133310332211100-3122130123002133-1012120131302020-0303030100312111-3212112130111002-1121223323111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213233303033002-3301112022112312-1023032123120003-2330231221312302-1201030001001223-3302110300101130-1123123111032022-1113010132121323"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order — shop_order / 002311013302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-2201213011333033-0213333310233111-1031120111331332-3322013010030100-2201011332100221-2023331320310013-0200022112320123-3333210311231101"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
shop_order = {}
```

<a id="canonical-3232113001003133-1131222211003201-2301000003213232-2333300102032300-0202333311212212-0310102331103202-0120320030312121-1003103301223122"></a>

## Direct properties — shop_order / 002311013302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113013102332202-2103110003102100-0132133033322201-1220300302230202-3132023113122313-2232122121031030-0002133310200131-0332033010333303"></a>

## Next pages — shop_order / 002311013302 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1330011013210322-3321011001101001-1023103132023230-1201110123211103-3232323202322332-2231212200030332-3233320030121220-1321121233312332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002230232111112-0011330010100230-3000110331321032-3233001031303113-0301131102210003-3322230131300003-3122132313202033-1121222330022301"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — shop_price_inquiry / 323133111311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-2223310111321002-2231332313133213-3012330333122000-2120123123312123-3331110132103131-1021001223230300-0323020223132201-0233211223101000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop price inquiry.

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

Terraform syntax:

```terraform
shop_price_inquiry = {}
```

<a id="canonical-0230133110003031-3213111312001100-1323110331301213-0103333003012211-2032311133301311-1133002011310300-0030231202202033-1311332130210012"></a>

## Direct properties — shop_price_inquiry / 323133111311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3332132211232320-3200300230222203-2011012132313200-2111320133010322-2333131223031002-1231222201332021-2321320221110032-2212310210323100"></a>

## Next pages — shop_price_inquiry / 323133111311 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3331130003130133-0113302011022001-3111013010132020-3133323213022300-3002131021232133-3321303301112121-2313123023033000-2023122131203212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011202213332003-0333030023032332-0213012201121221-0011213321323300-1013023111012231-0213121011323213-2333203132133231-1301200001100023"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — shop_promo_code_validation / 122003130101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-1210113133001210-1311311310033221-3213232330130110-1032022023303310-1323212331200120-1321102221131111-1222330033233132-0003313223013232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop promo code validation.

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

Terraform syntax:

```terraform
shop_promo_code_validation = {}
```

<a id="canonical-0200222300200022-3123110012012110-0212211233201022-2111112322113333-3000222032000203-0201001232303033-1103301230101003-2002322013222101"></a>

## Direct properties — shop_promo_code_validation / 122003130101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102311203112121-1121301121330012-0121110013131310-0030200113032123-1103331203120110-2312210320111113-0302130230232323-1111313103100011"></a>

## Next pages — shop_promo_code_validation / 122003130101 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1311312001112121-2113030033113022-0022203223122233-2131331110033331-3000033210003031-2103302000321130-3313130202300301-2231133103211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023332121113313-1103222232323200-2202233310220012-1123300102203211-3200003020000201-3122313310200301-2333003311203221-0113221320331231"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — shop_purchase_gift_card / 300212022001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-2203302032011201-1101122131203101-1312012130320223-2313000113131201-2010303100322020-3202000301130211-0223202001022023-3002113011213320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop purchase gift card.

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

Terraform syntax:

```terraform
shop_purchase_gift_card = {}
```

<a id="canonical-1002210222011320-3233111212122031-2010032220131332-3323323330020022-1130201131003021-0230132323101020-2133031123222003-2313203320222000"></a>

## Direct properties — shop_purchase_gift_card / 300212022001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001130212020322-2213012120112233-1010230311210130-2003222031211111-3120322203212322-3133033223231133-1011020021100111-1213231230023233"></a>

## Next pages — shop_purchase_gift_card / 300212022001 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1010222323231020-0103010320210313-3221223233333120-0000122131303311-3302311133013100-1132202331333233-1321222300332001-3120032022320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301210330022021-3011330313330223-2233303033220120-0121013310323321-2032230212222210-3011231020133332-1112333321320033-0211021231020320"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — shop_update_quantity / 133110233131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-008.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-2301220021320123-1223230131103210-2231312300100101-2210333020333023-3000013320211212-3103220331302202-2100002302230132-3132010300300132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop update quantity.

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

Terraform syntax:

```terraform
shop_update_quantity = {}
```

<a id="canonical-2110112013210120-2302120111231310-0230032230321010-2213302020323111-1133012103100030-2312122010301313-1233111010213102-1131011330211310"></a>

## Direct properties — shop_update_quantity / 133110233131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331113330020233-1332021031123033-2201133132120322-2031220021003312-1231133010102202-3333213201003102-3311211013113121-1322001023001002"></a>

## Next pages — shop_update_quantity / 133110233131 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213002211030123-3320030232111023-2311103303010011-3323210322332033-1020303032111111-0122201221323220-1303020111131002-0003132302331311"></a>

## bot_defense.policy.protected_app_endpoints.headers — headers / 003222300300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.headers

<a id="canonical-3102222020223230-3310003230032122-2301123003301113-0213102211203121-3113202333121320-0102103303333122-1113020211130311-1013321000211111"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for various HTTP headers that need to match. The criteria for matching each HTTP
header are described in individual HeaderMatcherType instances. The actual HTTP header values are
extracted from the request API as a list of strings for each HTTP header type.

Upstream description:

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032321203103302-3100313032210100-1012210230301121-3031002303022111-1101212113011303-2012213121033220-0101003203210202-0120101211021013"></a>

## Direct properties — headers / 003222300300 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-0122133301130111-0112232021201013-3020233100102110-0212022130313113-2220232310102110-1331233220131032-1123121333200210-3220122030320233): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-3131020100232021-1123311131013322-1002103010021012-1022201202200331-2310200232001010-2130311103003123-3033102302020110-3321001112231220): complete subsection reference.

<a id="canonical-2131011031210100-0300132030333003-3300230232120310-1210112021331211-2212202032122132-1331333113020103-2131303111200223-3130333303221112"></a>

<a id="canonical-2323133032131232-2123212101032001-0033102330021321-2112032130201303-0030131203311323-1331312131102000-1322132000113010-0201313121333202"></a>

## invert_matcher property — headers / 003222300300 / 4

Type: `"bool"`. Optional.

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

- [item](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321203021032210-3330200101033230-2201111101320001-2130020001023130-2103033030011112-2221113010032303-2313032100023310-2120230313333013): complete subsection reference.

<a id="canonical-3110303001023332-1101032111233323-2330122011012122-1221111303330312-2130333000103221-1212230232032311-1213122310031002-2332300133012322"></a>

<a id="canonical-1232031102310220-0212230033122222-2030333020013220-2200111212220022-1212032320332113-0023213131123122-2202130310333132-0213131012032123"></a>

## name property — headers / 003222300300 / 5

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0011132020011201-3313200013303012-2031223130110232-1123032302323021-3112222133112133-3100212300012001-3133313120320203-1132111203003120"></a>

## Next pages — headers / 003222300300 / 6

- [bot_defense.policy.protected_app_endpoints.headers.check_not_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-0122133301130111-0112232021201013-3020233100102110-0212022130313113-2220232310102110-1331233220131032-1123121333200210-3220122030320233)
- [bot_defense.policy.protected_app_endpoints.headers.check_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-3131020100232021-1123311131013322-1002103010021012-1022201202200331-2310200232001010-2130311103003123-3033102302020110-3321001112231220)
- [bot_defense.policy.protected_app_endpoints.headers.item](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321203021032210-3330200101033230-2201111101320001-2130020001023130-2103033030011112-2221113010032303-2313032100023310-2120230313333013)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0122133301130111-0112232021201013-3020233100102110-0212022130313113-2220232310102110-1331233220131032-1123121333200210-3220122030320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320211212200011-1313322320120011-1200033010212202-2032023132112001-2321131300022330-2122302130003102-0011012211001010-2101011032302312"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_not_present — check_not_present / 233123300113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- bot_defense.policy.protected_app_endpoints.headers.check_not_present

<a id="canonical-2030311011202300-3330113031331232-0321203030101022-3302332332210322-1233032022223200-2323233212300111-3002300012020102-1112103021332031"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

<a id="canonical-2302122100223122-3311003013101022-3303123010303311-2010021102012023-2211323333001121-3031212213201002-0133112010110022-0203002223101103"></a>

## Direct properties — check_not_present / 233123300113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201123323101122-1230323310301123-3201223221301211-1020123132122303-1132211223131322-2010101123233002-3203212112003202-1210031030100021"></a>

## Next pages — check_not_present / 233123300113 / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3131020100232021-1123311131013322-1002103010021012-1022201202200331-2310200232001010-2130311103003123-3033102302020110-3321001112231220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331001120103123-3101110220032112-3320322300020301-0132132302021033-3130002330112323-1312113330322312-0202002002022210-3120231213003203"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_present — check_present / 122222332333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- bot_defense.policy.protected_app_endpoints.headers.check_present

<a id="canonical-1013021122003231-0331023302010303-0123030111330330-3033132130031033-3232302331313230-2230121130211212-2333111103312123-2013213120211330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

<a id="canonical-1100333233000202-0012033211223120-2203120030202121-3223313232033123-1030333200210020-3223201011312323-0022322103021100-2002102221210013"></a>

## Direct properties — check_present / 122222332333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322301001310112-0003223331312330-2122300012033310-2213203313032010-1233212120222222-2321323113200301-1211231002203320-1310213222013311"></a>

## Next pages — check_present / 122222332333 / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1321203021032210-3330200101033230-2201111101320001-2130020001023130-2103033030011112-2221113010032303-2313032100023310-2120230313333013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
