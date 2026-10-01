---
page_title: "xcsh_azure_vnet_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_azure_vnet_site reference."
---

# xcsh_azure_vnet_site reference

<a id="canonical-0222220002211003-2121121211222112-1333031122211230-0101223011020002-0021203000003012-3232013023031033-3012011211021332-2210000021310113"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections — connections / 110002030103 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.connections

<a id="canonical-2122001311113203-1221301130030332-1110211000032211-0201111103011311-1203113103110203-0132032312133312-2102031220120013-2322202220012022"></a>

Type: `"list"`. Computed.

Add the ExpressRoute Circuit Connections to this site.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3113302213212231-3233300102103001-0220231003023110-3003320303132311-1000232002231230-1220200120032111-1113011233000003-3333112320320122"></a>

## Direct properties — connections / 110002030103 / 3

<a id="canonical-1330110101010132-3031120310300123-3312033123322311-3311212020313111-3200311232202221-1212011211130110-2003103300032233-3210303200020011"></a>

<a id="canonical-2212031013230211-0022321121230221-0301201022101310-2320323011001022-3332201023332333-3200332110112120-1220221311203101-3111133111101130"></a>

## circuit_id property — connections / 110002030103 / 4

Type: `"string"`. Computed.

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Upstream description:

Exclusive with \[other\_subscription\] ExpressRoute Circuit is in same subscription as the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

- [metadata](data-sources--azure_vnet_site--reference--group-006.md#canonical-3101321023312132-0322011333012022-3223110121001300-2310020113132230-0333300003111223-1301312203120011-2302030023123302-0221101231201222): complete subsection reference.

- [other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000): complete subsection reference.

<a id="canonical-3112230223223102-2311101230310201-2032101201303220-3101030113112331-1212210101022323-3312233022123122-3132220203112332-0030033201232202"></a>

<a id="canonical-3011032130203321-1301310301023212-2030133221100030-3231321002120312-2133100021210111-3103110320021211-2311201013120000-3332231220202302"></a>

## weight property — connections / 110002030103 / 5

Type: `"number"`. Computed.

The weight (or priority) for the routes received from this connection. The. Defaults to \`10\`.

Upstream description:

The weight (or priority) for the routes received from this connection. The default value is 10.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
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
  }
}
```

<a id="canonical-3012131223013001-0103010303132001-1011330203033220-2211113311122102-3301112030002110-2222311300201132-1213112021032003-3231013010120211"></a>

## Next pages — connections / 110002030103 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](data-sources--azure_vnet_site--reference--group-006.md#canonical-3101321023312132-0322011333012022-3223110121001300-2310020113132230-0333300003111223-1301312203120011-2302030023123302-0221101231201222)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3101321023312132-0322011333012022-3223110121001300-2310020113132230-0333300003111223-1301312203120011-2302030023123302-0221101231201222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001301211210023-0122111330010313-0133320203221310-0100012111223331-1010033230320113-1321320121000103-2303030000330323-3102202110010303"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata — metadata / 322301211003 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata

<a id="canonical-2330030110132122-0133211332101230-2030022003130010-0213320133031103-3113023232013111-3211101232022033-2233330133202311-2131320132301323"></a>

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

<a id="canonical-0212202111311320-2132232220211101-1311132121232133-1010021010333030-0323202012033212-3130310300203013-0322312322023023-0132122000113233"></a>

## Direct properties — metadata / 322301211003 / 3

<a id="canonical-3201301102310022-0232131032301301-3303300220213300-0110102132310323-1301311302103203-3323023123300121-0102330130000332-2301013313110330"></a>

<a id="canonical-2033011223213020-3233311223312100-0012212222332010-3021121302333303-1113332323203112-2323322103300102-0211132013232213-2102302333201110"></a>

## description_spec property — metadata / 322301211003 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3102103211010003-2323120021020003-1032130133013102-2030203303201103-0320133232222301-0222131113030310-2133130203030003-0112313020333011"></a>

<a id="canonical-3221230031030202-3201330313232333-2312111111021112-1222131321223021-3200111113022300-3201020320203110-0332213333311132-0021110101023320"></a>

## name property — metadata / 322301211003 / 5

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

<a id="canonical-0032031020132102-0120320033120210-0230132020011220-3323032110223103-0101330222122313-0101332230311212-3113123322212131-1010313011322130"></a>

## Next pages — metadata / 322301211003 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011130033201012-0310111333320230-0133222131022112-0321332210031322-3011220301300332-2311313021132020-0113331003220230-2311031321210333"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription — other_subscription / 113013113300 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription

<a id="canonical-2111233200132113-1333230302230022-2232231133321312-1222121213310323-2121322121111113-3000001032302103-0332302122220121-2231222210201330"></a>

Type: `"single"`. Computed.

Express Route Circuit Config From Other Subscription.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232310320202312-2232233221201023-0221223332333032-3301220030000001-3030310313212302-3213332331313301-1303200113210001-3021023222013330"></a>

## Direct properties — other_subscription / 113013113300 / 3

- [authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121): complete subsection reference.

<a id="canonical-0230113033011312-0331103013212020-1210033110113023-1322200003311221-0330202302011021-3333023231311023-0322232212333113-1001232112110210"></a>

<a id="canonical-2330012001111033-3201220332130102-1333203112232333-0330112201110132-3223032222111131-1233121113222213-2203033301103303-0303212303311211"></a>

## circuit_id property — other_subscription / 113013113300 / 4

Type: `"string"`. Computed.

Circuit ID. Circuit ID.

Upstream description:

Circuit ID.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-2330112231110013-2030020022223110-3131122032023112-3102311100231231-2313121333030331-1220022223332333-3330200313122303-0123002131200223"></a>

## Next pages — other_subscription / 113013113300 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232213333103113-0210121210023223-3023030311323002-1330301221322132-3301221232012302-0233202231333113-3223221203330332-3000123220122010"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key — authorized_key / 023220321331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key

<a id="canonical-3301222332200202-1211120011001033-2211302000030301-0223320010213220-1321100011032202-3122202201210310-1033230121203222-0121120323320231"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

<a id="canonical-3332232332321301-2012022330212032-1123311232202003-3022133332110321-3310032123220120-2110113000303113-2032303033023000-1022130221232111"></a>

## Direct properties — authorized_key / 023220321331 / 3

- [blindfold_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-2320131313120110-2023132023100112-1202332121032020-3003131020333220-0123103100320033-2233222110012313-0221010302123313-1200001201203000): complete subsection reference.

- [clear_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-1221131313131313-1200103010012233-3123312010213103-3011113232002001-3302221230100012-2111100012032102-1210220003103023-1010210010220012): complete subsection reference.

<a id="canonical-2132310112231220-1223103100333032-3033321011201113-1303013223002320-0313300032031103-1332001033302000-1201121312211311-1123003321222333"></a>

## Next pages — authorized_key / 023220321331 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-2320131313120110-2023132023100112-1202332121032020-3003131020333220-0123103100320033-2233222110012313-0221010302123313-1200001201203000)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](data-sources--azure_vnet_site--reference--group-006.md#canonical-1221131313131313-1200103010012233-3123312010213103-3011113232002001-3302221230100012-2111100012032102-1210220003103023-1010210010220012)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2320131313120110-2023132023100112-1202332121032020-3003131020333220-0123103100320033-2233222110012313-0221010302123313-1200001201203000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000121023312120-1312033212320332-2103110121332301-0102021331123001-2220223232102130-3122021331203022-1031201102112321-1231320121030320"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info — blindfold_secret_info / 302232212001 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info

<a id="canonical-1120122022112111-0113021132131223-3213031120121330-2233000022110223-2100030321021103-1303010303033001-1223022300103212-2000300320222033"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030301100021121-3123213221330231-0302020331223203-0021122012212020-2210010220332103-3010012112221201-3223013213220010-3023122022203021"></a>

## Direct properties — blindfold_secret_info / 302232212001 / 3

<a id="canonical-2333230203313130-3332021311300002-3013030220313211-0333033020222102-1130120233103011-1202110310303032-3332223113211210-1010101302123003"></a>

<a id="canonical-0332323023301310-2303133320301103-0032020332320010-2012231231132310-1031220000303232-1103001002032113-0203132311003231-1030012223120212"></a>

## decryption_provider property — blindfold_secret_info / 302232212001 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-1220132233030100-3330012312033231-1132010331313321-0012121310103102-3122000212003300-0311100222333231-3333003021322231-1113103221332123"></a>

<a id="canonical-1033031012130212-0121132000000302-0132031231303303-1223112100101330-0223222030312030-0333230003012233-2102111011031233-0300331310121032"></a>

## location property — blindfold_secret_info / 302232212001 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3113000301130303-2122010221200211-3132201313103110-1003222223331032-2301230022011200-3301330233022232-3303012201230313-3000200200323232"></a>

<a id="canonical-3011121211233002-0121133310313120-1003002302330320-1213300201131233-2211222021323132-0031123123133011-0233023301203030-0131323230021010"></a>

## store_provider property — blindfold_secret_info / 302232212001 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-1001012230232200-1222322032103022-3321003320032130-2332211221003101-1320221212133213-3301032113212100-3131011220211021-2202000330131112"></a>

## Next pages — blindfold_secret_info / 302232212001 / 7

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1221131313131313-1200103010012233-3123312010213103-3011113232002001-3302221230100012-2111100012032102-1210220003103023-1010210010220012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030022301112223-2201202221110021-1331120022021310-1032202002101202-3223000101232201-1002220001110021-3303232033113220-1321103023122111"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info — clear_secret_info / 121302012221 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections](data-sources--azure_vnet_site--reference--group-005.md#canonical-1230322032112312-3310310003131321-3032300000112032-1330232230123131-3231000201032131-2102330130310322-0112021110130210-3013121110332323)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](data-sources--azure_vnet_site--reference--group-006.md#canonical-1101230133331033-2120323320220323-1010110201231201-3012220033023012-2010013022012013-3222032113320322-1222112003321333-2322203332010000)
- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121)
- ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info

<a id="canonical-2121011103112020-0022103000233231-1013110012223002-0032000302013032-3202320120201310-0300133112223031-1030333300132111-1211222333022030"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1223201332311230-2113331000230020-3111033222022322-0133012212021200-0230000313222211-1003321201100111-3031001302302012-2301130120112011"></a>

## Direct properties — clear_secret_info / 121302012221 / 3

<a id="canonical-2001122032322002-3132332221132303-3131123212003333-1131032012100323-3311333102301230-2020030221110203-0021232031332312-2102132223200300"></a>

<a id="canonical-2012232221003103-2312302003103310-0030311020210021-1110312103210202-0010312030102123-3121322311121233-0000222032310021-0303132030331010"></a>

## provider_ref property — clear_secret_info / 121302012221 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3000123020331200-1032211013232020-0221202102011032-3030330102323120-0221333220121320-0000002102220212-1311303003333113-0011321212001110"></a>

<a id="canonical-0331100012101332-3330321213223210-3220100210210221-2303133331122021-2020222211030201-2103102113231111-0133100013310113-3132133010230131"></a>

## URL property — clear_secret_info / 121302012221 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3013221021122211-1013323222031230-1103320220303023-3112330221122230-2302310012101312-3100211331230101-3233222232020231-1200313313030001"></a>

## Next pages — clear_secret_info / 121302012221 / 6

- [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](data-sources--azure_vnet_site--reference--group-006.md#canonical-3133023123331033-2000222100301201-2002001210321212-3130322222322301-0310301102223112-1010023200001211-3313101130202132-1322021100111121)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2023033331210222-2230333023012012-2001203323310220-2100322210231333-2212333100032331-1022332113301211-2222220012113010-3123232322303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110201210023033-0233112323021331-1302022110023032-3311323013202120-0113012233132222-3320123033000132-2133213301331020-2331122211011201"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server — do_not_advertise_to_route_server / 001130213000 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server

<a id="canonical-3123101203123211-2232321322131300-1121013121332213-1303220120220320-0123333331202123-3120022202321203-0133301133201311-3111311102121313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise to route server.

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

<a id="canonical-0121222211202313-3311210333310333-2331001120120122-1022300020121033-0102100003030120-2232223133300320-1303201000202332-3020011103310022"></a>

## Direct properties — do_not_advertise_to_route_server / 001130213000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220302011100122-3111122123101012-2230000220301021-1013223222323313-1323311111111002-1130021010120332-3220033102230223-0133202313130320"></a>

## Next pages — do_not_advertise_to_route_server / 001130213000 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032013302303303-0112221330230221-3211211301012123-1203200232221302-3221203030301013-1100012102220322-3300032310032330-1033230100002213"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet — gateway_subnet / 002223232113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet

<a id="canonical-1301322111233010-3301231131003021-3023312002003012-0020003121102032-1023102201222101-0112000122133121-2103310203231220-0333221031330112"></a>

Type: `"single"`. Computed.

Configuration parameter for gateway subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-2023112303311330-3033012230222012-1010202101231030-2101202023002002-2033231023221330-0132221133330230-3331223032233022-1311230002130201"></a>

## Direct properties — gateway_subnet / 002223232113 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-2112121001020221-2010003131320331-2020331230131030-2010120133101302-3330132011310330-3311010332023220-0012310222331120-1010131320012120): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2120330112221233-3200131113333102-1133020103021200-3223222301110220-0321021103233033-1012112213103100-0002003201100321-3121331302303023): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-1132332311210032-1230233022222321-0330121120121300-0130330121112333-0210003000231223-1231031323331230-0023223202223203-3130212110233222): complete subsection reference.

<a id="canonical-3222023012211311-1322323113132013-0032113000321132-1200231222213232-0222221000000320-1333012111301130-1302220110333103-1300130301220312"></a>

## Next pages — gateway_subnet / 002223232113 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-2112121001020221-2010003131320331-2020331230131030-2010120133101302-3330132011310330-3311010332023220-0012310222331120-1010131320012120)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2120330112221233-3200131113333102-1133020103021200-3223222301110220-0321021103233033-1012112213103100-0002003201100321-3121331302303023)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-1132332311210032-1230233022222321-0330121120121300-0130330121112333-0210003000231223-1231031323331230-0023223202223203-3130212110233222)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2112121001020221-2010003131320331-2020331230131030-2010120133101302-3330132011310330-3311010332023220-0012310222331120-1010131320012120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013002132003030-0103222133001231-1130213223100202-0002301213123003-1122332231123212-2303331012202002-0102032322003331-1110301131112101"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto — auto / 031212220312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto

<a id="canonical-0012112333222212-0013110013122301-1313112021000212-1001030102211200-1200202030031303-3032021002332330-1223330132121023-2303332002331130"></a>

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

<a id="canonical-1132110321203223-3102111001302211-3001220103230203-0120330131202110-3002131111322203-3021131130323100-2103212100130300-0301302131213033"></a>

## Direct properties — auto / 031212220312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323231103020033-0301332121032232-2231331011220222-0301333201010331-3210303022002310-2120121011210333-0212220120100300-1100122301223321"></a>

## Next pages — auto / 031212220312 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2120330112221233-3200131113333102-1133020103021200-3223222301110220-0321021103233033-1012112213103100-0002003201100321-3121331302303023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000302112033110-1023101131100121-0300100032132023-0312000222333320-0120201212113322-0313321113232100-3221232223221101-3122231113033002"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet — subnet / 001112130313 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet

<a id="canonical-0332031301103030-1211202133100000-0113321232133131-2131311203220203-2231221003010010-1021121121202002-2211003121023212-3133122222212102"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-1303013320330323-3110212021130231-1313120300021032-0322232332022022-2333133203120211-1333133033201332-2300032301011020-0301021201210100"></a>

## Direct properties — subnet / 001112130313 / 3

<a id="canonical-2103200103112210-3202000301221030-1023330230011021-1213000012000331-1333320331320202-0210032212032211-3122021031333112-3130120032302222"></a>

<a id="canonical-0222102123010212-2030312223032121-0220300002012301-3032011103002112-0023121222203000-0312001321133332-0233023020333002-3003222002323122"></a>

## subnet_resource_grp property — subnet / 001112130313 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-2231231022211223-1023200313132111-3230202112010031-1211100102032123-0123112111110223-3202111111213000-3101332001203013-3213032113230011): complete subsection reference.

<a id="canonical-2012211121103332-2111221222020301-1133313132210332-0311211203303330-3130001331233321-3112221120103110-2300333132230233-0211002303003311"></a>

## Next pages — subnet / 001112130313 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-2231231022211223-1023200313132111-3230202112010031-1211100102032123-0123112111110223-3202111111213000-3101332001203013-3213032113230011)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2231231022211223-1023200313132111-3230202112010031-1211100102032123-0123112111110223-3202111111213000-3101332001203013-3213032113230011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313201223221001-0022322133313333-0230123023101323-3000313211203122-0231103020310322-1221121323301323-3203310022123203-0011102030201321"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group — vnet_resource_group / 220320011102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2120330112221233-3200131113333102-1133020103021200-3223222301110220-0321021103233033-1012112213103100-0002003201100321-3121331302303023)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group

<a id="canonical-1102102001233201-0201323023322120-2332031311102202-2221330113011130-1033003123323001-1201200322001002-0210321110111323-2200033100300300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-1310233113132112-0303211020022330-0313220220131023-0203132023000132-2030133022332031-3331220331232032-3121031032000020-2222330203230330"></a>

## Direct properties — vnet_resource_group / 220320011102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033300302020201-3302013101210213-1300133011130311-0221030032330122-3220002312010103-2020320231031232-1300022311103123-0213302230201230"></a>

## Next pages — vnet_resource_group / 220320011102 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-2120330112221233-3200131113333102-1133020103021200-3223222301110220-0321021103233033-1012112213103100-0002003201100321-3121331302303023)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1132332311210032-1230233022222321-0330121120121300-0130330121112333-0210003000231223-1231031323331230-0023223202223203-3130212110233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003200102131021-1210101233302202-2212323010033031-3313130223112000-2331222303332101-1202021101232131-3220230122033103-3313032302132232"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param — subnet_param / 210212310031 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param

<a id="canonical-0233120120103133-1101300111030013-1233223212013123-0010122031010333-3030100101001233-3211112013221231-2213121133320232-2133233222300112"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3322202303122133-1323013301100131-2203310302332111-0021122232130021-3220122101320233-1302121333301232-1011122012211201-3030233003231330"></a>

## Direct properties — subnet_param / 210212310031 / 3

<a id="canonical-3102331322230221-0220121233202302-0331222133212222-0012023231122213-2300300230302123-2331302313331022-3220302330313202-3331130332101231"></a>

<a id="canonical-3023230230113320-1221231110211313-1211330232210222-3223200012321322-1220010231310220-2211310333222220-2002012122212230-0000032211130102"></a>

## IPv4 property — subnet_param / 210212310031 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0011111300230001-1332303223330323-3313312201220221-1323321202210223-1022122200333020-3320223011323012-1213230013311113-3133100210133331"></a>

## Next pages — subnet_param / 210212310031 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-0232123213033320-0013013221100103-2113210133031322-0133330213000120-1000232022013201-1210232022323301-2102211113213033-0032130001122010)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300133233130233-1202033020123220-2022233221102031-3103121200011021-1312323002111320-1202101033120300-3023330321222331-1120311221321033"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet — route_server_subnet / 010202332120 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet

<a id="canonical-2201233012321032-3000320020210103-1011331023121121-0132130112221200-1022011023313303-1031333133000210-0033323012313313-0011100220111223"></a>

Type: `"single"`. Computed.

Configuration parameter for route server subnet.

Upstream description:

Parameters for Azure subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"auto\",\"subnet\",\"subnet_param\"]"
}
```

<a id="canonical-3230020331330031-1000110201320201-1003033210303012-3230030032120222-1021012321323031-1130310211200022-1303332310231030-1033020223012330"></a>

## Direct properties — route_server_subnet / 010202332120 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-1102231121022022-3232030000103103-1333131220210000-3122313133330323-0111012303222003-2300212010311233-1201033133322311-1331121011112331): complete subsection reference.

- [subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3211200100020231-2313133320312123-0322020000030000-1222313203203213-0133233133221022-1323031232313220-1130212301203130-2113302220110311): complete subsection reference.

- [subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-2023021311103200-0301302210010031-3021233121023013-0203003111221300-0013333213111313-2233021313333112-0111311200203031-1012212110100300): complete subsection reference.

<a id="canonical-0103302303320130-1232330010332331-2111313033120002-1201210221323001-2132311101212211-2313013300310312-3313233120030202-2302100111320103"></a>

## Next pages — route_server_subnet / 010202332120 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-1102231121022022-3232030000103103-1333131220210000-3122313133330323-0111012303222003-2300212010311233-1201033133322311-1331121011112331)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3211200100020231-2313133320312123-0322020000030000-1222313203203213-0133233133221022-1323031232313220-1130212301203130-2113302220110311)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](data-sources--azure_vnet_site--reference--group-006.md#canonical-2023021311103200-0301302210010031-3021233121023013-0203003111221300-0013333213111313-2233021313333112-0111311200203031-1012212110100300)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1102231121022022-3232030000103103-1333131220210000-3122313133330323-0111012303222003-2300212010311233-1201033133322311-1331121011112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122301230020113-2223110101321230-2233303123210203-3000212133220022-1122330221030332-2231213101112031-1223122333311023-2001213013230312"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto — auto / 103003013122 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto

<a id="canonical-0133033321323212-0302223323133213-3010220100010032-2110220112330231-3301231021001113-0312021131321322-0111301231111231-3022223223231033"></a>

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

<a id="canonical-0022110020200322-2133030131002121-0211010101122011-3301031111303210-3031333111010003-1120012203112103-0023200122301313-3230321331030030"></a>

## Direct properties — auto / 103003013122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211331202021010-3310321310030300-3331022122233100-3023203302033013-2000032003131221-1330000233100101-2303200321311032-0003301113232103"></a>

## Next pages — auto / 103003013122 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3211200100020231-2313133320312123-0322020000030000-1222313203203213-0133233133221022-1323031232313220-1130212301203130-2113302220110311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303332312001231-1022222311332120-2333322010330303-1203011230231230-0012313301222033-3232320110311000-0112102020131302-1122103223130113"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet — subnet / 332212112002 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet

<a id="canonical-1001233202202313-3131232101332010-3301002200013003-0120030001100323-2310330201130233-2010031030123012-0022001301331010-0020123000321222"></a>

Type: `"single"`. Computed.

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or
RouteServerSubnet).

Upstream description:

Parameters for Azure special subnet which name is reserved. (i.e GatewaySubnet or RouteServerSubnet)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-resource_group_choice": "[\"subnet_resource_grp\",\"vnet_resource_group\"]"
}
```

<a id="canonical-0211222313133303-1332032331211313-0201100202110312-3131212102022222-3310331133312221-1233321001311130-2310010103022203-3133311332023121"></a>

## Direct properties — subnet / 332212112002 / 3

<a id="canonical-0203310211322210-2222210013103301-2211132103100323-3220200220301013-0332133033211332-0101111110033212-3030012111221230-3111123213100330"></a>

<a id="canonical-3033233121130312-2103023011003333-0112231120030330-0020312100330232-3031310300130323-1100200302233231-3322231330300201-2222111203310302"></a>

## subnet_resource_grp property — subnet / 332212112002 / 4

Type: `"string"`. Computed.

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Upstream description:

Exclusive with \[vnet\_resource\_group\] Specify name of Resource Group.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-1302230323121032-2312022320301220-2320202211131033-0313012330211130-2331202323303132-2333033210013103-0212102113000032-1200213302101332): complete subsection reference.

<a id="canonical-1130211222122212-0103003012113102-3102123013322202-1220000003320011-1223102022321202-2213121111032130-0211113113203331-3320301303201130"></a>

## Next pages — subnet / 332212112002 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](data-sources--azure_vnet_site--reference--group-006.md#canonical-1302230323121032-2312022320301220-2320202211131033-0313012330211130-2331202323303132-2333033210013103-0212102113000032-1200213302101332)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1302230323121032-2312022320301220-2320202211131033-0313012330211130-2331202323303132-2333033210013103-0212102113000032-1200213302101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200011301303213-1223000212300122-2201310010312112-0012302131303213-1030232111220312-0303031333331210-0110200303112033-0102130101333023"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group — vnet_resource_group / 301000111220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3211200100020231-2313133320312123-0322020000030000-1222313203203213-0133233133221022-1323031232313220-1130212301203130-2113302220110311)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group

<a id="canonical-3120103102303110-0312023110302222-3303201122122120-2011332123110212-2202322011130223-0223100122211222-0311301121110333-1002101121121122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vnet resource group.

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

<a id="canonical-0220323023312122-2301100332223331-0202102002113023-3121210230011000-1213231000113212-1233030020331200-2031301303013311-1032110310101311"></a>

## Direct properties — vnet_resource_group / 301000111220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033101232002021-1003300021001120-3233113211033033-3333201323020313-2011112232320311-3312033011010020-0300321210100002-0221121111032323"></a>

## Next pages — vnet_resource_group / 301000111220 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3211200100020231-2313133320312123-0322020000030000-1222313203203213-0133233133221022-1323031232313220-1130212301203130-2113302220110311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2023021311103200-0301302210010031-3021233121023013-0203003111221300-0013333213111313-2233021313333112-0111311200203031-1012212110100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231222131101020-3103331330132013-0101220012002101-0322321301200132-0331200202301330-1320312022103112-3333022021213320-0101011213303222"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param — subnet_param / 300001212100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param

<a id="canonical-1032311312211220-2000030202310321-1223013220011312-0233221212311123-3323113323200021-0032322023331322-1331131130313003-2303120302331101"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0330311100321110-0101100113132012-2210120332021112-3220230201100103-2330201231313031-1100100331032110-0120320220220100-0332030101103133"></a>

## Direct properties — subnet_param / 300001212100 / 3

<a id="canonical-2021203332003223-1033010122312012-1101032212131220-0211303113103223-2230022012200312-3120100111023230-0321120233230313-2010122133322311"></a>

<a id="canonical-3231320302223031-3313102301113020-3031222133232110-2300022121201022-0101021112000301-3201101300102303-1200100112230232-3332310130132120"></a>

## IPv4 property — subnet_param / 300001212100 / 4

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0002311300302223-3132112322301302-2302332133013332-0330012212130022-3332203013111122-1003010213132111-1222012332100202-0333333222330221"></a>

## Next pages — subnet_param / 300001212100 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111222100023331-0201221313321132-3313000130300221-3320220303130312-2030210020013132-0000202202013110-0300000331330231-3120320230201202)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0221210213203222-0012113111230331-2303312003212102-0222102310133123-0310332030312123-1311033112100312-0133112203120001-0200302031023221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221023303023100-2331333021200020-3202130320212121-3103311131113223-3102333001311323-1300303112103312-3200103120320321-3012103202231223"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route — site_registration_over_express_route / 120331302320 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route

<a id="canonical-0103333120212201-1103322230130000-0320220221111332-0113213300031032-2320221202023313-1012100202030302-3330323231012302-2321201113102210"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2321120102031130-2331133021321001-1133332332111011-2121321020222033-1032311231231220-2113221132103031-3021023023231231-1033222032011003"></a>

## Direct properties — site_registration_over_express_route / 120331302320 / 3

<a id="canonical-1300011323033321-2232131323023302-1312133312331330-0030311033301012-0012102302131300-0332113213131303-3013023101010022-3002123230203112"></a>

<a id="canonical-3313230133101210-3003330312123303-3100220130101331-0011203231233201-2331311021311033-0102301220032202-0312201003122033-3303220002300122"></a>

## cloudlink_network_name property — site_registration_over_express_route / 120331302320 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1132330110011203-3200131332031110-0212311321011131-1132330231112231-1302222300311302-2333222220221030-2003233103021103-0133310132220311"></a>

## Next pages — site_registration_over_express_route / 120331302320 / 5

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3322112221103022-1322032123023103-1233300213112123-2101013332232100-3203120023210100-1131102111120331-3321310121112303-3020232121203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010131232022201-0232223111202221-1330111010311100-0132322221311020-0333300130333302-3010312030213322-0232201102023133-0231321211330033"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet — site_registration_over_internet / 032221230203 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet

<a id="canonical-3110233021020133-2112223300032220-2331220031332113-3233313220001330-3101300010110303-2333010003012033-0313320331020200-0312102131020123"></a>

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

<a id="canonical-3310022021132222-1200230300003311-1102010030110033-1002302311332211-1221021120201032-0112131130003231-0030012133222000-1013302032230123"></a>

## Direct properties — site_registration_over_internet / 032221230203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011023300121233-3133010312231030-3120210232133333-0022130223012223-0011022321032032-3302011011220010-1002121210123023-2112233311032220"></a>

## Next pages — site_registration_over_internet / 032221230203 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3011100300322203-0212310322003212-0333213120000101-3112233012222022-2210121212200313-0131103002013221-2001133020311021-1013131000230232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320102031021131-1202321000310010-0232310121111022-0320322010110232-1311003312133310-2103122303120031-1003022312313200-0302013030021113"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az — sku_ergw1az / 311322120301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az

<a id="canonical-3311233003302031-3022031103031231-0233011220000221-3122000121112310-2103233010321001-3101123331103131-2130300111003201-3200130001220213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw1az.

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

<a id="canonical-3123021203103112-2201323230122130-1030211300212033-3200111311310100-0012011033210110-2123120211202101-3222133201010203-1210103100111201"></a>

## Direct properties — sku_ergw1az / 311322120301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121132230202030-1021020013021233-1220330210213300-2112101132200203-3203110121313030-2020203121333132-2010120320233023-3012023132123002"></a>

## Next pages — sku_ergw1az / 311322120301 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2222311000330111-3000222023332301-2311333110200220-0123102130323230-1031310102321111-1031000201221003-0130002130113011-1300232321133322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223321101100133-0301203203102130-3321313030301131-3101330120330130-0233322213202102-3030221100013202-0201123221330203-3020102031213132"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az — sku_ergw2az / 020010130111 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az

<a id="canonical-2211000031012213-1213221021222212-3003030011121203-0220201001232200-2321330113101012-1330020312130010-2231123133200132-2021302203131003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku ergw2az.

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

<a id="canonical-0310330110331313-3100223331010122-3330030012213331-0030102132100130-2013220013122001-0330001133133310-3322323002130112-0332210112032100"></a>

## Direct properties — sku_ergw2az / 020010130111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113313223200221-1203322303332021-3030303233203220-2201330000321112-3010013020233010-1001201222312331-0121313003211001-0310220321003120"></a>

## Next pages — sku_ergw2az / 020010130111 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0211212312303130-0012322203100312-2011200011033121-3011100032023122-0113332120113011-2113302022010020-2201111300101233-1301011121211120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012001322002012-1021323220101003-0011323303201010-2323231123133030-2202333233200003-2200333123212331-1221002112102120-0333300032011201"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf — sku_high_perf / 010201301032 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf

<a id="canonical-1121211030011233-0123013302313301-3330321323102022-3222200031113110-0330110110032331-0301220003131131-3101121133222100-3321231312323112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku high perf.

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

<a id="canonical-1320313330313321-1210021020300321-1032211112322113-3122022212301302-0021102112212022-1201011211001213-0132313220120130-2203332030302102"></a>

## Direct properties — sku_high_perf / 010201301032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302110320222131-2323032321230303-2221320113313321-0110013320232021-3330102301113210-1112321322312003-0301211021103230-0001231032131001"></a>

## Next pages — sku_high_perf / 010201301032 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1210133331031012-3233002103231131-3133110300112012-2211120200022111-0321110333200100-0203120003322011-3312001001323113-2310132223020220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320010012223311-0033313030221130-0301010301002101-0030131100211331-2012320020232021-1312100311212222-2313223212321210-1031133323123022"></a>

## ingress_egress_gw_ar.hub.express_route_enabled.sku_standard — sku_standard / 203312132022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- ingress_egress_gw_ar.hub.express_route_enabled.sku_standard

<a id="canonical-2123111302003312-0011022033203223-1232131110113000-1003101010000333-0321100233101332-2300323230221302-3200330123331102-1002130101212203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for sku standard.

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

<a id="canonical-0300010123101112-3231312023233122-1210130301133123-2120122302312232-1322101000333233-1130100102120013-0132023320121311-2030301103013322"></a>

## Direct properties — sku_standard / 203312132022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1221232011311230-2222313032132131-1031132130200012-3203110203311121-0201012333003310-1113131303212313-2302202011103231-0233221113001222"></a>

## Next pages — sku_standard / 203312132022 / 4

- [ingress_egress_gw_ar.hub.express_route_enabled](data-sources--azure_vnet_site--reference--group-005.md#canonical-3122201231320123-3222202101013223-2131133300302133-0230333333200310-1123132213011313-0110330112322030-1023331310302130-2230032033121311)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121030133302333-3212212312303211-0231112130031211-1110300123313121-0322111232012320-2103332210221323-2211230020231321-0031031311001133"></a>

## ingress_egress_gw_ar.hub.spoke_vnets — spoke_vnets / 003000203113 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- ingress_egress_gw_ar.hub.spoke_vnets

<a id="canonical-1013300013312313-2100202233122223-2123110122113002-3013103130123221-2113213031333211-1230010300213313-0330210332102302-3311032233013220"></a>

Type: `"list"`. Computed.

Spoke VNet Peering (Legacy). Spoke VNet Peering.

Upstream description:

Spoke VNet Peering.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0200001011123200-2201100322013011-0031020312100120-3300200013213032-2301021102022333-0021221330230130-1132101331113120-0222213020100301"></a>

## Direct properties — spoke_vnets / 003000203113 / 3

- [auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-2220232201031212-2210210132202231-2122321033323122-3323211321321313-2320130002023233-2112010313221111-3300010121032123-0232230303100113): complete subsection reference.

- [labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-2011112002232222-1032012200102303-0331023233032320-3332223130333003-0003200131221000-3030211313322022-0300001231230211-3221201102223021): complete subsection reference.

- [manual](data-sources--azure_vnet_site--reference--group-006.md#canonical-0332031020033320-1313023330022021-2230131033110200-1130200300131210-1003021201322023-3230222331302031-0312112233122322-1020230223231130): complete subsection reference.

- [vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033): complete subsection reference.

<a id="canonical-1210110012033101-1120110200023212-0112231001201113-3111321101132333-0213123231331200-3232102310120203-0310323321330022-3032033202313231"></a>

## Next pages — spoke_vnets / 003000203113 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.auto](data-sources--azure_vnet_site--reference--group-006.md#canonical-2220232201031212-2210210132202231-2122321033323122-3323211321321313-2320130002023233-2112010313221111-3300010121032123-0232230303100113)
- [ingress_egress_gw_ar.hub.spoke_vnets.labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-2011112002232222-1032012200102303-0331023233032320-3332223130333003-0003200131221000-3030211313322022-0300001231230211-3221201102223021)
- [ingress_egress_gw_ar.hub.spoke_vnets.manual](data-sources--azure_vnet_site--reference--group-006.md#canonical-0332031020033320-1313023330022021-2230131033110200-1130200300131210-1003021201322023-3230222331302031-0312112233122322-1020230223231130)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2220232201031212-2210210132202231-2122321033323122-3323211321321313-2320130002023233-2112010313221111-3300010121032123-0232230303100113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310012220201230-3232330221033201-3220101301310102-0321110223111302-1133220003121211-3312001312022102-0310033103211231-0021020233103010"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.auto — auto / 002023031301 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- ingress_egress_gw_ar.hub.spoke_vnets.auto

<a id="canonical-2222333301313233-0303012012313221-1232210032000321-3010123121303030-1310303202011123-0030010320113100-2002213011021011-1100310031130013"></a>

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

<a id="canonical-3130003030330312-3022233330122231-2001331231100002-3101313301223111-0001113221130110-3332220001331222-2012102002203310-2133010122220110"></a>

## Direct properties — auto / 002023031301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301122300131212-3323223300113133-3102303323031301-0231112301001032-1233132213313303-0102122303022101-3102122310211311-2133213133323032"></a>

## Next pages — auto / 002023031301 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2011112002232222-1032012200102303-0331023233032320-3332223130333003-0003200131221000-3030211313322022-0300001231230211-3221201102223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011103231032301-3230231031221210-1013231123310130-0203012323012131-2210101032200231-2212010332033233-3130212201301331-1011203111111333"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.labels — labels / 121320221003 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- ingress_egress_gw_ar.hub.spoke_vnets.labels

<a id="canonical-2202011031022332-0223132333010002-0203102120231200-0011331223003130-0032131102012100-1120020102131210-3123313312122132-3321311131022202"></a>

Type: `"single"`. Computed.

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Upstream description:

Add Labels for each of the VNets peered with transit VNet, these labels can be used in firewall
policy These labels used must be from known key and label defined in shared namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0002332023323110-2020200213313321-1313101032203131-3121222013231303-2213003113003111-3132110212220021-1010323201320112-1022122011332021"></a>

## Direct properties — labels / 121320221003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113210213213201-2230202332031332-2130223223211320-1213303112331221-1122211100330102-3321123002200300-2333101212101013-2320332113333233"></a>

## Next pages — labels / 121320221003 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0332031020033320-1313023330022021-2230131033110200-1130200300131210-1003021201322023-3230222331302031-0312112233122322-1020230223231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232002323122131-3313223303233302-1233222222022000-3013010030221012-3210331310310013-2230113300211230-1330001123322300-0012303201310113"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.manual — manual / 112311302102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- ingress_egress_gw_ar.hub.spoke_vnets.manual

<a id="canonical-3103230203310031-3000132121311121-0000103001220020-3013121233330322-1011010322333031-3001301110330233-0202010021222233-3000122233202220"></a>

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

<a id="canonical-3103123312002212-0212020313222310-1311023321233202-1122012320310113-0021311230323110-3333120033110333-1211330210121030-1122200331310122"></a>

## Direct properties — manual / 112311302102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031331221013300-2100332311021013-0113020302132332-1031123300320022-3010331030321310-1302003120231210-1130223121131102-3022123210322320"></a>

## Next pages — manual / 112311302102 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213122220130311-3300310110011110-1331303113320311-2010312130302023-3011213110003111-3230032021111132-3303123030312121-1221330323121001"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet — vnet / 023000113303 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet

<a id="canonical-3230210230320100-3213023131202102-0123312110220320-2233132301133131-0112101103230110-3330201001311102-3020021322233103-0332011222222103"></a>

Type: `"single"`. Computed.

Resource group and name of existing Azure VNet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-routing_type": "[\"f5_orchestrated_routing\",\"manual_routing\"]"
}
```

<a id="canonical-2310233331200112-2300332103303222-3322231231103201-2001010121100320-0222200200333313-2030221011313310-3223321222121300-0021330322222311"></a>

## Direct properties — vnet / 023000113303 / 3

- [f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-0002201030102320-2021230102232102-0230332212323100-2303020033121233-1130313210313110-1033312031210331-2330132321221231-3032110120113023): complete subsection reference.

- [manual_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-3231021013221130-1022003103220202-3200321031123123-2303100312000122-2002120003033322-2012022122333301-2303311303231202-0330220101303333): complete subsection reference.

<a id="canonical-3322333010021302-1311211312103100-2012333333200130-1132133100221012-0330131311111123-3021323131231113-0112302131303320-0112031023222001"></a>

<a id="canonical-3331112000210230-2303313331310300-2222203113003030-2020300133030013-2102310301133212-2000033323103110-0221210301321233-3110220013201331"></a>

## resource_group property — vnet / 023000113303 / 4

Type: `"string"`. Computed.

Existing VNet Resource Group. Resource group of existing VNet.

Upstream description:

Resource group of existing VNet.

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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2130301322001122-1312222203332313-3110121313200123-2231321301220231-1121101111320130-0011123121000020-1203332031033303-1213023202011013"></a>

<a id="canonical-1313300011311123-1300313021023323-0123110012103121-1101112010031310-1110110110120211-2110011230012202-3203323100000322-1002001202200133"></a>

## vnet_name property — vnet / 023000113303 / 5

Type: `"string"`. Computed.

Existing VNet Name. Name of existing VNet.

Upstream description:

Name of existing VNet.

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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0210300210323302-1001202323233100-3310211302112302-0230220013111202-3112201133103231-0000022121000030-2013122033310101-3332311020101333"></a>

## Next pages — vnet / 023000113303 / 6

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-0002201030102320-2021230102232102-0230332212323100-2303020033121233-1130313210313110-1033312031210331-2330132321221231-3032110120113023)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](data-sources--azure_vnet_site--reference--group-006.md#canonical-3231021013221130-1022003103220202-3200321031123123-2303100312000122-2002120003033322-2012022122333301-2303311303231202-0330220101303333)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0002201030102320-2021230102232102-0230332212323100-2303020033121233-1130313210313110-1033312031210331-2330132321221231-3032110120113023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323023220223030-1031110023023210-2011111111011121-2113120103000330-1120313112321120-0212032120103022-2001011220030212-3001022202023121"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing — f5_orchestrated_routing / 112302020202 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing

<a id="canonical-0031130330201220-1310220320110011-2313010021213321-0310012303313103-3210103020022203-2030102101113232-1312221120222313-1233302120333323"></a>

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

<a id="canonical-0122223003001301-2000012132321211-1103111133310313-2031112220121200-1322331133100112-1203220231233232-0121033230201021-2012101301301330"></a>

## Direct properties — f5_orchestrated_routing / 112302020202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0103030222301003-2003320102023231-0100023023011022-0311210313301030-3130312132012022-3223003031312202-2123022301002231-1213131101232311"></a>

## Next pages — f5_orchestrated_routing / 112302020202 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3231021013221130-1022003103220202-3200321031123123-2303100312000122-2002120003033322-2012022122333301-2303311303231202-0330220101303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121121330300221-1022031221020110-3133032221121110-2000002120323321-1223032313330220-0030303121121321-2222132100320032-3032102311031213"></a>

## ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing — manual_routing / 020311330112 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.hub](data-sources--azure_vnet_site--reference--group-005.md#canonical-1021313213130313-2200023001301030-1132010102322300-0132313200212233-2120131322010032-3023020130010331-2221113011130330-2232022233133012)
- [ingress_egress_gw_ar.hub.spoke_vnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-1321200130322100-3313002220022201-0333322100332302-1113313222130012-1302300303031103-2001133133021032-1033323303321001-2233310320222011)
- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033)
- ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing

<a id="canonical-3312222112201230-0312120022201130-1313011221111012-2320102310120013-2003313010131122-1202302310201210-2122001002010212-3230103010330333"></a>

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

<a id="canonical-0223200202232030-2133323001232203-2203333032032233-1022122010223133-3032230233312122-2122120123302031-1231313032113132-1213332322012311"></a>

## Direct properties — manual_routing / 020311330112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130130002222213-2102301012320100-3133111301103210-2210223210211313-3233000203130001-0013331222030030-1301010033131232-0102333211222030"></a>

## Next pages — manual_routing / 020311330112 / 4

- [ingress_egress_gw_ar.hub.spoke_vnets.vnet](data-sources--azure_vnet_site--reference--group-006.md#canonical-1201103302120203-0020212312203010-2323013103220101-2112010122112122-1002222102312110-0113100220201103-0011011233233220-1212033011222033)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102122321111313-2132203201120230-2320012333300130-1003232130103032-1132031033232321-3021230110130232-1013101203121202-1223111131213313"></a>

## ingress_egress_gw_ar.inside_static_routes — inside_static_routes / 012220201330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.inside_static_routes

<a id="canonical-3132123121231322-0221230023311323-1320220301121131-1133203013220021-2223121302120031-0112132213222312-0013131000103102-0102311300120312"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Upstream description:

List of static routes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3321130010331003-2310102020321201-2311322131230210-3323322013300131-3311221001101223-3000012310312112-2022112100223330-1221213230332002"></a>

## Direct properties — inside_static_routes / 012220201330 / 3

- [static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001): complete subsection reference.

<a id="canonical-0022230321332200-1032103301132101-1032221012333303-3111123103021320-0221120031331122-2232003221022110-0210323130131200-2232110100112122"></a>

## Next pages — inside_static_routes / 012220201330 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222222133122232-0112223302323033-3030203223313111-3013112012213231-2000110002231023-3120220020132123-3313032131113032-2123000033200221"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list — static_route_list / 120013313330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- ingress_egress_gw_ar.inside_static_routes.static_route_list

<a id="canonical-2202201100213332-3232322230012020-2332022222030302-2033212333132202-0100012030332301-0033233113013222-0023032121332203-2310120023121310"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

Upstream description:

List of Static routes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-0300132300101302-0313000211303213-2321103333222302-3030333323100303-3231313013320013-3101332321110222-3220313322312211-3311202023220012"></a>

## Direct properties — static_route_list / 120013313330 / 3

- [custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023): complete subsection reference.

<a id="canonical-3132233203121101-3312030123021230-0300011310230022-2121220033112020-1230100132203002-1103032121202000-3011200221111333-3313232021113122"></a>

<a id="canonical-2011331001102233-2221102232131101-0103320032311332-3213331121221000-2100110132020122-1020012313113310-0123113211203112-1102321010011132"></a>

## simple_static_route property — static_route_list / 120013313330 / 4

Type: `"string"`. Computed.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

Upstream description:

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0020212210233102-3003130121212023-2201310323322023-0032322131032210-0331331002122231-2002312230211330-1331010303202300-3132201203313231"></a>

## Next pages — static_route_list / 120013313330 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000102113130130-3010030211222203-2332223322021232-3303230231221123-0131021122023313-2330121230002030-3301230232132312-0023031302220123"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route — custom_static_route / 012331333333 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-0112213121220001-1301022322130002-3132001313323120-1221312121233302-2332303130232003-1101031221120231-3200120223012301-0312301233013231"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3320030310031120-1313213000221321-2030320011103320-2222330220012020-0321331323331023-1322220012113220-0100210200303101-3012323103213230"></a>

## Direct properties — custom_static_route / 012331333333 / 3

<a id="canonical-3230102221101102-2000330221010202-0022010021301322-3011000310323020-2101322303122222-3001333110222131-2330032313301322-2020222022222313"></a>

<a id="canonical-1301031211023033-0103222203311103-2323131230331321-0022030210122033-2310311101020212-2220030032220331-1311312100231013-0011212022320201"></a>

## attrs property — custom_static_route / 012331333333 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of route attributes associated with the static route.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111303012333120-1003003311330111-3311021322112321-1122303020003100-0112200230332333-3221021232223033-1232123003012333-1200131310010111): complete subsection reference.

- [nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213): complete subsection reference.

- [subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232): complete subsection reference.

<a id="canonical-2102212123333302-2101003110233000-3310210031322021-3123023121232201-3312222112011010-2110220130310032-0310210202330011-0010333330101122"></a>

## Next pages — custom_static_route / 012331333333 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](data-sources--azure_vnet_site--reference--group-006.md#canonical-3111303012333120-1003003311330111-3311021322112321-1122303020003100-0112200230332333-3221021232223033-1232123003012333-1200131310010111)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3111303012333120-1003003311330111-3311021322112321-1122303020003100-0112200230332333-3221021232223033-1232123003012333-1200131310010111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202231313330103-3301322233020013-0333322203333310-1200312321131322-2313003202220132-3333333030122320-3232111111320211-2112133212131320"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels — labels / 310012013100 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1303200201120302-1102032003132122-1100133002110300-1121230213203330-2332020133330021-0201332000122330-1320033001000303-0112231032112221"></a>

Type: `"single"`. Computed.

Add Labels for this Static Route, these labels can be used in network policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301233131233221-0033110211123331-0210120022222221-0210223012112311-3032011012300203-0200220321030022-1102003000300332-2201032232030001"></a>

## Direct properties — labels / 310012013100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110320122010032-3313230300100320-2202331331121012-0123110200230333-0100113023102013-2033301210130112-0113201210320321-0101221022130223"></a>

## Next pages — labels / 310012013100 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123332131123203-2020030013321130-0101013321200021-2203013111021132-1013132103321031-2323202122220332-1011211333203012-1022201231102001"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop — nexthop / 201002121330 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-1202302333211213-2012013333231130-3203032112302022-0121221200320123-3002321220231320-3102331200332332-0020020310203012-0032300323310323"></a>

Type: `"single"`. Computed.

Nexthop. Identifies the next-hop for a route.

Upstream description:

Identifies the next-hop for a route.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1311130313100212-3303300310311313-2030132233002031-1211030211002121-1332013011321201-1101202211133020-1103220222113311-3120332023210132"></a>

## Direct properties — nexthop / 201002121330 / 3

- [interface](data-sources--azure_vnet_site--reference--group-006.md#canonical-3021233300313120-0313320031111211-1322310330030022-2113220232012033-1032100311310031-0222033330202012-0000011323012101-2032233303002300): complete subsection reference.

- [nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333): complete subsection reference.

<a id="canonical-1012212303002011-0200131320100202-1120133021201222-1123332321103022-0123112033030023-3303202233032112-2112301131110000-0032303222310020"></a>

<a id="canonical-0333220220201232-2232200122322211-2313323120103100-1311200303031133-3213220210320102-1003313213203101-3122212110122331-1122130112223121"></a>

## type property — nexthop / 201002121330 / 4

Type: `"string"`. Computed.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Upstream description:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Assumes there is only one local
interface on the virtual network. Use the specified address as nexthop Use the network interface as
nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN private virtual network.

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1320310033221131-1233122303202322-3231213102322110-0133100202300330-1221303202300233-3213322033121111-1303010100230002-0210033013303231"></a>

## Next pages — nexthop / 201002121330 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](data-sources--azure_vnet_site--reference--group-006.md#canonical-3021233300313120-0313320031111211-1322310330030022-2113220232012033-1032100311310031-0222033330202012-0000011323012101-2032233303002300)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3021233300313120-0313320031111211-1322310330030022-2113220232012033-1032100311310031-0222033330202012-0000011323012101-2032233303002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022101112321123-1312210231321231-3032020321001233-0302313123121130-2011221120230311-3010222302312121-3330103320133022-2021013030023110"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface — interface / 101113210333 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-1012023130010201-3230030110102001-3100323201132222-2311022222021300-0021021132300322-0120311301322021-3323122130021101-3223030023002213"></a>

Type: `"list"`. Computed.

Nexthop is network interface when type is 'Network-Interface'.

Upstream description:

Nexthop is network interface when type is "Network-Interface"

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

<a id="canonical-2331233211232301-1221001320300021-1110133023112202-2320112130300131-0033103211113220-3323103132123230-0102301101211301-0331233033032203"></a>

## Direct properties — interface / 101113210333 / 3

<a id="canonical-2020100221100101-1310132201101000-1102022320212210-3220110103121212-3000323203023332-0112323031323020-1220001031200030-1202220022032331"></a>

<a id="canonical-2222302103200203-1123003110111102-0200201310102210-2122100220211200-3111112000132133-1132302313202322-2020203322322020-2212313111121111"></a>

## kind property — interface / 101113210333 / 4

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

<a id="canonical-1120131031030230-1323323313210320-3200232233302221-3320312010320032-0322330201221231-2111232133102210-2122031030330300-3000111222012101"></a>

<a id="canonical-1310323122131121-3311011301213320-2311223200302322-3212330031301000-1311321033212303-0133211032021101-0312302022223332-1220210210220131"></a>

## name property — interface / 101113210333 / 5

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

<a id="canonical-3020123133202022-2131233102330002-3232321133321100-3323333213120013-1320200322202312-0333002302200221-2111120112322033-2132311110023311"></a>

<a id="canonical-3220321213220013-0220113012130021-1300112200333333-2033201212020221-1033003120030122-3211022023330101-0113031103301330-1310231120002023"></a>

## namespace property — interface / 101113210333 / 6

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

<a id="canonical-1311210102321210-0331232033230021-3020330113330231-0220010220020232-0301311300123331-1230320010122330-1201122020330321-3131223201312111"></a>

<a id="canonical-1311113012313331-2301210130131003-2303030333211122-0023033033312313-2131322110300230-2030002110130123-0201300012213301-0113121010121123"></a>

## tenant property — interface / 101113210333 / 7

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

<a id="canonical-3032301023001113-3020221202232332-2013312011113220-2001033230010223-0200203310221303-2310010321332021-0111030022320231-2310023131220023"></a>

<a id="canonical-2203311112302112-0013203323222202-2211312321310020-3132002210022031-1300000332122223-2032303211102020-0030311011232002-2233131322203000"></a>

## uid property — interface / 101113210333 / 8

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

<a id="canonical-2130120121331210-2220212101232130-2100300332020200-1303121322121100-1030313221001300-1031333231113002-3200113321122230-3211121103120003"></a>

## Next pages — interface / 101113210333 / 9

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021232113113331-1021120021323230-3123321022013133-2101101313211200-1222311022113020-3233030211121211-1320302220102112-3022131203133010"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address — nexthop_address / 213101233331 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-2233212103211032-1333222233300213-2320112323002030-2332321323331212-0120212331220302-0132232003000021-3221122321131222-2133311222301233"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

<a id="canonical-0032332203213132-3223202030333031-2110020223113333-3103022201023311-0000303222220113-3101022311303123-1021103202331000-2211221122323330"></a>

## Direct properties — nexthop_address / 213101233331 / 3

- [dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333): complete subsection reference.

- [IPv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-1003320010300300-0312330001122130-3222330330000333-3200131102322132-2033011212311303-2301103130223222-3122122012222033-0310000223123233): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-0003222022131223-2122020231002121-2312201213222133-1021110010120211-0012321021220022-2021300001330021-2330003321212112-2100022231031210): complete subsection reference.

<a id="canonical-2003302100332310-1231330321312221-2011221311030222-1312200230313102-1330001010033000-3220113010111323-3202211330222011-1131121112031220"></a>

## Next pages — nexthop_address / 213101233331 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-1003320010300300-0312330001122130-3222330330000333-3200131102322132-2033011212311303-2301103130223222-3122122012222033-0310000223123233)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-0003222022131223-2122020231002121-2312201213222133-1021110010120211-0012321021220022-2021300001330021-2330003321212112-2100022231031210)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221301313100010-3000322102110130-3222031302221120-1113123113102110-1312202321010221-1130310011020320-1011131231323002-1220133210013001"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack — dual_stack / 113210021131 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3131311023321302-2323100033000003-3302100113012130-3021030021332102-0002031003100032-1111333313313203-0230000333313030-0300103330033102"></a>

Type: `"single"`. Computed.

DualStackAddressType represents both IPv4 and IPv6 together.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132320112001000-0212210301131233-0103133202023123-2023213020333201-2301201303332132-3210010032113020-2300023013103310-1231312222231010"></a>

## Direct properties — dual_stack / 113210021131 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-1100312112131310-1020032103013013-1033321120331233-2023233022111112-2230323013102130-2011332300010222-2103030122001332-3122323002101123): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-0113223323333233-0133012120311202-0130312211120333-3112320103113233-2011313221101332-0101021311033320-0123312111110002-2110313120232003): complete subsection reference.

<a id="canonical-3103132033311321-2100012003331330-3312123213210200-3012112001201102-1021221012231332-2000013132133231-2300110213021222-0111320201303220"></a>

## Next pages — dual_stack / 113210021131 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-1100312112131310-1020032103013013-1033321120331233-2023233022111112-2230323013102130-2011332300010222-2103030122001332-3122323002101123)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-0113223323333233-0133012120311202-0130312211120333-3112320103113233-2011313221101332-0101021311033320-0123312111110002-2110313120232003)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1100312112131310-1020032103013013-1033321120331233-2023233022111112-2230323013102130-2011332300010222-2103030122001332-3122323002101123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012010011202110-0000101311102203-0000302310222310-2123012222031123-2312313331020021-0031130230320012-2020131321232132-3320032113102310"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4 — IPv4 / 121232201111 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3102300303033011-1120220232231013-2103102131131213-3323310210120333-0221323030211101-2320031112100122-3131101222133031-1312231322210301"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2313333221020033-2230212301232011-3223113033001112-1323323001201223-1210300012232010-1323103011133222-0230131013112032-2122023311033302"></a>

## Direct properties — IPv4 / 121232201111 / 3

<a id="canonical-1013232101223110-0011321101013301-2203101313230012-1210132000111332-0102231000120102-2103002311233302-1002000101002220-2103100331310302"></a>

<a id="canonical-0023300213022300-0330311120122122-1231213311310231-0110121031332010-1330210303231202-1120033313313130-1020000321213132-3230322131310300"></a>

## addr property — IPv4 / 121232201111 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0310131311032023-3300321033031003-1011100322003100-3111030123102013-2011130303133223-3123233312220030-1220203101220211-3021103101111121"></a>

## Next pages — IPv4 / 121232201111 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0113223323333233-0133012120311202-0130312211120333-3112320103113233-2011313221101332-0101021311033320-0123312111110002-2110313120232003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221200021111301-0232220032030130-1212202200132103-0300130311133322-0100232022012012-0011313112103113-1221210311231312-2013131321320113"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6 — IPv6 / 201200012312 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-0310302313200201-1103320100102232-0222213112033032-3221101303221101-3300133313212222-1020300203032303-2300210012110221-1123313123321213"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0110122000112323-1010331320100222-0033123223202122-0111233113312231-2030031100221021-2021012103000311-3303221100132330-2001323212030302"></a>

## Direct properties — IPv6 / 201200012312 / 3

<a id="canonical-0010213310303210-3330131223133221-1022200312221102-0322313132313020-0202202220000303-1222020031031121-1013231022102311-0231323023231111"></a>

<a id="canonical-2133031300221231-0333313321311031-0302011030233302-3112113133210310-3002111331101330-3012232330332302-1323031302331113-1230212113213121"></a>

## addr property — IPv6 / 201200012312 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-2110100001001110-2212023012021333-1231132212022331-0131210230100033-3322332233032102-0323122221111301-3202032330303220-1003012121031011"></a>

## Next pages — IPv6 / 201200012312 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--azure_vnet_site--reference--group-006.md#canonical-2010113223313322-3022112220202002-3130332313232101-3330122311232203-1300103011231033-0111313031132303-2001302121302222-0121301300330333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1003320010300300-0312330001122130-3222330330000333-3200131102322132-2033011212311303-2301103130223222-3122122012222033-0310000223123233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220221301201012-1131332321232231-1231311300100323-0110303022221103-2231200201000323-1022101021230311-2302132230201001-2211212113130200"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4 — IPv4 / 010313121123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1332220230202022-0323010110030200-0231321300132312-0302230022003130-3121213203101232-2020222013012011-2321032013130030-0213112223320213"></a>

Type: `"single"`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Upstream description:

IPv4 Address in dot-decimal notation.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1220312311030132-3333011322310123-3303311110323233-1103011212102211-2101332003211222-0101231233311202-1103100101201213-1130012001332333"></a>

## Direct properties — IPv4 / 010313121123 / 3

<a id="canonical-2001132223233032-2123123202111110-0123102113201233-3212021030100012-1321011311100221-1023330331013301-3102123200013230-1331020313010120"></a>

<a id="canonical-1001102223231323-2123121221133022-0131323233031011-1231120101030203-0130112031133300-1122133201232023-2322022133032021-1312002022320222"></a>

## addr property — IPv4 / 010313121123 / 4

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0223311112323120-3022111233321032-0321110323200010-3023103231323222-3133001232201122-0130311213130311-0203122230103030-1231020231222122"></a>

## Next pages — IPv4 / 010313121123 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0003222022131223-2122020231002121-2312201213222133-1021110010120211-0012321021220022-2021300001330021-2330003321212112-2100022231031210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022010103132100-0223223101031211-2110100121023200-0312103213331031-2012332220130000-2100001231000103-0310221132110002-2020113111332100"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6 — IPv6 / 113331211220 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--azure_vnet_site--reference--group-006.md#canonical-0330213221330200-2111310303221331-3221313233031220-2120210302133113-2000200303210102-2131211302321223-2332201223210203-1312211000323213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-1313011231232103-0200033320122211-0303122101001123-2303120220201333-1013220312031203-3201311330100010-3231332002113330-0103310032202103"></a>

Type: `"single"`. Computed.

IPv6 Address specified as hexadecimal numbers separated by ':'.

Upstream description:

IPv6 Address specified as hexadecimal numbers separated by ':'

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012002200000110-1322100212233032-1302330331113033-0210200212112223-2001231021022300-3333102121023223-0303031313210033-3232211332003120"></a>

## Direct properties — IPv6 / 113331211220 / 3

<a id="canonical-3113321320113332-2112201010122303-3022210120030331-0012322103213213-0332121232232012-0101010232022001-2133320003001221-0331021120031211"></a>

<a id="canonical-3332202021000001-2310133023312211-0133321232011030-3301211021101101-0020332230002321-1110331330202122-1003000033222232-1002031002000322"></a>

## addr property — IPv6 / 113331211220 / 4

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

Upstream description:

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0323210310221130-0323013310001201-3200130223033002-3311123011211032-3313220313022003-2101130133113131-3321010030101131-2300310302110133"></a>

## Next pages — IPv6 / 113331211220 / 5

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--azure_vnet_site--reference--group-006.md#canonical-2113110103012102-2022213033032223-2222021221021132-3020030120003320-1201212231103103-0332130021031020-3301203223020002-2102012310303333)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120211231110112-1311021113303133-0201311211131121-0120112310301333-0013120222302013-1122011110003033-2131320022113201-2322301232332133"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets — subnets / 202233230123 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3320211322300321-2213320010010202-1333311032332332-2322133310213013-3331300010211112-3002000032211120-2133011211231123-3312302231332220"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

Upstream description:

List of route prefixes.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

<a id="canonical-3200020011311322-1123221121101223-1312210223122132-1020330203212020-2312323223223213-3230011230102313-2200203100302031-1202210021010133"></a>

## Direct properties — subnets / 202233230123 / 3

- [IPv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-0203313330300223-3230133200313132-2120310023310003-3001202011212201-3001302333210210-3121221213111300-1332321001132312-3201011303312130): complete subsection reference.

- [IPv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-1320023212221200-2322221122103232-3010123100320220-3010111310230213-3313233030210003-1111120221311321-3221320121132012-0312102312002220): complete subsection reference.

<a id="canonical-1100130322203222-3002211233300130-2230220132113330-2332201203312232-1122302333011222-3133322302023021-2310003122220033-3101313033201321"></a>

## Next pages — subnets / 202233230123 / 4

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](data-sources--azure_vnet_site--reference--group-006.md#canonical-0203313330300223-3230133200313132-2120310023310003-3001202011212201-3001302333210210-3121221213111300-1332321001132312-3201011303312130)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](data-sources--azure_vnet_site--reference--group-006.md#canonical-1320023212221200-2322221122103232-3010123100320220-3010111310230213-3313233030210003-1111120221311321-3221320121132012-0312102312002220)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0203313330300223-3230133200313132-2120310023310003-3001202011212201-3001302333210210-3121221213111300-1332321001132312-3201011303312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333223030122301-2130233313021120-1110112233212122-1000230010100202-3133120212203013-0230331212122322-1300302211320031-3012331331010220"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4 — IPv4 / 303202332022 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1000021110133301-0312120312011330-2303032100023130-0233110103201300-3023031110233220-3311212200133221-0121211210323310-1220230212231320"></a>

Type: `"single"`. Computed.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2211021221000100-0302023323130213-3122022010010202-0020222233221131-0231030033133021-3332022121212132-2020121003312330-0330101120333201"></a>

## Direct properties — IPv4 / 303202332022 / 3

<a id="canonical-1111010303210112-0013322033203222-0211013000113201-3132030031323102-2221222233001122-3111320103000333-3100101332210200-0113322032030110"></a>

<a id="canonical-3223321310203302-0202120023312201-1310101101220020-2011300201112003-2312313112320032-3221102031113233-1300233020100323-0213221302102301"></a>

## plen property — IPv4 / 303202332022 / 4

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1302222230231103-2101320311200132-2031213002102201-3323012323031002-0201012130201011-2303323231012101-2330313320003203-0131203001030233"></a>

<a id="canonical-3031231220230331-1201203210001220-1203223203213003-1323333222020031-2003003233103102-3101200032331330-1330120233331321-3132210201022322"></a>

## prefix property — IPv4 / 303202332022 / 5

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1211332032303122-0010320002023100-2201323113313132-3330120032000120-2133032203202132-1112112211233111-1231111312333210-3231203031012212"></a>

## Next pages — IPv4 / 303202332022 / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1320023212221200-2322221122103232-3010123100320220-3010111310230213-3313233030210003-1111120221311321-3221320121132012-0312102312002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1201322001211101-1313000131231012-0221320330000003-0002013320323112-1210322303020011-0010123210023222-3312000001233110-0123311221310123"></a>

## ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6 — IPv6 / 221303300133 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [ingress_egress_gw_ar.inside_static_routes](data-sources--azure_vnet_site--reference--group-006.md#canonical-2133110020230102-0032002222232002-3013210032132211-0323223100202131-1123000022223021-0123212032312032-0310130123232213-2321012113221213)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list](data-sources--azure_vnet_site--reference--group-006.md#canonical-0202000311113213-1113222130133221-3012000201230011-2223000331122202-3231100301023201-2303001310102203-2210313213303321-3302022303110001)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](data-sources--azure_vnet_site--reference--group-006.md#canonical-3031233022110112-0233231213001223-2020333223210302-3121133020210200-2020321033202223-1310121123202023-3323023202133131-0222123310211023)
- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232)
- ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-0221322222121121-1331300320212120-1230201331330320-3103333322003123-3102130033333302-2031132231320031-3200010132012332-1300002232322112"></a>

Type: `"single"`. Computed.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2010113212112022-3101210321102331-3311003202133132-3110333123333023-2002030220112112-0012100002003100-1032313130221300-3131232000132231"></a>

## Direct properties — IPv6 / 221303300133 / 3

<a id="canonical-3213220000311313-3212302011130211-2123021202131231-1132233221111102-0022300232030130-2333220330333110-2032121101200112-0011112103001233"></a>

<a id="canonical-3331110330231031-3323332033112312-2213210301321111-0003101300033331-3011101020301003-0321213131021123-0333123331231111-1131200313113003"></a>

## plen property — IPv6 / 221303300133 / 4

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1120220002101230-1022330010123301-3001100200213220-0310320201310010-1210123332002233-1333331121233012-2232103300121323-0013322020103331"></a>

<a id="canonical-0021130122222023-3120213111223201-1102022013010120-3022103101221022-2111023101012012-3320000020221020-2310011002221211-0031022311223101"></a>

## prefix property — IPv6 / 221303300133 / 5

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Upstream description:

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0" The address can be compacted by
suppressing zeros e.g. "2001:db8::2::"

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-3230011133131022-2131323123332330-0312103210102200-0030300000310202-0100100211111022-3300120032221222-2333200322323212-1200310121312221"></a>

## Next pages — IPv6 / 221303300133 / 6

- [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--azure_vnet_site--reference--group-006.md#canonical-2222103232000102-0100200011013123-1212111303101010-3201132013312023-1312303200013000-3210202110330200-0333030122232013-3112310112331232)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0122313011113033-1032211210323022-2001323122102232-2100030323111021-1320002130003131-0232100113332232-3211201112331302-3113210102210312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311321313101101-0033030132131003-3333203132100201-3300330132023223-2032220111201223-2111013130302113-0120230121203031-2112112233300020"></a>

## ingress_egress_gw_ar.no_dc_cluster_group — no_dc_cluster_group / 133211302102 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_dc_cluster_group

<a id="canonical-2231212021223322-0331211002300201-3333221203031023-0222300231132103-2132231320232212-1223323100111301-3311233123322113-0331302110213020"></a>

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

<a id="canonical-0211321003232020-1332200312203022-3102311020113303-2210130110103010-3033220120301001-2013001021302202-1220023023123212-3012212310213201"></a>

## Direct properties — no_dc_cluster_group / 133211302102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110312333013320-3000003103131032-1032232330311010-1320301302313203-0110203323100012-1333320331030323-2202320201012202-0202132002333003"></a>

## Next pages — no_dc_cluster_group / 133211302102 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2101001023122203-0033102212111023-3321233031301033-0112330111120123-3303230032203232-1311222010033233-0330013132102111-2230120032221332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330111010021322-2222031023002132-1121213033300213-3332110113203000-1131321331111031-3233211200311330-1130013310211331-1313233213003022"></a>

## ingress_egress_gw_ar.no_forward_proxy — no_forward_proxy / 013313000332 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_forward_proxy

<a id="canonical-3301302102033221-3300201130013122-1033133200223323-1023330101122103-3002133010123233-1231301121302002-0301022013121033-2001020032012020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-3213030132201301-1210202033333301-1332331323321032-0132011330303101-2023212100012311-0121113132301333-1013011103322221-2131102113332332"></a>

## Direct properties — no_forward_proxy / 013313000332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322002311310022-0330332003223301-0033223333332311-3322122030012010-1200202332322300-2221323031322220-3020210123201111-0333123003220233"></a>

## Next pages — no_forward_proxy / 013313000332 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2232112110113101-2323121302332100-1211003200203012-2020123331102113-2011101102020333-0332203211013120-0210311311132122-3213310233001310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331322233023320-3321331033330210-0202013203233123-0001331123001002-2032102032010131-1213022331211230-3122011003233320-1133112130331033"></a>

## ingress_egress_gw_ar.no_global_network — no_global_network / 211033022021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_global_network

<a id="canonical-1100330110013123-1032312311022313-2121303321301020-0203322230330032-0300203312000312-2033222221102222-0301013230313113-2231210312333221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-0300230020202111-0131130013320101-0023000313123223-3221102103203303-0230231132311301-1112310020033000-1231231001122322-0321330030013311"></a>

## Direct properties — no_global_network / 211033022021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113003323011212-2322130132121222-3213122111212213-0333023232031233-2233202023030310-3201202033320330-3333003000213302-2332031000122312"></a>

## Next pages — no_global_network / 211033022021 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-2301333331223003-0020022320230012-2200002031110230-1000130230302230-0222210021000323-0103131213302220-2033010212231030-2212010320212120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130231213031130-1111011113223221-2122022023221203-2012331311103021-1002103110020200-3300002200001003-0001303232002113-2031301032130012"></a>

## ingress_egress_gw_ar.no_inside_static_routes — no_inside_static_routes / 030303001302 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_inside_static_routes

<a id="canonical-1332310002222103-3202202001301011-1323213200120312-0022032210303323-3211301010133311-0203020311232101-3112030212302201-1300310122323011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no inside static routes.

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

<a id="canonical-3111031032121103-1010210133012333-3331320013322013-3330330200320133-3011310112112232-3310030201323210-0121121031221132-1310130133300223"></a>

## Direct properties — no_inside_static_routes / 030303001302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123003303032223-2302002310220222-1130002222223200-1110121333322103-2212301223201212-2221333331220022-1002113232002001-0332110201003322"></a>

## Next pages — no_inside_static_routes / 030303001302 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-3331312201030030-2102232110000103-2312000101122210-2323232200100100-3212200013122310-2210133012313013-0001233030022123-2323012323202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303002311321322-1321311203001233-3222000111211302-1010113311123022-2303301023102030-0012230212300223-3201330201100101-1333013313320331"></a>

## ingress_egress_gw_ar.no_network_policy — no_network_policy / 020133013021 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_network_policy

<a id="canonical-1232220001201003-0012001220230010-1121100302022100-0330331300000102-0313201220203121-1333022202230302-2221231112100133-1213133033130231"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-0303123311233233-3122310110132022-3111230011230333-1032212001313333-0122100213211031-2003202123310230-2023122312220132-2333123000332130"></a>

## Direct properties — no_network_policy / 020133013021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320210133212330-1320321022203320-2113333230011301-3113130130021023-3203021313012132-0301002203123202-2322230210300023-3202221102111313"></a>

## Next pages — no_network_policy / 020133013021 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-1222031321110021-3122213001323211-2212013221020221-3323033000300023-1202211110032303-2200030231102133-3311202032023120-0233210203332303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102011333030313-1011110232202032-2323330300003201-0033321302321012-2333023231032312-1333222211231201-0331133312203223-1002133311033023"></a>

## ingress_egress_gw_ar.no_outside_static_routes — no_outside_static_routes / 311301031333 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.no_outside_static_routes

<a id="canonical-3011302012233033-1313332031301303-3302321230333212-1023231211130223-3231321020100013-1021232001300213-1331330112310231-2013321030100232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no outside static routes.

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

<a id="canonical-3023210023032023-0110013131132120-2211103302323003-0320113320301123-0121201212231223-3222300213030101-2011313030132313-2321123101230100"></a>

## Direct properties — no_outside_static_routes / 311301031333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111310303211232-2133223011010331-0221121002222330-3233232303300020-2031020030203010-1312313010012220-1221022110103303-2111123002323120"></a>

## Next pages — no_outside_static_routes / 311301031333 / 4

- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)

<a id="canonical-0301332221030113-2231320220003100-0121223000301012-0100011110302031-0120322021303223-1103033132113331-0130013200110001-1200020003302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020200131012103-3300131033332031-2202233331302300-1033303001312123-3003220022302130-1020211023330131-0011232100202321-3211321202031133"></a>

## ingress_egress_gw_ar.node — node / 310311313322 / 2

Breadcrumbs:

- [xcsh_azure_vnet_site](../data-sources/azure_vnet_site.md#canonical-1330232102102121-2311132211103111-1010023103323200-1323210323002133-2233321031322213-1333221233013013-0332233233111031-2332010010210310)
- [Property reference](data-sources--azure_vnet_site--reference--group-001.md#canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113)
- [ingress_egress_gw_ar](data-sources--azure_vnet_site--reference--group-005.md#canonical-2301120333003113-2003322233223332-3022013100223130-0210232002323221-3330102133112033-3332121202112233-3321203100031002-1120211021323111)
- ingress_egress_gw_ar.node

<a id="canonical-0211300103330102-3313332103002230-2321111013103133-0213223133122300-2003132010103013-1000130122010133-1132021301012131-2020131122001301"></a>

Type: `"single"`. Computed.

Parameters for creating two interface Node in one AZ.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3332032102222110-1220013132002312-2232201232012211-0033000303030313-1232120000032002-1113210300111033-2321121031122113-0312321300323322"></a>

## Direct properties — node / 310311313322 / 3

<a id="canonical-2200123010211122-0332130232112033-3002323210113233-3330130023031133-1232312222013220-3022033013321112-2233113213213332-2022322223310312"></a>

<a id="canonical-2233302323023131-3231210321303032-0133102132011322-2132223321110321-2132120033010223-3231200222220301-3331323011222302-1011003113231113"></a>

## fault_domain property — node / 310311313322 / 4

Type: `"number"`. Computed.

Namuber of fault domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3,
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
    "ves.io.schema.rules.uint32.lte": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "3"
  }
}
```

- [inside_subnet](data-sources--azure_vnet_site--reference--group-007.md#canonical-3330030013323030-2230320232322330-0111110010112011-2103000333233233-1022000330022023-1032102021023223-1010012012322133-1001022100211133): complete subsection reference.

<a id="canonical-2301311211131101-1212312103331032-1111302012332332-0123001323223211-2110123300123201-3022133111032213-0112332331330013-2230111030332201"></a>

<a id="canonical-0332121121002221-1213232010202110-2133002001110300-1122333302210211-3011022120031212-1002002321122111-3133030330110221-2000103211102321"></a>

## node_number property — node / 310311313322 / 5

Type: `"number"`. Computed.

Number of main nodes to create, either 1 or 3.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.in": "[1,3]"
  }
}
```

- [outside_subnet](data-sources--azure_vnet_site--reference--group-007.md#canonical-1232231023232320-2220202001233120-1020101121031232-2013233130313322-2313013003103001-3302032001302000-3010220011322121-0022130110333120): complete subsection reference.

<a id="canonical-1220300200102021-3021113232310222-2021312120310133-1212100333033002-0130112011220202-0101022213212233-1200133312221203-1123211302013032"></a>

<a id="canonical-0003211013212320-0003320212123013-0220223333232103-3233222111133213-1131021031301121-0212322221003211-3231211312313203-1220301110120130"></a>

## update_domain property — node / 310311313322 / 6

Type: `"number"`. Computed.

Namuber of update domains to be used while creating the availability set.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```
