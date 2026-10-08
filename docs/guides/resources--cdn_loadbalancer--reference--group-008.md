---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0321130032220302-2322003033030333-2332311132320213-0310332231001123-2321200010302230-1032011103233021-2333102302033221-1330031000110310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](resources--cdn_loadbalancer--reference--group-007.md#canonical-2112013023232200-2231110000021121-0231333312321300-1333300003302102-1121033120030020-1133301102321011-3013312100220111-3200310111230032)
- bot_defense.policy.protected_app_endpoints.flow_label.profile_management.view

<a id="canonical-2331332223331101-2331203203210311-0110131031020002-0201013331010310-3022101233212330-2303103032231231-0202030230223013-3031101012202130"></a>

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

<a id="canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="canonical-1011131100202302-1023003320223223-0130032212110020-3202310212123201-0021032313301213-0310130132320202-1201220302312200-3020210023221200"></a>

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

<a id="canonical-3313313310310323-1302211322101301-1122323222111032-1000122130202223-2302330100332033-3122323230003202-1220313212003031-3313001202000210"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.search`

- [flight_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-3312300230020231-1230221111102010-3131200230302102-3303302012313030-0212131321310212-3110013300130010-1121121132332333-2133111022221030): complete subsection reference.

- [product_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1331200222232021-1010121202101102-2122330331020202-2232132220113030-3112202223021110-3212131112331320-0130001330020203-2230112130000023): complete subsection reference.

- [reservation_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1300333230313320-3332010101002012-2010230332313111-1003101211113021-2302011300130010-3112002221103332-1202222102023031-1030121110301231): complete subsection reference.

- [room_search](resources--cdn_loadbalancer--reference--group-008.md#canonical-0200000033211031-2110022331222333-0300321133131001-3132220222022020-0001030121103122-0120300102002011-3031322010333130-1320330301102010): complete subsection reference.

<a id="canonical-3312300230020231-1230221111102010-3131200230302102-3303302012313030-0212131321310212-3110013300130010-1121121132332333-2133111022221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search

<a id="canonical-2131130313033223-0300323110312122-1121320112133301-3311020122223230-1113013002222200-0001220110233301-1113310323003221-2111330131312202"></a>

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

<a id="canonical-1331200222232021-1010121202101102-2122330331020202-2232132220113030-3112202223021110-3212131112331320-0130001330020203-2230112130000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.product_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.product_search

<a id="canonical-1103030111323210-3110233203020023-3322100230322323-2212333200100230-1200010312031300-2302302021100211-2111120133131211-1323030001030322"></a>

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

<a id="canonical-1300333230313320-3332010101002012-2010230332313111-1003101211113021-2302011300130010-3112002221103332-1202222102023031-1030121110301231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search

<a id="canonical-1002133332300013-0011031122322220-1213111031203200-1323011220121001-0122310212230102-3123021331002132-2203000331332301-3031102201203311"></a>

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

<a id="canonical-0200000033211031-2110022331222333-0300321133131001-3132220222022020-0001030121103122-0120300102002011-3031322010333130-1320330301102010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.search.room_search` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](resources--cdn_loadbalancer--reference--group-008.md#canonical-1222133011200110-2323313313021321-3320231110330002-3330113033312032-1213101202001223-1101103120112010-0130231022001303-2301122220113013)
- bot_defense.policy.protected_app_endpoints.flow_label.search.room_search

<a id="canonical-1131202200031301-2200323232312322-0132002013111023-2030132230321230-1323111202022320-1313123322132333-1221103020213330-1313323012333320"></a>

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

<a id="canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="canonical-1200320023030131-1012230032333313-1232311022030220-3233033003230131-2230200010223200-2032311120003331-1133022220122210-1132323331330002"></a>

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

<a id="canonical-3132212330301122-1123130333300123-3013230213133133-0323220320323032-0100100120010320-3103221002312023-0203112332033220-0102310031330023"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards`

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

<a id="canonical-3221302321130012-1021323010330332-0232233321122331-1031002022233001-0132311210322320-2131330300122220-0311231320233312-2012230202321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card

<a id="canonical-0300131133132312-0032110312020221-0200111320233230-1110013102212200-3203323112200110-3302231002030112-1012003323013130-1020210032110321"></a>

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

<a id="canonical-0230222010002321-3122033232111200-0232200333223211-0221113020201100-2300300010133322-1222332012230211-2230233210220222-3212120303113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation

<a id="canonical-2233110211300212-0131001300211131-3232201331112201-0030210033200023-2223302100120222-1130303321200021-0012300003023320-2200001121122320"></a>

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

<a id="canonical-0102332233313221-3230201311013231-3332113122213000-0033233131231223-1213202012010322-2300110212332222-3200112212102230-2301201313103221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart

<a id="canonical-1201020110002233-2321322330023222-2021023012112332-1001213112332131-1202323020201021-2121202322302013-1310132200010210-0220311313032011"></a>

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

<a id="canonical-0312012131201110-1220012100103001-0230233212300222-2231131111020300-1000322200210303-1011301113333132-1332011032100132-1233000123033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout

<a id="canonical-3011102211003230-2033012321032232-1030013132113020-1021123230321000-0301303311303010-1210102132200312-0110320300211103-0232222022013223"></a>

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

<a id="canonical-3000132111103122-0012000020321120-0022121331002201-0121102003030310-2020223213122023-1321230302000231-1301230120200301-1021210030213100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat

<a id="canonical-3131313001221131-1132003101033233-0021000211321011-2012230331011302-1333213322013220-3020102023202012-3002203210323001-3210013031220321"></a>

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

<a id="canonical-0103220311012301-3332112213030231-1213330311123213-0133100320032320-2220002000310113-3333123130202223-1130321200202132-0213020210013121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission

<a id="canonical-3300100013031323-0212123121330231-0122013133013313-1111203201330231-0020310131003303-2133211031021103-0032030111302330-1220100202203111"></a>

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

<a id="canonical-3000122332103313-2023021003231010-3100122011033233-2200013320233303-2202211000212000-1012013102112223-3020110210001323-3123211311031223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment

<a id="canonical-3022311302233011-2101132131033230-1301102223110321-0321300232121320-1331330311223231-1300033123100333-1310132300322100-1220103223013103"></a>

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

<a id="canonical-1112310200102333-2132020132311100-0133310332211100-3122130123002133-1012120131302020-0303030100312111-3212112130111002-1121223323111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order

<a id="canonical-2201213011333033-0213333310233111-1031120111331332-3322013010030100-2201011332100221-2023331320310013-0200022112320123-3333210311231101"></a>

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

<a id="canonical-1330011013210322-3321011001101001-1023103132023230-1201110123211103-3232323202322332-2231212200030332-3233320030121220-1321121233312332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry

<a id="canonical-2223310111321002-2231332313133213-3012330333122000-2120123123312123-3331110132103131-1021001223230300-0323020223132201-0233211223101000"></a>

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

<a id="canonical-3331130003130133-0113302011022001-3111013010132020-3133323213022300-3002131021232133-3321303301112121-2313123023033000-2023122131203212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation

<a id="canonical-1210113133001210-1311311310033221-3213232330130110-1032022023303310-1323212331200120-1321102221131111-1222330033233132-0003313223013232"></a>

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

<a id="canonical-1311312001112121-2113030033113022-0022203223122233-2131331110033331-3000033210003031-2103302000321130-3313130202300301-2231133103211311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card

<a id="canonical-2203302032011201-1101122131203101-1312012130320223-2313000113131201-2010303100322020-3202000301130211-0223202001022023-3002113011213320"></a>

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

<a id="canonical-1010222323231020-0103010320210313-3221223233333120-0000122131303311-3302311133013100-1132202331333233-1321222300332001-3120032022320030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--cdn_loadbalancer--reference--group-007.md#canonical-0313230121221201-3202111221220312-0132103303131201-1032023111333221-2110132301012100-1020311003222310-0212010312132322-1113102233330130)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](resources--cdn_loadbalancer--reference--group-008.md#canonical-2122231000130033-0112021222033023-2120013230101130-3012112013102222-0210220233202113-1123001021223111-1032222032333120-3022102332013102)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity

<a id="canonical-2301220021320123-1223230131103210-2231312300100101-2210333020333023-3000013320211212-3103220331302202-2100002302230132-3132010300300132"></a>

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

<a id="canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.headers

<a id="canonical-3102222020223230-3310003230032122-2301123003301113-0213102211203121-3113202333121320-0102103303333122-1113020211130311-1013321000211111"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0213002211030123-3320030232111023-2311103303010011-3323210322332033-1020303032111111-0122201221323220-1303020111131002-0003132302331311"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.headers`

- [check_not_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-0122133301130111-0112232021201013-3020233100102110-0212022130313113-2220232310102110-1331233220131032-1123121333200210-3220122030320233): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-3131020100232021-1123311131013322-1002103010021012-1022201202200331-2310200232001010-2130311103003123-3033102302020110-3321001112231220): complete subsection reference.

<a id="canonical-2131011031210100-0300132030333003-3300230232120310-1210112021331211-2212202032122132-1331333113020103-2131303111200223-3130333303221112"></a>

<a id="canonical-2032321203103302-3100313032210100-1012210230301121-3031002303022111-1101212113011303-2012213121033220-0101003203210202-0120101211021013"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-008.md#canonical-1321203021032210-3330200101033230-2201111101320001-2130020001023130-2103033030011112-2221113010032303-2313032100023310-2120230313333013): complete subsection reference.

<a id="canonical-3110303001023332-1101032111233323-2330122011012122-1221111303330312-2130333000103221-1212230232032311-1213122310031002-2332300133012322"></a>

<a id="canonical-2323133032131232-2123212101032001-0033102330021321-2112032130201303-0030131203311323-1331312131102000-1322132000113010-0201313121333202"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0122133301130111-0112232021201013-3020233100102110-0212022130313113-2220232310102110-1331233220131032-1123121333200210-3220122030320233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.check_not_present` properties

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

<a id="canonical-3131020100232021-1123311131013322-1002103010021012-1022201202200331-2310200232001010-2130311103003123-3033102302020110-3321001112231220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.check_present` properties

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

<a id="canonical-1321203021032210-3330200101033230-2201111101320001-2130020001023130-2103033030011112-2221113010032303-2313032100023310-2120230313333013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.headers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-1100110031020322-1021330213010003-0313300320311220-0101003231320232-2021332111310001-1330021302111232-3020223001131222-0310111230322302)
- bot_defense.policy.protected_app_endpoints.headers.item

<a id="canonical-0000022102332312-3312113303122120-2322112321300132-0000011222101333-3213222013333123-1112001013031012-3033032101330233-0320233303122003"></a>

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

<a id="canonical-3102121211210100-2321031112221031-3332121031020212-0221332022333322-0121222111232033-3002133121101100-3313113102121223-2133232111221313"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.headers.item`

<a id="canonical-2030113323010210-1210210220002032-1033323003113123-3332110301312031-0122310122230331-0021200112302021-1021223000313131-2122111030232002"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3321221000101103-3232020001030213-3132010201020330-1101001200211110-0121121121221122-3231213321231332-2312013210101022-3313002022211210"></a>

<a id="canonical-2003322003112031-1030331123113332-3203230031010102-0012232000103001-1013023312103322-0333222212100321-1220132002311001-2303212300223100"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2120310000310101-1201331203320111-0313302111310030-0030232103000212-0233312121210002-1213010212323233-1100213132123102-2130112011121322"></a>

<a id="canonical-0230010003321223-3202102021010212-2132302031313123-0013331301031303-2123303303331311-0112213322332130-3231303321102133-1133233332012100"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3310130131321203-3331122133211030-1121020112102101-0213323211311213-0102323212323210-3022231033021330-1000312320331032-2330022020312100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.metadata

<a id="canonical-3323121012011302-1002221011223033-3303130323033013-1133101300101123-0213000133111313-2211323303120232-0200103221320103-1313233013110022"></a>

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

<a id="canonical-2333100313221303-2303121221232331-3022232101222331-3100003211000032-1203103212101302-0110121221310301-1122031120201022-1330130212010003"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.metadata`

<a id="canonical-3110013123302312-0112311333133020-2011033021113002-1212103103011030-2300010310022002-3103010123123202-0220322210112103-1010131300013100"></a>

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

<a id="canonical-2303310300330213-0123030323120032-3232113330321300-2003221011303023-3321021102132021-0121221132211221-1310320211102311-2220102220221223"></a>

<a id="canonical-3223232323120223-0202103320122003-0012221003312013-1013001003031101-2023212320102310-0123331301111031-3100001303211200-3323121031100301"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0311023021301223-0213131330202333-3031023111101230-0230200120220300-0032012201030313-3202221121302331-3031112300021232-2220312200203022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigate_good_bots` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mitigate_good_bots

<a id="canonical-1013302330330003-0123313002013210-2123031030203131-1111313123121222-3001032303013320-2030213111103033-2202212010102000-3113130200132231"></a>

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

<a id="canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="canonical-3120013201021020-1230100333010221-1220101222133200-3333230330012101-3210033001301312-3302332202232313-1023201220212033-0032332231131321"></a>

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

<a id="canonical-2330313233230120-0033003213020022-1221333003320022-3113300221030200-2213301020121011-2331002300231002-1210223130313303-0333320211102220"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation`

- [block](resources--cdn_loadbalancer--reference--group-008.md#canonical-3302302022001231-3010031103003012-3133003223203201-1222003000111013-3020303020323011-2101021023230301-3212130200232123-2221312212222312): complete subsection reference.

- [flag](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110): complete subsection reference.

- [redirect](resources--cdn_loadbalancer--reference--group-008.md#canonical-1313313330213113-1313101230021132-1120203311013321-3211102113010322-1100103012220022-0323132213332203-2132220310032211-1113203122030233): complete subsection reference.

<a id="canonical-3302302022001231-3010031103003012-3133003223203201-1222003000111013-3020303020323011-2101021023230301-3212130200232123-2221312212222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.block

<a id="canonical-0323030001213121-2133101120032000-1321222331302110-1000111003221032-0310202233221012-3330110103122123-2330320013023013-3312200100310132"></a>

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

<a id="canonical-0311103331310310-2221030330112321-1013222320220331-3122012012302113-1023111301321210-3330332132121322-3132210032111213-2220311302030330"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.block`

<a id="canonical-0110112102103012-0131113133333331-3001012011213303-3202233313013100-3200131333322302-0021300230033111-1203011123000231-1121030000023310"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2302312332023201-2122100230012212-3221102201013001-0111013132213111-3231312112201122-2021202231302332-1013320022331230-2333330211302222"></a>

<a id="canonical-1123011323322222-0300021332011303-1200301122323032-1233220112122020-1201002010233201-1312131030103313-2112123331232221-1120022011001321"></a>

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

<a id="canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="canonical-0122001120010023-1001223032122011-2220320331210223-1033303123010310-2333133112332031-3202131330233311-2232330232321202-2200013112012002"></a>

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

<a id="canonical-2223302031323203-3003211310113312-2330001101002301-0023110011101312-0033323310003111-0233223130222121-2222123012322333-0212033333031222"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.flag`

- [append_headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-2021023122003030-3032103120101121-2110000210020320-0120132233233113-3103022001101000-3033222000102221-3320231003123303-2001111023022230): complete subsection reference.

- [no_headers](resources--cdn_loadbalancer--reference--group-008.md#canonical-2232022031111012-2123120231300022-0230012213330022-0013201210111203-0320113121001121-1221212111201023-1102322003200313-2102003100103220): complete subsection reference.

<a id="canonical-2021023122003030-3032103120101121-2110000210020320-0120132233233113-3103022001101000-3033222000102221-3320231003123303-2001111023022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers

<a id="canonical-3222321320000133-0211121331230002-1100003132013313-2233230123013023-1100120111212232-0133112312001303-3332322332030020-2233100302322100"></a>

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

<a id="canonical-2203233203312223-3223022021000210-2012211122021201-2132032300330123-2332230301121131-0321011131103302-1003223123203023-3002103133103010"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers`

<a id="canonical-1311312212020101-2123001310301013-1133112000033320-2212223222330012-2003210022100202-3311320233301332-1202111002223201-3321033110321312"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2233303033211200-3100031102033320-0131212001310210-1313302300320123-2233110103203010-2011313332101323-3321220120133321-2000220212131330"></a>

<a id="canonical-2303300300302223-0022323021013212-1021212201102113-0011331120030332-3321130231113132-3030302001203313-3201113101302132-3100321123232023"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2232022031111012-2123120231300022-0230012213330022-0013201210111203-0320113121001121-1221212111201023-1102322003200313-2102003100103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](resources--cdn_loadbalancer--reference--group-008.md#canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110)
- bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers

<a id="canonical-0330230120000312-1033111321333311-3103332013330210-1020121331013122-3113100111013011-0113310230320033-2011102023120130-0333112201322221"></a>

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

<a id="canonical-1313313330213113-1313101230021132-1120203311013321-3211102113010322-1100103012220022-0323132213332203-2132220310032211-1113203122030233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mitigation.redirect` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.mitigation](resources--cdn_loadbalancer--reference--group-008.md#canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="canonical-2202121020212130-0031330113001033-2102111001103232-0203002112020112-2333030123320032-0300313132030111-0200332120220211-1200022100232033"></a>

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

<a id="canonical-2002221312131033-3333201313333120-1210301010000200-0133100121221021-1233111130103113-0122023110323200-0201320123320101-0300330121303010"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.mitigation.redirect`

<a id="canonical-2201302121231000-0221002320320203-3213321312020231-3210101200130113-3222013022102233-3101012310220310-2312233112311003-2021133220202222"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2320231113201121-2020300211032022-1102102103223120-0111130001113120-3301232221202322-3003200222210111-3331112021232103-2000212121221231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.mobile` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.mobile

<a id="canonical-1231131102032000-2023133131033223-0001210110013110-3023311211312033-1312101111111220-2300311321200122-3013130131331221-1211221330111101"></a>

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

<a id="canonical-2223313223123202-2233212130010311-3312012021323023-3002302323032331-3131221133330101-2202011230010131-3013113021220322-1223300321232230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.path

<a id="canonical-2122100031203313-3232303022000332-2121223211032220-3331310221311232-3122231312211332-3212012122110122-1010003222310232-2001101033010013"></a>

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

<a id="canonical-0030312002203003-3230003110122303-1313232021130122-0130303103103020-3003200330202230-3130032031332013-0022322331213233-0313101221313102"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.path`

<a id="canonical-0113123112200103-3003330212200101-0023310322303303-1220331313003230-0312322111232202-1032200301002202-2020121131202102-2200022323303001"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1233212010212312-2022003132303231-1322322132311102-2012312202122311-0310123032200021-1102233201002312-0010311001031111-2212020320302100"></a>

<a id="canonical-1202033031023013-0110023011120302-3023201303223102-1221013112123301-2130233300131102-1321131300111231-3310113000223210-3332103213202013"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0011111310101223-3110232230120200-1301021322321031-2310301121111023-1220110133032032-0101123321320231-1312020001220321-0213030233222132"></a>

<a id="canonical-2133321001132033-3312030311113022-0221220321113102-1330002023130202-1211113031101231-1001213020202201-2222031201111023-0322032120230012"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.query_params

<a id="canonical-2031331220312022-2012111002021311-3201132002110103-3131111122312201-3222101102211000-1011010232030121-1310031023033311-3232111010102231"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1303111211232203-0203223120331000-2003002123321003-2201320303310010-1032212321301230-1222003311030300-1320233013113003-1013100100020202"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.query_params`

- [check_not_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-3031013010131032-2002031330002310-2120213101302321-1120103322331222-3330232000112310-1220103013313211-2133022212030332-1231232102310222): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-008.md#canonical-3022002210001313-3002300223003212-0020330310222223-2202022123021200-2312031310222110-0231012020102021-1100100113000002-1312001201303103): complete subsection reference.

<a id="canonical-3332331232011003-1310023210311112-3121300033103301-3320120233330133-0230331333331003-3331000021321311-0122111022203320-0222321010112000"></a>

<a id="canonical-0311321320121210-3001101121011210-3203121001013121-0133000111320200-2120203131221111-0331112322223332-0033213010303011-0012110303310022"></a>

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

- [item](resources--cdn_loadbalancer--reference--group-008.md#canonical-3211223223120223-2320210230221112-0301311210130120-1303000122132012-0021120313233110-3020001121123120-2021110232322001-0123113103320013): complete subsection reference.

<a id="canonical-3231333230023131-2113103220203301-2221232311000301-1021310023101011-2000031120030013-1000032300303333-3123332001311102-3020011330311003"></a>

<a id="canonical-3231103120222133-3113231100021013-0300210020122131-1012231313202120-0112002111120333-2212321320302233-0202013203030230-0012122022202122"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3031013010131032-2002031330002310-2120213101302321-1120103322331222-3330232000112310-1220103013313211-2133022212030332-1231232102310222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-008.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.check_not_present

<a id="canonical-0001030012111310-2332332113031230-2231223101120332-1000032223322132-0300112100122210-1310322210132021-0021200031202012-1100030133310021"></a>

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

<a id="canonical-3022002210001313-3002300223003212-0020330310222223-2202022123021200-2312031310222110-0231012020102021-1100100113000002-1312001201303103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-008.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.check_present

<a id="canonical-0032230330000320-1233212100213131-3101031003333333-1002000122131130-3313133323202133-1321322201121111-1002010322003032-3231100311300020"></a>

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

<a id="canonical-3211223223120223-2320210230221112-0301311210130120-1303000122132012-0021120313233110-3020001121123120-2021110232322001-0123113103320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.query_params.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- [bot_defense.policy.protected_app_endpoints.query_params](resources--cdn_loadbalancer--reference--group-008.md#canonical-2222021120300201-2020223023102321-3221032331210030-3302030312102232-3312302330311213-2130033013233102-3100313111202202-0211201310111123)
- bot_defense.policy.protected_app_endpoints.query_params.item

<a id="canonical-0130331303033033-0112111030002201-2133311023230313-3312300231223032-3201332302030012-0002200112011133-0230211312311000-3222331222003123"></a>

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

<a id="canonical-3133113001011312-3211002102212310-1301031101201300-1000333020233232-1333131312101232-0203331301222232-3102201133121113-3300111230232103"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.query_params.item`

<a id="canonical-3101310023133231-2101110022022331-0030221213230103-3021011032022022-3012131211232230-3000001330200301-0212120123210001-1023003111103113"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2022322023011001-0123033101203302-3220011222123332-0231313311302020-1123300000222001-3213101310320010-0113222103123123-0103102131202231"></a>

<a id="canonical-1302211301300221-1303333323011333-3330020201222222-1011113211311200-0322230310232010-1310301023302000-2103022200030331-0121310202101123"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2310212223302230-0102322010133003-3000313312030021-3221330230222310-3332210310313121-3123122330310000-3322213223120012-1100331000330230"></a>

<a id="canonical-2000212011103001-0132000301110130-2222112023131011-3103333113101233-0012322201202200-0011203331132003-0020312201121021-0303001212322231"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1103122121102201-2310232002301010-2300221230313032-0023230013011103-2102023331202303-1133033313011223-2120000031313003-1320100210123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.undefined_flow_label` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.undefined_flow_label

<a id="canonical-2332300202221301-0312322022201032-2130002330010201-0020222130213231-0313101133223010-0230310022321032-0102233010310210-2322212120103102"></a>

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

<a id="canonical-0113200330202102-0131023200210002-0120011001302123-3003313020231311-0232312022111032-0210130032310120-0320301113013123-0130001112002010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.web` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.web

<a id="canonical-1333322122333021-1132211132002112-0203230121320312-2311030103221131-0132323312230223-1323330323113030-0211113003322130-2331212302130323"></a>

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

<a id="canonical-1033303310333201-2322202210330231-1303323202223200-0121131311331201-2331312332001013-3100331211333112-3213330112110230-3230321011112100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.protected_app_endpoints.web_mobile` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- [bot_defense.policy.protected_app_endpoints](resources--cdn_loadbalancer--reference--group-007.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013)
- bot_defense.policy.protected_app_endpoints.web_mobile

<a id="canonical-3223133323220033-0003211312031211-0033100222000010-0122312112201212-3331020021020123-2131023310002231-0232312130133322-1123131020132301"></a>

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

<a id="canonical-2031322110131132-0233121000120321-1332022132320213-1300001232021133-1320022100011011-0221133211003332-1322311033122223-0002233032210301"></a>

### Direct properties for `bot_defense.policy.protected_app_endpoints.web_mobile`

<a id="canonical-2310021023310113-1132331323311031-3030220130000032-1101212220021232-2223323132320213-3233202103132333-0003332300111020-1000330322222031"></a>

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

<a id="canonical-0130021111111203-2032310323002002-0122300313003322-0110301130313122-2232111201113100-3201003313113133-0220131013102102-1223311220131210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `captcha_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- captcha_challenge

<a id="canonical-2300021221221121-3320002311031123-1112212300010133-0223113230131030-2120000032200201-2210002122002320-1101331010022310-0302201032122202"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Additional upstream details:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha. When loadbalancer is configured to do Captcha
Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page
will have captcha challenge embedded in it. Client will be allowed to make the request only if the
captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid
Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only
human users can pass through. Generally, computers or bots are not capable of solving a captcha. You
can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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

OneOf alternatives in this subsection:

- [captcha_challenge](resources--cdn_loadbalancer--reference--group-008.md#canonical-2300021221221121-3320002311031123-1112212300010133-0223113230131030-2120000032200201-2210002122002320-1101331010022310-0302201032122202)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-1201312312020123-3230111211203203-0323103120212310-2002022312313232-2200320302313111-0002122232302200-2233200202311231-0212233210001213)
- [js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-1102301031312322-1122311201230301-2021213310112100-3011201130322112-3031202232133333-1230111121313021-2312100212122201-2010022231031120)
- [no_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-3122322231130001-0110033223312111-0211313102301220-1122203023011032-2300330122000333-2313023000331020-1121211003020210-1230333230203311)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-3110233220333200-3212201333323201-2020113221230221-3322221202220102-2132220310202212-2231113002233113-3013031220202301-3203221133222020)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312111213120210-1231013222130110-1110323223220200-2323323001113202-2020000301120332-3010301003331201-1302201211323133-1333033332021012"></a>

### Direct properties for `captcha_challenge`

<a id="canonical-2030232213300110-2222201010123011-0211022030231200-1213120220320031-3131301123122013-3312000012011302-3100203230120323-0212013233110010"></a>

#### `captcha_challenge.cookie_expiry` property

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2123032030231021-1112001101300133-2333311122312112-3320010130020110-0112101132313020-1030122222300213-1110332100322002-3202031012333330"></a>

<a id="canonical-1303332213011011-2212312212232323-1102103321102011-3220203310011030-2013130220202102-2000113201300333-1121210302011300-0102120121011231"></a>

#### `captcha_challenge.custom_page` property

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- client_side_defense

<a id="canonical-2101210301333033-0320000320311100-3201232120000003-0222313110211113-2323001021102312-3310231102100213-0131220011312311-3102233221201121"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: client\_side\_defense, disable\_client\_side\_defense; Default:
disable\_client\_side\_defense\] Defines various configuration OPTIONS for Client-Side Defense
Policy.

Additional upstream details:

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

- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2101210301333033-0320000320311100-3201232120000003-0222313110211113-2323001021102312-3310231102100213-0131220011312311-3102233221201121)
- [disable_client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-2323022203312322-3012211311100012-3031300132031113-3310223011233211-0120103103100233-1101032030002322-1221023301030013-2002330013332333)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
client_side_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302023020100313-3013011022202202-3120222013103102-1011310201002122-0300203222133321-3311230112010033-3020222311202112-0102322031320303"></a>

### Direct properties for `client_side_defense`

- [policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011): complete subsection reference.

<a id="canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- client_side_defense.policy

<a id="canonical-1133131220002112-3130113012120313-1302213332101112-0003131211102100-2111120030001303-0210030023313103-2020302210021301-3122230122021002"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Client-Side Defense policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insert_all_pages_except"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212131022302101-1031123320012321-1113011121003112-2000130313311312-1200113201000101-3313212001311133-3231331100103230-0331211100311322"></a>

### Direct properties for `client_side_defense.policy`

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-008.md#canonical-2130300320210021-3202220131032010-1230100320310002-2323321113200213-1110220232032020-0311120013201303-1100011220313032-1002330301022113): complete subsection reference.

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-008.md#canonical-2201311330313320-1232113110111220-2223231323030033-3320030202232020-2202301102020001-0011312003220012-2111022000112000-0111112132212221): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303): complete subsection reference.

<a id="canonical-2130300320210021-3202220131032010-1230100320310002-2323321113200213-1110220232032020-0311120013201303-1100011220313032-1002330301022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.disable_js_insert

<a id="canonical-0213112322220203-0310233202232221-2230313003232302-0210220121333122-2021302033003221-1320212013012002-2101210132200120-0312202223202032"></a>

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

<a id="canonical-2201311330313320-1232113110111220-2223231323030033-3320030202232020-2202301102020001-0011312003220012-2111022000112000-0111112132212221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insert_all_pages

<a id="canonical-2121220033211002-0233333332013100-2322121200020033-0301232000001020-0222221001331131-0311021122100101-1021202002121202-3230322301211020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for js insert all pages.

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
js_insert_all_pages = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insert_all_pages_except

<a id="canonical-3022312310012312-3220213003200033-3023021110223323-3103300333103312-2212032213232033-0213313331020332-2110003200300032-1031112220030203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
js_insert_all_pages_except {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033200300102030-3123123333220130-3222101212101201-3130110232321210-0331132311133020-0012320212031331-0210230111210300-2222023221132230"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except`

- [exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020): complete subsection reference.

<a id="canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-2003312100300120-0031320302101323-1000212101123100-2303312101311020-2111332112301013-1012133113331203-3031333321011311-0132310233233201"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3101230201131123-0223323002100233-2310331123323123-1122312201231232-3331113313230200-0132331310001222-3311323313233320-3033021321232120"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-0233222230112133-1021222222312133-0100031110310123-1131002012220222-1111213330011213-1121030313011130-2313231233320303-0331131310212310): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-1311320022132201-2220101011200313-0300021221032302-1230121331022101-3022330003311200-0313221232201302-2121010112231133-0000201220321011): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-2221033023222032-0121123212121332-1301303032103003-0202021211132110-1001022112100333-2301200022020012-3133013333312020-3213003301012111): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-008.md#canonical-2210033010113310-2031231013213022-3023300301231331-1322202021312311-3331311033121332-2130111022203022-3203102121120233-1021113311213012): complete subsection reference.

<a id="canonical-0233222230112133-1021222222312133-0100031110310123-1131002012220222-1111213330011213-1121030313011130-2313231233320303-0331131310212310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-1130300003131313-0103122113331002-2213333311203221-3303012111011220-3302231130212120-0301333303222121-2010103211020110-0032103121111202"></a>

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

<a id="canonical-1311320022132201-2220101011200313-0300021221032302-1230121331022101-3022330003311200-0313221232201302-2121010112231133-0000201220321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-3313022330020111-2133321333201211-0132202031003023-1000110030102213-0003222213232033-1013201112023021-1003012110231201-1320121312103110"></a>

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

<a id="canonical-3203030020333233-0232313330030203-1010212202301033-2111033201233333-3212201320022102-1021323330303121-1101001221333203-2112233232122210"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain`

<a id="canonical-3031232201300203-2213222003112000-0330323300122211-1020321223133111-3002112103002321-1313121021102302-3322201322221212-0321000021122133"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2222012312200333-1123101102132003-1103020130112032-0012133032130032-1323320201230230-3022220023313332-2011332303002232-3203320132300213"></a>

<a id="canonical-3013223230013103-1101201220212221-1131101211130230-3000111103011302-2111320200320122-1233222100013210-0113031010322313-0300223311020000"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3323212322200030-1210202321130001-3310120223012203-1311013121230220-3111030123012322-1013112213031310-0100100010032130-2330131210002022"></a>

<a id="canonical-2212330223310021-1303123201012221-3301313121032102-2300312221203230-1110331311013313-3330333300201011-1320122021300311-1003222131013013"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2221033023222032-0121123212121332-1301303032103003-0202021211132110-1001022112100333-2301200022020012-3133013333312020-3213003301012111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-1333312112123212-1230011203202302-1300030112311020-0103021120002232-1322213202002011-0121321230123112-3020303220312201-1313112323231110"></a>

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

<a id="canonical-0113003100123221-1020320302331013-3012123233031212-1221013330013200-2020313213122322-2112212202213103-0010122032112300-1220101020133101"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata`

<a id="canonical-0231003011131203-2120020330021030-1300031011201033-1311030110200131-3213003130132302-3120301003311102-3100103020213020-1312320221131312"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3211112033010001-1320031313211302-3123302330203201-1021200033202300-2320303231320311-3321030301201103-3230111332333332-2120111300012331"></a>

<a id="canonical-2133302223322211-3223030313312022-3222003122200212-0133003231111223-2000300302213022-1312303222320030-1231211132311132-2003300200033210"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2210033010113310-2031231013213022-3023300301231331-1322202021312311-3331311033121332-2130111022203022-3203102121120233-1021113311213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-1333130200023213-3121322131220300-3101113022323112-3202331300231300-3120112300301311-2023120231031301-2231100213231030-2312203221021110)
- [client_side_defense.policy.js_insert_all_pages_except.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-2121232220222302-3110001021203220-3033301202023200-1131112212301111-1011132110032020-3032112203133020-2012223331112121-1221003302210020)
- client_side_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-1213131310333101-2322120110001030-0011132222013320-1210002122100011-2000033032201222-1022132323122210-3303110322101231-3321133123112011"></a>

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

<a id="canonical-2022132112003321-3013200112303313-2101120331201331-2230123313023120-3231133103022201-2310221232233332-3112231012010303-1031323100320323"></a>

### Direct properties for `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path`

<a id="canonical-0311300300323311-2001233312101312-3222232222320030-3032123120121220-3233223303323203-0212122332301011-0101000223103100-2010223132021321"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3203213220001312-0022131212132032-2110002132100202-3103123233200203-2210202323233313-2213002213001013-1330122011210133-1011120321001201"></a>

<a id="canonical-0323323323123231-0011333301211313-2303033112313120-2213303220333133-1011311311100100-1220232110122111-1021331232300313-3203311110002112"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0112122122200121-3213100113323323-1220332211320223-0310322303113321-1201001313230131-1122022002221312-2231331312323001-3023233013013333"></a>

<a id="canonical-1210121220223002-0131231112212100-3202211323132302-0110112303233130-2202311020122112-3100301001033333-1123112100221001-2132233220121201"></a>

#### `client_side_defense.policy.js_insert_all_pages_except.exclude_list.path.regex` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- client_side_defense.policy.js_insertion_rules

<a id="canonical-0032123211201322-1203331031033202-0330112130003233-3131003101312202-3122010122211332-1011012231212020-3131233000000301-3121023303313023"></a>

Type: `"object"`. single nested block, Optional.

This defines custom JavaScript insertion rules for Client-Side Defense Policy.

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

<a id="canonical-1000100322120130-1312331200233121-3320022002232300-1203101021112021-2322122223010122-0110323233000213-0323131320022120-0003103212101322"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules`

- [exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231): complete subsection reference.

- [rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331): complete subsection reference.

<a id="canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- client_side_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-0033021021103030-3110320120300010-1122011303020221-0021330230023023-0322032222322102-0022031122332200-2113211301221130-3311233000001233"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1020231332002210-3213220322012332-1332020103013313-0211332011213302-0333233303301331-1312323023021002-2001133133332132-3202230002000031"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list`

- [any_domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-2112322300003322-1231031202111323-2123300201312323-1221203333221323-0323111022202232-1302100100013210-0331231111021000-0003213221322332): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-2322123032311121-1222301003302223-1201000321312313-2021323231210223-0302012213331013-3302232021032223-0203300101322100-1202012031110121): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-0331132301132103-0021110300121031-2221103332210322-3200133303110302-3323201121112011-0023322210033132-0132213000113303-0013332122001320): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-008.md#canonical-0331030011303300-1311003002101020-1030002303213032-1331121301120132-2233212033130012-2300200020230001-0323010022110202-0002222120010302): complete subsection reference.

<a id="canonical-2112322300003322-1231031202111323-2123300201312323-1221203333221323-0323111022202232-1302100100013210-0331231111021000-0003213221322332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-0122333333321202-1001111023302220-0220002310310012-2301212111233330-0230200112032231-1010331133013013-1031213113332110-2001202023322121"></a>

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

<a id="canonical-2322123032311121-1222301003302223-1201000321312313-2021323231210223-0302012213331013-3302232021032223-0203300101322100-1202012031110121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-2302222233200123-2310123201220120-3322003123221331-0000110223010012-3332231200012221-1102331033031023-2022300330020331-1333320003113121"></a>

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

<a id="canonical-1302002303310210-1103032213123233-1130301021031321-2220333101112332-1101123220203102-1120031200131120-0311201210302020-0133323013202012"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.domain`

<a id="canonical-2203132331100031-0232203032010003-0122203133231020-3212332012111201-0110322312311003-3033131012122100-2121230302330111-0202320120032012"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.exact_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3003222111223032-0213020211121220-2213030201333032-0213121030310110-1021333130220020-2111132103030323-3001033203200222-1303001231233032"></a>

<a id="canonical-2030103313012331-1320233210301300-3323130131233102-2223232333131131-0010310020110013-0323323313300012-0023301332221002-1112312113233320"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.regex_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3300302002310002-2321231110120132-1133332333231311-0322023123011033-1120302022332032-1033120023203013-3000223030231011-3221021300110103"></a>

<a id="canonical-2301001300300130-1212002331311103-2323313201313020-0133220033113230-2323310111203123-3130331321103103-2222312311222212-1121023032201023"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.domain.suffix_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0331132301132103-0021110300121031-2221103332210322-3200133303110302-3323201121112011-0023322210033132-0132213000113303-0013332122001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-2332121222032311-3323130233322121-0332222320222031-3032332302010130-0030110223021123-2230220222211030-3010113310301300-2333232211201212"></a>

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

<a id="canonical-2012312303221202-0233013330130021-3233222002113122-0201301313210222-0032130103110330-2202323201213202-2003320102212103-2112131132021331"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.metadata`

<a id="canonical-2033222001310103-2110021031213101-3022000320113220-1213322112013321-3232330212301131-1233023133230133-1100211230311233-1223023333100230"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1130230332310103-3011303331030303-1001000002031021-0300232301222211-3223332102223212-3333031320123032-3102202230002030-2230133012313101"></a>

<a id="canonical-0331303221020321-0110310033303113-3203321311201233-3010221122223233-0102121312031223-3120101220133112-0211102121323323-3000300113111221"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0331030011303300-1311003002101020-1030002303213032-1331121301120132-2233212033130012-2300200020230001-0323010022110202-0002222120010302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.exclude_list.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.exclude_list](resources--cdn_loadbalancer--reference--group-008.md#canonical-0320231222001331-0003122023230111-1312301203100033-0300202231320230-2111012332101032-2223100032202320-1221011030102001-1332003033230231)
- client_side_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-1010303203000310-3301211222303301-3230012010303132-3112021013023021-1312322301102113-3300103001100311-2200030311111211-2231313221333032"></a>

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

<a id="canonical-1201131033221011-3132102103000003-2321331301130310-2110013213033101-1300131220231230-1013333013332110-3032102013220213-0110203321113211"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.exclude_list.path`

<a id="canonical-0203023003023120-2130003023021122-3310032220022312-3222213330211301-0002010130232110-3123131112203001-1331010302001221-1331211220202031"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3213310002001320-3312302331123122-2201112010333330-2303233113320002-0330021321333001-1230002223233130-2133300002220131-3030220000011130"></a>

<a id="canonical-1011302030131203-3120133123303211-1200100203302130-2330322121233101-3122220203221103-0020333003333300-2213321010121121-2311133312110213"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2201010323011010-0001002002301220-2311313323201031-0221020032301132-0120301231323221-0221131303122301-0321031231022211-2112010001123122"></a>

<a id="canonical-2002123331112101-2231303022233230-1121122120201112-0132031222301303-0212130132100202-3101002200113210-2222210320002300-1013311113302333"></a>

#### `client_side_defense.policy.js_insertion_rules.exclude_list.path.regex` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- client_side_defense.policy.js_insertion_rules.rules

<a id="canonical-2113231100130120-2203032320230203-3232002223231003-2200321232311320-0013032321302030-0201220321103001-3031003300230001-0211321332131013"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Client-Side Defense client JavaScript.

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
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1001112322322232-3003020031033302-3130120011113003-3221100000211100-2023320003013331-3002112122301122-0100232333230201-3333030103310012"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-2110102120002011-3132103223032223-1223230211332301-0333000212113300-3111011223123300-3131320001210310-3311132323121032-0023310010322003): complete subsection reference.

- [domain](resources--cdn_loadbalancer--reference--group-008.md#canonical-0022003231313222-1230033102113003-1301032020211022-1130332002032233-2022100333320303-1200201202213233-0120131200011211-0321023332212020): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-008.md#canonical-3222013222323100-0313302221020131-1021202113310121-2121123020112312-3130121112230031-0123000322121111-2221112013302112-0012130222222213): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-008.md#canonical-2003333022203031-1121112303022232-2331322031223212-1123103020213010-1210100002023102-2211202302302300-1200131301023033-1311100233200000): complete subsection reference.

<a id="canonical-2110102120002011-3132103223032223-1223230211332301-0333000212113300-3111011223123300-3131320001210310-3311132323121032-0023310010322003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-3003200331231231-2033231203133300-0103201222320033-3221212323302222-3023203202001131-2221112210103023-3313011032300332-0031313023312031"></a>

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

<a id="canonical-0022003231313222-1230033102113003-1301032020211022-1130332002032233-2022100333320303-1200201202213233-0120131200011211-0321023332212020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-3112032003220220-0132023321312013-2332222302033200-0010303110332231-2231300102232203-1021202231001000-2231322310032021-2112222012200133"></a>

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

<a id="canonical-3122022303001211-2213120033000303-0331211321232002-2230331222002233-3032113333322120-2232011021003231-0223332121111211-0233132032002302"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.domain`

<a id="canonical-0211201333031123-0121202332131233-2121302130301001-1130323013212230-0102020133230221-2023331300110100-0031331001101210-1033200210113230"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.exact_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2232121011202312-3031232013232020-3211002133032231-2103222322002321-1320201330333333-2213122101010112-2211202202301130-2322203101200112"></a>

<a id="canonical-2033031120321032-3031020133310121-3202302232320103-1102031230313303-3311213113111220-0223231131031120-0020101223033022-1211120332121020"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.regex_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2202213302301200-2110113313201033-3111213303323333-0132200310113022-1001323102221130-3112222001030202-1221131132031102-0320130222113032"></a>

<a id="canonical-3321303313220200-2031010232132012-0032120001312213-3321213111202313-1301222101332012-2002123312220303-1232133223221000-2132100331101330"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.domain.suffix_value` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3222013222323100-0313302221020131-1021202113310121-2121123020112312-3130121112230031-0123000322121111-2221112013302112-0012130222222213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-3112211113001233-3122232313020030-1120300233221102-3021303000203333-3120013120103001-3030021123330200-1302100331132201-2330120022301010"></a>

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

<a id="canonical-1100110233022010-2302310332022130-0223332030033201-2202121223003223-1223021001030202-0013300123012021-0331330113133222-3010011130133111"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.metadata`

<a id="canonical-1033200021301103-3003320100331232-1300302232221120-3311120032212030-1120033332130132-3322100111202022-2200021220120101-2200023222223231"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0103011023303321-1133010032021202-1102310211010123-3123032112230002-1002022310200131-2202213120333200-0313120000132221-3003022201312303"></a>

<a id="canonical-2230221003123130-0103121112202300-2200303322010230-0222001231212330-0220300103113101-3323112331331211-0020232333102203-3130033111033302"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2003333022203031-1121112303022232-2331322031223212-1123103020213010-1210100002023102-2211202302302300-1200131301023033-1311100233200000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `client_side_defense.policy.js_insertion_rules.rules.path` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112)
- [client_side_defense.policy](resources--cdn_loadbalancer--reference--group-008.md#canonical-3332312332120331-0123322313301003-2111102202302323-3132222100011310-3303321202232001-2230300232022312-1323120131333112-0310301121131011)
- [client_side_defense.policy.js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-0021111123200121-1111102322000122-2121121130332032-1003122101001100-0021232323021031-0332023032232202-1210032320213303-3011102231020303)
- [client_side_defense.policy.js_insertion_rules.rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-1102132001030102-2130310321210221-1112300001020133-2332321311220320-0112332022322200-3010001330323213-3313331213202012-3232313302100331)
- client_side_defense.policy.js_insertion_rules.rules.path

<a id="canonical-1022320212001002-1331132122202203-3211012120310110-2210101300113101-1311131032210103-0301000203301330-3002310211010010-2033001313022330"></a>

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

<a id="canonical-0010001131112332-1301233031132310-3232120003001121-1132311011300130-1130203130130313-0311333112110003-0302022310102021-3030233203211022"></a>

### Direct properties for `client_side_defense.policy.js_insertion_rules.rules.path`

<a id="canonical-1132131031330031-1101201031101132-2212101133031113-0232032010300313-3131103011200111-1020101230030033-1303313323203000-0231202311103310"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.path` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3212132311111110-2122301310332323-1012000123113003-3121121201311130-1221210231210222-2220031010112121-3001032230230201-0321232221031033"></a>

<a id="canonical-1132300120002010-0222320331102221-3012211022103322-3313310320112103-2221021011331112-0121311022103131-0310322330022323-3033211133013130"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0223120020230003-3130032322033301-2200012333033200-1203213001000013-1233122022020310-1303010321312330-0301332130331212-0031232022200021"></a>

<a id="canonical-0022333100022110-2110002032030022-1310032100302230-3330222020312103-0031111323030110-2033321013100102-3002123203323211-2021012011231300"></a>

#### `client_side_defense.policy.js_insertion_rules.rules.path.regex` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
