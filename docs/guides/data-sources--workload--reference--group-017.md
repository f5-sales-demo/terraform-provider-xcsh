---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-3021302021012012-1313032111213313-0012000330200310-2012132323022003-1201132300101233-2123321033133120-2223131231323230-2130321211212303"></a>

## simple_service.container.readiness_check.tcp_health_check.port — port / 300302021000 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.container](data-sources--workload--reference--group-016.md#canonical-3120213033223022-3001202202000120-1331313132130330-1102012032232032-1332103223123321-2212010010003232-3123223220322033-1032222002231223)
- [simple_service.container.readiness_check](data-sources--workload--reference--group-016.md#canonical-0312322121032303-0120231311113002-3232323333022313-0313301223202010-0113120313112233-3202000230211322-1223203103333003-2223013222202203)
- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-2003212223310021-3331301230033133-2110300232002230-1202301130320132-3231030132313333-2123332103111333-2332101310010223-2302122111302123)
- simple_service.container.readiness_check.tcp_health_check.port

<a id="canonical-0012332202330310-0002023210211033-2031330322333001-1300102322120131-1121123002210120-2021231132002033-2332210333123202-1323011122123020"></a>

Type: `"single"`. Computed.

Port. Port

Upstream description:

Port

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"name\",\"num\"]"
}
```

<a id="canonical-1001330200013110-1011210032033032-2000232221223010-0233100302332313-0013211013101120-0020010011112001-1233323333020113-3231000230023133"></a>

## Direct properties — port / 300302021000 / 3

<a id="canonical-1223022312202202-0313203022322002-1103023303333212-2333303133323230-2312310210020330-2200012101201310-1311131231032200-0213201220210120"></a>

<a id="canonical-3323200012332220-3110011310213133-2110301313101210-0113020210211101-3101033001132010-2221133321201122-0023322103222022-2233112333130103"></a>

## name property — port / 300302021000 / 4

Type: `"string"`. Computed.

Port Name. Exclusive with \[num\] Port Name.

Upstream description:

Exclusive with \[num\] Port Name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.iana_svc_name": "true"
  }
}
```

<a id="canonical-3232030210322330-3110323013222130-2030331212323020-3031113023310201-2211131233113032-0122123302123113-0220013211120100-3012321212331120"></a>

<a id="canonical-0130310202201002-2023032203310223-1032111313032123-0201212330021302-3303111323002013-1022222012011301-1223223023121120-1031213303103011"></a>

## num property — port / 300302021000 / 5

Type: `"number"`. Computed.

Port Number. Exclusive with \[name\] Port number.

Upstream description:

Exclusive with \[name\] Port number.

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

<a id="canonical-0301102303121002-0010130330100202-3001230000233000-2212020233003313-1230021032032220-3220201022300222-2202213203111302-2331232130313333"></a>

## Next pages — port / 300302021000 / 6

- [simple_service.container.readiness_check.tcp_health_check](data-sources--workload--reference--group-016.md#canonical-2003212223310021-3331301230033133-2110300232002230-1202301130320132-3231030132313333-2123332103111333-2332101310010223-2302122111302123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2112103133232330-2212313300022031-2132030012302003-0033223033113333-2311301323022003-0110122323101031-0101130200003103-2203110002233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331311030012010-1323200132233220-1331020213300232-3333112213232101-3322113033230002-2113333023111321-0013213112031133-3013033313011011"></a>

## simple_service.disabled — disabled / 321032032323 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.disabled

<a id="canonical-1230223131002211-3031103301233131-0310231020102201-2131313210102210-3303321202330111-1220223210213313-2232003201202212-0310301130123300"></a>

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

<a id="canonical-0110002101123022-2000113322000323-3033333001121113-3013203223121212-1203121030031212-2321213332100022-0202212100223110-3333233221013211"></a>

## Direct properties — disabled / 321032032323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0200133121302313-1203311031122321-1330132023013223-1310111232121301-1333201131103002-1130101011222112-0011020220011332-1310133033032322"></a>

## Next pages — disabled / 321032032323 / 4

- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1210220313300213-3013331300310122-2232001332132003-2302230111202103-0230012033213001-3213221121020312-3303110313013123-1101012020030220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123012020332101-3233022320220301-1112232003000111-0211302333131110-2320302310012332-0000032303200003-3111302220223200-2100313132310230"></a>

## simple_service.do_not_advertise — do_not_advertise / 323321122101 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.do_not_advertise

<a id="canonical-1330003310220120-0230312011113200-3223131133322232-3033310030032111-1032030010133312-3101021220122300-2130020310301320-0021132302232233"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for do not advertise.

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

<a id="canonical-0301221120120310-1023233221100031-0211011030233201-0233331102110302-0101131030221010-2132321200102130-2301312312121320-1101233321031121"></a>

## Direct properties — do_not_advertise / 323321122101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002213211302023-3031210320323322-1332330323320121-3031133110112020-3123111132022010-2310103033222133-0112310203232012-1230332131003111"></a>

## Next pages — do_not_advertise / 323321122101 / 4

- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112301033220330-3221220203203031-3110213102031333-2101331310132120-1230332121000123-2130332201113331-0231232223002123-0323032301233322"></a>

## simple_service.enabled — enabled / 033001221011 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.enabled

<a id="canonical-1322110013113212-0213001333002132-1330211012032023-0322333022321312-2100330122102130-1202011111103203-1103122331123212-1033302000002210"></a>

Type: `"single"`. Computed.

Persistent storage volume configuration for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1013232101030202-1330100002332011-1102221003220201-0003313033313312-1032033110311110-1013302202322112-2333300021202033-0220231300302033"></a>

## Direct properties — enabled / 033001221011 / 3

<a id="canonical-2202200001010330-0332313010010121-2101122010231133-0211332202130231-3303322102002331-3210012032033211-0031222013021033-0131232022202111"></a>

<a id="canonical-2013033130131112-0210102030321301-3103230312131303-2220212131010030-3101230212133311-3302011010221110-2213322011213023-1323031112003120"></a>

## name property — enabled / 033001221011 / 4

Type: `"string"`. Computed.

Name. Name of the volume.

Upstream description:

Name of the volume.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
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
    "ves.io.schema.rules.string.dns_1123_label": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.dns_1123_label": "true"
  }
}
```

- [persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331): complete subsection reference.

<a id="canonical-0133031131221010-2312313110202123-2113012222101210-1133030010333123-0313321121132132-1032123030120333-2302031102000031-2023002112303120"></a>

## Next pages — enabled / 033001221011 / 5

- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223311232123113-1212022322023222-0201223020323331-0003001311002132-2321320222313332-1022000200132201-0322232332133031-3103212212112322"></a>

## simple_service.enabled.persistent_volume — persistent_volume / 220121301032 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-017.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- simple_service.enabled.persistent_volume

<a id="canonical-0102001303002020-2223302120021122-1223332231230032-0203202222000101-2320311120322222-1322320131021030-3221131012103133-3210131012122000"></a>

Type: `"single"`. Computed.

Volume containing the Persistent Storage for the workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2110032313003301-3032003332121032-3023120011100102-2310111103011303-1332310012231212-1211230301332100-3011312003121112-2230301213020222"></a>

## Direct properties — persistent_volume / 220121301032 / 3

- [mount](data-sources--workload--reference--group-017.md#canonical-1121310022032211-3322100203123030-2130032220221232-1003303032222002-2233101013131002-2032133200010313-0033221033000303-1032322332101333): complete subsection reference.

- [storage](data-sources--workload--reference--group-017.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233): complete subsection reference.

<a id="canonical-1023132301233020-3003333110230221-0232313210020111-0220330202211232-3222322202031201-1033102220003023-2122231003030021-2231232103112320"></a>

## Next pages — persistent_volume / 220121301032 / 4

- [simple_service.enabled.persistent_volume.mount](data-sources--workload--reference--group-017.md#canonical-1121310022032211-3322100203123030-2130032220221232-1003303032222002-2233101013131002-2032133200010313-0033221033000303-1032322332101333)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-017.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233)
- [simple_service.enabled](data-sources--workload--reference--group-017.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1121310022032211-3322100203123030-2130032220221232-1003303032222002-2233101013131002-2032133200010313-0033221033000303-1032322332101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330012313132101-2302222220111211-1020110211210002-1231303122323021-0113010003021123-0111232231111022-2022201310021020-0330201212202102"></a>

## simple_service.enabled.persistent_volume.mount — mount / 113212303122 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-017.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- simple_service.enabled.persistent_volume.mount

<a id="canonical-1130201020220121-2313130002312300-3331011100003021-1222030220213300-2103103313000023-2123011320110130-1031002112303001-2101033330203211"></a>

Type: `"single"`. Computed.

Volume mount describes how volume is mounted inside a workload.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030212101021333-1223210110033122-3303111231011210-2200230031323112-2313032122333003-2100011302103111-2031030313203121-2213221010122031"></a>

## Direct properties — mount / 113212303122 / 3

<a id="canonical-3001213221110001-1203121211303001-1001120230322211-0013210311031211-0210031223233210-2310022211021300-2331202120113300-3320131200320122"></a>

<a id="canonical-2011103202010211-3321013323130220-3011320321112033-0302220230011101-1320023132101123-0110031103030100-0313213113233300-1103302231102033"></a>

## mode property — mount / 113212303122 / 4

Type: `"string"`. Computed.

\[Enum: VOLUME\_MOUNT\_READ\_ONLY|VOLUME\_MOUNT\_READ\_WRITE\] Mode in which the volume should be
mounted to the workload - VOLUME\_MOUNT\_READ\_ONLY: ReadOnly Mount the volume in read-only mode -
VOLUME\_MOUNT\_READ\_WRITE: Read Write Mount the volume in read-write mode. Possible values are
\`VOLUME\_MOUNT\_READ\_ONLY\`, \`VOLUME\_MOUNT\_READ\_WRITE\`. Defaults to
\`VOLUME\_MOUNT\_READ\_ONLY\`.

Upstream description:

Mode in which the volume should be mounted to the workload

&#8203;- VOLUME\_MOUNT\_READ\_ONLY: ReadOnly

Mount the volume in read-only mode &#8203;- VOLUME\_MOUNT\_READ\_WRITE: Read Write

Mount the volume in read-write mode.

Receipt-pinned upstream constraints:

```json
{
  "default": "VOLUME_MOUNT_READ_ONLY",
  "enum": [
    "VOLUME_MOUNT_READ_ONLY",
    "VOLUME_MOUNT_READ_WRITE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2313311310320331-1211303011313033-2222300202321312-3120200020132031-0311331331332312-3120313001211212-0000313000101233-0313221132210002"></a>

<a id="canonical-2132312302010000-3011201330221101-2101331121200032-1232033131311233-0203320222200131-1313011033301201-0021120321310213-0120331101322001"></a>

## mount_path property — mount / 113212303122 / 5

Type: `"string"`. Computed.

Path within the workload container at which the volume should be mounted. Must not contain ':'.

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
    "pattern": "^[^:]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^[^:]*$"
  }
}
```

<a id="canonical-2102012230331121-2010303122100322-0330032303100120-3031330021111202-2112232100232320-3002210232121120-2222333311321322-3022113223112013"></a>

<a id="canonical-1133010301030320-1203301201031122-1132101221312302-0120311220331110-1322121020021120-1100101203221310-3123232203321322-3203302310102313"></a>

## sub_path property — mount / 113212303122 / 6

Type: `"string"`. Computed.

Path within the volume from which the workload's volume should be mounted. Defaults to '' (volume's
root).

Upstream description:

Path within the volume from which the workload's volume should be mounted. Defaults to "" (volume's
root).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0021110230223322-2000300013203112-3222122210010222-1220332100322223-3121232120033013-3022112222233022-0233112021102001-1332110021331231"></a>

## Next pages — mount / 113212303122 / 7

- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233302122330230-1103102302003133-0321330233202110-1121221112133023-0031010320032201-0013222103321021-1103233331303023-3131302111331120"></a>

## simple_service.enabled.persistent_volume.storage — storage / 110133012033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-017.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- simple_service.enabled.persistent_volume.storage

<a id="canonical-3200132300201303-1130203302131102-0102231320112310-2000201302120011-0223123231103233-3023133120302022-0332311330210231-3013111112001021"></a>

Type: `"single"`. Computed.

Persistent storage configuration is used to configure Persistent Volume Claim (PVC).

Upstream description:

Persistent storage configuration is used to configure Persistent Volume Claim (PVC)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-class_name_choice": "[\"class_name\",\"default\"]"
}
```

<a id="canonical-1101301030013010-0300210121013223-3121010212010330-0033222230211213-0321131131131031-1022001222331101-0213123101100033-1310001131320210"></a>

## Direct properties — storage / 110133012033 / 3

<a id="canonical-0230303230311022-3302003220120332-3133303130133020-3223310213101301-3231232031010320-3233000122122203-1321311032321221-2010130230322100"></a>

<a id="canonical-1101010301323201-3312130110113102-1123222012103333-1201333130123001-1122203232122133-1021210321333110-2010020002330301-1311003002001230"></a>

## access_mode property — storage / 110133012033 / 4

Type: `"string"`. Computed.

\[Enum:
ACCESS\_MODE\_READ\_WRITE\_ONCE|ACCESS\_MODE\_READ\_WRITE\_MANY|ACCESS\_MODE\_READ\_ONLY\_MANY\]
Persistence storage access mode is used to configure access mode for persistent storage -
ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once Read Write Once is used to mount persistent storage
in read/write mode to exactly 1 host - ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many Read Write
Many is used.. Possible values are \`ACCESS\_MODE\_READ\_WRITE\_ONCE\`,
\`ACCESS\_MODE\_READ\_WRITE\_MANY\`, \`ACCESS\_MODE\_READ\_ONLY\_MANY\`. Defaults to
\`ACCESS\_MODE\_READ\_WRITE\_ONCE\`.

Upstream description:

Persistence storage access mode is used to configure access mode for persistent storage

&#8203;- ACCESS\_MODE\_READ\_WRITE\_ONCE: Read Write Once

Read Write Once is used to mount persistent storage in read/write mode to exactly 1 host &#8203;-
ACCESS\_MODE\_READ\_WRITE\_MANY: Read Write Many

Read Write Many is used to mount persistent storage in read/write mode to many hosts &#8203;-
ACCESS\_MODE\_READ\_ONLY\_MANY: Read Only Many

Read Only Many is used to mount persistent storage in read-only mode to many hosts.

Receipt-pinned upstream constraints:

```json
{
  "default": "ACCESS_MODE_READ_WRITE_ONCE",
  "enum": [
    "ACCESS_MODE_READ_WRITE_ONCE",
    "ACCESS_MODE_READ_WRITE_MANY",
    "ACCESS_MODE_READ_ONLY_MANY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1332222100302022-2320303122300130-0310020001323133-3303332212210222-3100221323300101-1001330011112220-0002111321121312-1210213111221110"></a>

<a id="canonical-1031233232132011-1032300200000330-3322231220130013-0233101023222210-2211220131313111-2302223301232010-3311213313233313-1233322003030311"></a>

## class_name property — storage / 110133012033 / 5

Type: `"string"`. Computed.

Exclusive with \[default\] Use the specified class name.

Upstream description:

Exclusive with \[default\] Use the specified class name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [default](data-sources--workload--reference--group-017.md#canonical-2010000222012032-0002313022130011-3123231302300212-0221121100202132-2233032213121330-1130033321022310-3200020030203002-2320103331230221): complete subsection reference.

<a id="canonical-0323213130210212-1101312012012013-1200331133230021-2003120011310311-1333202321201121-3021022022022131-1233331301002132-0121131020301223"></a>

<a id="canonical-0200332302303231-2212101230133331-1000112202202112-1222013031122103-2123312033032130-2320230233132111-2033220123023312-0231320010332111"></a>

## storage_size property — storage / 110133012033 / 6

Type: `"number"`. Computed.

Size (in GiB). Size in GiB of the persistent storage.

Upstream description:

Size in GiB of the persistent storage.

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
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.gte": "0.004",
    "ves.io.schema.rules.double.lte": "256",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2111001021021201-1230102323301110-0203333002010301-0301002003130100-2133321200000313-1220123302222001-1103031110321122-3022112313030333"></a>

## Next pages — storage / 110133012033 / 7

- [simple_service.enabled.persistent_volume.storage.default](data-sources--workload--reference--group-017.md#canonical-2010000222012032-0002313022130011-3123231302300212-0221121100202132-2233032213121330-1130033321022310-3200020030203002-2320103331230221)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2010000222012032-0002313022130011-3123231302300212-0221121100202132-2233032213121330-1130033321022310-3200020030203002-2320103331230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210113311333010-3222102022312110-3032312223032120-1120232101210302-2013111310130210-1130132001010013-3023020322103021-3112331111010211"></a>

## simple_service.enabled.persistent_volume.storage.default — default / 233002201202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [simple_service.enabled](data-sources--workload--reference--group-017.md#canonical-3001213300210321-2011120130303023-2211001300322013-1332022123332020-0232131121303302-0131000112110201-1133003122121203-3212202000010030)
- [simple_service.enabled.persistent_volume](data-sources--workload--reference--group-017.md#canonical-2123111321221000-0222301112113233-3033212121211311-2222113101333011-3111222331002110-1021332110233023-1113010330312332-2232102112321331)
- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-017.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233)
- simple_service.enabled.persistent_volume.storage.default

<a id="canonical-0201223002010212-3110102220213203-0222230320310113-2102012102232230-1130133322201020-1130133212230300-0310202332201112-1313201011022323"></a>

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

<a id="canonical-3103313112111301-2021332120133312-0231132001310031-1311202331013231-2030022133010311-1203230002022013-1213130233021032-3123330310302322"></a>

## Direct properties — default / 233002201202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020201133230301-1331021333132113-3202002031201112-2031013111232232-3031121330132000-1122112013211031-0232000133102232-0100210201310102"></a>

## Next pages — default / 233002201202 / 4

- [simple_service.enabled.persistent_volume.storage](data-sources--workload--reference--group-017.md#canonical-3222112122312030-0100131331212110-3311100330320223-2221122302332032-2221333323323122-0210301032111313-0001233130123103-3111000232322233)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1332023020220230-0231312012221313-1132321323222002-2223303033022101-2220302230211001-0323332011222312-2312021312021221-1032021121111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200213031222231-2301323213221003-3313301231320233-1333312101003313-0321200112302113-1313232211302131-3212320221100122-3203231213211131"></a>

## simple_service.simple_advertise — simple_advertise / 213213220320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- simple_service.simple_advertise

<a id="canonical-2322232111322303-2221013322111133-2202332330023003-2321130113230202-0111120120132233-1203332010213001-2102213320012002-3133012213113121"></a>

Type: `"single"`. Computed.

Configuration parameter for simple advertise.

Upstream description:

Advertise OPTIONS for Simple Service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0233033021123012-3003032020012120-0312312300111203-1021200021223123-2121123213321313-3120310003233313-3311013021212113-0230201313002211"></a>

## Direct properties — simple_advertise / 213213220320 / 3

<a id="canonical-1222302213332212-0303231030022002-2120322232333111-1320002331212032-1021132200203002-2200212200321320-2000030303033232-3201312220020033"></a>

<a id="canonical-3030220101002302-0301022231332113-1011223300101113-1323320011331000-0200001203311300-2223313320132123-0003223323321030-2303322133123011"></a>

## domains property — simple_advertise / 213213220320 / 4

Type: `["list", "string"]`. Computed.

List of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form Supported Domains and search order: 1. Exact Domain names:
www&#46;example.com. 2.

Upstream description:

A list of Domains (host/authority header) that will be matched to Load Balancer. Wildcard hosts are
supported in the suffix or prefix form

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Load Balancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3012203233110003-0322302302023231-3203001032331001-3210023223023133-1200213002110103-1301222023331320-2103210322100200-0232033321321321"></a>

<a id="canonical-2100023133202123-0320202313001112-2200103002012111-3121000202033113-0212331330212200-1010311311023111-3003002031200310-0213211012330023"></a>

## service_port property — simple_advertise / 213213220320 / 5

Type: `"number"`. Computed.

Service port to advertise on internet via HTTP loadbalancer using port 80.

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
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2302103011233200-1220130011021303-1322321111011132-2331200301230131-3230230023111111-3003320031200200-2203223310130022-3000003312333100"></a>

## Next pages — simple_advertise / 213213220320 / 6

- [simple_service](data-sources--workload--reference--group-016.md#canonical-1133330001222133-0023013333013330-3330013322032021-3302213310300033-2312032031201020-3233031203301231-0303033100232200-0223111033220010)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212123223033322-0222131012302323-1222020132333011-3331221000110112-0030333113311022-0133011333311202-3122103310033210-1033100110312032"></a>

## stateful_service — stateful_service / 023202320112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- stateful_service

<a id="canonical-3132102321212312-1030011200002201-0312221231111333-3000303301000002-2111003212203322-3013233131032031-0201123032301323-2201123002233101"></a>

Type: `"single"`. Computed.

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like Cassandra, MongoDB, redis, etc.

Upstream description:

StatefulService maintains per replica state and each replica has its own persistent storage. Each
replica has a unique network identity and stable storage. Stateful service are used for distributed
stateful applications like Cassandra, MongoDB, redis, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-scaling_choice": "[\"num_replicas\",\"scale_to_zero\"]"
}
```

<a id="canonical-0002131321331033-3200320233012111-2131131130012131-2213211201232110-1112211123300013-2303231300012000-1021001213023030-2303331022200001"></a>

## Direct properties — stateful_service / 023202320112 / 3

- [advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220): complete subsection reference.

- [configuration](data-sources--workload--reference--group-027.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100): complete subsection reference.

- [containers](data-sources--workload--reference--group-027.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003): complete subsection reference.

- [deploy_options](data-sources--workload--reference--group-028.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133): complete subsection reference.

<a id="canonical-0211202302123311-0002302301110323-2121323221310032-1102131221120231-0021022131232020-0033031300212031-3230322112212110-1322232212032232"></a>

<a id="canonical-3231022032102211-0333110313220210-2223333313221132-2313021120033330-2023323322312010-0310301130000032-1123302012023222-1000333103311120"></a>

## num_replicas property — stateful_service / 023202320112 / 4

Type: `"number"`. Computed.

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Upstream description:

Exclusive with \[scale\_to\_zero\] Number of replicas of service to spawn per site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gt": "0",
    "ves.io.schema.rules.int32.lte": "5"
  }
}
```

- [persistent_volumes](data-sources--workload--reference--group-028.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303): complete subsection reference.

- [scale_to_zero](data-sources--workload--reference--group-028.md#canonical-2130310200302331-3301113001002211-0231300130002032-0320200012113302-2010020220101022-0221323130323300-0010123130122332-1232023022230312): complete subsection reference.

- [volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212): complete subsection reference.

<a id="canonical-0322213001000023-3123221221020223-3203320201130100-0212303011002011-0223110332310120-0312113010321011-0320110021233023-0121202332101201"></a>

## Next pages — stateful_service / 023202320112 / 5

- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.configuration](data-sources--workload--reference--group-027.md#canonical-2032211221320222-3233032312121000-0321111221221122-3331201113022312-0020000223303230-0020033200320022-2222023002102320-2200212222310100)
- [stateful_service.containers](data-sources--workload--reference--group-027.md#canonical-3101311130033212-3010211101012322-2332300122313021-2300111123200012-3222222201300313-2322010011032230-2110031201133132-3103131301302003)
- [stateful_service.deploy_options](data-sources--workload--reference--group-028.md#canonical-1013001210313000-3001012211030201-1210322313302221-3132213121000320-0012111130102110-3332131123011111-1230211023121203-1131013202130133)
- [stateful_service.persistent_volumes](data-sources--workload--reference--group-028.md#canonical-1132102002130121-0202013313231301-3123121303133333-3111103021121020-1113123203313322-3332020232010320-2113332300313132-2131322221012303)
- [stateful_service.scale_to_zero](data-sources--workload--reference--group-028.md#canonical-2130310200302331-3301113001002211-0231300130002032-0320200012113302-2010020220101022-0221323130323300-0010123130122332-1232023022230312)
- [stateful_service.volumes](data-sources--workload--reference--group-028.md#canonical-2011330200003232-3232210131033220-1331201211101011-0113112011320103-3212220101322221-0100131112031303-3331333222200110-3212321201321212)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203231022303232-3121213122031011-2010030223123110-2202321013202113-0120131321301001-2002003110110213-3121210001311022-0230330000331020"></a>

## stateful_service.advertise_options — advertise_options / 031321232123 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- stateful_service.advertise_options

<a id="canonical-0111221301333012-1330211132212122-2102332121012301-1222122213133322-2001112100000333-1200220223303322-3302212113001302-2302223221302012"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

<a id="canonical-0131303330332102-2122221333211110-0312200212033033-0120112133113010-0200211221013111-3220023201232231-1222312013320200-0130331233333322"></a>

## Direct properties — advertise_options / 031321232123 / 3

- [advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323): complete subsection reference.

- [advertise_in_cluster](data-sources--workload--reference--group-020.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122): complete subsection reference.

- [advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220): complete subsection reference.

- [do_not_advertise](data-sources--workload--reference--group-027.md#canonical-2130301222010233-1223132131303222-0322121021211321-2133011133112020-3312223303312212-3213023022123103-1032021220021202-2003132313121010): complete subsection reference.

<a id="canonical-2112313321112230-0101010200103020-3123111222030133-0132212103011030-3231213101133330-2002333021112023-0212102200312323-1333312332132331"></a>

## Next pages — advertise_options / 031321232123 / 4

- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_in_cluster](data-sources--workload--reference--group-020.md#canonical-2000012203330031-3013213320003321-1130203203021012-0301212320010201-2013312012132313-2201113011320310-2123010110120120-3030200031122122)
- [stateful_service.advertise_options.advertise_on_public](data-sources--workload--reference--group-020.md#canonical-1022213130002111-1013210322013321-1133330230031233-1133100000003020-0332013302021302-0032122302100211-2023110210133002-1211310111330220)
- [stateful_service.advertise_options.do_not_advertise](data-sources--workload--reference--group-027.md#canonical-2130301222010233-1223132131303222-0322121021211321-2133011133112020-3312223303312212-3213023022123103-1032021220021202-2003132313121010)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022302203223122-0203003322313102-2222303323311332-3110001102032212-0111200011310131-1233222023301211-1002211102123221-2212333013300101"></a>

## stateful_service.advertise_options.advertise_custom — advertise_custom / 220322030130 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- stateful_service.advertise_options.advertise_custom

<a id="canonical-2202303311011230-1323020101303130-2233010130123230-0213033131123000-1113233022133222-3303020021030133-3121112110020212-0122312130311121"></a>

Type: `"single"`. Computed.

Advertise this workload via loadbalancer on specific sites.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102203111323313-1210031301232103-3023330302303023-0322113203203220-2213020300113121-2113200101100111-1101223110232301-1333121200301200"></a>

## Direct properties — advertise_custom / 220322030130 / 3

- [advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123): complete subsection reference.

- [ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221): complete subsection reference.

<a id="canonical-2213101020323112-2311322133120312-2313013210033031-2202303222331211-3330220110130312-3222200100002312-3221310220222313-3212301021201303"></a>

## Next pages — advertise_custom / 220322030130 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110122211300022-2231000201330321-2330023113131211-3033233313031333-0100300000322310-3213300200102103-1300100200301131-0202322122011233"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where — advertise_where / 031002113100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- stateful_service.advertise_options.advertise_custom.advertise_where

<a id="canonical-2213103203011233-2321231203232002-2003221030021302-1100133330020120-3211300331102113-3303031310020302-2320103112131132-0122001210010333"></a>

Type: `"list"`. Computed.

Where should this load balancer be available.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1030203232332030-3301102133302030-0301221310012012-0100233323000132-0303200022333321-2332001022222312-1233233200001010-0322031213112122"></a>

## Direct properties — advertise_where / 031002113100 / 3

- [site](data-sources--workload--reference--group-017.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-017.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331): complete subsection reference.

- [vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303): complete subsection reference.

<a id="canonical-1000131002312202-1030013203020133-3023010110131111-1302211302211120-0010121020333132-0033130232302131-2013232201332321-2123301022332302"></a>

## Next pages — advertise_where / 031002113100 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-017.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-017.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212011000003000-3310331233311211-3033320132130213-0003232321101002-3203230031101102-3221302021013122-3023113312203302-0131011212110012"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site — site / 021022211230 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.site

<a id="canonical-2321120210020110-0010102111230011-3330311021122332-3211222221130230-0123110211223021-1011330300212123-2101031323023333-1103020213000210"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0131301002333212-3122020333031233-2223300032200130-3201201313030111-1321332002110031-0211331000022012-2310323110020122-0202321313301313"></a>

## Direct properties — site / 021022211230 / 3

<a id="canonical-3310311013212020-0300120022032111-3220122212032021-1323212311001020-3320121212330131-1132220311130020-2022312132301312-3020001013231120"></a>

<a id="canonical-1113232011120112-1112231311300110-3012310101122003-3321021120331323-0101233213010011-3121011120010030-3303112220321330-2203303010210232"></a>

## ip property — site / 021022211230 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

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

<a id="canonical-3120102131120100-0020333300301213-3220201321203303-0022111311310320-1300310202222322-3021103121202232-1301132033112322-0233331223323011"></a>

<a id="canonical-3121132223330322-2022110130100102-0121131200121303-3013133212333230-3333110201300221-2110301223212001-3010300303312330-0002131210101322"></a>

## network property — site / 021022211230 / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--workload--reference--group-017.md#canonical-0311323121320121-1121012301100112-1233122132103231-2202232321201330-1300031303100112-3332030020233302-1120023312101132-3022221031130011): complete subsection reference.

<a id="canonical-0202213110233202-1001133330221220-2302131112102123-0130300010022003-1211333110202021-0002213331320131-1211320312222120-2333032201222301"></a>

## Next pages — site / 021022211230 / 6

- [stateful_service.advertise_options.advertise_custom.advertise_where.site.site](data-sources--workload--reference--group-017.md#canonical-0311323121320121-1121012301100112-1233122132103231-2202232321201330-1300031303100112-3332030020233302-1120023312101132-3022221031130011)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0311323121320121-1121012301100112-1233122132103231-2202232321201330-1300031303100112-3332030020233302-1120023312101132-3022221031130011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312133311213330-3310131232302332-1210001120211110-1302311300022323-0313331030131202-0131312132201233-2230213013202302-3023212331303131"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.site.site — site / 212202321212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-017.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013)
- stateful_service.advertise_options.advertise_custom.advertise_where.site.site

<a id="canonical-2110023130200132-1212200321022322-2230133122200310-1033030131303031-3023231012101030-2233331302211100-3100010111303302-2102020010333103"></a>

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

<a id="canonical-3200230221300233-0220011103310021-0012310333122111-1220111331111000-3201110320321332-2120232223300031-2232031023013303-0202320102230200"></a>

## Direct properties — site / 212202321212 / 3

<a id="canonical-1311130300100313-0221100332013233-3103331331131133-3212101131303000-3111101023101112-0031213012130202-0203100203122120-2212213122123302"></a>

<a id="canonical-1022030100300321-1222002220001220-3330032300221211-2310110032121020-1012001320210002-3130333323120202-2111102331110321-2330033011311013"></a>

## name property — site / 212202321212 / 4

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

<a id="canonical-0110131230120301-1221321110212120-2203011212300110-3033323221132232-1020222103312030-1202020031121103-1122210333311220-2230200021233100"></a>

<a id="canonical-2312201012120221-0003123021321022-1130033210211200-2323233103223332-2132030103030231-1130002200111130-2011131103133121-0322203013130313"></a>

## namespace property — site / 212202321212 / 5

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

<a id="canonical-2323221213120100-2122230131100212-0132000331201023-3120331212323203-2033230120112111-0321030100122301-3001020210333010-1021301112002002"></a>

<a id="canonical-0230021333310022-1021013130303221-3211030310213231-1100033332322203-0331113321030000-0123121133231120-1022100331213333-0211111300010020"></a>

## tenant property — site / 212202321212 / 6

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

<a id="canonical-0021031130222021-1011202020132211-1132322122203122-0221302101200123-0312312003301023-3221003303031030-3010122121030210-2011132323022131"></a>

## Next pages — site / 212202321212 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.site](data-sources--workload--reference--group-017.md#canonical-3320202201130032-0203220333223120-2030320330020021-0221300332330021-1030120320110000-2222122233221333-0132001012231210-0133323000302013)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123212310133333-1022020223031232-0033033133313132-2100211102322233-2030311330222032-2210030311131000-0232223032222333-2321021233133022"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site — virtual_site / 101202330111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site

<a id="canonical-0223123031021033-0000033123311213-3320332121312303-0100103312303031-0003313301012120-0322102212121330-1133131102002332-2200000130333033"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121232321110002-1213323032300022-2201100311333002-3202233010103100-2230030032011121-3021012032000022-1200203111211300-1222021322121320"></a>

## Direct properties — virtual_site / 101202330111 / 3

<a id="canonical-1023032331321300-3313033000233133-2222230220302301-3011122212311010-1212202000311321-1330103312310130-1303020010031030-1210130022210220"></a>

<a id="canonical-0132002230113131-3021103212132100-1111020300022213-0011213213321313-1213122321120213-1230012311020133-0032212323210331-3320012101030002"></a>

## network property — virtual_site / 101202330111 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--workload--reference--group-017.md#canonical-3110123121231321-3121211321023220-2233013112201330-2121100330320032-0012123012203020-3130100111122332-0123223000133331-3031333313133110): complete subsection reference.

<a id="canonical-1021100033222021-2031011022002002-2202320211312120-0002312332331000-3110322030003332-1202222031132032-0221032323021100-2003012101330012"></a>

## Next pages — virtual_site / 101202330111 / 5

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--workload--reference--group-017.md#canonical-3110123121231321-3121211321023220-2233013112201330-2121100330320032-0012123012203020-3130100111122332-0123223000133331-3031333313133110)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3110123121231321-3121211321023220-2233013112201330-2121100330320032-0012123012203020-3130100111122332-0123223000133331-3031333313133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331111312200213-3130331032303301-3120311123013023-0201011023110010-2131012120210023-3303232231130100-0101301302320202-0031222221103010"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 101102133213 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-017.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331)
- stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-3232023033111213-1110232213210203-1101333233012212-1223022222131101-0103020310211021-0221313203130313-0000231020020012-0223111222333020"></a>

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

<a id="canonical-0032010232313113-3221001100220303-0010122122100132-3332311123032113-3233033331331332-1001301132100322-1030122120121323-2132011021331112"></a>

## Direct properties — virtual_site / 101102133213 / 3

<a id="canonical-1310100322133320-1132100103333233-2001200223113222-3022232303323212-3100221312221331-2323310232220011-1110030102311101-3030123121010331"></a>

<a id="canonical-1331230231220233-3111131220333012-0322200020212023-1200011001002003-1001222111001023-0330320032100201-1332232133100121-1313320110213130"></a>

## name property — virtual_site / 101102133213 / 4

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

<a id="canonical-2100203120322223-1033320202100131-2100201111000113-0222203113132100-0120323010231231-3103002002222213-3321111233031210-1312002312313022"></a>

<a id="canonical-0103100003320312-3111002000202112-0300202010110230-3310013332200120-3210232010212332-2333313123100201-0000023322221131-1000223202200233"></a>

## namespace property — virtual_site / 101102133213 / 5

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

<a id="canonical-2210210232010130-3221132320123313-0231130122121311-1332122302132311-3121312100011312-0331031230313120-3002311022300110-1212122032132203"></a>

<a id="canonical-0131020010020101-1212121202312111-1302020212333111-0320132023223013-1003310112200132-1003020101220333-3203102330032210-2221021330101203"></a>

## tenant property — virtual_site / 101102133213 / 6

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

<a id="canonical-0130112330110221-3303021023210001-2310000332130331-0112310223221332-2012201031133311-0032331113310203-3023131232121010-3123031003312313"></a>

## Next pages — virtual_site / 101102133213 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.virtual_site](data-sources--workload--reference--group-017.md#canonical-2032022022213011-2330313321021313-1332303311331122-3303001132120112-1231213133110121-1001301121011121-0022023232231103-1202232213310331)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301103212001123-3033203323311101-3132002122232103-1213123230200200-0000001033220303-2330233211320213-0020311010333022-0101233021233112"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service — vk8s_service / 000133223133 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service

<a id="canonical-1201310232312011-2310012112333322-3130102301103210-1212322301213032-0031210332130333-2020022000112031-2122330310133020-2111332012223311"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1021311213221330-1222131031122003-0310200233003012-3230002231330010-2121123201131031-0302331033011202-2313201022010032-2202212202131322"></a>

## Direct properties — vk8s_service / 000133223133 / 3

- [site](data-sources--workload--reference--group-017.md#canonical-0110130013013011-1110302030103030-0321201230331211-3232303310333121-3310112222313131-1022011002122101-1132220313123223-1032121300032111): complete subsection reference.

- [virtual_site](data-sources--workload--reference--group-017.md#canonical-2222303112221113-0331322312133021-3033323021222002-2232332300210110-3121100200032000-2333032111222110-1120002312333303-3322122122201213): complete subsection reference.

<a id="canonical-2322032212013201-2302302320101021-3221320011232032-0230210311213301-2212320323030031-3310101121231210-1333331012112030-2122330121220033"></a>

## Next pages — vk8s_service / 000133223133 / 4

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site](data-sources--workload--reference--group-017.md#canonical-0110130013013011-1110302030103030-0321201230331211-3232303310333121-3310112222313131-1022011002122101-1132220313123223-1032121300032111)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--workload--reference--group-017.md#canonical-2222303112221113-0331322312133021-3033323021222002-2232332300210110-3121100200032000-2333032111222110-1120002312333303-3322122122201213)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0110130013013011-1110302030103030-0321201230331211-3232303310333121-3310112222313131-1022011002122101-1132220313123223-1032121300032111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310232131322020-3201123300303310-3032310000300100-0103321210212030-2021012231210322-3132113202221210-3332003233122022-3121222230000312"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site — site / 121112130133 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2111110233131321-1320220132223301-0212120323231211-2211202000032331-3321213321100010-3200230220220030-1233113310202110-1010010133210020"></a>

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

<a id="canonical-0100321031030203-0033001302131211-0212232002023111-2302320232020303-2102302301303000-1131202003221102-1333133221132213-1333111300211001"></a>

## Direct properties — site / 121112130133 / 3

<a id="canonical-0322302020001021-3231232032301011-0230112331220113-1310021103020221-0013111030131323-2123200202303302-2120130323202120-0021332113100323"></a>

<a id="canonical-3003013212112322-2300230323201222-1103330132212123-0112012100123203-3112311230111032-2013321211312122-3131030201030000-0222233112200311"></a>

## name property — site / 121112130133 / 4

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

<a id="canonical-3103100301000023-3332313002203001-1213120113322222-1302003222201032-2103331322313322-0102012221223203-3020313022032032-3012010330133001"></a>

<a id="canonical-0310120032230330-0200101121011102-2310203023233303-3133322320303330-1333233102010301-3102200033113221-0030030202333233-3110030113010312"></a>

## namespace property — site / 121112130133 / 5

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

<a id="canonical-2011031030301012-3133102100130030-0111132310230300-1131111200023220-3333031112230031-2300022213002332-3221003221120031-3012121133301313"></a>

<a id="canonical-3222232311011302-3131021321312113-0011323103201320-1000332322023033-0000031332213231-0002301132110012-0222330332233323-2022230231302202"></a>

## tenant property — site / 121112130133 / 6

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

<a id="canonical-1201130301013203-0001212020010032-1101203010021100-0112002103021310-1301211213213220-3001023030033010-2322110110010333-0203313302323111"></a>

## Next pages — site / 121112130133 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2222303112221113-0331322312133021-3033323021222002-2232332300210110-3121100200032000-2333032111222110-1120002312333303-3322122122201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021302100121222-2100011023120213-1332123003332020-3202112102311113-3112301133332301-3221131302012122-3232212013203102-0222203331011211"></a>

## stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 202200131033 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.advertise_where](data-sources--workload--reference--group-017.md#canonical-2023022020012320-2131122313011201-2300000120322300-1332100012103313-0221311213321232-0302220100100100-0302230032132330-0021310212030123)
- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1001331030330331-1223200200201333-1321331330020020-3220331320211002-2301223300031003-1201011200330103-1233322010230223-0010310231132321"></a>

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

<a id="canonical-1203130300000320-3113002310102012-1121323133233120-3111223310130001-1230011001033323-2332321021000223-0310111132320300-1312111322330201"></a>

## Direct properties — virtual_site / 202200131033 / 3

<a id="canonical-2002221300202301-3023313302220213-0101200303011100-1312023020203132-2300013333301100-1212233312121030-0100200331232130-0101023211031021"></a>

<a id="canonical-0033013300111002-2030100133232222-2021321001012230-1301023303012203-0003230031111012-2001133212232210-2200332303311311-0012201231021032"></a>

## name property — virtual_site / 202200131033 / 4

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

<a id="canonical-0300303201102131-2331023330012011-0210320222213002-1123023010120033-0120302113131320-0121222222101110-0310210101100010-3003022323211222"></a>

<a id="canonical-0302303202332133-0103231311103031-0022211210020231-0330031231003001-1001130200311100-3221232311123021-3332121223303212-1203003221130330"></a>

## namespace property — virtual_site / 202200131033 / 5

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

<a id="canonical-0300200302203012-2321023133233203-0013131131213300-1102301030213121-2102201213110200-2131220131221201-3111312012222103-2223110233020012"></a>

<a id="canonical-3302231320101323-0203021010210100-2102331301121201-0113311300323022-3212133211121123-1022313012020310-0323212110002021-2103113011220233"></a>

## tenant property — virtual_site / 202200131033 / 6

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

<a id="canonical-1000223302010332-0302012102021113-2120222221212111-2020321331011300-3100313102232302-0110103213000231-2030012032320221-3232011002021103"></a>

## Next pages — virtual_site / 202200131033 / 7

- [stateful_service.advertise_options.advertise_custom.advertise_where.vk8s_service](data-sources--workload--reference--group-017.md#canonical-3011201020212332-2011332133320103-1300120313022330-3010230001020133-2322002100330210-1233003022123011-1021320030320131-2111111310111303)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003022000111201-3011030013031311-0110230020003321-2201032101132032-0230200103331132-1320220331300310-1121302101122221-2132033023103010"></a>

## stateful_service.advertise_options.advertise_custom.ports — ports / 000223311121 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- stateful_service.advertise_options.advertise_custom.ports

<a id="canonical-2332031023012310-1211012112123322-0230032031121211-3231301231203332-3022010032300331-0323330323120321-1023131332101033-1123302030303131"></a>

Type: `"list"`. Computed.

Ports. Ports to advertise.

Upstream description:

Ports to advertise.

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

<a id="canonical-0021302323230323-1323331302133211-3121203312003110-0231131213123220-0003323003023011-3231320231132120-3300101331030030-2012230013331003"></a>

## Direct properties — ports / 000223311121 / 3

- [http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201): complete subsection reference.

- [port](data-sources--workload--reference--group-020.md#canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101): complete subsection reference.

- [tcp_loadbalancer](data-sources--workload--reference--group-020.md#canonical-3021012100101331-3130231312320112-3123200110222222-2102132131213021-0312223301023023-0100202023203122-0211120233110023-0332201103002102): complete subsection reference.

<a id="canonical-1002310321121312-0111033231013233-2323001013203311-0102212203323232-3011003132003021-2012230101133321-0322321133310113-1231320232213032"></a>

## Next pages — ports / 000223311121 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.port](data-sources--workload--reference--group-020.md#canonical-3020203222030212-3033113130221110-2320320300332011-1001030302320102-2022321230310000-0112020113321222-3002102322302000-3221222122031101)
- [stateful_service.advertise_options.advertise_custom.ports.tcp_loadbalancer](data-sources--workload--reference--group-020.md#canonical-3021012100101331-3130231312320112-3123200110222222-2102132131213021-0312223301023023-0100202023203122-0211120233110023-0332201103002102)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032300123313101-2233003220033103-1033223200232212-2102110302013001-1110021031000003-1021002013023111-3033031320013212-2312001000103103"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer — http_loadbalancer / 032031231320 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer

<a id="canonical-2022122301122103-3000320010020112-0122200002001233-2031231101310101-1223030013300123-2120321210212123-3112121031201101-3110312113333023"></a>

Type: `"single"`. Computed.

Configuration parameter for http loadbalancer.

Upstream description:

HTTP/HTTPS Load balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-loadbalancer_type": "[\"http\",\"https\",\"https_auto_cert\"]",
  "x-ves-oneof-field-route_choice": "[\"default_route\",\"specific_routes\"]"
}
```

<a id="canonical-0102332101103010-1023013331203230-3203220200202021-1231102022312233-3103212210031321-2232200103203300-0030320202023012-2112031232000330"></a>

## Direct properties — http_loadbalancer / 032031231320 / 3

- [default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220): complete subsection reference.

<a id="canonical-0330302213321211-2111203033222023-2002303332300311-2000303132023011-1102231133113203-3001023021322100-3003133021320101-3222132310103302"></a>

<a id="canonical-0102330022122033-3010313020132221-3131103211311201-1012102332222030-2312220112011210-2133001230331021-1112200112030102-3202213323022312"></a>

## domains property — http_loadbalancer / 032031231320 / 4

Type: `["list", "string"]`. Computed.

List of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form Domain search order: 1. Exact domain names: \`\` is invalid
Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the..

Upstream description:

A list of domains (host/authority header) that will be matched to loadbalancer. Wildcard hosts are
supported in the suffix or prefix form

Domain search order: &#8203;1. Exact domain names: \`\`www&#46;example.com\`\`. &#8203;2. Prefix
domain wildcards: \`\`\*.example.com\`\` or \`\`\*.bar.example.com\`\`. &#8203;3. Special wildcard
\`\`\*\`\` matching any domain.

Wildcard will not match empty string. E.g. \`\`\*.example.com\`\` will match \`\`bar.example.com\`\`
and \`\`baz-bar.example.com\`\` but not \`\`.example.com\`\`. The longest wildcards match first.
Wildcards must match a whole DNS label. E.g. \`\`\*.example.com\`\` and \*.bar.example.com are
valid, however \`\`\*bar.example.com\`\` or \`\`\*-bar.example.com\`\` is invalid

Domains are also used for SNI matching if the loadbalancer type is HTTPS Domains also indicate the
list of names for which DNS resolution will be done by VER.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [http](data-sources--workload--reference--group-017.md#canonical-0210000123110012-0203033330020312-0012332023112212-1310121212222023-3210023231230233-1213201111200133-1221001123112101-1011211000210303): complete subsection reference.

- [https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003): complete subsection reference.

- [https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100): complete subsection reference.

- [specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010): complete subsection reference.

<a id="canonical-2330022102313302-1323100033003312-2322122330033103-1312211120222312-1210102032131311-1123120221303102-1022233021210303-0202021001311000"></a>

## Next pages — http_loadbalancer / 032031231320 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http](data-sources--workload--reference--group-017.md#canonical-0210000123110012-0203033330020312-0012332023112212-1310121212222023-3210023231230233-1213201111200133-1221001123112101-1011211000210303)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](data-sources--workload--reference--group-019.md#canonical-2222310212003321-3311011031323313-1033123102310023-1131322023022203-0310203302121223-0212302033021012-1111313132200003-2303112232301100)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](data-sources--workload--reference--group-019.md#canonical-1013211100312123-2303001111231323-0313131132122000-0333201330302230-1031323200021000-1112233102332030-3102023323212313-2010030013121010)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303030230012203-0300003213313012-2302021332010201-1222032303023322-2123203333331301-1032011003003200-0131110233123201-0222300111300020"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route — default_route / 202323333221 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route

<a id="canonical-1331001032123101-1232131002331201-3012212220101231-3303000331300212-0023321111032120-0030233301312201-3023230301032110-1202203331102313"></a>

Type: `"single"`. Computed.

Configuration parameter for default route.

Upstream description:

Default route matching all APIs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-2001012001022320-3212213103230302-0002031022203303-1211101221201323-1220031210303312-3210130233311210-0201322011321331-2131022201112101"></a>

## Direct properties — default_route / 202323333221 / 3

- [auto_host_rewrite](data-sources--workload--reference--group-017.md#canonical-3020030333232332-0113230332320033-0203210331002213-1200033131111332-0111322130123223-1323233013002303-0232323031301101-0322221310113002): complete subsection reference.

- [disable_host_rewrite](data-sources--workload--reference--group-017.md#canonical-3133001210012120-1000122332022200-3132023300022300-2211300033302131-3131012322223130-3131032012210310-2121301122033202-2222323213321303): complete subsection reference.

<a id="canonical-1121330220210123-2302103130113233-1011012021110210-0011330031133232-1210123102010323-1131333312001012-3132312003201301-0202202030002212"></a>

<a id="canonical-1220300310010010-1022233012032030-3113031212310133-0010322012130131-0111302111112212-0312213212103231-1111310011010122-1110012011301220"></a>

## host_rewrite property — default_route / 202323333221 / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1202332010012233-1000111211210221-1011323030132211-0321133202230222-0021032011001201-3000010333211321-0000122021121232-1110022003120012"></a>

## Next pages — default_route / 202323333221 / 5

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite](data-sources--workload--reference--group-017.md#canonical-3020030333232332-0113230332320033-0203210331002213-1200033131111332-0111322130123223-1323233013002303-0232323031301101-0322221310113002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite](data-sources--workload--reference--group-017.md#canonical-3133001210012120-1000122332022200-3132023300022300-2211300033302131-3131012322223130-3131032012210310-2121301122033202-2222323213321303)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3020030333232332-0113230332320033-0203210331002213-1200033131111332-0111322130123223-1323233013002303-0232323031301101-0322221310113002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031312003000301-2323032112133110-3130311333100133-2112301102331331-0212030312211331-1310223222202033-1220120302111202-0200212013023200"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite — auto_host_rewrite / 100333213202 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.auto_host_rewrite

<a id="canonical-1031322330221100-3010131221113312-0312212133222203-0312131302031310-1030030222130131-3112012100300300-2331130302222221-2323002211102300"></a>

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

<a id="canonical-3311300333213313-1021220200303012-2332000330111130-2302031231313123-2010002221002332-0023313221110303-0012003310023012-2302222130132233"></a>

## Direct properties — auto_host_rewrite / 100333213202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311013002201003-3323103022321233-3301022231303201-1133300201133213-1200221301323300-3013121201111111-2330203303213232-2230030322020323"></a>

## Next pages — auto_host_rewrite / 100333213202 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3133001210012120-1000122332022200-3132023300022300-2211300033302131-3131012322223130-3131032012210310-2121301122033202-2222323213321303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311011311022123-0010320121211021-1131322013031312-3002203200002212-2310111030302023-2032021022122123-1203003303031302-0012301223012002"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite — disable_host_rewrite / 233132130210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route.disable_host_rewrite

<a id="canonical-0011213122113022-1113132022020021-1211320213133213-1003133113011012-0233320311303030-2213203112020302-0023313232330213-1013333300011210"></a>

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

<a id="canonical-1123321102221122-2131122302112110-2210033022032210-1031231020030332-1131022220001101-0030332030103122-0000233031330121-0021310021032203"></a>

## Direct properties — disable_host_rewrite / 233132130210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031133211312302-1230331001133323-0233230100210223-0120212210132200-0033301303003333-2323120223122230-0303101202100312-0131311101311030"></a>

## Next pages — disable_host_rewrite / 233132130210 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.default_route](data-sources--workload--reference--group-017.md#canonical-2020222323220111-1331211203202000-3201020021101201-0111321110223013-1230323223112113-0223023100222323-0302123123231102-2203231302311220)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0210000123110012-0203033330020312-0012332023112212-1310121212222023-3210023231230233-1213201111200133-1221001123112101-1011211000210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313333012103000-2113000210011310-0331233221200123-2203310322233110-2212330301202020-1210003031311312-2031233121031323-2001323321302112"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http — http / 130122212232 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.http

<a id="canonical-2131221122103110-2132101202331321-0113230103212231-0312010022332110-3302120211232230-0022102012022231-1022000221133313-3023210113333111"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-2021322032113220-2000312301311030-2230100333032232-0323331131000001-2322203113023033-0022010221230120-3013113232000100-3202222312233021"></a>

## Direct properties — http / 130122212232 / 3

<a id="canonical-0100113010111322-3312032011311323-1232313301113332-1111201103222311-0232123331200231-0112012030222113-0011310031111201-2031333101010120"></a>

<a id="canonical-0211022330002032-0130322322020132-3102321103211122-1032210132021302-3320303220133131-0130031010312212-1112110131311121-0000033103113013"></a>

## dns_volterra_managed property — http / 130122212232 / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0112231201101032-0113212012320101-0031020330200200-2202001000121002-2022100013210011-3213320231100303-0320003201020321-0013213200023133"></a>

<a id="canonical-0101123323232012-3200012102023311-1323012222203233-1302201323102203-0102300331203310-1333102121020030-0122323333103300-1301013001020220"></a>

## port property — http / 130122212232 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0300020030233032-1010223203221332-2302220121102332-1221202313332222-1203110020033003-1311032001110102-3023102100030220-1131320132300310"></a>

<a id="canonical-3202033122101113-2022300121312330-1201310200230031-2103020332003000-3130123201013113-0111103111222103-3131222020212222-3010003313220201"></a>

## port_ranges property — http / 130122212232 / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2231122312011330-0020121232301333-0330212020010332-0031132210200032-1133130220301330-3112322201200123-3102320130212020-2022312200030001"></a>

## Next pages — http / 130122212232 / 7

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312222220301231-1002313220320120-0201211001213111-3203002002233220-1103322033310212-1321113331023120-3101323010102121-0231122113211330"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https — https / 201323331310 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https

<a id="canonical-0102303302013230-1330320233233310-2113122031113012-3130002023203133-2013130231201011-2332120303302122-1322302211021320-2001313330201211"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-3030120210313213-1012220321032003-2230330322013213-1320001310030120-2023312011301310-1013011213013300-3301113033003011-0111201133100012"></a>

## Direct properties — https / 201323331310 / 3

<a id="canonical-2023211030203023-0332211033031000-3113001110013310-3233332332303200-0322212201120211-3303001121031001-3031130200121303-2222110302003113"></a>

<a id="canonical-2002010222121212-0210320213302122-2103322020133303-3023210123233321-3213023100202202-2133031311013331-0323212210021032-0323003031213103"></a>

## add_hsts property — https / 201323331310 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0333131301310121-1113202230332201-0102101202110101-0000230011233302-2101323222031312-2300331321121321-2200132031303122-0203032023303131"></a>

<a id="canonical-2013312231110302-0030011103003332-2221331010100210-1123123213111123-1313310010222231-0131022233023130-1323222322102302-2311220110331020"></a>

## append_server_name property — https / 201323331310 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203): complete subsection reference.

<a id="canonical-1202230333231220-2320221010310103-0010310113130021-1323113233231003-0010210312321210-3102221330311201-0320221210101102-2220130230223013"></a>

<a id="canonical-3222120111111001-3330312303010323-0210033202233023-3033112032211113-2220332310002303-3022230121131303-0331222300131020-2200103103130222"></a>

## connection_idle_timeout property — https / 201323331310 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--workload--reference--group-017.md#canonical-0023232310300200-3223031333313121-2212120213302110-1333203102110320-3003312012200002-2202003121222200-1031312300201132-0031122320322313): complete subsection reference.

- [default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0012001103030201-3203200132211322-3232213203333223-3312110000013103-3010201310230301-0300213202122012-0103012122313002-1023213201333300): complete subsection reference.

- [disable_path_normalize](data-sources--workload--reference--group-017.md#canonical-3233221102122003-1130211310313313-2323010000022332-1312131231102103-3101321021113300-1323200323023012-1230101113333022-0202333112310002): complete subsection reference.

- [enable_path_normalize](data-sources--workload--reference--group-017.md#canonical-2013223221202330-3132232032332332-1132202330202010-3210121032132010-2033202003222023-1321220120101313-2201212323132112-0331002302300020): complete subsection reference.

- [http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331): complete subsection reference.

<a id="canonical-2323233222130220-3222032011200211-3223310000021123-3211020332212011-1012331300011200-2232321230210001-1332032012221223-0133001333023032"></a>

<a id="canonical-3233002132112201-1002310111111113-3212230231101312-1223323320322233-3322100210311212-2011132302310231-1233212211101223-3132120002212100"></a>

## http_redirect property — https / 201323331310 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](data-sources--workload--reference--group-018.md#canonical-2233212220331030-3333132301233302-1131111310300010-1020030131001213-2022003333122002-1023230213320033-0102023112233011-0323230323302323): complete subsection reference.

- [pass_through](data-sources--workload--reference--group-018.md#canonical-3210111122032323-1101022110110331-1023201230033121-2011003332013100-1002021032212313-0320000121303232-3101312000213323-3232133323310302): complete subsection reference.

<a id="canonical-0312003211101322-3200003333232333-0021102011221020-3203313002313021-0221103131023013-1101113301101100-1223220333232112-0331203131231020"></a>

<a id="canonical-1033013110301300-1121221123112231-1322221120032100-1111120011312231-1331313133202100-2233120233112310-0100112133023310-3103220321010031"></a>

## port property — https / 201323331310 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1203111001210101-3230123313322322-3132030210020303-0101320020220232-1231002112311233-3233233330022101-1123033023222021-1103032120323213"></a>

<a id="canonical-0230033211201310-2013033201033303-2102102013002312-2312310120211212-1331111330323131-1121123221202130-3013211022202213-0222202303300012"></a>

## port_ranges property — https / 201323331310 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-1303010020132212-1202002030101021-2213112100111012-0022123103233033-2233010000031101-3200322032230202-0320211113111201-3231123033031232"></a>

<a id="canonical-0201323332011210-2301102030000230-0302223102002312-0023012012100130-3030322012210220-3103202002312033-1101303110010330-1231221210213100"></a>

## server_name property — https / 201323331310 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--workload--reference--group-018.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320): complete subsection reference.

- [tls_parameters](data-sources--workload--reference--group-018.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302): complete subsection reference.

<a id="canonical-1000000300310213-2233102330230300-1130233223222321-2310132030130033-3312213100001332-1000132103102311-0221302322120003-2200130312011212"></a>

## Next pages — https / 201323331310 / 11

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header](data-sources--workload--reference--group-017.md#canonical-0023232310300200-3223031333313121-2212120213302110-1333203102110320-3003312012200002-2202003121222200-1031312300201132-0031122320322313)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0012001103030201-3203200132211322-3232213203333223-3312110000013103-3010201310230301-0300213202122012-0103012122313002-1023213201333300)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize](data-sources--workload--reference--group-017.md#canonical-3233221102122003-1130211310313313-2323010000022332-1312131231102103-3101321021113300-1323200323023012-1230101113333022-0202333112310002)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize](data-sources--workload--reference--group-017.md#canonical-2013223221202330-3132232032332332-1132202330202010-3210121032132010-2033202003222023-1321220120101313-2201212323132112-0331002302300020)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.non_default_loadbalancer](data-sources--workload--reference--group-018.md#canonical-2233212220331030-3333132301233302-1131111310300010-1020030131001213-2022003333122002-1023230213320033-0102023112233011-0323230323302323)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.pass_through](data-sources--workload--reference--group-018.md#canonical-3210111122032323-1101022110110331-1023201230033121-2011003332013100-1002021032212313-0320000121303232-3101312000213323-3232133323310302)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_cert_params](data-sources--workload--reference--group-018.md#canonical-2010001011230201-3020111320131112-2333021133221211-0200112033223200-3003002233220012-1123230332003001-0103023230030232-1202122133301320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](data-sources--workload--reference--group-018.md#canonical-3230312021230322-3310001132303313-1110012112202213-3330322321030323-1121221221133120-2321102112023333-2011212100221003-0220120111022302)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203121220012233-0231300121011232-0221200211133032-3022032330032300-3030101000131312-0102312202201120-0320010330121210-0010010331031232"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options — coalescing_options / 003310211331 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options

<a id="canonical-0011120122220210-3300331323112202-3111121200130022-2201011323330020-1022111012122223-2331123132332331-3313300313231231-0310110103110001"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-1110332210313123-2313103102222300-1322131023121011-3103312311120213-2112021031321100-2131131003310130-2322123200033302-0300203331333001"></a>

## Direct properties — coalescing_options / 003310211331 / 3

- [default_coalescing](data-sources--workload--reference--group-017.md#canonical-3302211123200123-3303023033113022-0310023022321220-3301022002322232-2121001320100231-2111032202120100-0132133100231101-1001313330303202): complete subsection reference.

- [strict_coalescing](data-sources--workload--reference--group-017.md#canonical-1021033211222331-3220122130223113-1230202132011320-3032113310232333-0302213321012102-0232220222320001-3310023320113101-0333002002031212): complete subsection reference.

<a id="canonical-1002002220102100-1122120020023030-0302311133001001-2201313131030003-3331103103020100-0310231010331323-1300033231201230-0232332300003102"></a>

## Next pages — coalescing_options / 003310211331 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing](data-sources--workload--reference--group-017.md#canonical-3302211123200123-3303023033113022-0310023022321220-3301022002322232-2121001320100231-2111032202120100-0132133100231101-1001313330303202)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing](data-sources--workload--reference--group-017.md#canonical-1021033211222331-3220122130223113-1230202132011320-3032113310232333-0302213321012102-0232220222320001-3310023320113101-0333002002031212)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3302211123200123-3303023033113022-0310023022321220-3301022002322232-2121001320100231-2111032202120100-0132133100231101-1001313330303202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112123310000322-1133303313213222-1333322202322010-1031132033131031-1010012311322223-3213322113222303-3200102201323230-2122011000123203"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing — default_coalescing / 212012213003 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.default_coalescing

<a id="canonical-3000023113021233-3313310221312132-1100013300012132-2103221133011202-1200000102321000-3332323223221022-0303122031111312-0231003003000120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-0212233001201122-2113110110010211-0000223102233111-1102033000202123-1003101122313201-2133121130330323-3100212202211102-1103313123131321"></a>

## Direct properties — default_coalescing / 212012213003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302003322132201-3133200320302003-2101101212231113-2230103311111013-1022123013233213-3321022133313213-3120323233013311-3212133001221313"></a>

## Next pages — default_coalescing / 212012213003 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1021033211222331-3220122130223113-1230202132011320-3032113310232333-0302213321012102-0232220222320001-3310023320113101-0333002002031212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302031101322000-2013320023303202-0033013312020223-1010010220212202-0133203120330331-2112022230123110-2313230202331033-3030321030221112"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing — strict_coalescing / 010312023112 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options.strict_coalescing

<a id="canonical-1103030131202011-1332212233021202-0023200200232010-2133303321332211-3110133033223130-0333320100032322-0333233333323111-2032332111212220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-2311130332230131-0331200032300310-3232013113130003-3133001220120221-1011330103223323-2220302102031001-0010200221113230-1302232132231211"></a>

## Direct properties — strict_coalescing / 010312023112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020303112220002-3220131322032123-2210323312002012-3121010131212300-3012211032001211-3202310323111022-1322110302000022-2322031230212102"></a>

## Next pages — strict_coalescing / 010312023112 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.coalescing_options](data-sources--workload--reference--group-017.md#canonical-3222230311111111-3020113330012230-0002302032113022-3221131030313033-2131321231133022-0132211032010123-0101033132301003-2200000120321203)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0023232310300200-3223031333313121-2212120213302110-1333203102110320-3003312012200002-2202003121222200-1031312300201132-0031122320322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000103031030033-2200210023201203-2100123112231132-3120033100231022-1120333111012133-0312113200221221-0330010212200122-1311030302031033"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header — default_header / 013103233113 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_header

<a id="canonical-0202120112230021-2330211320002031-2020211310301333-1110202220210121-2012310322221200-2012021320102013-2311200031131121-2231203031023301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-3313031113103132-1333030232013212-1103301012203321-3221131120203131-2321220030111322-3101031302311330-2100212031011102-1021320231230122"></a>

## Direct properties — default_header / 013103233113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212110021220000-2201223311132011-0033020203231020-0201303131030220-0330132131233303-2003222313033133-1332021101330302-3303022310322003"></a>

## Next pages — default_header / 013103233113 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0012001103030201-3203200132211322-3232213203333223-3312110000013103-3010201310230301-0300213202122012-0103012122313002-1023213201333300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103220103002232-0122011322101132-3213311222211302-0330311003330112-1331211020002123-0030233102121003-2022303132203210-0213321003312312"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer — default_loadbalancer / 303230130100 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.default_loadbalancer

<a id="canonical-2013333302102123-2201103223101032-3132320311301202-2312232233031031-3312300123303332-2310003023110301-1021002201231332-0333001312302301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-3332121120122022-0020031331311110-1031010223202011-0301310121212032-1333210333302013-0023233323121232-3011131000332003-3110111200113302"></a>

## Direct properties — default_loadbalancer / 303230130100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321322002021003-3022232031212002-0003122002231011-2303120322001300-0023023213112310-1200121001313203-1031013103201311-2012103312033300"></a>

## Next pages — default_loadbalancer / 303230130100 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-3233221102122003-1130211310313313-2323010000022332-1312131231102103-3101321021113300-1323200323023012-1230101113333022-0202333112310002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031320131213231-0013113332322112-1301313030113000-1003130223003112-0010030233332213-2230130121231100-3302233010022113-1111113122003000"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize — disable_path_normalize / 110030323020 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.disable_path_normalize

<a id="canonical-3130333013230322-3321013313020312-1111133211031213-0130213030223330-3022303322110200-2113012212113210-2230000123112232-3113222101031010"></a>

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

<a id="canonical-1322231001311121-1033113013011001-0032203112133110-3332102122021203-3310333331213302-1021330300323002-0002122132231312-2130002111322210"></a>

## Direct properties — disable_path_normalize / 110030323020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013020130101302-2312213120332102-3032312120032120-0101131033000312-3111222032010212-1302321221322011-0123132200022331-2313323331322302"></a>

## Next pages — disable_path_normalize / 110030323020 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2013223221202330-3132232032332332-1132202330202010-3210121032132010-2033202003222023-1321220120101313-2201212323132112-0331002302300020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022003330333203-2322012130012322-2131232131220112-2102230102200333-2212300010012313-2102031122311023-2102302011100311-2112330202331001"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize — enable_path_normalize / 130213110210 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.enable_path_normalize

<a id="canonical-3203102120213132-3220323212200121-3002102321212130-2331013210130220-3003122323301013-1311123222311121-3310310201023330-0103031112103101"></a>

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

<a id="canonical-1102212311311021-0233333301331002-0210332001310012-0320121002222311-1321003230320302-0110103033311310-1031132033120221-1023023332010032"></a>

## Direct properties — enable_path_normalize / 130213110210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022311203321320-1121100002211333-1020202300110230-0002000323232233-2201213010101332-0133311302031121-0211210000310332-0100232032210203"></a>

## Next pages — enable_path_normalize / 130213110210 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112122211012233-3131010221010231-1323020210131131-3203333320011112-2213331322122032-0020200030211021-1313233102131300-3010201113103000"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options — http_protocol_options / 023230033111 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options

<a id="canonical-1100300322212023-1103310203022101-2313321131110032-2222221311201103-2112212101123111-3210202301113001-0200300103232121-3203033103021131"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-0222333101110312-0202100131200132-1230132220132332-0311331310112111-3303032033113132-0133113111100033-1231212131213013-1311033232230220"></a>

## Direct properties — http_protocol_options / 023230033111 / 3

- [http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--workload--reference--group-017.md#canonical-1120022312332203-3300020131303230-3333113203222313-3332321311320301-1103311333100132-3320203113100031-0220132331312223-1013030020310030): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--workload--reference--group-018.md#canonical-2002032323000231-2311320210121132-0203220023130321-0313122310230003-2230102333222001-1330130233002122-1002200201123013-1232113113210320): complete subsection reference.

<a id="canonical-1223000230000120-1032212113111120-0201113232210022-3212223200212313-3302210233323230-2210301121202331-2010210120302123-2212320030003223"></a>

## Next pages — http_protocol_options / 023230033111 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--workload--reference--group-017.md#canonical-1120022312332203-3300020131303230-3333113203222313-3332321311320301-1103311333100132-3320203113100031-0220132331312223-1013030020310030)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--workload--reference--group-018.md#canonical-2002032323000231-2311320210121132-0203220023130321-0313122310230003-2230102333222001-1330130233002122-1002200201123013-1232113113210320)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131021010220003-3331013100303122-2121202221223023-3322002303021231-2220202001121302-2133012212111120-1311030121032331-1121101131302011"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 221010010030 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3323323110312321-2123331132323201-3331213102132330-2013313031213323-2311233133032300-1200022312132311-3132202032030231-2212332301110003"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3202130103210122-2331331220331012-1010221103012332-2033330322121323-3211300132022233-0202201102301322-3111333013333312-0101303112321232"></a>

## Direct properties — http_protocol_enable_v1_only / 221010010030 / 3

- [header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300): complete subsection reference.

<a id="canonical-3022313111223101-1220133321131011-2213131200312211-3031103010130032-2313002130201011-0320132111120330-3133011322333111-1111132130203200"></a>

## Next pages — http_protocol_enable_v1_only / 221010010030 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212003331213311-2020021303323013-0121020322101313-1022311230220130-3310011310300002-1132111311013103-1300122311033230-0112331213330202"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 200202031212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0012320031220201-0103233013103012-3131300313101000-0323101131102232-1031012112012111-1211213012320221-0020330123300121-2231223122200223"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-1212303122001131-0100100231230203-3012013003130002-3231130333010201-1331031111120011-3020123330302012-1333202103031002-0122231210122123"></a>

## Direct properties — header_transformation / 200202031212 / 3

- [default_header_transformation](data-sources--workload--reference--group-017.md#canonical-1320331202000003-1323303202003022-3131303223233211-3111011131200010-0133122002030230-3201201100212023-2222333220223011-1020300112200221): complete subsection reference.

- [preserve_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-0103212202233220-2332113321232210-1122020103202310-3220032001031210-0012332323221332-1321133202231220-0211003210100001-1233202320113332): complete subsection reference.

- [proper_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-2133201310333311-3030302202212131-0003112221112301-0223321231302320-2230011022230212-0021221031133000-2303220232003220-2212222003230223): complete subsection reference.

<a id="canonical-1132131133132001-2313113010020232-2102210123010202-2031231323122030-0002033111201220-1332112232213030-3101313113312120-3123021012113102"></a>

## Next pages — header_transformation / 200202031212 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--workload--reference--group-017.md#canonical-1320331202000003-1323303202003022-3131303223233211-3111011131200010-0133122002030230-3201201100212023-2222333220223011-1020300112200221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-0103212202233220-2332113321232210-1122020103202310-3220032001031210-0012332323221332-1321133202231220-0211003210100001-1233202320113332)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--workload--reference--group-017.md#canonical-2133201310333311-3030302202212131-0003112221112301-0223321231302320-2230011022230212-0021221031133000-2303220232003220-2212222003230223)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1320331202000003-1323303202003022-3131303223233211-3111011131200010-0133122002030230-3201201100212023-2222333220223011-1020300112200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123220312030302-0313322320031110-3312212222101212-0331222331332001-0312321331000323-0330122023330310-3030303033120001-2303020012100231"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 021322220212 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1113313133310002-1223100012302300-0213032030332132-0031221100333112-2013123010303031-1033013030323303-1022230213221010-2000202021232112"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0133010312003210-1102033301102310-2211202011230313-0333023320000120-2212123301131132-3031120130203020-0122333002011210-0001133133203032"></a>

## Direct properties — default_header_transformation / 021322220212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211331013313021-3011132221301022-0223230000130102-0003320021312212-2120001033311020-1132321013202231-2320030330330211-2303111332311223"></a>

## Next pages — default_header_transformation / 021322220212 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-0103212202233220-2332113321232210-1122020103202310-3220032001031210-0012332323221332-1321133202231220-0211003210100001-1233202320113332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130100210210220-3120021030300110-1030022220010030-1021123033122312-2230200113200011-0033220310331013-3030102020100100-1032131300311212"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 212000220302 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1210302133332333-3013010011032123-0102022120011010-1300303021111001-2013321232010311-0333000313121100-1330100332213000-3212000112110211"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1001301233013202-3031121312103203-3113110330312323-1211230300032231-1011133100213030-3303103021200030-2232120200130132-0302213030032202"></a>

## Direct properties — preserve_case_header_transformation / 212000220302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010121113011131-1303331210132211-1310030021300003-0032323213222013-2102213003133031-3122120223012122-2303131312003213-3333023001130201"></a>

## Next pages — preserve_case_header_transformation / 212000220302 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-2133201310333311-3030302202212131-0003112221112301-0223321231302320-2230011022230212-0021221031133000-2303220232003220-2212222003230223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002232132323012-3112202320032101-0231113031330000-3211332300133202-0132103201132222-1332232332212013-0232312200133010-1110132120132333"></a>

## stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 021000113222 / 2

Breadcrumbs:

- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)
- [Property reference](data-sources--workload--reference--group-001.md#canonical-0301022022220103-1212310020010033-3233011010231130-1311133001133002-1331212010221022-1102300330221113-1112020030323222-3211222302103031)
- [stateful_service](data-sources--workload--reference--group-017.md#canonical-2132303011012103-3220213303333100-0323112032103022-2013230102321010-3022210221311323-0233232113013303-1322202323231301-3212322330102113)
- [stateful_service.advertise_options](data-sources--workload--reference--group-017.md#canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220)
- [stateful_service.advertise_options.advertise_custom](data-sources--workload--reference--group-017.md#canonical-1111211210302023-1130301301031112-0133023021111122-2200322020203101-0223203003301101-2022121313331313-1323300011032311-0202020122011323)
- [stateful_service.advertise_options.advertise_custom.ports](data-sources--workload--reference--group-017.md#canonical-0013330003233312-1103212221031032-2233120012231121-3101311130002333-2222211213103202-3221232132020020-1233011032133130-3130112100210221)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer](data-sources--workload--reference--group-017.md#canonical-0323333020010100-1132122030202301-0220233131023222-0121213003321102-3000102211200300-2122203321110300-0131101323132000-2220030303220201)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https](data-sources--workload--reference--group-017.md#canonical-1033211310210331-3122131003202312-2011303010012203-1210330231200011-0103302332022003-0002020201221022-0330311011330033-0101322031121003)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options](data-sources--workload--reference--group-017.md#canonical-1120222011000333-0211100013303031-0300010010332233-1031300033323310-3030100021331110-1303033230310102-2010030333330102-1101231301332331)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--workload--reference--group-017.md#canonical-2030232200121110-2100031231111330-2023320230132100-0300321031312321-0222130321013122-2202030033211322-1330310030303102-2312021102321131)
- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-3311112121000322-0002013333033031-2203123223101302-3313011003021031-3301121033212320-2030122030303120-1332130113020201-3121001120200223"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2310120211231320-0230022011003031-3032000033032030-0202300210012332-0132312230211013-2001211310032011-1032221312101210-1232131200333002"></a>

## Direct properties — proper_case_header_transformation / 021000113222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3333311322300111-1230302221120301-3222300300333231-3031230301122023-2213221201331120-1123013232321021-1211310211302020-3331032321022122"></a>

## Next pages — proper_case_header_transformation / 021000113222 / 4

- [stateful_service.advertise_options.advertise_custom.ports.http_loadbalancer.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--workload--reference--group-017.md#canonical-2200203320303131-2023111232202232-1211300022202210-1000312310010200-2031131200010101-0320121333322101-0022110332101033-1213011011100300)
- [xcsh_workload](../data-sources/workload.md#canonical-1002113323301123-1000231011222222-2130003220201313-1332320021201220-0322102223332102-3110303113322020-1203310220131002-3110013201311100)

<a id="canonical-1120022312332203-3300020131303230-3333113203222313-3332321311320301-1103311333100132-3320203113100031-0220132331312223-1013030020310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
