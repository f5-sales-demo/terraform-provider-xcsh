---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1010100213022302-3302232130322211-0220320222312200-1302220032103031-1212200332331000-3232222111001123-3303032231011111-0112031200212220"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — authentication / 313332130332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-1133332031311322-1311020313030110-0010112310320222-2322320121033113-2020211130333300-0003001023102332-3023200030110310-0311003112130133"></a>

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

<a id="canonical-3321323222131020-3102222111310311-2203213111030223-2030213123301012-2213320011330022-2300011230012111-0313113302132322-3332212132011011"></a>

## Direct properties — authentication / 313332130332 / 3

- [login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133): complete subsection reference.

- [login_mfa](resources--http_loadbalancer--reference--group-012.md#canonical-2310113000131023-2200111020211101-2100210032102203-0332300332120322-3000122313320123-0202223030203021-1120231203232311-3003102313121100): complete subsection reference.

- [login_partner](resources--http_loadbalancer--reference--group-012.md#canonical-3220333122100331-3330111221312311-0133330123031210-1101111222312112-2011112023103020-1013310232003130-0323010222132310-2103111003002202): complete subsection reference.

- [logout](resources--http_loadbalancer--reference--group-012.md#canonical-1021233303223023-3111113131122032-2132210010322203-0320212131213331-3213030221323113-2211111022330123-1301300102012333-0132323333331321): complete subsection reference.

- [token_refresh](resources--http_loadbalancer--reference--group-012.md#canonical-2213022330021103-0212013331310113-2202010210023233-0020032032230122-0331132032133310-1113321111003033-2200321032022301-0113131200100033): complete subsection reference.

<a id="canonical-1203311303220113-3321202000313201-0032303232201131-3011213023021121-1230123333220103-3000223333223322-1021130231200320-2201021133213130"></a>

## Next pages — authentication / 313332130332 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](resources--http_loadbalancer--reference--group-012.md#canonical-2310113000131023-2200111020211101-2100210032102203-0332300332120322-3000122313320123-0202223030203021-1120231203232311-3003102313121100)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](resources--http_loadbalancer--reference--group-012.md#canonical-3220333122100331-3330111221312311-0133330123031210-1101111222312112-2011112023103020-1013310232003130-0323010222132310-2103111003002202)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](resources--http_loadbalancer--reference--group-012.md#canonical-1021233303223023-3111113131122032-2132210010322203-0320212131213331-3213030221323113-2211111022330123-1301300102012333-0132323333331321)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](resources--http_loadbalancer--reference--group-012.md#canonical-2213022330021103-0212013331310113-2202010210023233-0020032032230122-0331132032133310-1113321111003033-2200321032022301-0113131200100033)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121002102333331-0322220003121113-3210313310003311-2013103231022211-2213021310132010-0201332202122300-3311110211120031-0103311301222212"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — login / 101102303321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-2000023133232122-0031031300230132-1003333032112100-3223101333310311-1022032123222101-2301102323203112-3231311003322220-1123131312302120"></a>

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

<a id="canonical-0102010233112100-2213203321203132-0001321331131112-0320302122100211-3110212101201131-1232000302321110-1323210110022132-3321022033330222"></a>

## Direct properties — login / 101102303321 / 3

- [disable_transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2002233110222300-3023110131333200-3002232311133002-0300020211021100-0310313001231321-2012123332113020-2201223333213103-1123011202211201): complete subsection reference.

- [transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100): complete subsection reference.

<a id="canonical-0032123001300122-3110203130130310-1223201201010131-0123003332200011-2113310323202021-0031212013133033-2123303013220312-3203131110120231"></a>

## Next pages — login / 101102303321 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2002233110222300-3023110131333200-3002232311133002-0300020211021100-0310313001231321-2012123332113020-2201223333213103-1123011202211201)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2002233110222300-3023110131333200-3002232311133002-0300020211021100-0310313001231321-2012123332113020-2201223333213103-1123011202211201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030210001011032-3013230031212221-2100113121132321-3112312002300121-3221012132313130-2331013321002332-3203203302220333-1332120220221011"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — disable_transaction_result / 310031033103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-1230212012331220-3201000120232113-2323032010302323-2331112010203031-1233201333111210-2323333012222311-0223202122112221-0230302011033023"></a>

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

<a id="canonical-3331210023000103-3232001302112331-3233323303331112-3313010212231111-3323000233030232-1331222223211132-3003030030123331-1221012030122220"></a>

## Direct properties — disable_transaction_result / 310031033103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310231221003112-2210323031021332-1032100001123022-0203231213132012-3210330003022200-0300030113001022-2223120232321021-2032120031100020"></a>

## Next pages — disable_transaction_result / 310031033103 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211230112033101-1010202103333120-3023200033221031-3031332312211000-2010033212313321-2133021233113120-0221220113321310-2222031010003212"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result — transaction_result / 113021230130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-2213203032213033-0033231122310222-2100321111000002-0102121120010033-1323313131312103-1220332301321210-1002001323031123-3211001322003003"></a>

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

<a id="canonical-3222220030111223-1021102323030201-2310030220222301-2021003130003322-3111330002330102-2012032301320220-2123012210132031-0230231213212101"></a>

## Direct properties — transaction_result / 113021230130 / 3

- [failure_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-3232211032320313-1130031322033023-2221012313221033-0201130023033302-3311022132223311-2132101303331210-0033233300233213-3222103313032301): complete subsection reference.

- [success_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-2121212102313303-2313010110313312-2123211221312000-0233121300030033-0020323300200312-1232031220320310-0003323033212203-3233233103331030): complete subsection reference.

<a id="canonical-0102323120033003-0113222122010032-0101120302233122-3330030022003320-1222100313022012-1120033230100220-2112330330212122-0133322300200020"></a>

## Next pages — transaction_result / 113021230130 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-3232211032320313-1130031322033023-2221012313221033-0201130023033302-3311022132223311-2132101303331210-0033233300233213-3222103313032301)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--http_loadbalancer--reference--group-012.md#canonical-2121212102313303-2313010110313312-2123211221312000-0233121300030033-0020323300200312-1232031220320310-0003323033212203-3233233103331030)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3232211032320313-1130031322033023-2221012313221033-0201130023033302-3311022132223311-2132101303331210-0033233300233213-3222103313032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203301111220322-1201011212213333-3331320133222130-3202321303032302-0210010220011103-1233212232223031-1230222012001310-3323031132331213"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — failure_conditions / 113003012230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-1220001002212120-0013022033000122-2312323233202321-1210300112101311-0202110302312002-2013003102120001-1032001010030201-2023210011010201"></a>

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

<a id="canonical-2120000011033203-3311023103302232-2020121331302211-2012201321212021-1333103311100023-1130033100031333-0131130131212013-3032231222112132"></a>

## Direct properties — failure_conditions / 113003012230 / 3

<a id="canonical-0100221213133020-0003123221310330-3003220010231321-2030303132200303-2101010022030331-0100110213311121-2213332110213110-2310320123212132"></a>

<a id="canonical-2021011232031023-2111101003032300-0012111302000223-3113112231112223-1222003100130321-1320321113313333-3122012033213120-3320202231132133"></a>

## name property — failure_conditions / 113003012230 / 4

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0331100330021000-3121231121202220-0320210110301311-0021312200300302-3021010213012001-3011133233022212-3201233111330031-3202032032032102"></a>

<a id="canonical-0200203103131311-2112110212331121-3333120113013010-2133322310312110-2330203302030300-3232100232221002-2123122033320003-2300231120032110"></a>

## regex_values property — failure_conditions / 113003012230 / 5

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

<a id="canonical-2112121011212212-1022111322200100-2033312113132220-3003001201132310-0000222331220131-0213322303310231-3000000010123020-1232333033103320"></a>

<a id="canonical-2033101020231013-1103123231022103-0132013110220001-3333122000312221-2322303111103011-1233302332103011-2020330313300120-0122002222032202"></a>

## status property — failure_conditions / 113003012230 / 6

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

<a id="canonical-1022311033013332-1013303230213331-0003312130313002-1310033321032113-2333111033031331-3110021211330022-1120301002113022-2023033030101031"></a>

## Next pages — failure_conditions / 113003012230 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2121212102313303-2313010110313312-2123211221312000-0233121300030033-0020323300200312-1232031220320310-0003323033212203-3233233103331030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322001201211010-0123321200330331-1001112030000230-0100102011012212-0102230301012302-0101010313201022-1320233100313203-3300311313102300"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions — success_conditions / 123020232221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](resources--http_loadbalancer--reference--group-012.md#canonical-0130122103111102-2111313002210101-1012020303103332-3130331011230100-3313332232231331-3102233322311323-0233021111211131-0010332000203133)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-0132123133133003-0222312331000231-3002010302113221-1330132003002113-1233133100122023-1212330102211221-3011220030321000-3230113020001222"></a>

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

<a id="canonical-0022331032211202-2302221033023121-1322022111122312-0033202203211311-2132330121112033-3311201132302202-1110003303233201-2102310220212022"></a>

## Direct properties — success_conditions / 123020232221 / 3

<a id="canonical-1310212002110102-3021310131312302-1012002302002301-2210300000001200-2132111312103132-0110311323332122-2010333033013023-3201133121201123"></a>

<a id="canonical-3021123211013230-0213212112302233-1022321222201132-0203003333320132-1202022301301212-2212121333223120-1031032100311003-3231212201111032"></a>

## name property — success_conditions / 123020232221 / 4

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-1300212111021120-0332201233233100-3323202310101322-2030002220223011-1213320133323032-1113032110201211-0300330331313220-1122032110331233"></a>

<a id="canonical-0320312222020013-2300101300033213-0220320320231331-2113102121100022-0122102002322021-3232313121030220-2203101212321200-2110010000231310"></a>

## regex_values property — success_conditions / 123020232221 / 5

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

<a id="canonical-0013322233020021-3103302023212010-0331201000111122-1122313301032233-3212333311300311-1012203302131032-1103133101210012-3130023331211302"></a>

<a id="canonical-3103113232011221-1301130032222003-2000133303310223-0200133112201203-2111200103022102-0333321231021210-2232231121320233-0032231002212132"></a>

## status property — success_conditions / 123020232221 / 6

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

<a id="canonical-0210012302001233-0303202321213011-2211120012333000-2313010332031033-2302231111000301-2101032321302311-1222132031002221-2310122213022330"></a>

## Next pages — success_conditions / 123020232221 / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](resources--http_loadbalancer--reference--group-012.md#canonical-2011231221330302-3312302222301223-0200002303013032-0111032033103123-3033021333312021-2312022303311313-2331301023311331-0323011213130100)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2310113000131023-2200111020211101-2100210032102203-0332300332120322-3000122313320123-0202223030203021-1120231203232311-3003102313121100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333213120133313-2023221223001133-2000213133131103-2303202111120122-2130122030230003-0131010302110323-3333233102023312-2222200101103210"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa — login_mfa / 020320002013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-3231230111201131-1132221302212231-3013122133130222-2021033100001001-3132123211211133-0221030103002112-1002330031020213-3011233321323332"></a>

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

<a id="canonical-1020213032201110-2320220002130112-3010013010112212-0220320021330022-3123130302113203-3213121221122003-2312120231011303-3123331302111121"></a>

## Direct properties — login_mfa / 020320002013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312220030133022-3331333321013310-3322112111102101-1101102030222102-1033330233331131-2332331031213312-0322300121120210-3203210313133113"></a>

## Next pages — login_mfa / 020320002013 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3220333122100331-3330111221312311-0133330123031210-1101111222312112-2011112023103020-1013310232003130-0323010222132310-2103111003002202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320320331000223-0101012300001321-1120312320120302-2212211112002103-1232121330023132-3031320110331313-2102030032210122-2321322222310330"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner — login_partner / 102131033201 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-1322110233302311-3102112301010030-3123313110123113-1011300232231331-0321131211233230-0100230211001200-1120330232033210-0020103010213122"></a>

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

<a id="canonical-0031211323131332-3212212001232020-1010233231122000-1330022023102330-1023333111212333-3022023213312221-2033001301120323-3213113001121313"></a>

## Direct properties — login_partner / 102131033201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122211020123132-3331131010201322-0223332220133322-3231021323013003-2203313112320221-2111112110030002-3023101112031203-3010132210102021"></a>

## Next pages — login_partner / 102131033201 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1021233303223023-3111113131122032-2132210010322203-0320212131213331-3213030221323113-2211111022330123-1301300102012333-0132323333331321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332033200200212-2021230031331230-1231011220311300-1331012303012201-3303213132223023-2131033010311000-3232231112002010-1230233220221133"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout — logout / 232303023230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout

<a id="canonical-1112113202210331-0311013230101212-2213011310031302-3332103201331220-1211320202303322-2130112301331202-2033310322221332-2133113110032012"></a>

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

<a id="canonical-3313202311312320-3011231302210032-1020103222123111-0220120120132202-0213023102032232-3232003110032102-1100121023321213-1311012121321022"></a>

## Direct properties — logout / 232303023230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231022231312131-0320010023231211-3323111210002202-2332213210321221-0322002130131013-0013023333003231-3200310331021203-3001123010003001"></a>

## Next pages — logout / 232303023230 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2213022330021103-0212013331310113-2202010210023233-0020032032230122-0331132032133310-1113321111003033-2200321032022301-0113131200100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011321223202121-3103330111102123-2311312201033223-2020233330323033-3200001111331031-2311333230120312-2322222012213220-1220210033003013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh — token_refresh / 230123110200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh

<a id="canonical-1320323000202010-2121100233003322-3013121322021303-2012311211101301-2103322111131212-2123201002102303-1021233020233032-1311130001301112"></a>

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

<a id="canonical-1210301222120310-0233012331123232-2223100132303201-1313113133213213-3110010212333312-3112110310331001-0331121230133003-0310102233123120"></a>

## Direct properties — token_refresh / 230123110200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332331220330132-3030112132232012-0101022330202311-1122220203110300-1101102030113311-3330331201212101-1310322013011210-0101003311331131"></a>

## Next pages — token_refresh / 230123110200 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](resources--http_loadbalancer--reference--group-011.md#canonical-3100100102113232-2222112221110221-3031301000022320-1013103323013100-2213321113300323-0200332211203332-0133323123121222-0012203331303302)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222203100322230-2011021211333321-0113320101011311-3023013310223230-1321121203233020-0031233321311012-2333130212000210-2230230322011222"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services — financial_services / 310110312323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services

<a id="canonical-0023313203202312-3331201300113301-0020121131213012-1113121221002220-0122011123130321-3311100231132101-0223121301330300-3303332233222001"></a>

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

<a id="canonical-1002102132123311-0311132132101132-3003301013100001-1002003013202301-1013112323120232-1332123330022210-2301113113213202-3221003013222111"></a>

## Direct properties — financial_services / 310110312323 / 3

- [apply](resources--http_loadbalancer--reference--group-012.md#canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321): complete subsection reference.

- [money_transfer](resources--http_loadbalancer--reference--group-012.md#canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001): complete subsection reference.

<a id="canonical-0320121032232310-1301010322101130-2232021321332333-0201121103012311-3130133021320313-0110233023102031-1321301333232230-1320113223102023"></a>

## Next pages — financial_services / 310110312323 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply](resources--http_loadbalancer--reference--group-012.md#canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer](resources--http_loadbalancer--reference--group-012.md#canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210110300010200-0110020202110011-1220013200221233-1130233202032312-3220311022133232-3211232232011112-3122001102202231-3202031230113312"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply — apply / 233100133032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-0303023201021203-1213331020131330-2121232130223013-0023102213212322-1232323320320101-1003212211102112-2333233020020303-2110111300023020"></a>

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

<a id="canonical-1230223310031120-2300111013320030-3331312010022330-0033302101131031-3112030331333011-1232223301200102-0120333200030212-2203011231301110"></a>

## Direct properties — apply / 233100133032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120002212330103-0132000231002332-3030201031232312-0011102031303103-3301310111333312-1100201211120310-2203221200222023-1112012300221223"></a>

## Next pages — apply / 233100133032 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012210332003202-1031001101230223-2213113220210333-0022230110010311-2230311300313111-2023202332202332-0312020133031122-0333303133201013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer — money_transfer / 002201021010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0000000033330222-2133212001221133-3120230302002120-0112000331223313-0220013330232103-0322031210302120-2232222330001331-1021230033132212"></a>

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

<a id="canonical-0230311302121311-2330002131330332-3011222002230331-1023310310022113-0003133031211313-0112310232331131-3321230311222000-0101030012211011"></a>

## Direct properties — money_transfer / 002201021010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013300101231031-0002332303333120-1303101313030332-2222122301002100-1330220120212113-2310211223323233-2320210113121112-0032123121321312"></a>

## Next pages — money_transfer / 002201021010 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-012.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000131123232020-3331200202002101-1130022221023032-0202201221021201-3012120000020113-0321220233001021-0130012011003102-2320212322220222"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight — flight / 012113122032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-0112222123130203-0300003200212330-0222031122230301-0303100303302031-3032310120131132-2220033133001120-2002311232131032-3303122222032103"></a>

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

<a id="canonical-1030332203102120-3223120031000010-3021010012300222-0222010133000032-3322233302323313-2222111201011213-3213010030021100-3130222122022232"></a>

## Direct properties — flight / 012113122032 / 3

- [checkin](resources--http_loadbalancer--reference--group-012.md#canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200): complete subsection reference.

<a id="canonical-2323111200302102-3003212301203003-3021022322202120-0000202232030313-0200023001221123-0303303211201123-2031002310131020-3001333221210012"></a>

## Next pages — flight / 012113122032 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin](resources--http_loadbalancer--reference--group-012.md#canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001021331130220-3322320120130032-0302002022030002-2220001111331300-2301321033212213-0132200312210232-1211210121302230-0101031332100001"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin — checkin / 021013332021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-0022211211033231-3330201033232230-2302333012201023-1310113121121123-3122322102130112-1210331131233323-0201011221210222-0321231200003220"></a>

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

<a id="canonical-3313013232322233-1113223111202101-2010021321121301-3230112021212003-0101100103313033-2211203302111010-0221233110100010-1302323002032121"></a>

## Direct properties — checkin / 021013332021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121202221021012-1021101210211031-2200122023102133-2320132203223320-3332000131301123-0112203300230023-0323103223212313-1330110201333020"></a>

## Next pages — checkin / 021013332021 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-012.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003213132023302-1013011011302233-2112233000231303-0232021301321330-2133011033130113-0223233201222310-0313310213113032-3132301330331311"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management — profile_management / 332111200321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-1123331100331221-3233022320303222-3010002023010313-0333302033211002-2233022023130030-0131222133203300-3331212323101213-3222122011321301"></a>

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

<a id="canonical-0322200300010213-1032301103200323-1011103123130100-0010121322210123-0313101012221011-1033032220233131-1230212012020120-0220220303311222"></a>

## Direct properties — profile_management / 332111200321 / 3

- [create](resources--http_loadbalancer--reference--group-012.md#canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131): complete subsection reference.

- [update](resources--http_loadbalancer--reference--group-012.md#canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201): complete subsection reference.

- [view](resources--http_loadbalancer--reference--group-012.md#canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203): complete subsection reference.

<a id="canonical-1312113110110231-3230321032033100-1232103220200133-1201312231323301-2030200313232301-3300113023202113-3021121013012113-2223213033100131"></a>

## Next pages — profile_management / 332111200321 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create](resources--http_loadbalancer--reference--group-012.md#canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update](resources--http_loadbalancer--reference--group-012.md#canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view](resources--http_loadbalancer--reference--group-012.md#canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031221202213133-2213130032012122-2311103203112221-0302211000103211-3331101121313211-3222030020003311-1310113301131111-3110312232323312"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create — create / 312321330313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-3112132212320123-0132131233323130-3000333221101133-2230201003300011-0203300023001230-0100233203213230-2221202331202232-1223021133010323"></a>

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

<a id="canonical-2010322023312201-2301011302322100-0102003231122231-2101220223211113-1330231203222313-2332120131203001-1133103002330130-2200011301111210"></a>

## Direct properties — create / 312321330313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302012122303033-1202010133103130-1030032331020122-2311022030213100-1220132212133120-2313100032220133-0231310123021333-1021102012130000"></a>

## Next pages — create / 312321330313 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232320031202210-0023332100130002-2233201131112130-0200020133303032-1000310230032222-2200310022020002-2121212102131030-1200300103331013"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update — update / 303220221311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-2211033010302312-2231010230330030-1012013003110302-1232021101132131-2121300210021132-2000031213012223-3022322220003310-3001120203003110"></a>

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

<a id="canonical-0210312110113312-2003201011001031-0223021202230022-1131212210130222-2220001211132230-0332302010112210-1003223023331202-1112303003111021"></a>

## Direct properties — update / 303220221311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110313001120110-2232002221331223-3321332122321003-3320003112300131-3113011320012311-3132111323203220-2230312300100133-0001232221332302"></a>

## Next pages — update / 303220221311 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131130122001033-2333300220013221-3331331112212101-0200033200303220-2332030110122313-3323201001103302-2222002013010102-2112221010200002"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view — view / 211220301230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-3101203232122311-2011111033011310-1013100001021311-0231131032121003-3221320013200030-2323122132311302-0222110233031330-3101230201111132"></a>

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

<a id="canonical-3033323332221023-0303131210130223-3230231330003200-3203131020330101-3221120120123233-1220131132310012-3122330311032331-1221012020303212"></a>

## Direct properties — view / 211220301230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313232320233303-3022311120133211-2322101020300200-1032203201200210-2323021022230231-1211233030211001-0130312230010331-3101103312020333"></a>

## Next pages — view / 211220301230 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-012.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212312203033233-0113311320023313-1211232222033201-3101020100333320-2230330312330213-3002320100020013-0210103301010022-0201330203031131"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search — search / 302003210130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-0230313123323323-0220201122122023-0301030122333211-0020003122122023-0210023231211332-3122202303211332-0132113310210201-1012011012210022"></a>

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

<a id="canonical-2213232101201010-1202003033002103-3330030123210323-3133200233302010-3322211123230300-3330033021223101-2232311322322033-1121321013030322"></a>

## Direct properties — search / 302003210130 / 3

- [flight_search](resources--http_loadbalancer--reference--group-012.md#canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032): complete subsection reference.

- [product_search](resources--http_loadbalancer--reference--group-012.md#canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102): complete subsection reference.

- [reservation_search](resources--http_loadbalancer--reference--group-012.md#canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212): complete subsection reference.

- [room_search](resources--http_loadbalancer--reference--group-012.md#canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133): complete subsection reference.

<a id="canonical-3001203131103213-3112230310021100-3211110110322200-1331133102221331-2220003211111332-3031030233333131-2103202300211030-3020302202022301"></a>

## Next pages — search / 302003210130 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](resources--http_loadbalancer--reference--group-012.md#canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](resources--http_loadbalancer--reference--group-012.md#canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](resources--http_loadbalancer--reference--group-012.md#canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](resources--http_loadbalancer--reference--group-012.md#canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011020120113032-2201100010113213-2322020302222212-0020120213133213-1123303110311123-3132012300302331-0331223332001320-1230101200332102"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search — flight_search / 303213213222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-3020011233310330-0323311030233213-0020003113013100-3330002031312103-0032303312101203-1012002113320330-1103011123221302-0032333222100100"></a>

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

<a id="canonical-1032113313032223-1133012312301333-1130103203021210-2122103210122301-1122130320132330-0120112001312121-1123333123221310-0212023312100003"></a>

## Direct properties — flight_search / 303213213222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0100113132101011-0201020322013323-2201230231112010-2232000303211000-0123213020213121-2201003220201033-3301302003303222-0103202101312200"></a>

## Next pages — flight_search / 303213213222 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201122101102232-0032303130233323-0202323332312333-0302110133032301-0130211000100223-3223113303100211-1011312113120203-2012301003321100"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.product_search — product_search / 333321302121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-1031211310300013-1000320303030231-2202322331203212-1021021131030310-0202331223320030-3003320022333312-3111120232010300-0213110120223330"></a>

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

<a id="canonical-3113230311310322-2020121000121302-3111302222121201-2020302110102201-2211212200002313-1311332321332233-2312000210300120-3123100101002311"></a>

## Direct properties — product_search / 333321302121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333111001132123-2220202111101133-3132203130110111-2233112311131112-3333123103020011-3120223033220330-0220030022120121-0231220021323021"></a>

## Next pages — product_search / 333321302121 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000331212031102-0323221003300200-0031011100111001-2120101012311330-0313330301220122-0011203130311002-1000230131303321-3011132132103132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search — reservation_search / 232022131013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-1010113203001113-3112203002030112-1222020112332232-0003232302223322-3221330112221021-1213302000313331-3202330032001110-0220022103230212"></a>

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

<a id="canonical-3203331133233002-3023313330003021-1220302211100300-1200333031011320-2120120222201031-0030023300330030-1031102111011032-2002022212000303"></a>

## Direct properties — reservation_search / 232022131013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223312213123001-2321311021331131-2303312223101330-1031012200213313-3231101132002312-1222310020113202-1322132201012201-0203011133030223"></a>

## Next pages — reservation_search / 232022131013 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132101223232031-2312033000031302-2331111313000020-0320003221020200-0101203130322311-0210030113212030-1232221322212003-2320331012331333"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.search.room_search — room_search / 030221012230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-1103211223102301-2002230123113121-2321131010301210-1223233300302213-1130110321131321-1200223331303033-2213333212332130-2232322012100101"></a>

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

<a id="canonical-2121220020022313-0113133302221033-0313101012130012-2220113120101310-3022300312211002-1311201013031121-0110023320013333-1132301322331111"></a>

## Direct properties — room_search / 030221012230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110131000120101-1203013123011101-3112013312301020-2012123102222031-0133311202122323-0211222301303102-0302003110010222-0100101233033203"></a>

## Next pages — room_search / 030221012230 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-012.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132001231321303-2320210012021323-1222331232123013-3300230013333333-0121123111031302-0211323012312230-2331121330101313-2330221313112232"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards — shopping_gift_cards / 032032103210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-2312330022202233-0031211231120002-3010130013100312-2331212132230123-2100132331223010-2331213111201011-1210202211111330-1131320203012003"></a>

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

<a id="canonical-0232221023030232-0001201332202221-2312302030331312-1220130302021203-3010231121121100-1213133102131100-0202120031123222-1220313111311220"></a>

## Direct properties — shopping_gift_cards / 032032103210 / 3

- [gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103): complete subsection reference.

- [gift_card_validation](resources--http_loadbalancer--reference--group-012.md#canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223): complete subsection reference.

- [shop_add_to_cart](resources--http_loadbalancer--reference--group-012.md#canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203): complete subsection reference.

- [shop_checkout](resources--http_loadbalancer--reference--group-012.md#canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320): complete subsection reference.

- [shop_choose_seat](resources--http_loadbalancer--reference--group-012.md#canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300): complete subsection reference.

- [shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-012.md#canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301): complete subsection reference.

- [shop_make_payment](resources--http_loadbalancer--reference--group-012.md#canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012): complete subsection reference.

- [shop_order](resources--http_loadbalancer--reference--group-012.md#canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222): complete subsection reference.

- [shop_price_inquiry](resources--http_loadbalancer--reference--group-012.md#canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101): complete subsection reference.

- [shop_promo_code_validation](resources--http_loadbalancer--reference--group-012.md#canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211): complete subsection reference.

- [shop_purchase_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123): complete subsection reference.

- [shop_update_quantity](resources--http_loadbalancer--reference--group-012.md#canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231): complete subsection reference.

<a id="canonical-2232231200221331-2220020231202002-3330303323233201-2333132313313101-3012210112132330-2013223233132200-1023303012032310-3112102022300311"></a>

## Next pages — shopping_gift_cards / 032032103210 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--http_loadbalancer--reference--group-012.md#canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--http_loadbalancer--reference--group-012.md#canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--http_loadbalancer--reference--group-012.md#canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--http_loadbalancer--reference--group-012.md#canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-012.md#canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--http_loadbalancer--reference--group-012.md#canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order](resources--http_loadbalancer--reference--group-012.md#canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--http_loadbalancer--reference--group-012.md#canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--http_loadbalancer--reference--group-012.md#canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--http_loadbalancer--reference--group-012.md#canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--http_loadbalancer--reference--group-012.md#canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311000113031221-1331021131133002-1222322102013023-0231121201323122-2002210113330032-2312103210213313-1130010021001121-0332220011201112"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card — gift_card_make_purchase_with_gift_card / 001113322121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0203201032022321-1101122012013302-1123023201002231-2333310320323031-1131211011333011-1203222223312020-0010021132303133-0023201002032122"></a>

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

<a id="canonical-3333321002001312-0001232031221003-0131220202111330-0021223121123122-3112010213202232-0301222123131312-1021231311330003-3100101322003110"></a>

## Direct properties — gift_card_make_purchase_with_gift_card / 001113322121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101022131213310-3313200132210323-3122133112332320-3213332122111132-2311101112110132-1221023121123212-2023103313230002-3301233133333113"></a>

## Next pages — gift_card_make_purchase_with_gift_card / 001113322121 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102203023011312-0333122303320030-1030130023022300-3202200021333230-3120211012233303-1023333312210102-0233211100110303-2123031122033200"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation — gift_card_validation / 210202330110 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2302300101313113-0221230010311211-3330323220321021-3211313301233112-3232210001312323-3120120320020131-0000031002120312-3122221203333311"></a>

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

<a id="canonical-2310121021311220-3201001221010032-0302132311103112-1233101220113023-2232020101220020-2033223200121312-2232302230013312-0133213212303302"></a>

## Direct properties — gift_card_validation / 210202330110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231223122330330-0210020331131231-3033113331123123-1121123223133230-3000133303222103-3331302011212010-1031201032010002-1330210012232310"></a>

## Next pages — gift_card_validation / 210202330110 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101213013303303-0023233332032233-3130322111300110-3103300023100220-2013130321113020-0321310231132033-0233223031122303-3301103122211000"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart — shop_add_to_cart / 221013331130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1312213101332122-3222210331132012-2233003203030201-2021000320011213-2321010022312103-0120032001201033-1203210212320020-3002200233222111"></a>

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

<a id="canonical-0221022310100121-2231112230031300-1333002231303232-2030203011013320-0010033222301133-3121331012301231-0233311331110003-1212312031333300"></a>

## Direct properties — shop_add_to_cart / 221013331130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333231303100012-1002001120200122-3231301310211222-1133302330103310-0100202230103001-0002123102221320-1020313102231002-3110032300310330"></a>

## Next pages — shop_add_to_cart / 221013331130 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333003321120120-3031221232031022-0322021212331232-3312122313211321-0332203212203001-2112222210112123-2031100322211300-1222211203031102"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout — shop_checkout / 131220311212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0331000022302323-0303022031133110-0230123001303113-0120010230233031-3111222330131011-3132200321233022-0122230102102132-2011013333313201"></a>

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

<a id="canonical-1113220222203012-2032311012321313-3002033001133323-0003220010322210-1001223020201330-3012203133002221-0201030023020122-2011021112100000"></a>

## Direct properties — shop_checkout / 131220311212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202320310021231-0121220023213033-1002311321232222-0223110200013123-2332223112021021-2322223001021000-2303302133330333-2313210320331201"></a>

## Next pages — shop_checkout / 131220311212 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300211102322213-1131330213213132-1132130331331232-0003200021103220-0331032031300321-3221200221101103-1312030302022222-0123210213102022"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat — shop_choose_seat / 103233232123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-0302331033201212-2003132210102222-1022222202200131-0131131333000011-0012031000021121-1331331122213203-0233000313121112-2122202222012301"></a>

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

<a id="canonical-3232212313100131-0330332133311032-0300323123300021-3131122102010331-0101323200012100-1203312230103310-1301033231113012-3013022100202221"></a>

## Direct properties — shop_choose_seat / 103233232123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312133201131122-0233322023302103-1003132033020323-0313232123302201-2220311332302210-3212301233333222-2022323000300113-3200113033130223"></a>

## Next pages — shop_choose_seat / 103233232123 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003210103110311-1121033032322123-0330003233222031-1012020221110323-1020100313321031-3100023330311303-3301333233322202-2100133123301010"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission — shop_enter_drawing_submission / 030032221130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-2330200123133311-1301121320203012-2332131301122121-0131301023231102-0121131123313322-3310333100130320-0131331312033301-2333302023002120"></a>

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

<a id="canonical-2022330321231120-2323233212033102-0103230322313133-1130322033333103-0032223001321333-0230203211020323-2132120122003000-0110111000102133"></a>

## Direct properties — shop_enter_drawing_submission / 030032221130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033011000102331-1310330332222311-0232200331010220-0212123023102132-0030032222203111-2212132131221123-1123021020111221-1001032301230330"></a>

## Next pages — shop_enter_drawing_submission / 030032221130 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200313313211120-2010223223223211-2122123013032110-3200020121012131-2321031222210030-1132312010311020-2312000300101330-1023003312210132"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment — shop_make_payment / 030130032023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-2200001333010330-3230121212022011-3111310033211322-3112302333012023-2000133212033001-1001022033000020-1332220100222302-3123122123310321"></a>

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

<a id="canonical-1000220303113303-2311321000131210-2013213020202300-1113322322203210-2221010103000033-3131032111332310-0230301131302031-0303112203321110"></a>

## Direct properties — shop_make_payment / 030130032023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302021130221311-1022013100102031-1322121233320322-1310322113121301-3010200103003200-2300203320232001-1312100010301101-1220322013022113"></a>

## Next pages — shop_make_payment / 030130032023 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230301010121122-0310311101303000-3023101100210322-0211230322231331-3102330211133130-1200223022233121-1200300233220011-1233222232213211"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order — shop_order / 223220020113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-3220302200030232-2021121332323313-3023203333012232-3002110202203002-0231032202030210-1233102223130311-1322101111220301-1102101310110023"></a>

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

<a id="canonical-1103203111222303-1223313333222113-0211023130111231-0032033030000203-1331131203200013-0210112313020021-0000300320201130-0333001110122202"></a>

## Direct properties — shop_order / 223220020113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301110130310221-2322222002200012-2312111311033032-1332300123330130-0012121113331333-1212102010131332-2111230030102203-1320213030223033"></a>

## Next pages — shop_order / 223220020113 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133120330002011-3001100322220110-2101021223230121-0220020323222202-2200002222201212-1022200311003113-0113211212001331-3112021112321301"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry — shop_price_inquiry / 311022121132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-3022011222111023-3223311313102010-0203013300121111-0031101331002313-2322310322020120-1031200030312102-1132310111301323-3213232223201101"></a>

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

<a id="canonical-2320322221010201-0321311321201133-0212210121120002-3100220112330012-1013322201210311-1322322111221231-3333100330111220-0322212021123011"></a>

## Direct properties — shop_price_inquiry / 311022121132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100001032232100-1302031132212110-1321311100013130-3033121010212231-2300133103103211-3232011213300220-1310313012110212-0120023110203222"></a>

## Next pages — shop_price_inquiry / 311022121132 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123310000303130-3201001203000210-3003310133321121-1231210011022020-0301122031232013-0331203031332100-1332332203221300-2031110313231230"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation — shop_promo_code_validation / 331121002322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-1313120301111002-0021323121222320-3101231003312211-3203010312330301-2232002230123310-3120133001323113-2120033323110111-3033010322031221"></a>

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

<a id="canonical-3300120202002203-3031233033203020-0221032203231000-0231103213331010-0332302331332132-1113112101133313-1023311200102201-0331300320103012"></a>

## Direct properties — shop_promo_code_validation / 331121002322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332202002233021-1213111201213011-3002302203213102-0132002202123120-2300112301222221-1211001011301220-0111323111210201-0121223301031012"></a>

## Next pages — shop_promo_code_validation / 331121002322 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313202222211022-1100112232201231-1033130031012213-0221023111100103-1000022322333200-3001002013320312-1222030103013221-1302100132030101"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card — shop_purchase_gift_card / 032130233302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-2103031332302032-3013332112121013-0111102031132311-3012111013331133-0223003223302200-0332020310323131-0232301302023200-3323310312103223"></a>

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

<a id="canonical-0301012210212300-0103320100233330-3213233001123313-0032233002003030-3310031301303200-2113213231213312-1021022332310321-0133321012333111"></a>

## Direct properties — shop_purchase_gift_card / 032130233302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121302003133213-2103103302111203-3332133303022121-0111012103220032-2213033101223232-2221200131001023-2131320021322203-0120132232232330"></a>

## Next pages — shop_purchase_gift_card / 032130233302 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013010223213022-3322002211202212-0023102033313113-3321221122222030-1003223132210110-2132010203303231-0202112210113300-3133032223002112"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity — shop_update_quantity / 012110321002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-011.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-0213230030220332-0323201101101032-0022003131013320-3212023302310203-1013221013110233-3103002101221031-3123321203332113-2130031002120003"></a>

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

<a id="canonical-2111010230210321-1331331130103200-0103321030103112-1131011012031031-1130130231310013-2203101301121023-1222032232220311-0020031232331130"></a>

## Direct properties — shop_update_quantity / 012110321002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233300222010020-3332333020200020-1102130213012323-3301103001303332-1013103002230010-3131203332103303-3010233010311120-2121013131030313"></a>

## Next pages — shop_update_quantity / 012110321002 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-012.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322302222311030-0333330001032211-2203220100330011-1030003103313300-3022213203210000-0221210023033103-2102331231003131-3101213202300223"></a>

## bot_defense.policy.protected_app_endpoints.headers — headers / 302230210013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.headers

<a id="canonical-1332322330232113-1310021101202323-2232101220133003-2001012323220010-1013203013123301-2211230000303313-2233002130010121-0100230333013200"></a>

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

<a id="canonical-3130331023301203-2102013120310000-3323123203001120-3013332333322300-3311020220002232-3123313221012022-2022230211332232-1302312221100301"></a>

## Direct properties — headers / 302230210013 / 3

- [check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-2232200010013313-2232132010233122-2222330110102011-2301103212102220-0202321220000321-1300332333012210-0320232113233301-3200230222101323): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-012.md#canonical-0112231303022003-2112223020302322-1003111201213312-1113312330101330-0321002303113231-3121112233201111-1300330210000003-3002130003020121): complete subsection reference.

<a id="canonical-1213022301322211-0210220110100100-3300300320001003-3203032013322313-2222233321010213-1123113021121010-1013013222013330-0020323310110221"></a>

<a id="canonical-1103322320312123-1022320003122011-0213110121000301-1133130131110212-3122113220233201-2223311312010211-0312332133201032-2332311323021332"></a>

## invert_matcher property — headers / 302230210013 / 4

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

- [item](resources--http_loadbalancer--reference--group-012.md#canonical-0312123300323102-0331001312331110-2013210121020302-1123221333001330-3000213312012320-3231311010020201-3003311330112112-3132300302131111): complete subsection reference.

<a id="canonical-1323302102233101-2030303310312232-1010002221213112-2221212332021330-0321231221113121-3023311120112020-2103003101320210-1011003212130232"></a>

<a id="canonical-1132120122201331-2010323301102130-3203203132021112-1031322110220131-2211100323222130-1011100112231330-1123203130011303-0232323310322021"></a>

## name property — headers / 302230210013 / 5

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

<a id="canonical-0230001310021011-0120232110320030-0200311132212202-0100303312310131-1002112000103332-0222311000132201-0231302132320320-2131200002131131"></a>

## Next pages — headers / 302230210013 / 6

- [bot_defense.policy.protected_app_endpoints.headers.check_not_present](resources--http_loadbalancer--reference--group-012.md#canonical-2232200010013313-2232132010233122-2222330110102011-2301103212102220-0202321220000321-1300332333012210-0320232113233301-3200230222101323)
- [bot_defense.policy.protected_app_endpoints.headers.check_present](resources--http_loadbalancer--reference--group-012.md#canonical-0112231303022003-2112223020302322-1003111201213312-1113312330101330-0321002303113231-3121112233201111-1300330210000003-3002130003020121)
- [bot_defense.policy.protected_app_endpoints.headers.item](resources--http_loadbalancer--reference--group-012.md#canonical-0312123300323102-0331001312331110-2013210121020302-1123221333001330-3000213312012320-3231311010020201-3003311330112112-3132300302131111)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2232200010013313-2232132010233122-2222330110102011-2301103212102220-0202321220000321-1300332333012210-0320232113233301-3200230222101323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300000033122010-3101020311021223-3333002303333121-2022002313230200-2100102101322211-2200033203111012-3203003332112030-1333302113203013"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_not_present — check_not_present / 333331310002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.check_not_present

<a id="canonical-3101122231110210-3011020112333013-0022233132130111-1313033000231121-1320011132020211-1202022030113133-3120023201211331-2221101121101020"></a>

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

<a id="canonical-1302220312133122-2133233133331103-1113302030000022-3030231131030003-0133232300231220-3223001011102301-1233322112320303-3310031331310021"></a>

## Direct properties — check_not_present / 333331310002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020331320231222-1230101202312310-0330000121301120-3310032320321113-1131202213322111-1123133021230110-1233301312012103-3330020012312233"></a>

## Next pages — check_not_present / 333331310002 / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112231303022003-2112223020302322-1003111201213312-1113312330101330-0321002303113231-3121112233201111-1300330210000003-3002130003020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111001302222033-3002231211121320-3122101031332110-3111033103331003-3202300013003203-2213212223231332-3001232321222321-1232022333131022"></a>

## bot_defense.policy.protected_app_endpoints.headers.check_present — check_present / 031122022323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.check_present

<a id="canonical-1031300113112110-0031203033333332-3000310210113332-1000111212001111-2133020120001110-0300123111101013-2110221102201130-3233003021012211"></a>

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

<a id="canonical-0002321202301033-2130333212013203-0133323033311210-1031112313020000-3122133223222100-3003100313111203-3001212232121330-3213223300001112"></a>

## Direct properties — check_present / 031122022323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321303100122311-3003322101022130-1012230210330222-1230020033003331-2133003103322300-1120123013000313-2333313320021221-1333313132110011"></a>

## Next pages — check_present / 031122022323 / 4

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0312123300323102-0331001312331110-2013210121020302-1123221333001330-3000213312012320-3231311010020201-3003311330112112-3132300302131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120103130002001-2003221320103312-1100132023330102-1002022010102102-3010132313323123-3030030100333103-1312001330212011-1113121101211311"></a>

## bot_defense.policy.protected_app_endpoints.headers.item — item / 212100100003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.item

<a id="canonical-2223211010130223-2023130001021023-0131020231302301-3001123222011033-1230303112310113-1221013230033111-1000022010132202-0203202103203223"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123202322230102-1233122312132322-2333301012013320-2202122221021023-1321130312232301-2223220023323003-1320203211013313-2013201212302120"></a>

## Direct properties — item / 212100100003 / 3

<a id="canonical-3213201132331311-0030201120030003-1211102212132012-0000122031103011-0202032132231230-1230332111023032-1232323321003132-1103233131330013"></a>

<a id="canonical-2313300012110323-1212203133201230-1022023302330130-0332033010123203-0322211301330311-3000203322220220-0310123113001022-0120331203211002"></a>

## exact_values property — item / 212100100003 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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

<a id="canonical-3022123031310213-2322001033132013-2023123233312321-2023211321023103-3010323302322122-0213231303102213-3112210030001101-2133111310333020"></a>

<a id="canonical-1301223230121313-0031210112212330-3120311111201213-1312001332200330-3121123303022112-3233031222212311-1320013200322000-2302010001330220"></a>

## regex_values property — item / 212100100003 / 5

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

<a id="canonical-1211132131201023-0012332200320320-3010002211212323-0122331020031222-1012223232321001-2020113313102122-2130032333302322-0011322012230021"></a>

<a id="canonical-0201102330233113-3211013012210033-1331130133210321-2002012110223030-3110013333211012-1301030113221113-3231100302131232-2132010333312123"></a>

## transformers property — item / 212100100003 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

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

<a id="canonical-3231322302013311-2331220311001101-1323230330310330-1121200133223303-2320031012333132-3023012122203123-2021100210131332-1101022323123011"></a>

## Next pages — item / 212100100003 / 7

- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-012.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2023212323122011-1101323330132021-1310011202223312-0131122130110010-0101330133311212-3003232332003321-2322011030222200-1132132132002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332203311210112-1000121212311021-3011221103112133-0132133300123000-2313313301221333-3332122200132003-1322030301333013-0111133311013231"></a>

## bot_defense.policy.protected_app_endpoints.metadata — metadata / 030221310102 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.metadata

<a id="canonical-3113031211122233-3320301002122210-2032001233331312-2112121320301322-1302233003211223-1313013123123130-0122031010203320-0012121211102003"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113131111210033-3123211312111100-0103101020202330-2101322310131022-1123333330000022-0030000220323132-1002331221222331-2301221301010103"></a>

## Direct properties — metadata / 030221310102 / 3

<a id="canonical-0130000112213030-2101112122321200-1333021232231130-2031323022022123-2031203331132130-3220120233031332-3003332031012011-3122333020313230"></a>

<a id="canonical-0032102210220000-2011221123232231-0010111223030222-2213130110111120-3313002301133322-0123131111003001-1131032300110032-0321303202020232"></a>

## description_spec property — metadata / 030221310102 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1012303210131033-3032033202232200-3323010130011123-3301000113310231-0002002212130031-3003333013213212-2122220300213232-0101032102111123"></a>

<a id="canonical-1231303022111233-0321100210302131-3132321301021313-0333212331110022-1102011010130331-3130132101330333-0332311013113013-3202233012300320"></a>

## name property — metadata / 030221310102 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1012213210031012-1121330003212102-1121312212313221-1111112103202201-1033130131223121-1323003331213233-3331212100122112-2112333102321003"></a>

## Next pages — metadata / 030221310102 / 6

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200332103111311-3331103321320021-3300022321112003-2220210032301110-3121012100030020-3110320100033110-2213130032312123-2312232022333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121321222212220-1221120221110000-1132130201110133-3300103122312210-2031202220031132-2331322100300330-1101201012022312-3100132301303021"></a>

## bot_defense.policy.protected_app_endpoints.mitigate_good_bots — mitigate_good_bots / 200133320013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mitigate_good_bots

<a id="canonical-3000110330200220-3123311030311311-2210101121320130-2303033201301123-3102003232231023-0021130113123023-0030301310002003-3313332113010023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for mitigate good bots.

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
mitigate_good_bots = {}
```

<a id="canonical-3232210321212231-3122332020010121-1023312121130131-0231301003300203-1132321020331313-1230301101022123-2031020312003030-3003112332203320"></a>

## Direct properties — mitigate_good_bots / 200133320013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300330012320112-2322101323232122-3000201201332323-0032332130330333-2101303010232301-0233321002201120-1131013333203210-0201112121101220"></a>

## Next pages — mitigate_good_bots / 200133320013 / 4

- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131210313101323-2211030320222130-0320202312010323-2323231133203200-1300330102021213-0121201102012133-1222333011202030-3101002202120203"></a>

## bot_defense.policy.protected_app_endpoints.mitigation — mitigation / 120032122321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="canonical-2130232330320031-2112320310100230-2002110023010113-3012223112122002-0232132120001311-2321312303132311-0133011010232021-3233120200321112"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "flag"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("flag",
    "redirect")}
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
  "x-ves-oneof-field-action_type": "[\"block\",\"flag\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321212100312311-0211102213113301-3322101132113033-1001221111022022-3001221332103302-0211213130223100-3201213101311032-2123313113222003"></a>

## Direct properties — mitigation / 120032122321 / 3

- [block](resources--http_loadbalancer--reference--group-012.md#canonical-1110113102130312-1122032103010331-0313201330221213-0300002321113102-0202203002321220-2210302012301033-2333322300101331-0132303002311230): complete subsection reference.

- [flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330): complete subsection reference.

- [redirect](resources--http_loadbalancer--reference--group-013.md#canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211): complete subsection reference.

<a id="canonical-1320303321213023-2211000223210031-1133310212110200-3202121112023122-1232332211311302-2020301232220300-2102133300130203-1222020323310202"></a>

## Next pages — mitigation / 120032122321 / 4

- [bot_defense.policy.protected_app_endpoints.mitigation.block](resources--http_loadbalancer--reference--group-012.md#canonical-1110113102130312-1122032103010331-0313201330221213-0300002321113102-0202203002321220-2210302012301033-2333322300101331-0132303002311230)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-013.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](resources--http_loadbalancer--reference--group-013.md#canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1110113102130312-1122032103010331-0313201330221213-0300002321113102-0202203002321220-2210302012301033-2333322300101331-0132303002311230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033011121112202-1021201201102012-3133302332232131-2320323301011213-3103031120113130-0201021313133130-2103123221213032-0023022230230112"></a>

## bot_defense.policy.protected_app_endpoints.mitigation.block — block / 121021311131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-011.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-011.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-011.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-012.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- bot_defense.policy.protected_app_endpoints.mitigation.block

<a id="canonical-0131100203121021-0022130333113212-0112133212022103-2223120030122303-2130123200011333-1330201000020111-0330121213320110-3031310002230330"></a>

Type: `"object"`. single nested block, Optional.

Block request and respond with custom content.

Receipt-pinned upstream constraints:

```json
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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202031312202211-3120003200132013-2211213323011101-3002022322103323-2311331301323311-2102303220011001-3020213012311212-0120121121132303"></a>

## Direct properties — block / 121021311131 / 3

<a id="canonical-2003221232201011-3021101332002012-2131301232202201-2030100323120032-3010213323211100-2231133301230323-0311132303101231-2233021101031223"></a>

<a id="canonical-2233202102332113-2131223333200131-3131131023320232-1223221223311222-1212003131200130-3333313030231320-3220130032323103-2012231101110133"></a>

## body property — block / 121021311131 / 4

Type: `"string"`. Optional.

Custom body message is of type URI\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in base64 format.

Upstream description:

Custom body message is of type URI\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Your request was blocked" or it can be HTML paragraph or a
body string encoded as base64 string E.g. "&lt;p&gt; Your request was blocked &lt;/p&gt;". base64
encoded string for this HTML is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0033032010332233-2022302313032322-1011303000312111-0300210323121211-3331313023301003-2133332332130222-3001303333203013-2213112102310332"></a>
