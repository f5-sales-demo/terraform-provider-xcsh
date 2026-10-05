---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3021121032320301-2013123121002022-2310112313300031-2323030003200221-3210321003331302-0130100230111213-0302201003120232-3023030102223203"></a>

## Next pages — vlan_interface / 223011300311 / 6

- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201031103112231-0330110211313320-2123013321011213-2000002123110033-0230111033002310-2123302332100312-3002021220313132-0332030313023100"></a>

## offline_survivability_mode — offline_survivability_mode / 310130220023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- offline_survivability_mode

<a id="canonical-3203213033022033-0233302332212131-2202111210010301-2132011323032220-2130030113131311-3211113010120211-2223111003303213-2022320213320300"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300223033312100-2311023111102030-2033130231302012-1110033120120302-2233223311310130-3312011303030222-0012201203011031-3111110013010131"></a>

## Direct properties — offline_survivability_mode / 310130220023 / 3

- [enable_offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2203323101230022-1133103332103223-2322333231131013-3322223200313333-0320203103020033-3313002110322230-0310000220213131-1211112300030132): complete subsection reference.

- [no_offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-3123323000230212-3102121330200201-0201030120033300-2323100013322001-3323132303023302-3231120012133100-3002031313310001-2130112101330300): complete subsection reference.

<a id="canonical-2030021123100010-3112101330212032-1330003030022122-3022303213133202-3233111211220122-2331231303332121-0333221202221213-2200322103312031"></a>

## Next pages — offline_survivability_mode / 310130220023 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2203323101230022-1133103332103223-2322333231131013-3322223200313333-0320203103020033-3313002110322230-0310000220213131-1211112300030132)
- [offline_survivability_mode.no_offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-3123323000230212-3102121330200201-0201030120033300-2323100013322001-3323132303023302-3231120012133100-3002031313310001-2130112101330300)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2203323101230022-1133103332103223-2322333231131013-3322223200313333-0320203103020033-3313002110322230-0310000220213131-1211112300030132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301103210232330-3112001132113101-0121213110322002-0330320311311111-1313012312013230-2320123101212001-2121301231220211-3102012300300230"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 201111302322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-0033012223320232-2103233222131121-3211012013110302-2103323232202310-3300021013213031-0231330113302021-3322202232002202-0331302123012123"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-2021312102201301-0302200123220101-1012332222131101-2302111213312223-2013131321000122-1230220233130233-3301120101123301-3013100002331312"></a>

## Direct properties — enable_offline_survivability_mode / 201111302322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010312122321213-1213122222300102-3111313202311233-2101312321330001-3211031103322101-0232322301232330-1001130133130213-3303113020021202"></a>

## Next pages — enable_offline_survivability_mode / 201111302322 / 4

- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3123323000230212-3102121330200201-0201030120033300-2323100013322001-3323132303023302-3231120012133100-3002031313310001-2130112101330300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133033310102021-1233121303021100-1130230220000310-0120300321221222-1233222300310202-0123302310000101-2220330133010020-0113312003031011"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 332322032121 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-1120311012102221-3311200211322022-3132100332303102-0310312200233122-3023213203322210-2213301303222021-3222222303313202-0013233300000022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-1333020011102332-1132131113130333-3133221333313023-0202101113301322-1303232200233303-3221323223230302-0231321302313321-2302023132333323"></a>

## Direct properties — no_offline_survivability_mode / 332322032121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113131000211021-2220131330301233-0302002323010032-3300322110030320-3012011112021000-2000103232010011-1123230211320333-3323221203300030"></a>

## Next pages — no_offline_survivability_mode / 332322032121 / 4

- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-015.md#canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323033031133112-2031000003333201-0112333203122211-1002000233021010-3023031113013033-0010112223300331-1301213033233131-3333321110031320"></a>

## openshift_virtualization — openshift_virtualization / 021002222210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- openshift_virtualization

<a id="canonical-3003123030210310-0301231331323231-0021331202123232-3303321030311102-3012130232302131-0202020101313000-2300110212321333-3123322121330230"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for openshift virtualization.

Upstream description:

OpenShift Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
openshift_virtualization {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020201312311203-0133022313211111-1000212330222332-0201131223302113-0002123010030033-3001302003121300-1031210301233100-1312003113102312"></a>

## Direct properties — openshift_virtualization / 021002222210 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013): complete subsection reference.

<a id="canonical-2111322322013300-1013101332122201-0123013100131102-1321322210011022-0100323113223033-1320322012103221-3311030112333012-1223320312213023"></a>

## Next pages — openshift_virtualization / 021002222210 / 4

- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123111132301111-2121230102133212-0332232133211210-1333003111020133-2213200003231301-3310220333220332-0100011321103001-3232212321012111"></a>

## openshift_virtualization.not_managed — not_managed / 322033333011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- openshift_virtualization.not_managed

<a id="canonical-1121312132013120-3113222101223331-3002202122110303-2211232102202131-0012003030001303-1102032220323023-0200332313201131-1322213033230103"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

Receipt-pinned upstream constraints:

```json
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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320103020230220-3232311301211000-1023013200321032-3323122021300320-0230233110012323-1223233123312212-3301122120010121-3023002131122103"></a>

## Direct properties — not_managed / 322033333011 / 3

- [node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213): complete subsection reference.

<a id="canonical-0020223021011321-3103233000000332-1302323103213301-3113011000310013-2221131031033130-3133302332320320-3303300000112322-0120130033201012"></a>

## Next pages — not_managed / 322033333011 / 4

- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032033101232011-0032001003033323-3112311100011310-2113330310122003-3112202330232313-0303012013323101-3213013220102311-1331311331110230"></a>

## openshift_virtualization.not_managed.node_list — node_list / 212130302021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- openshift_virtualization.not_managed.node_list

<a id="canonical-3323103211230202-1203223332203131-0110121212010110-0230020122331221-3133312211301102-1011120121203103-3032230130222200-3332033113132300"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030220033122123-1000313223021332-1122130330103210-3203021013223102-3100000022031130-0301131020133111-3113303123200133-3232332330333130"></a>

## Direct properties — node_list / 212130302021 / 3

<a id="canonical-2103011133220100-0010331233012221-2333233221013220-2102110030031330-1230130322233322-2232211110221322-0320211131120120-3302330322122103"></a>

<a id="canonical-2311313123033232-2223103100323032-2003003030123300-3233212203220201-1102330330200231-3121330231210030-2321200201302301-1100103310133300"></a>

## hostname property — node_list / 212130302021 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012): complete subsection reference.

<a id="canonical-1122323122103313-1103220222201012-0320332131230112-2300230112030133-0213003200331003-3011003211200311-1020131113312101-3131112001311112"></a>

<a id="canonical-1302130200310223-1213303112120131-3212300221002233-0203302321333021-0121030233131200-2320023233231122-0323202223223330-1321022323122020"></a>

## public_ip property — node_list / 212130302021 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1022132103330223-0333033301223111-0330001000203012-0321222331211311-1331213330132200-3331102032032230-1320123221213232-0130332032112012"></a>

<a id="canonical-2331002212113103-0120123001322013-2003100332221323-0232000220133220-0112323303330211-1330002312113221-2331001122231331-0130032310333232"></a>

## type property — node_list / 212130302021 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-3121201323033221-3133202313332000-1312202031320022-2311322232131123-2001112103231213-3112030213220202-3001011100230213-0100131232300232"></a>

## Next pages — node_list / 212130302021 / 7

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032222202132200-0321202213100332-3000201313113211-2332300121213023-0003310130030021-0313133301330033-1201303103031322-3111130111110131"></a>

## openshift_virtualization.not_managed.node_list.interface_list — interface_list / 031210011032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- openshift_virtualization.not_managed.node_list.interface_list

<a id="canonical-0133110031123132-0212113111332303-3320112132233213-3030032132313101-0213102111111313-2320120200032021-1032300103230212-3303133030232300"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111030200210103-2233033331103000-3010233111203223-3011213231230311-3111332201122111-3131233211031231-0311232030001111-2101331130232022"></a>

## Direct properties — interface_list / 031210011032 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322): complete subsection reference.

<a id="canonical-0312133301131232-3312302233001110-2113110100133021-2330202232213312-3321003230003103-1113023221011310-1121312332300133-0200331133120230"></a>

<a id="canonical-2003121231102131-1312312121333103-1102300303121102-1101002200310210-3313100132310111-3223111220120010-0213121111222101-3132202113220032"></a>

## description_spec property — interface_list / 031210011032 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-015.md#canonical-3120030010211121-3020323231113103-2102330210201023-0020121233001020-3123302220131333-0031310222210310-2222121123112312-3023231322012233): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-3211030230302131-3130002303002032-2001210201122113-0231133232100330-0113210301133023-2310223332332022-1330130120201221-2303012321020312): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200): complete subsection reference.

<a id="canonical-1333100131021021-3122011310000102-2331213120322001-3030111101231112-1033300010220110-0130203100002201-2233300000230011-0012202123201013"></a>

<a id="canonical-0322102101202123-3110032221010130-0122322311031202-1021302023303212-3113221230031003-2232010032321011-0030320313213230-1113122223000322"></a>

## is_management property — interface_list / 031210011032 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2331223002322300-3213310331210003-2213100322300112-3320021020020211-3130020023313233-3332023203230131-1220121201003211-2023021310212333"></a>

<a id="canonical-3210131233032323-3202112122032312-0021032333330333-1121112101132201-1313230331200233-0312123311313313-0002302301303023-0112323222321003"></a>

## is_primary property — interface_list / 031210011032 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-2230323123311012-1012110222111132-0220010203122131-0203302202012020-1230212310203131-2223313231131202-0020222022323010-3103102122112320"></a>

<a id="canonical-1212332031202201-0302323301103101-1233120202030212-2000201120202220-1110022130320032-3301011102300321-3232110113112330-1132220303331213"></a>

## labels property — interface_list / 031210011032 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-015.md#canonical-1213022213322102-3303133001122212-1230310310003310-3120313120012322-1033111331100220-3021211102013112-0233331131030233-1333203312131030): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-015.md#canonical-3323211302312000-3232023112321122-3023321012011011-1110301230322200-2331101002310313-0300011001120332-2221211310033210-0131321302233012): complete subsection reference.

<a id="canonical-3230303021101323-3201302220211110-2321132003100202-1133003223001330-2312231131111030-3032233303232111-3113302010121300-2201300131212322"></a>

<a id="canonical-0320202000233202-1200200131212232-0323212300303311-0201211010033130-2012302122011331-3001332310011302-1222201220220111-3321023330102122"></a>

## mtu property — interface_list / 031210011032 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-0211033211121133-2312110313311223-3331232021333300-1212010223230322-1132311133123323-1103232303211003-1200010113331203-2002130111302133"></a>

<a id="canonical-0000310011030003-1022000312120032-1103132121030030-1021110031121023-1020312313222023-2300010022112211-2201231303103112-3332233101013133"></a>

## name property — interface_list / 031210011032 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-1011011213323210-3122010003102300-0330331301102032-3333301223303111-1201212112100210-3332111231323310-3202020132312321-3022102002123201): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3000222332330022-2133212331113332-3312321101221300-3212001103011211-3331313203222203-3013001232002331-3301220303332120-3023210001123102): complete subsection reference.

<a id="canonical-2103023311312232-1321312103233020-3133230202101012-3330130323233030-2312113133300311-3220331133231122-1200123000200200-0300001220332200"></a>

<a id="canonical-1130211131132231-1102331022112231-0232333102331001-3203330200131021-1023230330320223-2232323000332011-0233203000121313-0223300013312001"></a>

## priority property — interface_list / 031210011032 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-1021203102231201-0231333013111010-1313030332200030-3022331203000230-3220220010313110-2231223230203220-1123111001231323-1033123310230230): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0330203123013021-2220332110311313-1132132123211300-3302302130110331-2120002110212333-0101220202112101-0333103021012030-3022221021223113): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-2011222002031302-3200302122333032-3123332200133332-1213200011123210-2110002011322231-3321210212222320-3231013110330213-3131112113021112): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-0030030233110130-3213101333223331-3101211022101203-1022313322122110-2223021203331030-2030210001231300-2333232113003101-1103112122103212): complete subsection reference.

<a id="canonical-3312320313312031-2013133112222213-2121303313313133-2202320112230232-0321020330012322-0330301003221311-2332011202332331-0103002121030121"></a>

## Next pages — interface_list / 031210011032 / 11

- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-015.md#canonical-3120030010211121-3020323231113103-2102330210201023-0020121233001020-3123302220131333-0031310222210310-2222121123112312-3023231322012233)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-3211030230302131-3130002303002032-2001210201122113-0231133232100330-0113210301133023-2310223332332022-1330130120201221-2303012321020312)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-015.md#canonical-1213022213322102-3303133001122212-1230310310003310-3120313120012322-1033111331100220-3021211102013112-0233331131030233-1333203312131030)
- [openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-015.md#canonical-3323211302312000-3232023112321122-3023321012011011-1110301230322200-2331101002310313-0300011001120332-2221211310033210-0131321302233012)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211)
- [openshift_virtualization.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-016.md#canonical-1011011213323210-3122010003102300-0330331301102032-3333301223303111-1201212112100210-3332111231323310-3202020132312321-3022102002123201)
- [openshift_virtualization.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-3000222332330022-2133212331113332-3312321101221300-3212001103011211-3331313203222203-3013001232002331-3301220303332120-3023210001123102)
- [openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-016.md#canonical-1021203102231201-0231333013111010-1313030332200030-3022331203000230-3220220010313110-2231223230203220-1123111001231323-1033123310230230)
- [openshift_virtualization.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-016.md#canonical-0330203123013021-2220332110311313-1132132123211300-3302302130110331-2120002110212333-0101220202112101-0333103021012030-3022221021223113)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-016.md#canonical-2011222002031302-3200302122333032-3123332200133332-1213200011123210-2110002011322231-3321210212222320-3231013110330213-3131112113021112)
- [openshift_virtualization.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-016.md#canonical-0121232330223030-1033213222121310-1330203213131333-0330121003032131-2102322013231232-3002323223301322-1012012001213201-3010033021333203)
- [openshift_virtualization.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-0030030233110130-3213101333223331-3101211022101203-1022313322122110-2223021203331030-2030210001231300-2333232113003101-1103112122103212)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321212100313302-0212102300120133-1312300231130302-1131210312000333-3020011202332230-1112311311312311-1031300212212010-2303220223222323"></a>

## openshift_virtualization.not_managed.node_list.interface_list.bond_interface — bond_interface / 200103222122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2122012313313212-0103233320222101-0020011223010311-3011111132002330-2021033111301113-2203210210013133-1133232213131103-1103123103310120"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303330130203233-2200333332122300-1311330030003002-3103131222113220-1311221221103231-0020102220111222-3230302003230311-0031333301321312"></a>

## Direct properties — bond_interface / 200103222122 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-015.md#canonical-2033310201130303-2332210213321010-1000222221121002-1200130210032320-2312232023311131-0221303022331110-3310230003202133-0303301213112131): complete subsection reference.

<a id="canonical-0320203333322013-2231120301031222-1303020002011330-1223223000221213-0322030331103320-3331130221120033-2233310201311332-2202312001100020"></a>

<a id="canonical-2120103021333213-2330212110001201-2221030033112310-0300120231130203-3213011333103201-2323203033311121-3031320221120230-1110102123312020"></a>

## devices property — bond_interface / 200103222122 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-015.md#canonical-3210132312220113-0202310201202011-2301210123103023-0133323331330202-2210201022023300-3103313033312102-1000021131102300-3122313302223021): complete subsection reference.

<a id="canonical-2130202023211212-1102220121221112-2013031130311012-0333132230012203-3011320330110321-0310203303032103-1022312110013330-0131310323320032"></a>

<a id="canonical-0123002033102333-3100303102133312-2031111212013232-3232001202100131-3132201011131121-0031202302000220-3120121233320003-3223013131301021"></a>

## link_polling_interval property — bond_interface / 200103222122 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-3302210122233321-2103110200103202-1333222030310103-1302123210301123-1302332133022301-2300130331032302-1033202030132323-2113221011233300"></a>

<a id="canonical-1303120120111133-1012331003230333-2301302302002113-1323232112300001-1232030030123230-2133102210200012-0032223131200033-2122031131222220"></a>

## link_up_delay property — bond_interface / 200103222122 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-1112312013203202-1230010320220330-2020013023333023-0222130232113033-1220031002320211-3310220312231200-2013222103321120-0200322322033021"></a>

<a id="canonical-3003320212122032-2213200203321310-0233131001110030-1232120201210121-2100213022212102-1110011133211201-2330011303021020-3210013133032121"></a>

## name property — bond_interface / 200103222122 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0110311110103133-1001331112023312-2001001300211110-2033230211030232-3101210201333310-0303103033221011-0101222330310312-0210222322231001"></a>

## Next pages — bond_interface / 200103222122 / 8

- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-015.md#canonical-2033310201130303-2332210213321010-1000222221121002-1200130210032320-2312232023311131-0221303022331110-3310230003202133-0303301213112131)
- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-015.md#canonical-3210132312220113-0202310201202011-2301210123103023-0133323331330202-2210201022023300-3103313033312102-1000021131102300-3122313302223021)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2033310201130303-2332210213321010-1000222221121002-1200130210032320-2312232023311131-0221303022331110-3310230003202133-0303301213112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120000303013200-1320202333311221-2330233111031203-0223110333022212-1203222030023111-0113031132122313-2031200323113032-1003023012203030"></a>

## openshift_virtualization.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 132200022013 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1022130202310100-0231113132220011-0210221020210313-1300311001011112-1233233333320122-0202030012000301-2121112231222323-3011130133003222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-2021202022122133-2102211031231133-0222100031310201-0131110023232032-3010100233331321-2202023003222330-3303233321322031-1333120211332220"></a>

## Direct properties — active_backup / 132200022013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131201020221233-1022103320133032-0120002300223221-2010110211022123-1302031120211202-2203300112312200-1120112310320032-1333102010112020"></a>

## Next pages — active_backup / 132200022013 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3210132312220113-0202310201202011-2301210123103023-0133323331330202-2210201022023300-3103313033312102-1000021131102300-3122313302223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023302223123021-3132333122003010-3102111110133310-2122302220233122-0123322311030113-2002013301100202-0120133301233001-0222022211331221"></a>

## openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 103332110120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- openshift_virtualization.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-2213032231012001-0121000130301120-2130322021100130-1323013212332211-3030020032032001-1220302111030000-2211221332033000-0100332310232230"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202211011123312-0301011221013001-1333233302100232-2303031133000110-3122311223220321-2123333202201321-3020203213210203-1013010333010103"></a>

## Direct properties — lacp / 103332110120 / 3

<a id="canonical-0120012312022332-1020211131202232-3000220323213322-1011120232331130-3030033000101300-1033032212133220-1213110211220331-1203212201120013"></a>

<a id="canonical-1333112332332120-1211223003320131-0031301322202330-1330211200201203-2111303022031132-1010300201031002-0321303101133020-3133131230321322"></a>

## rate property — lacp / 103332110120 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1211000121222022-1231321122202213-0200313323122221-2200101210010033-0313113123312132-1300213021101303-2001022021113132-3011030003120303"></a>

## Next pages — lacp / 103332110120 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-015.md#canonical-1000202111323010-0123302031331221-2033012212102233-1100122013000231-1012020033020123-2123302312000032-0021030203022220-3031222111302322)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120030010211121-3020323231113103-2102330210201023-0020121233001020-3123302220131333-0031310222210310-2222121123112312-3023231322012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312103311030221-1313330010031211-1100322121230202-1311302203122103-1202223231220310-3311332223211301-2310222222113111-1103101303121311"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 033021201320 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0222212231303101-0121300112310310-3203033003230011-3300310123230113-0100332223103301-3132031300203300-2002220133233133-2202311303202232"></a>

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
dhcp_client = {}
```

<a id="canonical-2031200123133101-1201332320230032-1110211102330222-2001332302110333-1311011120323312-2021021112010232-2322321132331311-0332000300000032"></a>

## Direct properties — dhcp_client / 033021201320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1123212221132331-1122213230003321-3202131002312020-2101102312121213-3001302121333223-0013101310023302-3000333100101001-2203311130000213"></a>

## Next pages — dhcp_client / 033021201320 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300103202213202-3010311001212310-0311121203213201-1203231001233220-2301331320031203-3223222131202113-2011120300302010-2331302032033112"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 230111202232 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-2000220210211231-3120302311332311-3000213233332031-2132303320130123-3132321212032202-0203231230020123-2111220220002321-0210100112331210"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302221130112130-3133232113100032-2320223110122220-2212031103300003-3321020320113112-0331310111132113-1231221201002203-0322330020202122"></a>

## Direct properties — dhcp_server / 230111202232 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-0323210112011302-3333133111330013-1311303320022201-1011322122200100-0021221110131111-3112111200203023-1103202220220033-0220310012233213): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-2021221123033000-3232132230203103-1201011301313320-1003020232200313-0113210202113232-2110010001023313-1103110310122113-2320112301010032): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133): complete subsection reference.

<a id="canonical-1213301230131000-0321110132201123-1003313131230231-3132213303311313-2131321010122231-1200023031312003-3021202111211323-0330210220210001"></a>

<a id="canonical-2132032211303001-2132231220111102-0212323232120021-1033101312123223-1021101011311232-3022230320323303-0033022231233011-0103033213312030"></a>

## dhcp_option82_tag property — dhcp_server / 230111202232 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0333111313122303-3223203023230322-1220023021011230-0000233102321110-0033133013122103-1332300222003332-1012110232322123-3121020002012102"></a>

<a id="canonical-2311310002011023-1232230013112111-1120203103021300-1211330302322223-2321322021312022-0323023102133111-1130330002233012-0200103102032020"></a>

## fixed_ip_map property — dhcp_server / 230111202232 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-2223230122102210-2333032312101132-2211310112202313-1203223211311122-0021030230323003-3332101001331233-2020320223212330-3103111020323011): complete subsection reference.

<a id="canonical-0033131123330123-0220023211303103-1210222012000331-3121000133303201-2131220323021212-3020332202302222-3021131023312123-1103323330221130"></a>

## Next pages — dhcp_server / 230111202232 / 6

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-0323210112011302-3333133111330013-1311303320022201-1011322122200100-0021221110131111-3112111200203023-1103202220220033-0220310012233213)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-2021221123033000-3232132230203103-1201011301313320-1003020232200313-0113210202113232-2110010001023313-1103110310122113-2320112301010032)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-2223230122102210-2333032312101132-2211310112202313-1203223211311122-0021030230323003-3332101001331233-2020320223212330-3103111020323011)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0323210112011302-3333133111330013-1311303320022201-1011322122200100-0021221110131111-3112111200203023-1103202220220033-0220310012233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111102020201200-0102010320102221-0213023201301330-3301210332333123-3220110013130100-1210123021211200-3332012300002012-1110233013030112"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 213333033233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-1202220313230323-3003202011222320-3312003230321210-1133021210122200-3031303313131302-1201121310301210-3122010032312302-2132203011132232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-3332231233302113-1112011002330230-3322201313223031-1322122003121311-2001110002233101-1132120132202122-1233113200233212-2100110132233100"></a>

## Direct properties — automatic_from_end / 213333033233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233133321121212-1033003213333120-0120132002111220-0311031032232120-1010312130032220-0321211333220301-3000303012002100-1311102030121001"></a>

## Next pages — automatic_from_end / 213333033233 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2021221123033000-3232132230203103-1201011301313320-1003020232200313-0113210202113232-2110010001023313-1103110310122113-2320112301010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102203303322211-2123230321022132-1302201003331021-1120102312021102-2233321332020310-3132222123011010-3101213333223331-2211131130311032"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 232221332000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-0223203111202001-3222001102313010-0101122210200112-0331112313023302-3202323101133020-2202303020221002-2133303222030033-1212021221211023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-2021301001203320-0000010300301303-0131232223111231-3021110312121102-3003302200330123-0011230231120233-2033233201332303-3301312231321220"></a>

## Direct properties — automatic_from_start / 232221332000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330101210322101-3002020020303032-1223002103221333-3210113001113010-2201322132013223-2001020331322121-0002303020223101-3123130001222112"></a>

## Next pages — automatic_from_start / 232221332000 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233001222313332-2123122222203213-1333331213131331-0002120310311213-0130102101131330-3031200012330022-0322210310132302-0131103300312002"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 102130321200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-0010300122212131-3131313330032112-0032320003113020-1210130010101000-3120321022322223-3210110222323320-3211201033113112-0302011301320321"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130122320301220-2102022232203301-0101010120233222-1313323332030130-3101010201032200-0333131200200010-0303323100300332-1001120222003300"></a>

## Direct properties — dhcp_networks / 102130321200 / 3

<a id="canonical-0033112233103100-2012020313112003-3033030300113333-1100230033330203-1031320110300202-2002112003311123-1312013313202301-1021120100321202"></a>

<a id="canonical-2002103133002200-2130010230232031-1333011030102031-3302331231120223-0200221022330201-1011212101320313-0331331102312111-3213011222210103"></a>

## dgw_address property — dhcp_networks / 102130321200 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3233021301211131-3223001010123322-2232031203010002-2330133001311331-2123232000223003-1231021303203210-2122323210021331-0003302120133122"></a>

<a id="canonical-3111031003223202-0011233322010112-0223012210212201-1113212032302120-2321032033101020-2101222202313101-1132131330113322-0312001101321311"></a>

## dns_address property — dhcp_networks / 102130321200 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0331200211020023-3330122302320202-1010302311122223-1023222021331211-3301021123312331-3332303013332233-2331002231121220-1300022211013313): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3222003331300113-2320331123001013-2300130320023200-1323030310331220-0321320103201312-3210213013222322-3012131302013332-2101103310320011): complete subsection reference.

<a id="canonical-1320332332030220-2332311302230330-2223110130100001-1323221031202121-0123201230123021-3230321300312330-3302102123032221-0100031032303303"></a>

<a id="canonical-1323120303230300-1000003222013300-3102112203323111-0100101133121230-2321111001231232-0011012232322101-0100202023212122-2010023231101001"></a>

## network_prefix property — dhcp_networks / 102130321200 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-1013212223120102-0123303033222033-0100110033322221-2032101210303320-1203112013013202-1022200230223003-0011021202231222-1013331321123203"></a>

<a id="canonical-3201321012220300-0123231121131223-2221303133130223-2020312211322332-2102223333103230-2203311112331233-2131232321312103-2000102230102230"></a>

## pool_settings property — dhcp_networks / 102130321200 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-015.md#canonical-3221032312233123-1211121022123310-3303002100132210-2002122012332001-1222033213110133-3203002033301002-0203310030111120-3111232010021213): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-015.md#canonical-1221122231010011-2031131320331303-3111132232220311-0133313033111330-1221130003132303-1031033301302101-2113311013101031-3201113330331203): complete subsection reference.

<a id="canonical-3231130310121311-2211020333201102-0322000233123023-1331003012113110-0123003223012302-2120011131301110-0001212000003312-3000202012022202"></a>

## Next pages — dhcp_networks / 102130321200 / 8

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0331200211020023-3330122302320202-1010302311122223-1023222021331211-3301021123312331-3332303013332233-2331002231121220-1300022211013313)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3222003331300113-2320331123001013-2300130320023200-1323030310331220-0321320103201312-3210213013222322-3012131302013332-2101103310320011)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-015.md#canonical-3221032312233123-1211121022123310-3303002100132210-2002122012332001-1222033213110133-3203002033301002-0203310030111120-3111232010021213)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-015.md#canonical-1221122231010011-2031131320331303-3111132232220311-0133313033111330-1221130003132303-1031033301302101-2113311013101031-3201113330331203)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0331200211020023-3330122302320202-1010302311122223-1023222021331211-3301021123312331-3332303013332233-2331002231121220-1300022211013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102133231221122-0202120000303321-0003332211322301-2122033011131112-1131010113323023-2101203122321201-3122233121210212-2100333331333020"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address — first_address / 211312102202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2212333212211020-0021001300030320-3131300100303132-1031023031123220-0210213023131001-2023212310332320-1112022202103021-0202203210320231"></a>

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
first_address = {}
```

<a id="canonical-3312000333130202-0223301233020300-2130230100210002-2321102122321000-2300310320233120-0010031000011321-3202311221203333-2013020031210133"></a>

## Direct properties — first_address / 211312102202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112210120303301-0130231220012131-1300313320303331-3021030022100123-3200011220323120-3020231003302330-0031222121033032-1233031223011023"></a>

## Next pages — first_address / 211312102202 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3222003331300113-2320331123001013-2300130320023200-1323030310331220-0321320103201312-3210213013222322-3012131302013332-2101103310320011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012221113021021-1113223221101221-1130232321013002-3203102231320323-2313011231011020-1010121033130111-1303323103332330-3300321111020213"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address — last_address / 112201211112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0223103003300202-2130302011211002-3210101323200220-0011001033013002-2321333110302202-2012212222000032-2133300031322100-0300012310031111"></a>

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
last_address = {}
```

<a id="canonical-2330000331121323-0122303102322211-0331120210212300-0003220131110112-0131021031331110-3120011021312022-2322020122031021-1312032232202103"></a>

## Direct properties — last_address / 112201211112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220211011012003-2032331310320321-3223133130003003-3233023311023313-3210323023022311-1203132131232131-3011230211030002-3210301112012022"></a>

## Next pages — last_address / 112201211112 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3221032312233123-1211121022123310-3303002100132210-2002122012332001-1222033213110133-3203002033301002-0203310030111120-3111232010021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032233321100031-1322230202222033-0300020233002002-1201032103313213-2121320112331320-3312032030232023-1232010112130231-2000202022032330"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools — pools / 223130010112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-2222321001132202-1332001231030003-3231231201200112-0320232312323230-1130333320110201-2303023122000300-3202023201300330-0130330201030110"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312112323331230-1312221131311201-1010210123231331-0232010122002131-3103303101000103-2221220001030220-2033023201210323-0101033201012313"></a>

## Direct properties — pools / 223130010112 / 3

<a id="canonical-3331213220303012-1301220312111020-3102212330332323-0221213122211011-2311203231033023-1021123013302033-1331111313313133-3021003202322200"></a>

<a id="canonical-1030102002200103-2033131111233300-1033221232020101-0111320031121133-3031312133130003-1333112123010130-1111310200320003-1311333311232002"></a>

## end_ip property — pools / 223130010112 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1220210101333202-2210232330130320-0313121121232001-1112312110231110-2110002002233132-3012223102012312-2121023312230233-1211102231232213"></a>

<a id="canonical-2203303222133121-2102033020313123-3310322021032232-1033132011310333-2000111312231232-2101320102203110-0103121000003031-0113223302000013"></a>

## exclude property — pools / 223130010112 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-2021232322103313-2310003002202331-2013032212130311-0101301113301122-0120011330201231-3211333313022101-3203133220230133-1312002330203112"></a>

<a id="canonical-0133121110203221-2132232013302311-1111223101233312-2110200121100210-1220212200223022-1313320013330221-3001023030130302-2332301102110300"></a>

## start_ip property — pools / 223130010112 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0312202210313131-0130311303110221-3032213310033030-3032320011131210-3212123133212232-3012212201312113-0110312120113123-0031133020232112"></a>

## Next pages — pools / 223130010112 / 7

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1221122231010011-2031131320331303-3111132232220311-0133313033111330-1221130003132303-1031033301302101-2113311013101031-3201113330331203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002330122323022-1200101231331221-0330031210220221-0110232330212102-1312011033122020-1131202201101201-3122332011331311-1231112322300120"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 032322220031 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3330003002303000-1132231113300121-2330101220002031-0232023302021002-3102211331131331-3200021010331301-3300201310313112-2312011200021230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

<a id="canonical-1213001000312330-2301120120110031-2101202110233122-0233331302223233-2220201000201311-3132203213102102-3120220302200322-3333111321213301"></a>

## Direct properties — same_as_dgw / 032322220031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300232120023032-3111120003030021-2322232012212332-1010121102202301-3201123013330032-3111003233023312-3133020200211213-2122332011212103"></a>

## Next pages — same_as_dgw / 032322220031 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-0031222121131213-3231222020033331-2332223033221233-2112202330200101-1203230033330231-1001012031111331-3000323110213223-0322311002012133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2223230122102210-2333032312101132-2211310112202313-1203223211311122-0021030230323003-3332101001331233-2020320223212330-3103111020323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000033111332300-3321200300200300-1032122132312020-0030001112121302-1321022030230112-2200102031113011-3122222030322001-1311020030333303"></a>

## openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map — interface_ip_map / 313033120003 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- openshift_virtualization.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-3232322300123021-2332323211221330-1300103203302101-2003123022311311-1001131032311200-3220120010123320-3112233322021113-2023130332313322"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

Receipt-pinned upstream constraints:

```json
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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132221212100223-0202021230230202-1132323130220010-0101031221330212-2032231001110000-1012213103100121-1213311111320022-2332123331202303"></a>

## Direct properties — interface_ip_map / 313033120003 / 3

<a id="canonical-0002310222132022-2302332012221200-2230213033232213-1030222102323111-1220310210203330-3113332233300231-1300021311100231-3302310313030220"></a>

<a id="canonical-2203133222132033-3132221230102110-1131312200103111-3132030121231331-0133202221121133-0123012311130030-1330020132300003-3212012101231110"></a>

## interface_ip_map property — interface_ip_map / 313033120003 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-1201320123330321-0120111021200012-3113311101223203-1203210223003133-3031333003132120-0233013112123203-2123102320301231-2112322333321030"></a>

## Next pages — interface_ip_map / 313033120003 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-015.md#canonical-3200300233110011-0000303003113101-3010131112010221-3101033013112120-3102203120012200-0032323121322130-3222211331331202-1030121203113333)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3211030230302131-3130002303002032-2001210201122113-0231133232100330-0113210301133023-2310223332332022-1330130120201221-2303012321020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223113103303233-2322202332002333-2202002123021113-3221303202213113-0001311220021302-2311123300302331-0301233101103122-2211121331022200"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 220013201022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-1110202022021132-2013321232002121-3233011101321102-3111101223332012-1220200112133220-2323203022223211-1000330010132010-0230232010233100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022113122110123-1121022210312210-2311131100210222-3233111002020133-2023112000013000-3001212221201012-3110002321222032-1100332211131100"></a>

## Direct properties — ethernet_interface / 220013201022 / 3

<a id="canonical-1122331230230312-1210131201010333-1311121221103332-1310121211332303-3133313332121312-2201212300223010-3003332212101112-3323012230322203"></a>

<a id="canonical-2321332020212003-1212003013310030-3332332120210321-3203311122200011-0213011033011223-1221102330221010-1203332323220001-2123031202133311"></a>

## device property — ethernet_interface / 220013201022 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0213302221111121-2020120213201023-3221113300102021-2220121132110110-0122220201202203-3130323320202111-0211323202313220-2012102003111012"></a>

<a id="canonical-1033313303001110-1232301131100231-0032203203101311-1031102310223132-3011233310013220-2011221222201120-0120313101122200-1022133212221100"></a>

## mac property — ethernet_interface / 220013201022 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-3300222021220310-1311220330033323-1223302012020322-1322303033030112-3013000230333033-1223233012300222-1011311233320303-0130222122013003"></a>

## Next pages — ethernet_interface / 220013201022 / 6

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323212101330031-1331123003023331-1003211330121333-2310333212331012-0310330110233302-0110132032230023-0111122033112121-3332310001302333"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 011231102230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1032311020103023-3133121330130220-0303323201212030-3231022323013321-1032201012231030-2112111201133101-2012122303123123-2102330000321231"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121112103322230-0123302101103331-3203022132321001-1103130130213313-0100133032212013-2102122300110130-0021012221020232-1033331210222102"></a>

## Direct properties — ipv6_auto_config / 011231102230 / 3

- [host](resources--securemesh_site_v2--reference--group-015.md#canonical-0100011021021321-2033021100023100-1232021332200320-2032012310211123-1213023333123121-2001032022231122-1212312230002132-0202210201203020): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010): complete subsection reference.

<a id="canonical-0312310233203200-0032230120320230-1311010122212101-3202010201302012-2301012110211323-2000012032122020-1110020230122103-3022330101010213"></a>

## Next pages — ipv6_auto_config / 011231102230 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-015.md#canonical-0100011021021321-2033021100023100-1232021332200320-2032012310211123-1213023333123121-2001032022231122-1212312230002132-0202210201203020)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0100011021021321-2033021100023100-1232021332200320-2032012310211123-1213023333123121-2001032022231122-1212312230002132-0202210201203020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131233213122220-2100130030301230-0102232312120200-0303321231000320-0133101023310200-1132203122021320-1033323110301123-0022323120331311"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 232031303012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2301102331321213-1333121200332212-1303313110202033-1330310200013033-2001101012331221-1030131002322332-2110202110030032-0020323212022303"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-1030021030023202-3203231233333013-2122323301230303-3121303021023021-2121203221313333-0121113210331303-1312211103000133-2202023121013131"></a>

## Direct properties — host / 232031303012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010003331100013-3002000202100332-1301100001321221-2223012223120310-3301030131132013-3002220133300330-1112130102022030-2111323110222313"></a>

## Next pages — host / 232031303012 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313112233030130-3222210011223320-3000111100233301-3301321302032320-2102312130300323-1313311021213122-3030013033233303-2310302003030233"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 120310101123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1221313221000302-3200120023110103-3012123231322032-2122011112322101-3030213232013230-1323302132323321-3023233101333102-2101221001331232"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302203331021023-2223120221122101-1032112013132010-0100010131012220-1330220200012001-2021031323233101-0133201000013020-3332021133322231"></a>

## Direct properties — router / 120310101123 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120): complete subsection reference.

<a id="canonical-0202103223021130-0123112200033121-0002330133132030-0333330003233303-3030133132223113-0120113221213231-2030332321310113-1102133121302133"></a>

<a id="canonical-2203113333111320-0231221003101200-0121012110110301-1123303320221011-0300312123012221-2201313210320302-3030202211013010-0010213002120013"></a>

## network_prefix property — router / 120310101123 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200): complete subsection reference.

<a id="canonical-2021112211222311-2012221303020102-0301333032020222-3021022001131233-2012210023032101-3100032302120131-2121330332003010-0132120021130132"></a>

## Next pages — router / 120310101123 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032321121332303-2103120133131133-2030011010020203-2031120111202132-1213332230123200-3232201131023233-0023122033131011-3313231031222012"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 120032310010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-1310310200221112-0223013311023230-2122230312302002-1330211202132233-3123213210022022-0100311011310013-0033013330232123-0300213222120012"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200023331010221-3012122010030310-1001011222030120-1132323220002302-3003232322031133-3300000231202021-2212310213002020-1201230330333213"></a>

## Direct properties — dns_config / 120032310010 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0003210330003222-3221133323332110-2203230221223302-0111203111321321-2131120202330121-0033022130232003-2022201213300000-0003302022113333): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230): complete subsection reference.

<a id="canonical-0321321033233021-1323100102111332-0131131233320303-0313221012333102-3330030321301022-1022001200333112-3203123333320133-1213310310310011"></a>

## Next pages — dns_config / 120032310010 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0003210330003222-3221133323332110-2203230221223302-0111203111321321-2131120202330121-0033022130232003-2022201213300000-0003302022113333)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0003210330003222-3221133323332110-2203230221223302-0111203111321321-2131120202330121-0033022130232003-2022201213300000-0003302022113333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312321011133301-1021113131203010-2300030210021321-3112002121211301-2012230221200110-0011233300021110-0201122330032213-2022112021201223"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 302220121320 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-2001303312301232-1102023132311122-3131102311010002-0233010111001300-1233321333021030-1003112321231031-0202212233011111-1323101213331132"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010112122013121-0312011103113320-0303121312002103-0322120133010022-3030202312200212-1332303201331103-2323100223032010-3100301100101110"></a>

## Direct properties — configured_list / 302220121320 / 3

<a id="canonical-0331331123030333-3000212003300011-3123132231010102-0223032020120330-3200013020031313-3110211200301121-3230223310301113-3210202021110323"></a>

<a id="canonical-3022231130201210-3322331101331111-2011201200320210-0132113032013011-2033011233321111-1033230330110310-3003131303102012-3301313322021211"></a>

## dns_list property — configured_list / 302220121320 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3023102210313313-3303330032301210-3331203320021300-2223333321210030-1031232233033231-2122210021130210-1303300312133332-3322012211113220"></a>

## Next pages — configured_list / 302220121320 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123210020021113-3011132220021232-3122301132030200-1023301132121011-1211213232212020-2013232020232211-1200101123122012-2220100212032111"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 313031022231 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2023213000031310-0011120011001232-2103311002212210-0330203212032130-0301003120131031-3020133320221030-1022012131200332-2120223030320033"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-0202223230331221-0030333303231312-1322211113123031-2012202230311303-3001333032320220-2013131133131000-0130001030000233-2230302021212123"></a>

## Direct properties — local_dns / 313031022231 / 3

<a id="canonical-2022102321120220-0232022013001320-1032003112003211-2303332200300013-0311220313312200-1032121222323101-2211312023121313-0223213300121113"></a>

<a id="canonical-0321003301313101-2002200013312023-0221230333310302-3100220201111021-0113212021130110-0201103210120211-3021200113123120-2233200031202021"></a>

## configured_address property — local_dns / 313031022231 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

- [first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033): complete subsection reference.

<a id="canonical-0211232323120223-3311210320030323-3212113210323102-3320112120003201-3221122020121323-1112123232300001-1021022202301002-3002331113230323"></a>

## Next pages — local_dns / 313031022231 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-015.md#canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-015.md#canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0201011121203210-0130103101231112-2111322110021113-2330222230012031-2033032300000201-2220002000121223-1033331110032220-1003023102002113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210033310323030-1112310322302300-2300230312210233-1220023111100001-1012320210331103-2211221131222033-1323303111103213-1130101231300033"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 103133201300 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0131000302230131-2000223012203011-3003000130113231-1133232132100232-1203303122022201-0221203031220022-0202133122121321-2013132022131112"></a>

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
first_address = {}
```

<a id="canonical-3222120331230200-3310212002303211-3320122131302100-2302333233110310-3032001322200332-1303303123220330-1212221113111130-1210011202210213"></a>

## Direct properties — first_address / 103133201300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102101232213001-2112210302013111-1132222313122232-1003113323113313-1112230332321320-2103223011123223-2010022101132233-1101011002202320"></a>

## Next pages — first_address / 103133201300 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3223231321103322-2302103321131331-2231032212211132-1221132320332030-2000011030310320-2113112333230031-1120300222010033-2113123021232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210112121321030-3310033230322031-2122100003210132-3201101322213021-1000010023221101-0212301212012113-0320132013102231-2132312030003030"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 020203233131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-015.md#canonical-2332233121201002-1012123221021323-3123123231113301-0111000023212000-2110230030120000-2203202220333131-1330021002302200-3013312011020120)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2111132102221232-1311321332030233-3211122301311030-3222211313032111-1012212221023212-0202132320121020-1230212231030233-0021022323000031"></a>

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
last_address = {}
```

<a id="canonical-1031232112210223-0012311103311133-3323302102232301-2102200133210010-0102323030112223-3323332101132222-2302313313322300-0203322310313132"></a>

## Direct properties — last_address / 020203233131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001203322211301-2132120113300100-2012000230333230-2322123003203210-1312311312300201-1233100333322200-0122121331220211-2311011200302022"></a>

## Next pages — last_address / 020203233131 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-015.md#canonical-3023112312233311-2032012303013332-2031320110333131-0310333103033201-0110310110210302-1332231321013120-3202031000220320-1202302113121230)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312320133210211-2030122232212023-1001030010320111-3022302301331001-2323032100331132-2010110210122020-0201113132211302-3312133310130333"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 003220111203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-1203003203230303-2012002030222300-1312221333200331-2120030231303133-3111221222110221-2221130010301112-3000113120303222-1033131220013323"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210312010023002-2021203221310202-0211312320331133-2112102220011120-3003000333103001-1121100023320030-1001100021200201-3320333200131220"></a>

## Direct properties — stateful / 003220111203 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-1321123213122322-2300323310330011-2311022323013232-1211200021210331-1300230102122213-3232311012121132-3200230033102302-0101121202011222): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303): complete subsection reference.

<a id="canonical-1123033031232300-3020103321120122-1013221100132302-0100132333113303-2121000222132000-3300012003012210-3030310123112332-3232211122221212"></a>

<a id="canonical-0131302022322130-3123210111000302-3330021100122033-0113112122331021-3200313102213313-0200320220232002-3211321233013333-1213223121320113"></a>

## fixed_ip_map property — stateful / 003220111203 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231): complete subsection reference.

<a id="canonical-2320032002303020-1111310100230113-2003123003001302-3203010232300010-1002100301333001-1120323110020200-0320211322102220-3000103320300220"></a>

## Next pages — stateful / 003220111203 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-015.md#canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-015.md#canonical-1321123213122322-2300323310330011-2311022323013232-1211200021210331-1300230102122213-3232311012121132-3200230033102302-0101121202011222)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-015.md#canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0003022002013022-3301212213301013-3110132221313300-2111303002120101-2032111210032222-2031021112301331-3010001000233322-3120212301000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120031013002310-1010003332030122-2020220301310123-1222321001321031-1203331120313323-0120211203131010-2101323213012311-1113013133023112"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 321223001223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0101330300332111-3200311310020232-0211320203220012-2200310311210202-0303013301001201-0210133302110011-3210022021100103-3222301131102101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-2023223310321311-1222321102310200-0313003012302021-1122220023113233-2302330301333221-2230311012301032-2013013211300221-2323212313230221"></a>

## Direct properties — automatic_from_end / 321223001223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202113231110023-1232201002331310-1222022221112102-1133122021113323-3310021210211120-1333213312002031-1323201201312000-1202330321101203"></a>

## Next pages — automatic_from_end / 321223001223 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1321123213122322-2300323310330011-2311022323013232-1211200021210331-1300230102122213-3232311012121132-3200230033102302-0101121202011222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001301020013202-3230200112102323-0103302003013121-1100100201100113-3012300232221300-0001003333020232-0010203131213310-3121100310320233"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 102331031033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-3331130110200321-3122031010321023-1130010010121112-3002221022002302-2311212120122022-3200123003223333-2300303201332202-1020011200032132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-1233333220212130-2100102212020000-1023001232330032-2301021120222300-1001232232320131-1322302302011231-2122203100023032-0120123223112331"></a>

## Direct properties — automatic_from_start / 102331031033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121231201321030-0121021202133012-0030300202001322-3310100333201313-3210213213323202-0100220221003120-2011330100113303-3012310110011201"></a>

## Next pages — automatic_from_start / 102331031033 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200201230213223-2301113100033321-0323101232312302-2230211013222100-2203210102011201-3200000000322023-1331200312212011-1013021221033323"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 232121003102 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0112333322212313-0233332103200012-3031232321321003-1101233222201202-0300021333213132-2110221011202113-2103022330312112-0221112320102233"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230232002101100-2331130021012112-1310102232102121-3012233331230312-0033222110202302-2123022313121133-2310000232312031-3302300123221130"></a>

## Direct properties — dhcp_networks / 232121003102 / 3

<a id="canonical-0102023133312102-1122030103331131-0211300131013113-3022021133303112-2202023332320223-3102032321301233-2221211221101032-2110302110230213"></a>

<a id="canonical-0021322130111133-1020322121220033-1130001201030313-3220303232013101-1310033030111032-2211300231222021-2230023012033032-1022023001020321"></a>

## network_prefix property — dhcp_networks / 232121003102 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

<a id="canonical-2010202133113330-1120221312220323-1031033031320020-3321311220212101-1130101131203333-0122131021221302-1312322302210123-3300033121212202"></a>

<a id="canonical-0130010132202102-0032011323201103-0303010223200313-3220121111212010-2212230003113033-2332112322223000-2031032313121310-0232203321232212"></a>

## pool_settings property — dhcp_networks / 232121003102 / 5

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-015.md#canonical-2313331132321311-1132001221220132-0331222201103203-3013021030002312-2230033321133221-2103103201110223-3222131103110330-3130220113222121): complete subsection reference.

<a id="canonical-1102313120222112-0013211330032302-3233000320302010-3030033313021001-1103221123321233-1112233200300231-1003231330021303-1031210332303130"></a>

## Next pages — dhcp_networks / 232121003102 / 6

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-015.md#canonical-2313331132321311-1132001221220132-0331222201103203-3013021030002312-2230033321133221-2103103201110223-3222131103110330-3130220113222121)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2313331132321311-1132001221220132-0331222201103203-3013021030002312-2230033321133221-2103103201110223-3222131103110330-3130220113222121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002333332231120-3313301222303100-2012000320200201-1123303133301200-1001212012021203-1331201022330113-0121310302231013-1103000023231110"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 300303312033 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0222030101031120-2203232113010333-0031323121201323-1033211230202232-3031003323212223-1311220103112001-1123311100321031-2310030112121100"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
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

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031223112303132-2302303013130311-0202033300211012-1220332203033330-2111133021301313-2200331232103231-2032300132121021-2331001212112032"></a>

## Direct properties — pools / 300303312033 / 3

<a id="canonical-0212300233002111-3002130333313233-2101231333210011-2032022121030332-1013113020212012-0330213323211213-1312211232321000-0233322110232000"></a>

<a id="canonical-0202210313323002-3033320223113021-0023121313332200-0232002320201220-1033001110213030-2201211111001031-1232123013220233-2001210331310100"></a>

## end_ip property — pools / 300303312033 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1203321322310333-3102232203211111-3323103123121132-2330202002030202-1202021022232301-1002321103331012-3100133100110011-2130213222110232"></a>

<a id="canonical-0233300032021022-0111310132230310-0210110320312003-0120031223110201-0222023111110202-1030212312011000-2223032330122133-3103023020313000"></a>

## start_ip property — pools / 300303312033 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0103011130331031-2313333230002220-3230032323320011-0000222302202111-3033011230303200-1233320101203223-0121011230101301-1321103113222320"></a>

## Next pages — pools / 300303312033 / 6

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-015.md#canonical-3023113002233231-0122322022122011-0322133003230302-0300122223011320-0031331321220201-3101303130223012-1332013000322132-0132132122203303)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0023220101132012-1312022021323021-0122212202023003-1030220233032213-2221232323310110-1233122300113133-1103003233321130-3033121202011231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310121002013202-2231031232122230-0320332301302200-0232202220330031-0121020122011212-1033333123012302-3131003033211200-1311233033101331"></a>

## openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 120032110022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-015.md#canonical-1021303021330233-1310222233202012-0303020131021233-0113101333102111-2330200333113230-3322201022220300-3110322330010312-2323232002023200)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-015.md#canonical-2020211020001012-0101121103312233-0203333101133021-0001230033132022-3032213320133321-3112201020220120-3101033011111211-3202010120103010)
- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1320100020131231-3031030122003123-3122112320312000-1122020000233020-3021323023000330-2330213310102330-0300022023230203-1122100132010333"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

Receipt-pinned upstream constraints:

```json
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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103333222322112-1011130230303031-2000220213013013-1233131123133000-0210332130021221-1132321033131222-1000010312111232-2233232131221333"></a>

## Direct properties — interface_ip_map / 120032110022 / 3

<a id="canonical-2022021113023112-1220102003110003-0003220133202330-3133312101100031-3310220023021302-3201120003212212-0302120212330010-1302320333202220"></a>

<a id="canonical-0130032321023302-3231310301223021-2301013130133022-2110210000021010-0112201103012011-3120301011123311-0000032032333032-2100313203211111"></a>

## interface_ip_map property — interface_ip_map / 120032110022 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-2021201020012020-2021232210020111-3310310001200020-2322101030102330-0322210300203031-3133123112232310-3302120301030122-0222000322023110"></a>

## Next pages — interface_ip_map / 120032110022 / 5

- [openshift_virtualization.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-015.md#canonical-0010001013200021-1020010101111002-0003212312223112-2302010022000230-0122022231301311-0133010221211021-1011310303120313-3010033111113200)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1213022213322102-3303133001122212-1230310310003310-3120313120012322-1033111331100220-3021211102013112-0233331131030233-1333203312131030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202331200122122-0213011020002310-2003010133303133-1013022002301310-3131310131213001-3303330032230131-0222201003220033-2012123220133213"></a>

## openshift_virtualization.not_managed.node_list.interface_list.monitor — monitor / 013311313103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.monitor

<a id="canonical-3031111122310211-3302310023020210-2311313021023323-0202111322222101-1003333103103222-1223000301323122-0200200313311032-2220201121222130"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

Receipt-pinned upstream constraints:

```json
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
monitor = {}
```

<a id="canonical-0201113101112222-0303120310210123-1333120133332312-2133322101213022-3030102003001330-2013333321011201-1110233321132110-2121312233020202"></a>

## Direct properties — monitor / 013311313103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032011003222210-0011223100201302-3123330110120120-1013231022122333-0222023100113012-1313212010002033-2001322310211032-3123131121323201"></a>

## Next pages — monitor / 013311313103 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3323211302312000-3232023112321122-3023321012011011-1110301230322200-2331101002310313-0300011001120332-2221211310033210-0131321302233012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023102332133303-0202333111312300-2212020030100302-3112333121011220-3010120011121332-1121130123211103-1030010032010000-0320332212311221"></a>

## openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 123302120233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3333113010210320-0033330300013112-1032111100222230-3200321110031102-0003333323203131-0121100130202311-3333113333121210-3000303300232210"></a>

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
monitor_disabled = {}
```

<a id="canonical-2320100230331103-2000213111120211-2233320200222312-1320130201330213-2021302222321202-0010222103132013-1131223303231300-1130131331300020"></a>

## Direct properties — monitor_disabled / 123302120233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211332023113223-2232321130010212-3132030000022123-0211103120232212-1001032133010101-0212210230202213-3013103223312233-2303030112113110"></a>

## Next pages — monitor_disabled / 123302120233 / 4

- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110310101033322-1031322221302112-0102011032330323-3300003003332301-0022313021210213-3000232300301222-0301032001103331-0230103232003031"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option — network_option / 101110102310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- openshift_virtualization.not_managed.node_list.interface_list.network_option

<a id="canonical-1301032021211102-1213331211122003-1303320213313313-3020030301220200-3201300323310332-2323232330000122-0130212321222200-3303213032020312"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300233222003321-3030001322222131-1211233000032013-0001110011030020-2001320130323122-2331033230233032-3031312111312201-1020112310310023"></a>

## Direct properties — network_option / 101110102310 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-015.md#canonical-1203311300332013-2200303332030210-1223311321323021-2130302111211220-0213013012330212-2120112212100221-1230231131201220-0320011001121103): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-015.md#canonical-0203123203103313-1122110333123123-1023231222022003-2010033112013230-0301020022023132-2310201213001300-3301023000020313-3201133103311311): complete subsection reference.

<a id="canonical-3032032203112023-0331130202001111-2222033132332112-1100003120211002-2010021330113330-0113323022102313-0310101010102210-1312011121020222"></a>

## Next pages — network_option / 101110102310 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-015.md#canonical-1203311300332013-2200303332030210-1223311321323021-2130302111211220-0213013012330212-2120112212100221-1230231131201220-0320011001121103)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-015.md#canonical-0203123203103313-1122110333123123-1023231222022003-2010033112013230-0301020022023132-2310201213001300-3301023000020313-3201133103311311)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1203311300332013-2200303332030210-1223311321323021-2130302111211220-0213013012330212-2120112212100221-1230231131201220-0320011001121103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003000103131232-1232111210302322-0301221323112113-2110201132110020-2210101332201210-2223103221200232-1302321002030131-1021303310312021"></a>

## openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 313220223330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [openshift_virtualization](resources--securemesh_site_v2--reference--group-015.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213)
- [openshift_virtualization.not_managed](resources--securemesh_site_v2--reference--group-015.md#canonical-3312123010201223-0322231331231113-1001303330031101-0333001011110002-2013203012023212-3302312220022210-0233131332122310-2310031033301013)
- [openshift_virtualization.not_managed.node_list](resources--securemesh_site_v2--reference--group-015.md#canonical-0132300131131032-2021130132112331-0312003011233211-2101210212103030-1211222230002213-3032312132223331-0301102321220023-0333221223201213)
- [openshift_virtualization.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-015.md#canonical-2310102100123101-1311022222003201-2300021203211133-0123212010313320-0200223222033202-2121330212012032-0013031012001213-2101320202001012)
- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211)
- openshift_virtualization.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0231102121103130-2122332220023100-0211210100300033-1020122030312203-0211223103100013-0223310122020130-1112030233331033-2033322233022220"></a>

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
site_local_inside_network = {}
```

<a id="canonical-2001203003221123-0003223112123223-0023200221030030-0300013301202121-2301103003102010-1113110003100303-0331302110030132-1132132231221033"></a>

## Direct properties — site_local_inside_network / 313220223330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122332132010121-1231102021213333-3310102303311200-3113333212131111-1100011301220003-1232112031233302-2331331310332310-2202023133022223"></a>

## Next pages — site_local_inside_network / 313220223330 / 4

- [openshift_virtualization.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-015.md#canonical-0001231101221021-0233231122130012-2122031032023322-1210233023122311-1001201103030032-0323123012023203-3032321332110103-3020123301321211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0203123203103313-1122110333123123-1023231222022003-2010033112013230-0301020022023132-2310201213001300-3301023000020313-3201133103311311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
