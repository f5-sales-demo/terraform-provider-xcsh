---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2231202112221202-0020030313203302-2320202212313031-2320101122301023-1220213203102220-0112111023210301-0003011112320030-2023110210213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-010.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.apply

<a id="canonical-0303023201021203-1213331020131330-2121232130223013-0023102213212322-1232323320320101-1003212211102112-2333233020020303-2110111300023020"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
apply = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001130111333030-2333131211300221-2032130032132001-3121201112013113-2101230103222033-1021110001032013-1033323101021201-2220231211213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](resources--http_loadbalancer--reference--group-010.md#canonical-2313330222120110-2222031320011333-2210303331220110-1210112000113221-0120310301320110-3320203021001300-3323133332212333-0032021010221112)
- bot_defense.policy.protected_app_endpoints.flow_label.financial_services.money_transfer

<a id="canonical-0000000033330222-2133212001221133-3120230302002120-0112000331223313-0220013330232103-0322031210302120-2232222330001331-1021230033132212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for money transfer.

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

Terraform syntax:

```terraform
money_transfer = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.flight

<a id="canonical-0112222123130203-0300003200212330-0222031122230301-0303100303302031-3032310120131132-2220033133001120-2002311232131032-3303122222032103"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Flight Category. Bot Defense Flow Label Flight Category.

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

<a id="canonical-1000131123232020-3331200202002101-1130022221023032-0202201221021201-3012120000020113-0321220233001021-0130012011003102-2320212322220222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.flight`

- [checkin](resources--http_loadbalancer--reference--group-011.md#canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200): complete subsection reference.

<a id="canonical-0210132233133331-1021021322100330-2301003323100213-1121101221131333-3313111333332310-1133112102232000-3023222112123000-1200221211133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](resources--http_loadbalancer--reference--group-011.md#canonical-2131020221010102-2012122000002123-3132321003203233-1323312223200203-3033213210222121-1123012101220330-2201203221233110-3221101122221132)
- bot_defense.policy.protected_app_endpoints.flow_label.flight.checkin

<a id="canonical-0022211211033231-3330201033232230-2302333012201023-1310113121121123-3122322102130112-1210331131233323-0201011221210222-0321231200003220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
checkin {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management

<a id="canonical-1123331100331221-3233022320303222-3010002023010313-0333302033211002-2233022023130030-0131222133203300-3331212323101213-3222122011321301"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Profile Management Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1003213132023302-1013011011302233-2112233000231303-0232021301321330-2133011033130113-0223233201222310-0313310213113032-3132301330331311"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.profile_management`

- [create](resources--http_loadbalancer--reference--group-011.md#canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131): complete subsection reference.

- [update](resources--http_loadbalancer--reference--group-011.md#canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201): complete subsection reference.

- [view](resources--http_loadbalancer--reference--group-011.md#canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203): complete subsection reference.

<a id="canonical-3323111010311303-1101013331021322-1010020300123101-0130223003112113-1122033030030122-2122331021203303-0111200003302301-1320122131123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-011.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.create

<a id="canonical-3112132212320123-0132131233323130-3000333221101133-2230201003300011-0203300023001230-0100233203213230-2221202331202232-1223021133010323"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
create = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202231202011131-0131123002023212-3121312212212112-0331220201122330-0012020310131032-0312212012323003-2211210021203000-3130212230230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-011.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.update

<a id="canonical-2211033010302312-2231010230330030-1012013003110302-1232021101132131-2121300210021132-2000031213012223-3022322220003310-3001120203003110"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
update = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011021320013023-1330220123302101-0300123023313103-0203003310131113-2233011003301132-1030101221112202-0122001021033220-2031113100333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--http_loadbalancer--reference--group-011.md#canonical-0333133220303221-0122112213221002-3333302300203020-0120323102323011-2030032211312313-3031123103012122-3121230231000011-1300121020012021)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-3101203232122311-2011111033011310-1013100001021311-0231131032121003-3221320013200030-2323122132311302-0222110233031330-3101230201111132"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
view = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-0230313123323323-0220201122122023-0301030122333211-0020003122122023-0210023231211332-3122202303211332-0132113310210201-1012011012210022"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2212312203033233-0113311320023313-1211232222033201-3101020100333320-2230330312330213-3002320100020013-0210103301010022-0201330203031131"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.search`

- [flight_search](resources--http_loadbalancer--reference--group-011.md#canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032): complete subsection reference.

- [product_search](resources--http_loadbalancer--reference--group-011.md#canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102): complete subsection reference.

- [reservation_search](resources--http_loadbalancer--reference--group-011.md#canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212): complete subsection reference.

- [room_search](resources--http_loadbalancer--reference--group-011.md#canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133): complete subsection reference.

<a id="canonical-3003121201002221-2313221231103211-3231233123331101-1100323321202131-0232323332013021-0132323311031021-2011010130122323-3322012330322032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-011.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-3020011233310330-0323311030233213-0020003113013100-3330002031312103-0032303312101203-1012002113320330-1103011123221302-0032333222100100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for flight search.

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

Terraform syntax:

```terraform
flight_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122100303123213-3332022223222132-1312221330231002-0111220333110120-1121202101113001-2013012111023133-2323310201221103-3031011122031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.product_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-011.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-1031211310300013-1000320303030231-2202322331203212-1021021131030310-0202331223320030-3003320022333312-3111120232010300-0213110120223330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for product search.

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

Terraform syntax:

```terraform
product_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112330212220123-1012213000130302-1301021120213201-0121230233133211-0122300103023231-3303031011232212-2002120000313013-1310111022311212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-011.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-1010113203001113-3112203002030112-1222020112332232-0003232302223322-3221330112221021-1213302000313331-3202330032001110-0220022103230212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reservation search.

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

Terraform syntax:

```terraform
reservation_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020211203330201-2131011020310020-0022201313320200-1220203320031130-3113023023033332-1331313323211330-0111102033203330-2021332000333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.room_search` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--http_loadbalancer--reference--group-011.md#canonical-0200331320001212-2122333200220102-1210123221101031-2123022331313031-3111303102221230-0211022020221002-0231303103020000-1001323102021121)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-1103211223102301-2002230123113121-2321131010301210-1223233300302213-1130110321131321-1200223331303033-2213333212332130-2232322012100101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for room search.

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

Terraform syntax:

```terraform
room_search = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-2312330022202233-0031211231120002-3010130013100312-2331212132230123-2100132331223010-2331213111201011-1210202211111330-1131320203012003"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1132001231321303-2320210012021323-1222331232123013-3300230013333333-0121123111031302-0211323012312230-2331121330101313-2330221313112232"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards`

- [gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--reference--group-011.md#canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103): complete subsection reference.

- [gift_card_validation](resources--http_loadbalancer--reference--group-011.md#canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223): complete subsection reference.

- [shop_add_to_cart](resources--http_loadbalancer--reference--group-011.md#canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203): complete subsection reference.

- [shop_checkout](resources--http_loadbalancer--reference--group-011.md#canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320): complete subsection reference.

- [shop_choose_seat](resources--http_loadbalancer--reference--group-011.md#canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300): complete subsection reference.

- [shop_enter_drawing_submission](resources--http_loadbalancer--reference--group-011.md#canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301): complete subsection reference.

- [shop_make_payment](resources--http_loadbalancer--reference--group-011.md#canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012): complete subsection reference.

- [shop_order](resources--http_loadbalancer--reference--group-011.md#canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222): complete subsection reference.

- [shop_price_inquiry](resources--http_loadbalancer--reference--group-011.md#canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101): complete subsection reference.

- [shop_promo_code_validation](resources--http_loadbalancer--reference--group-011.md#canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211): complete subsection reference.

- [shop_purchase_gift_card](resources--http_loadbalancer--reference--group-011.md#canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123): complete subsection reference.

- [shop_update_quantity](resources--http_loadbalancer--reference--group-011.md#canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231): complete subsection reference.

<a id="canonical-0130300013310220-0020210202020301-0230222030231321-2311211333230333-3233201001032013-2102032100021022-0320111003322211-3223123122112103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0203201032022321-1101122012013302-1123023201002231-2333310320323031-1131211011333011-1203222223312020-0010021132303133-0023201002032122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card make purchase with gift card.

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

Terraform syntax:

```terraform
gift_card_make_purchase_with_gift_card = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002121030231031-1110201222120212-2033313223123301-2203321303231130-1203031010101310-0323000120310312-0000110221031011-0121123231323223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2302300101313113-0221230010311211-3330323220321021-3211313301233112-3232210001312323-3120120320020131-0000031002120312-3122221203333311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for gift card validation.

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

Terraform syntax:

```terraform
gift_card_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103130300000311-0112233120100213-3021221131011201-3332312023223130-2002120013013213-3021211030110021-0232310321032330-3221320030301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1312213101332122-3222210331132012-2233003203030201-2021000320011213-2321010022312103-0120032001201033-1203210212320020-3002200233222111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop add to cart.

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

Terraform syntax:

```terraform
shop_add_to_cart = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103213100212323-2301300123032122-0321223133212233-2322330120322210-0320330302121320-1231210110211130-1130023320122032-1223020320010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-0331000022302323-0303022031133110-0230123001303113-0120010230233031-3111222330131011-3132200321233022-0122230102102132-2011013333313201"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop checkout.

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

Terraform syntax:

```terraform
shop_checkout = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020000331112012-3111012312333312-3101301120100130-2310303303113203-2022333023031231-0303301313110011-2303111033210033-2202131323233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-0302331033201212-2003132210102222-1022222202200131-0131131333000011-0012031000021121-1331331122213203-0233000313121112-2122202222012301"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop choose seat.

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

Terraform syntax:

```terraform
shop_choose_seat = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320321333211010-1212130220211311-0130203033310000-0213311133200103-3202212033022122-3020112030223112-2013312333212011-2133303120121301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-2330200123133311-1301121320203012-2332131301122121-0131301023231102-0121131123313322-3310333100130320-0131331312033301-2333302023002120"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop enter drawing submission.

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

Terraform syntax:

```terraform
shop_enter_drawing_submission = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331011021232132-2131322033312302-2112202200031313-1323131332312222-0311303002002021-3331221011321313-3310322312333000-3011132122003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-2200001333010330-3230121212022011-3111310033211322-3112302333012023-2000133212033001-1001022033000020-1332220100222302-3123122123310321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop make payment.

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

Terraform syntax:

```terraform
shop_make_payment = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001131230111330-0323012132220320-2013323110323330-0001011033011023-2311201013130311-0223333133233202-2200003000003210-1000111130233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-3220302200030232-2021121332323313-3023203333012232-3002110202203002-0231032202030210-1233102223130311-1322101111220301-1102101310110023"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
shop_order = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322333210220201-1322223102230232-0211002311213100-3012030231213021-2112101333300120-0130022213001213-0133302300032112-0323031320213101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-3022011222111023-3223311313102010-0203013300121111-0031101331002313-2322310322020120-1031200030312102-1132310111301323-3213232223201101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop price inquiry.

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

Terraform syntax:

```terraform
shop_price_inquiry = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331112112202310-3311020330321310-3311212022102213-3000223301303322-3132110201023102-1032020323213021-0222313321001210-1300222203322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-1313120301111002-0021323121222320-3101231003312211-3203010312330301-2232002230123310-3120133001323113-2120033323110111-3033010322031221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop promo code validation.

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

Terraform syntax:

```terraform
shop_promo_code_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032113203003111-0333310323013200-0121113002301223-1320300010011203-1301011233132302-1123302333030023-1011011111320001-2022221001211123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-2103031332302032-3013332112121013-0111102031132311-3012111013331133-0223003223302200-0332020310323131-0232301302023200-3323310312103223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop purchase gift card.

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

Terraform syntax:

```terraform
shop_purchase_gift_card = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130222123200323-3310111201110123-0002021221002203-3233231123133220-1222333220220313-2131322231022030-0013122032023303-2122311210312231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--reference--group-010.md#canonical-0310210030323002-3211300333333221-1112113200201202-3222212102120220-2030332233030223-2113022203230032-2031201020000321-2001031122232002)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--http_loadbalancer--reference--group-011.md#canonical-3210103130303012-3000301112011033-1230030100030231-0102001030322021-2211330012122222-2130003000221110-0031121031233332-3302120303310220)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-0213230030220332-0323201101101032-0022003131013320-3212023302310203-1013221013110233-3103002101221031-3123321203332113-2130031002120003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for shop update quantity.

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

Terraform syntax:

```terraform
shop_update_quantity = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.headers

<a id="canonical-1332322330232113-1310021101202323-2232101220133003-2001012323220010-1013203013123301-2211230000303313-2233002130010121-0100230333013200"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for various HTTP headers that need to match. The criteria for matching each
HTTP header are described in individual HeaderMatcherType instances. The actual HTTP header values
are extracted from the request API as a list of strings for each HTTP header type. Note that all
specified header predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0322302222311030-0333330001032211-2203220100330011-1030003103313300-3022213203210000-0221210023033103-2102331231003131-3101213202300223"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.headers`

- [check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-2232200010013313-2232132010233122-2222330110102011-2301103212102220-0202321220000321-1300332333012210-0320232113233301-3200230222101323): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-011.md#canonical-0112231303022003-2112223020302322-1003111201213312-1113312330101330-0321002303113231-3121112233201111-1300330210000003-3002130003020121): complete subsection reference.

<a id="canonical-1213022301322211-0210220110100100-3300300320001003-3203032013322313-2222233321010213-1123113021121010-1013013222013330-0020323310110221"></a>

<a id="canonical-3130331023301203-2102013120310000-3323123203001120-3013332333322300-3311020220002232-3123313221012022-2022230211332232-1302312221100301"></a>

#### `bot_defense.policy.protected_app_endpoints.headers.invert_matcher` property

Type: `"bool"`. Optional.

Invert Header Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-011.md#canonical-0312123300323102-0331001312331110-2013210121020302-1123221333001330-3000213312012320-3231311010020201-3003311330112112-3132300302131111): complete subsection reference.

<a id="canonical-1323302102233101-2030303310312232-1010002221213112-2221212332021330-0321231221113121-3023311120112020-2103003101320210-1011003212130232"></a>

<a id="canonical-1103322320312123-1022320003122011-0213110121000301-1133130131110212-3122113220233201-2223311312010211-0312332133201032-2332311323021332"></a>

#### `bot_defense.policy.protected_app_endpoints.headers.name` property

Type: `"string"`. Optional.

Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2232200010013313-2232132010233122-2222330110102011-2301103212102220-0202321220000321-1300332333012210-0320232113233301-3200230222101323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-011.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.check_not_present

<a id="canonical-3101122231110210-3011020112333013-0022233132130111-1313033000231121-1320011132020211-1202022030113133-3120023201211331-2221101121101020"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112231303022003-2112223020302322-1003111201213312-1113312330101330-0321002303113231-3121112233201111-1300330210000003-3002130003020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-011.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.check_present

<a id="canonical-1031300113112110-0031203033333332-3000310210113332-1000111212001111-2133020120001110-0300123111101013-2110221102201130-3233003021012211"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312123300323102-0331001312331110-2013210121020302-1123221333001330-3000213312012320-3231311010020201-3003311330112112-3132300302131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.headers](resources--http_loadbalancer--reference--group-011.md#canonical-3011003033310101-0122020002121002-3031202212302210-0021210121111322-1202003203203333-3130233012020232-2110123113302102-1102220303022203)
- bot_defense.policy.protected_app_endpoints.headers.item

<a id="canonical-2223211010130223-2023130001021023-0131020231302301-3001123222011033-1230303112310113-1221013230033111-1000022010132202-0203202103203223"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1120103130002001-2003221320103312-1100132023330102-1002022010102102-3010132313323123-3030030100333103-1312001330212011-1113121101211311"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.headers.item`

<a id="canonical-3213201132331311-0030201120030003-1211102212132012-0000122031103011-0202032132231230-1230332111023032-1232323321003132-1103233131330013"></a>

#### `bot_defense.policy.protected_app_endpoints.headers.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3123202322230102-1233122312132322-2333301012013320-2202122221021023-1321130312232301-2223220023323003-1320203211013313-2013201212302120"></a>

#### `bot_defense.policy.protected_app_endpoints.headers.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2313300012110323-1212203133201230-1022023302330130-0332033010123203-0322211301330311-3000203322220220-0310123113001022-0120331203211002"></a>

#### `bot_defense.policy.protected_app_endpoints.headers.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2023212323122011-1101323330132021-1310011202223312-0131122130110010-0101330133311212-3003232332003321-2322011030222200-1132132132002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.metadata

<a id="canonical-3113031211122233-3320301002122210-2032001233331312-2112121320301322-1302233003211223-1313013123123130-0122031010203320-0012121211102003"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0332203311210112-1000121212311021-3011221103112133-0132133300123000-2313313301221333-3332122200132003-1322030301333013-0111133311013231"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.metadata`

<a id="canonical-0130000112213030-2101112122321200-1333021232231130-2031323022022123-2031203331132130-3220120233031332-3003332031012011-3122333020313230"></a>

#### `bot_defense.policy.protected_app_endpoints.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1012303210131033-3032033202232200-3323010130011123-3301000113310231-0002002212130031-3003333013213212-2122220300213232-0101032102111123"></a>

<a id="canonical-0113131111210033-3123211312111100-0103101020202330-2101322310131022-1123333330000022-0030000220323132-1002331221222331-2301221301010103"></a>

#### `bot_defense.policy.protected_app_endpoints.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3200332103111311-3331103321320021-3300022321112003-2220210032301110-3121012100030020-3110320100033110-2213130032312123-2312232022333022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigate_good_bots` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mitigate_good_bots

<a id="canonical-3000110330200220-3123311030311311-2210101121320130-2303033201301123-3102003232231023-0021130113123023-0030301310002003-3313332113010023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for mitigate good bots.

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

Terraform syntax:

```terraform
mitigate_good_bots = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="canonical-2130232330320031-2112320310100230-2002110023010113-3012223112122002-0232132120001311-2321312303132311-0133011010232021-3233120200321112"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1131210313101323-2211030320222130-0320202312010323-2323231133203200-1300330102021213-0121201102012133-1222333011202030-3101002202120203"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation`

- [block](resources--http_loadbalancer--reference--group-011.md#canonical-1110113102130312-1122032103010331-0313201330221213-0300002321113102-0202203002321220-2210302012301033-2333322300101331-0132303002311230): complete subsection reference.

- [flag](resources--http_loadbalancer--reference--group-011.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330): complete subsection reference.

- [redirect](resources--http_loadbalancer--reference--group-011.md#canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211): complete subsection reference.

<a id="canonical-1110113102130312-1122032103010331-0313201330221213-0300002321113102-0202203002321220-2210302012301033-2333322300101331-0132303002311230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-011.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
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

<a id="canonical-1033011121112202-1021201201102012-3133302332232131-2320323301011213-3103031120113130-0201021313133130-2103123221213032-0023022230230112"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.block`

<a id="canonical-2003221232201011-3021101332002012-2131301232202201-2030100323120032-3010213323211100-2231133301230323-0311132303101231-2233021101031223"></a>

#### `bot_defense.policy.protected_app_endpoints.mitigation.block.body` property

Type: `"string"`. Optional.

Custom body message is of type URI\_ref. Currently supported URL schemes is string:///. For
string:/// scheme, message needs to be encoded in base64 format. You can specify this message as
base64 encoded plain text message e.g. "Your request was blocked" or it can be HTML paragraph or a
body string encoded as base64 string E.g. "&lt;p&gt; Your request was blocked &lt;/p&gt;". base64
encoded string for this HTML is "LzxwPiBZb3VyIHJlcXVlc3Qgd2FzIGJsb2NrZWQgPC9wPg=="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3202031312202211-3120003200132013-2211213323011101-3002022322103323-2311331301323311-2102303220011001-3020213012311212-0120121121132303"></a>

#### `bot_defense.policy.protected_app_endpoints.mitigation.block.status` property

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

Additional upstream details:

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Accepted","AlreadyReported","BadGateway","BadRequest","Conflict","Continue","Created","EmptyStatusCode","ExpectationFailed","FailedDependency","Forbidden","Found","GatewayTimeout","Gone","HTTPVersionNotSupported","IMUsed","InsufficientStorage","InternalServerError","LengthRequired","Locked","LoopDetected","MethodNotAllowed","MisdirectedRequest","MovedPermanently","MultiStatus","MultipleChoices","NetworkAuthenticationRequired","NoContent","NonAuthoritativeInformation","NotAcceptable","NotExtended","NotFound","NotImplemented","NotModified","OK","PartialContent","PayloadTooLarge","PaymentRequired","PermanentRedirect","PreconditionFailed","PreconditionRequired","ProxyAuthenticationRequired","RangeNotSatisfiable","RequestHeaderFieldsTooLarge","RequestTimeout","ResetContent","SeeOther","ServiceUnavailable","TemporaryRedirect","TooManyRequests","URITooLong","Unauthorized","UnprocessableEntity","UnsupportedMediaType","UpgradeRequired","UseProxy","VariantAlsoNegotiates"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

<a id="canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-011.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="canonical-1111201203311012-0220020201120300-0210331212113230-3323211121033211-2010003113333112-2331110331202213-1231322332312222-3223200232113003"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
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
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000032111112200-3211003123231332-0011321123310302-3232010111231320-1130003210010012-0231222232030200-2313230233031010-3202121010233132"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.flag`

- [append_headers](resources--http_loadbalancer--reference--group-011.md#canonical-1213312221221213-2031002222203202-2021211223321101-1200032013313110-0202003213321101-2000333232320323-0132231003111201-1022010330023110): complete subsection reference.

- [no_headers](resources--http_loadbalancer--reference--group-011.md#canonical-0211301100002311-2002221301313232-1321120122303213-0100132032331320-2302302022002311-1223322321301211-2101310303113030-3323212202102222): complete subsection reference.

<a id="canonical-1213312221221213-2031002222203202-2021211223321101-1200032013313110-0202003213321101-2000333232320323-0132231003111201-1022010330023110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-011.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-011.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers

<a id="canonical-1320212321122303-0330022202031233-1020221132012032-1011231220200322-2312112210101031-0012201323201212-1320312220022001-3001020322133020"></a>

Type: `"object"`. single nested block, Optional.

Append flag mitigation headers to forwarded request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("auto_type_header_name",
    "inference_header_name")}
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
append_headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211120213012122-3311212332201221-2000220301210012-1223201031300201-1211130103132312-2331230103013031-3233110022212212-3200301121011111"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers`

<a id="canonical-3300112202021111-1211202230111020-1121101112322023-3123320010310310-3012303112301021-2222132132112221-0031233321213002-3331110131122213"></a>

#### `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers.auto_type_header_name` property

Type: `"string"`. Optional.

Automation Type Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2031131233231211-1212311312232120-1232120233001233-3200123312303020-3211132011010232-3122312001011120-2121302012113222-3012233333121010"></a>

<a id="canonical-1221210312131232-1122300320311221-3212011202310102-1310323203213312-1110010020203121-3132012233020120-1211123301130302-2233122232230312"></a>

#### `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers.inference_header_name` property

Type: `"string"`. Optional.

Inference Header Name. A case-insensitive HTTP header name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0211301100002311-2002221301313232-1321120122303213-0100132032331320-2302302022002311-1223322321301211-2101310303113030-3323212202102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-011.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--http_loadbalancer--reference--group-011.md#canonical-2033033123112223-0333320233033031-2232033202003133-3101111012101010-0021300032230233-1301110001123012-3201030303033330-2323132021003330)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers

<a id="canonical-1301201003033201-3301312200102033-0011102313321223-0131201203100221-3102312101222230-1112313020311232-3100113113223101-3032312301302220"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_headers = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.redirect` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--http_loadbalancer--reference--group-011.md#canonical-3222120100010123-1123101021000131-1333001123203112-1130332200123230-1202330130030002-3303230132133010-3203023112332001-2101223032333331)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-2110331300223000-2231230223320130-1102203331301330-2133001311221300-0113033302223202-2202023202312233-1002120002203332-3230113301013022"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213311020230121-2103112011112112-0231111001200101-2322233211001233-1300231102301121-3001131333013232-0033032100321300-1113312322131201"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.redirect`

<a id="canonical-0310212310321303-0122131332003130-1202301212103010-2311003302332332-0020302103232021-3131231320323301-2013101302022000-2300132202222133"></a>

#### `bot_defense.policy.protected_app_endpoints.mitigation.redirect.uri` property

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-0200310013012011-1313310020131023-0203321100301133-2101212302102010-3320231321202232-3031131303020130-0120030303011112-1220331131110320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mobile` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-3313303032331200-3232300021021302-3201202210123022-2123122011112212-0013110302223232-3302012020203332-0332120013032022-0100111110021120"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
mobile = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132020021201300-2213301200112303-1020320300322230-3032200333333311-3311101002332123-1010201023012022-0122220101130023-1320121003331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-2102102331201332-3031310220311310-0210310331230303-2131322131222330-1032213332201311-1031213020310130-1001231331121131-2133110323302003"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0013231102013202-3332003231130313-0122000122233110-0123303101210211-2123113003103232-1100113110221220-3301130310302213-3122231311301032"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.path`

<a id="canonical-0231120322232033-1222330301213110-2303203322232230-1230332311000201-0323312112213231-2220000033022011-1000300330302110-0301203233203302"></a>

#### `bot_defense.policy.protected_app_endpoints.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0210320230130222-2132013203023330-1232033232213320-0020121321232113-2333013033303112-2303010110231122-2332003300132203-0013223303030010"></a>

<a id="canonical-0331132321203223-1100303132111010-1321313330131102-0313002310222301-0203132230010021-2222211332030223-1213011230021130-1031331121210212"></a>

#### `bot_defense.policy.protected_app_endpoints.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2000323232003103-1303122001032101-2110100011012123-0110010221202122-2211023002023201-2103332213031303-0131220101011021-0301133210133132"></a>

<a id="canonical-0121103303012322-3331202020000022-2320012000220103-2001132211022313-1003030101123200-0010332020331030-1031002102132132-1323130000103111"></a>

#### `bot_defense.policy.protected_app_endpoints.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-0011121301220020-0322113321302121-2233221320233011-0033221120220332-3200313111112303-3102132331232110-0200010100330023-1123201203103132"></a>

Type: `"object"`. list nested block, Optional.

A list of predicates for all query parameters that need to be matched. The criteria for matching
each query parameter are described in individual instances of QueryParameterMatcherType. The actual
query parameter values are extracted from the request API as a list of strings for each query
parameter name. Note that all specified query parameter predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130130003012032-0000203233102331-3233320322103221-3231302301311320-1133302011012023-3002320211321213-1101330220001031-0321301233003211"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.query_params`

- [check_not_present](resources--http_loadbalancer--reference--group-011.md#canonical-2101332322221030-2100012222013201-1113313132003002-1331113222301200-1012332201201212-2130021021102202-1032321332312322-3012021031112223): complete subsection reference.

- [check_present](resources--http_loadbalancer--reference--group-011.md#canonical-2303010213110121-0113212220013123-0130010231301113-0132322301333233-0212030200331313-1212301133123000-2122013231332132-0020202330230001): complete subsection reference.

<a id="canonical-1220033123103110-3320010303031313-1201321012312322-3111330202112303-2032301102232211-2330222120011230-0121110222123322-2210310003213122"></a>

<a id="canonical-2002003133303221-0201223012132200-3111232310132013-3330002221023123-1133020220321313-1123323310002020-2221310232203321-1133233332121021"></a>

#### `bot_defense.policy.protected_app_endpoints.query_params.invert_matcher` property

Type: `"bool"`. Optional.

Invert Query Parameter Matcher. Invert the match result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](resources--http_loadbalancer--reference--group-011.md#canonical-3013113233301011-1230322322233002-2033320211013202-3320003010033202-3021232200131000-3310210332031313-1121113230030033-3130031023300023): complete subsection reference.

<a id="canonical-2122003301100110-3013202221322133-2331332032310320-0131221031001200-2323101322232101-3100211031011111-2313333112320002-1332021332223100"></a>

<a id="canonical-2001330231102012-3031113332330230-2010220023100111-3332212223302310-3301001201211211-2111030011200332-3030322211231323-2112332300033003"></a>

#### `bot_defense.policy.protected_app_endpoints.query_params.key` property

Type: `"string"`. Optional.

A case-sensitive HTTP query parameter name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2101332322221030-2100012222013201-1113313132003002-1331113222301200-1012332201201212-2130021021102202-1032321332312322-3012021031112223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-011.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-1132101020022110-1121013301110201-2322322022132001-2013011301323312-3021320131203101-1030010003103333-2211003101311102-1333133220311011"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303010213110121-0113212220013123-0130010231301113-0132322301333233-0212030200331313-1212301133123000-2122013231332132-0020202330230001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-011.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-2112033302211131-2212211113311013-3320101223033202-1013101013030232-0311320331130311-1230223100000300-1123211132132332-1230013200103313"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
check_present = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013113233301011-1230322322233002-2033320211013202-3320003010033202-3021232200131000-3310210332031313-1121113230030033-3130031023300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--http_loadbalancer--reference--group-011.md#canonical-1013232023221202-0101200322012222-1110023122333121-1203113113021223-0233213112333213-0331320110123211-3111131201203220-3313320112231132)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-0013301000320112-3131203021212303-2221033132301211-0033331323232201-2101033310100213-2033201130203312-3101100030323002-2101013213023223"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2012210020030130-1201303100130022-1013132132113310-1031001002203110-2030010321310331-0001101231210311-2211333212022310-2123011302221222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.query_params.item`

<a id="canonical-1012120213310332-1220000301320103-0103322321033031-0322213211130003-0303102132333213-0210202212210301-1302012103120231-1013003200002101"></a>

#### `bot_defense.policy.protected_app_endpoints.query_params.item.exact_values` property

Type: `["list", "string"]`. Optional.

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0100211110013221-3001123030100203-0200323110113012-0220302300030132-1132133210013230-2001132010103003-3111210213122301-1330321311231231"></a>

<a id="canonical-1123021013211021-2112331030022212-1202200003023002-2231032100302112-3002012010122111-0322011012102102-3113023110113123-1103331000120030"></a>

#### `bot_defense.policy.protected_app_endpoints.query_params.item.regex_values` property

Type: `["list", "string"]`. Optional.

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0130120321022212-1022230333231321-1330022213030110-0222202231001111-0330231320213320-2300203122202230-3021131232031313-0111033202133220"></a>

<a id="canonical-2110102110322101-0331032021012020-3312010132203131-1023211113300111-1010112200201013-2320230332312220-2301033023003220-2200211100110123"></a>

#### `bot_defense.policy.protected_app_endpoints.query_params.item.transformers` property

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1203210321222120-0033120203013313-2131212022012032-1233321301313330-3120311213033232-3302302000221200-3023302011011311-0232031130132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.undefined_flow_label` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-3231333310002230-2111320230302220-3003230031133120-0233301332231012-0203123211201132-0001312013032023-2123110201223011-3303230001201011"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
undefined_flow_label = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130020123113200-0230312201110222-1333301202301130-3302211302301012-2231221131222131-3321323201300333-0101100233311213-2020002023103320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.web` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-3332033113333131-0102010031313110-2003301203222232-0213332102233133-0303201321103330-2013200333310310-3223301221322331-0210112122320232"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
web = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101313121221202-1301032013003023-1220301222330001-2201122130031032-2313210123222031-3023312033232202-2101213232211113-0202330221311302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.web_mobile` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030)
- [bot_defense.policy](resources--http_loadbalancer--reference--group-010.md#canonical-3003102101201221-2220120002312132-1311311200202221-1222002031230111-2122111021030103-3111101330033221-2203020313112231-1312313122001013)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--reference--group-010.md#canonical-1300221231230010-3211322010122201-1301210202313022-3223001103000203-0000010003012111-2230313222320101-0013122332310301-0322233322113021)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-2310331132311233-3121223003002301-1030010212011311-3321313131330120-2333031302301020-0331101012213133-2331223011010231-1131310202231202"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile traffic type. Web and Mobile traffic type.

Receipt-pinned upstream constraints:

```json
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
web_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300120032121133-1302303030010101-1032323031022330-2132322010201131-3000130120331321-3313303323111023-3320322132003301-2201123001302011"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.web_mobile`

<a id="canonical-3002103320313313-3000321033103112-3132331002323023-0122013332133113-0213011211331201-0313313131133003-2001332333023103-0013320113110220"></a>

#### `bot_defense.policy.protected_app_endpoints.web_mobile.mobile_identifier` property

Type: `"string"`. Optional.

\[Enum: HEADERS\] Mobile identifier type - HEADERS: Headers Headers. The only possible value is
\`HEADERS\`. Defaults to \`HEADERS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["HEADERS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("HEADERS"),
}
```

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

<a id="canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- bot_defense_advanced_protection

<a id="canonical-1331023323322330-0120310010111332-1123202023223302-3231330202031222-2112233201102033-2111320023203321-1120030130231113-3212030313203022"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Advanced Protection - replaces BotDefenseAdvancedType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("both_web_and_mobile",
    "mobile_only"),
  validators.ConflictingObjectAttributes("both_web_and_mobile",
    "web_only"),
  validators.ConflictingObjectAttributes("mobile_only",
    "web_only")}
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
  "x-ves-oneof-field-client_type_choice": "[\"both_web_and_mobile\",\"mobile_only\",\"web_only\"]"
}
```

Terraform syntax:

```terraform
bot_defense_advanced_protection {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223121030001112-0102000002332110-0102101031231100-2130122203112003-0013020002312232-0323333130123133-1330100202131000-3212231222111001"></a>

### Direct properties for `bot_defense_advanced_protection`

- [both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233): complete subsection reference.

- [mobile_only](resources--http_loadbalancer--reference--group-012.md#canonical-1313221110001232-1000111131110001-2120333202122021-0101332112201311-3303230010022023-1222103103001210-1000210111303323-2003031202312033): complete subsection reference.

- [web_only](resources--http_loadbalancer--reference--group-012.md#canonical-2232121232322332-1210220123031111-3332230323110120-0021301322100121-1101230230111002-2201001232201011-2100002232021012-0213233110120212): complete subsection reference.

<a id="canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- bot_defense_advanced_protection.both_web_and_mobile

<a id="canonical-0011111001010222-3031000001230011-0331010111323231-1132312320133201-3200030003211130-2133003121022102-2313033122133020-1031222202012023"></a>

Type: `"object"`. single nested block, Optional.

Both Web &amp; Mobile. Both Web and Mobile configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("js_insert_all_pages",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("js_insert_all_pages_except",
    "js_insertion_rules")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
both_web_and_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103232201111222-2222132213210221-1121200303013312-1000303210311022-0131102000020012-0001221301010133-1003012011200211-3312032013113213"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile`

- [disable_js_insert](resources--http_loadbalancer--reference--group-011.md#canonical-0313310300303030-0332030023231201-3100303131122331-2231221230211102-2222331033131311-2022032021031222-3233003002000020-1323231023031312): complete subsection reference.

- [disable_mobile_sdk](resources--http_loadbalancer--reference--group-011.md#canonical-1003101001321320-0233032302330132-0311303011302022-0111210200301000-0311213323331133-0020211303331000-2001332210101131-1033222210001110): complete subsection reference.

- [js_insert_all_pages](resources--http_loadbalancer--reference--group-011.md#canonical-0120102310320113-3323113113120310-3311231020303300-0030331100311200-1230011032000020-2011301210131032-0100213103123202-0022322122133201): complete subsection reference.

- [js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313): complete subsection reference.

- [js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110): complete subsection reference.

- [mobile](resources--http_loadbalancer--reference--group-012.md#canonical-3113113321230312-1011132221021022-2222200122232321-1121313322131023-3010032012223210-0112301112312333-2003213221232012-3022332221010201): complete subsection reference.

- [mobile_sdk_config](resources--http_loadbalancer--reference--group-012.md#canonical-0021022102130333-0233002330221112-1311012030002331-0321313122100100-2122112223222023-2200212300000310-2001022021132003-1100033212221121): complete subsection reference.

- [web](resources--http_loadbalancer--reference--group-012.md#canonical-3200000311231133-3323200203131332-0220323231120332-1111202111132333-0131021232301123-1131022130213101-1010110003233102-3213003220120002): complete subsection reference.

<a id="canonical-0313310300303030-0332030023231201-3100303131122331-2231221230211102-2222331033131311-2022032021031222-3233003002000020-1323231023031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.disable_js_insert

<a id="canonical-2100230222030030-1330131300220203-1033122213122100-1013102230222211-1222231002320301-1130202231323011-3110112100302330-0102121302130333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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

Terraform syntax:

```terraform
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003101001321320-0233032302330132-0311303011302022-0111210200301000-0311213323331133-0020211303331000-2001332210101131-1033222210001110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.disable_mobile_sdk

<a id="canonical-0333311120232000-0331332212030002-0001102000000320-2030121210010211-1303133002201220-3132213223310311-3233302121102000-0330213212200200"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_mobile_sdk = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120102310320113-3323113113120310-3311231020303300-0030331100311200-1230011032000020-2011301210131032-0100213103123202-0022322122133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages

<a id="canonical-3022023332201123-3011011102310101-1311131033002301-1020222013113302-3002223002300020-3230033013113330-3010132313103130-2313102230200101"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132333022132102-0121023310233223-0100000211003201-2220001121111033-2202101310030220-1110132021330303-1223031123133121-1221320332332312"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages`

<a id="canonical-2000131323110012-2202131212011000-2013301230200312-2101130213221101-2031301320211333-2001021133013131-3111112002313110-2130123132122012"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

<a id="canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except

<a id="canonical-2310302220102221-0221212010022122-0211302010200230-1313330303002323-2031131010322012-0023000300002123-3011103212033213-1222231010300333"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages with the exceptions.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232322223132312-1232233102103013-1313003030310012-2330121212231101-3031230201230230-1320030232311000-2112020013103232-1013200022311233"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except`

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123): complete subsection reference.

<a id="canonical-0233300213210101-3303221222222130-2300103032013201-1212332313101021-3100301320331012-3303102011003033-1030213311020002-0000131003331101"></a>

<a id="canonical-3311300203301021-3101113100002312-1303301021210103-2333000221011031-0033223111031023-0022031001312121-1100203331301230-1022113102330101"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

<a id="canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list

<a id="canonical-3300211223123133-2310330102320301-0230213223031313-0231223332232332-1110130231232001-3200113221330100-2131120033011311-1311023200202323"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032100113011311-0103001230210322-0333232301123112-3221033132333032-1112100202330202-3331023210230010-2013111113130320-0211320211200123"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-3331232331321231-1221000220032021-1303320301323033-3103230220000310-1111233003133033-3331021131033013-3202030001330303-1232100322202313): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-1002310310313002-1313312000211033-1133131121233023-3002320023122312-3221110013120120-0110120211231033-3211313223021322-1300100131321201): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-3310223023102312-0113032103112001-3230233230132310-2323022123331023-2220120001200003-1323213303100030-0011322010230312-2020101203300032): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-3223123302031303-2121331213002022-3030100231303021-0111222033122202-3100231200220022-0023302321101222-1031231113110012-0202121330000121): complete subsection reference.

<a id="canonical-3331232331321231-1221000220032021-1303320301323033-3103230220000310-1111233003133033-3331021131033013-3202030001330303-1232100322202313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1123102122021033-3001221033221033-3110113111333001-0032303023121021-1320203312302300-0033110010231230-1332011030210002-0200333011130201"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002310310313002-1313312000211033-1133131121233023-3002320023122312-3221110013120120-0110120211231033-3211313223021322-1300100131321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-1302200000121101-1203001023210313-0213232303132101-2132032223220212-3122212012032132-1331311121002013-3010110200333322-0003010002112130"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122022232322211-1223023113012123-0110232323011030-2120231210011223-1130311001202002-1021100022221013-1120302020330130-0001102300031333"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-1310110101011331-2331001133333232-3100303112232130-1102302320000132-2212000012032010-1310100101210000-3002022232233210-3020202003123031"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0021021201131212-3221220211120113-2211210120223221-0013031202301311-3301013011303000-3000130310200001-0323112321110203-0312112020020100"></a>

<a id="canonical-2301131030223031-1323231223101011-2032333102222201-1332322121310011-1312313131101321-0322012321012122-0121101112003303-1213202013312010"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2230121301003100-0033230201321232-1130131202133311-2123312021122010-3032230203033202-1322301103123310-2002301200031222-3330031112221032"></a>

<a id="canonical-0221201101303233-1313011201331302-3302200300112032-0320011333231033-3322122321100203-2212300122131331-3222133001200200-3323200131012330"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3310223023102312-0113032103112001-3230233230132310-2323022123331023-2220120001200003-1323213303100030-0011322010230312-2020101203300032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-3012113233000222-1222201100001220-0210201103021012-3022332000210232-1112112310222122-3110132332022233-2233023101101010-0023233212000032"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1033021121301201-1121201322223330-2201133021123231-3213003222021123-2011110332323120-3011112313320000-2103101002001201-2033222323130321"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-2101022300032300-0223133002123111-2310332012322230-1103230033222310-0320001320332110-0322213100112003-0322322232031211-2132222111200332"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3031102032131232-2031212303222133-0302001121310201-1323022110101132-3022222100102232-3311200322002200-1010133301312032-1333022211321200"></a>

<a id="canonical-1320311231123031-1300113212332211-2302222311232331-2202233202031301-1013101303300200-1010111000313000-1230000032230123-3013123200302023"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3223123302031303-2121331213002022-3030100231303021-0111222033122202-3100231200220022-0023302321101222-1031231113110012-0202121330000121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except](resources--http_loadbalancer--reference--group-011.md#canonical-1031323022203230-0123321130003310-1202223011330031-2310302130000322-1132331120311002-2212202003121322-0222122130103320-0201023033121313)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-0212032121300310-3132132113333323-3030323300203320-0132121233333232-3021301200301022-0222030121231331-3113303020303020-1031130323033123)
- bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path

<a id="canonical-0213232031122322-1002003110020101-0030102311330213-2030133122212313-0302031301321103-3003310010220112-2032001302310330-2032012003113100"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001121231303101-2113122131221031-2133021131122001-1212031021232102-0123010112212101-2122003103131103-1111023330001022-0332302130310000"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-1231013221223200-0331220022103123-0003213000212211-3121302102022000-2232303203012011-3002101033220101-2323312021023330-0311122210013102"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3000210323120332-0033232132001100-1032012103032222-1331231002131011-3010011233200011-2113231212002320-2231000020332221-3122031301120312"></a>

<a id="canonical-2302031222323301-3300302210203311-2010331312302330-3032331020303010-1312303102233323-3230102010323213-0132111322133331-0020200132123221"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1212123302030320-1313003303131300-0321132021220331-1233222200332020-2223203311223132-0002131022331322-3203311323203000-3031112010103101"></a>

<a id="canonical-0310312121132233-0213032222133021-1203303013221010-1131222130102232-1313302110100301-1102323031301221-3022211323010311-2233133110321331"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insert_all_pages_except.exclude_list.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules

<a id="canonical-1322323020010130-0212021132010313-0203031102023321-1311322102122130-0233010103232221-2332123332312232-0001122312012112-0001012211033233"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101331131000210-0222213203330033-1102313133001301-2123231022223233-0202011110222112-0200322031210223-3212322000023011-0221233223132220"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules`

- [exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233): complete subsection reference.

- [rules](resources--http_loadbalancer--reference--group-012.md#canonical-3312322233021011-1030322220233213-0332312130021312-0210222001301223-0223223002100023-2303333000003330-0212101011221123-1322200333120130): complete subsection reference.

<a id="canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list

<a id="canonical-1202330301101212-1220320202301032-0002330121110100-2102333200001021-0112231103230013-0021031031011312-1320023100130202-3032011323201211"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101303033202122-3323020300211022-3031332221011301-0330101222310102-0021013013120033-2031003101133223-3130303313011301-3210321200300012"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list`

- [any_domain](resources--http_loadbalancer--reference--group-011.md#canonical-3203220101033033-1221321022123231-2123112331212230-2200032312313001-1212020211212032-3101123211021002-0033313231120103-3330310313322023): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-011.md#canonical-2210202110212101-2031110200321121-2223102312220223-3120230131303210-3100001322021321-0230003233113232-3222101210322133-0210211332231000): complete subsection reference.

- [metadata](resources--http_loadbalancer--reference--group-011.md#canonical-0133330222010013-0230302203212302-2233003031002313-3222111301010011-0232212012030221-2113101311122233-0031332232100011-2222300211322211): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-011.md#canonical-2301312002033012-2233003003110011-0110200300122230-1301130332320001-2012333232213321-1010131331333022-1213332320321302-0333133232103312): complete subsection reference.

<a id="canonical-3203220101033033-1221321022123231-2123112331212230-2200032312313001-1212020211212032-3101123211021002-0033313231120103-3330310313322023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3210012210212101-0300212232101021-0301231121200321-2233101220330020-1300003212021322-1100113321211301-2032221130122320-0101022003332032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210202110212101-2031110200321121-2223102312220223-3120230131303210-3100001322021321-0230003233113232-3222101210322133-0210211332231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain

<a id="canonical-1330333303213002-3223022111323322-0321031321331232-1321310302003200-1303333113210001-1232013001012302-2222021022132221-3311300100321330"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000300113110302-0021201201321112-0113302132010210-1010031012222032-3013320303131003-3313002301110033-2320012333320310-2311201312332101"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain`

<a id="canonical-0121030313110330-1020002122300300-1222012211020323-2221102203222322-1221021111332210-2213030022013223-0321113011010000-2012031313221032"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3313211103312222-3010303131230033-2030112132323330-0223303333010223-3213112213220310-2212010133322210-1003321200330121-0100002020003112"></a>

<a id="canonical-3133032032201231-0210222021223303-2331112323001202-0010030133100332-3322330132201123-3300322213211223-3100312223331032-2003002223121020"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0201300113101211-2321212333223131-3112231100120133-2110331231332233-3101010001122020-3100231032303213-1021323001033031-2032231002212230"></a>

<a id="canonical-0233033200110333-1100221213020131-0211003023131031-3032032311302121-1010112023030010-0102200002331001-0220300300031312-3012300233233211"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.domain.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0133330222010013-0230302203212302-2233003031002313-3222111301010011-0232212012030221-2113101311122233-0031332232100011-2222300211322211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata

<a id="canonical-1021222212202100-1022100213010001-1110013301120122-2211132030303331-2120300020112230-2112322113131123-3120213313132123-3321211112011300"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1301301130222221-0031120032311233-0023001102033113-0101001003013331-3320131301002200-0331331221031012-1302122010312203-0310222322200101"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata`

<a id="canonical-0332010132033220-1033011232203031-2103100330300110-3201312032221001-2131133023002131-3300131132311233-2133231221102112-0131200130233132"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3002023322012021-2002022112232202-3020222231133211-0002021231313020-1001202031301223-3111101333013033-1200230322333103-2113021103022032"></a>

<a id="canonical-2222002122322031-0320133333321110-2322031330111310-1132103213110210-2022322332310010-2322131002020330-0213020102201220-3301111212011312"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2301312002033012-2233003003110011-0110200300122230-1301130332320001-2012333232213321-1010131331333022-1213332320321302-0333133232103312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331)
- [bot_defense_advanced_protection.both_web_and_mobile](resources--http_loadbalancer--reference--group-011.md#canonical-0211323020001103-3101110303333101-0011232000332331-0121032010320213-2100131012133022-1333203221230222-0321002132202220-2233013302001233)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules](resources--http_loadbalancer--reference--group-011.md#canonical-2211230013131102-1112000223020011-2311323203331321-2210122131203130-2200030322303233-0131321233313232-2000202223023230-2222321111100110)
- [bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list](resources--http_loadbalancer--reference--group-011.md#canonical-3332232200233322-0000133212030200-3011111312103103-3313110320010331-2133002130121001-3322312203231012-3000123032130310-1301030021231233)
- bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path

<a id="canonical-2110010211221331-1121132120122313-0231123320211210-0122130320012113-3122110031301202-2330101202122230-2003200110111322-1113121322103321"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023330310300202-0202200210001323-3111323033100032-2100333303213112-0212333021003103-0123210011121233-0321001333230210-1221233201322231"></a>

### Direct properties for `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path`

<a id="canonical-3221130022221113-2230001301232330-3013332130313200-2232201310230221-2101233202020001-2212122111300112-0022210220202320-1301103301122013"></a>

#### `bot_defense_advanced_protection.both_web_and_mobile.js_insertion_rules.exclude_list.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2332303112312010-1332202301303001-2001212232031201-2111320012332012-0030323101020221-3101123133100221-0310333223033013-2000201312321023"></a>
