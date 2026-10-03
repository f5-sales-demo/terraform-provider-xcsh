---
page_title: "xcsh_service_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy reference."
---

# xcsh_service_policy reference

<a id="canonical-3023321122200311-0122113303023133-0111010223332121-1101031223201220-2113103320120003-0103111103332302-3113232111133011-2112003031102302"></a>

## rule_list.rules.spec.request_constraints — request_constraints / 123230213121 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.request_constraints

<a id="canonical-2112123331021132-0331130211222303-1220320200120130-2123332130032032-1232330002003032-1003011220032133-0200332210010013-3303233302321003"></a>

Type: `"single"`. Computed.

Configuration parameter for request constraints.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_cookie_count_choice": "[\"max_cookie_count_exceeds\",\"max_cookie_count_none\"]",
  "x-ves-oneof-field-max_cookie_key_size_choice": "[\"max_cookie_key_size_exceeds\",\"max_cookie_key_size_none\"]",
  "x-ves-oneof-field-max_cookie_value_size_choice": "[\"max_cookie_value_size_exceeds\",\"max_cookie_value_size_none\"]",
  "x-ves-oneof-field-max_header_count_choice": "[\"max_header_count_exceeds\",\"max_header_count_none\"]",
  "x-ves-oneof-field-max_header_key_size_choice": "[\"max_header_key_size_exceeds\",\"max_header_key_size_none\"]",
  "x-ves-oneof-field-max_header_value_size_choice": "[\"max_header_value_size_exceeds\",\"max_header_value_size_none\"]",
  "x-ves-oneof-field-max_parameter_count_choice": "[\"max_parameter_count_exceeds\",\"max_parameter_count_none\"]",
  "x-ves-oneof-field-max_parameter_name_size_choice": "[\"max_parameter_name_size_exceeds\",\"max_parameter_name_size_none\"]",
  "x-ves-oneof-field-max_parameter_value_size_choice": "[\"max_parameter_value_size_exceeds\",\"max_parameter_value_size_none\"]",
  "x-ves-oneof-field-max_query_size_choice": "[\"max_query_size_exceeds\",\"max_query_size_none\"]",
  "x-ves-oneof-field-max_request_line_size_choice": "[\"max_request_line_size_exceeds\",\"max_request_line_size_none\"]",
  "x-ves-oneof-field-max_request_size_choice": "[\"max_request_size_exceeds\",\"max_request_size_none\"]",
  "x-ves-oneof-field-max_url_size_choice": "[\"max_url_size_exceeds\",\"max_url_size_none\"]"
}
```

<a id="canonical-3030310330131331-2310020110220302-1110322312100301-2123301231223002-2022023300033333-1002033233230232-1330202211200331-1022302031003201"></a>

## Direct properties — request_constraints / 123230213121 / 3

<a id="canonical-1131210031223213-1132313020110003-1011223120333302-2100330202123121-3223011130121221-3201303130330123-0203131231131233-1331023213220011"></a>

<a id="canonical-1320121030133213-3111201013121020-2332022322033002-3202312111122131-1121031222331313-0120210203033103-0022223323231112-2022300313330321"></a>

## max_cookie_count_exceeds property — request_constraints / 123230213121 / 4

Type: `"number"`. Computed.

Match on the Count for all Cookies that exceed this value. Exclusive with
\[max\_cookie\_count\_none\]

Upstream description:

Exclusive with \[max\_cookie\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_count_none](data-sources--service_policy--reference--group-003.md#canonical-1233023322123212-3210231301202012-3232321110232302-3303331333320223-0000211210010011-3020010300232033-1020132130321000-1121233123310122): complete subsection reference.

<a id="canonical-3031223001320213-2021311110120300-3032323002000221-0031033113031331-2101111233203233-0012311311100212-0111322132031123-0220330111110331"></a>

<a id="canonical-3302210312220302-3331023000332131-2322332032223100-0312030013300100-0333122022311211-0120201030111101-0101232203322313-1001222320321021"></a>

## max_cookie_key_size_exceeds property — request_constraints / 123230213121 / 5

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_key\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_cookie_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-0231212010201010-0223201300213000-2202331323101213-0310120010202330-3012303020202332-0002030121103011-0311200120002022-0130200302313333): complete subsection reference.

<a id="canonical-2121130112321031-1313331103233102-2020121322012102-0301033212030200-3301123312331201-3131123103231230-0001330132231032-3120133300102203"></a>

<a id="canonical-1330320011113201-0031323031223220-2132213331013323-0131031012111323-3322000032103013-1211030011223213-2200232302322101-1303113031323102"></a>

## max_cookie_value_size_exceeds property — request_constraints / 123230213121 / 6

Type: `"number"`. Computed.

Exclusive with \[max\_cookie\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_cookie\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32768,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32768"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32768"
  }
}
```

- [max_cookie_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-2231100030011131-1130113322112212-2210332232213102-2023102200203201-1033302101013133-3113210120333302-0000232331321022-0013221232231122): complete subsection reference.

<a id="canonical-2230132211021102-3303102212231301-2122030321302023-3021323033330122-0131231202013311-2302202323232223-2021201232131232-1132223133210203"></a>

<a id="canonical-0210002212121120-1113010212320131-0101023122323310-2310303300120123-2210201110012002-0210300312230303-0030021203013022-1012001221303031"></a>

## max_header_count_exceeds property — request_constraints / 123230213121 / 7

Type: `"number"`. Computed.

Match on the Count for all Headers that exceed this value. Exclusive with
\[max\_header\_count\_none\]

Upstream description:

Exclusive with \[max\_header\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "40"
  }
}
```

- [max_header_count_none](data-sources--service_policy--reference--group-003.md#canonical-0223013130031300-2131021200313021-3200122220002011-1103222320211221-1330021221203101-0303311221020313-2033221331010032-3030233200013100): complete subsection reference.

<a id="canonical-3302231322211230-0022112201303001-3103031030301013-2032031303130020-1331313020302123-1321222133121203-0013010232032212-1020013221012012"></a>

<a id="canonical-2311213223130003-3222221113310003-2303333013301102-3122333312002110-3222310320121000-1013123130020011-1202110312123313-0001302122101200"></a>

## max_header_key_size_exceeds property — request_constraints / 123230213121 / 8

Type: `"number"`. Computed.

Exclusive with \[max\_header\_key\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_key\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_header_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-1023233320313320-2030010120301220-0321110122003322-1022010102033132-1330230222320103-0112122133101001-1320222312022112-3023312322023121): complete subsection reference.

<a id="canonical-3331202323031333-1320202023223301-3001120321133101-3100313222110002-2233021033000331-0101331233102332-2001120112020010-1002022013112131"></a>

<a id="canonical-0331231300130323-0011320100020112-3030121021231302-0131132333310011-3302111200011023-1111133130123221-3331011011212201-2233101210132103"></a>

## max_header_value_size_exceeds property — request_constraints / 123230213121 / 9

Type: `"number"`. Computed.

Exclusive with \[max\_header\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_header\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "64000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "64000"
  }
}
```

- [max_header_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-1222121330131233-3031102030122221-1230031313133222-0031223303313321-1231023011033113-0301121023330213-1022031331323200-2133232122112020): complete subsection reference.

<a id="canonical-0101110133003132-2032222120310332-3002311033110331-1333330113322011-1300032121032031-0101322011033000-3320233021023122-1201022030031223"></a>

<a id="canonical-1010210000122033-0210220012302021-3322022110121200-0113311121022011-3331003102330010-1130333311122210-2132110123130202-1330122312223012"></a>

## max_parameter_count_exceeds property — request_constraints / 123230213121 / 10

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_count\_none\].

Upstream description:

Exclusive with \[max\_parameter\_count\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_count_none](data-sources--service_policy--reference--group-003.md#canonical-2123322111023022-1112333301122232-0103312312133313-0100232123231020-0233102100310221-1300100012131300-0233003200130022-1221121211322313): complete subsection reference.

<a id="canonical-2032232301003310-0000322220332033-1020311210122010-1001023223331112-3121120321302001-2022200311210120-3133313102013032-3123033213120001"></a>

<a id="canonical-3120313211331033-0113311101313120-1230231313122212-3013131323322213-3231103000003330-2312103030200010-2313222302202322-3130300003333232"></a>

## max_parameter_name_size_exceeds property — request_constraints / 123230213121 / 11

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_name\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_name\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1024"
  }
}
```

- [max_parameter_name_size_none](data-sources--service_policy--reference--group-003.md#canonical-3200031212011103-0120221233031230-0321321321201211-3322311002313301-2222121133100311-1212132210011122-3320000223021100-1032333013122123): complete subsection reference.

<a id="canonical-2033122330120011-2122021120213110-2100110310323022-2231120103003212-1310000210331302-2223101111230132-0021100223110221-3303113300121132"></a>

<a id="canonical-0100223232022222-1120220230223121-2210121210102210-3021312332230000-0110310131120003-2230303210203012-1000032301123303-0300321232103322"></a>

## max_parameter_value_size_exceeds property — request_constraints / 123230213121 / 12

Type: `"number"`. Computed.

Exclusive with \[max\_parameter\_value\_size\_none\].

Upstream description:

Exclusive with \[max\_parameter\_value\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1073741824,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1073741824"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "1073741824"
  }
}
```

- [max_parameter_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-1013302000001312-2130301022132020-1111003112120101-3132221102122221-1230011211011133-1321123133323131-2233202312121330-1300323300210221): complete subsection reference.

<a id="canonical-0032010330202320-3002001301210231-3323121111223121-3033200010213011-2231113121002333-0210211102212200-2022212130022111-3312221233131001"></a>

<a id="canonical-3232011011300313-3100111021120200-3133221300003122-2312030121220101-2331233321011103-1320020003130123-0102110021220322-3202013301001323"></a>

## max_query_size_exceeds property — request_constraints / 123230213121 / 13

Type: `"number"`. Computed.

Match on the URL Query Size that exceed this value. Exclusive with \[max\_query\_size\_none\]

Upstream description:

Exclusive with \[max\_query\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [max_query_size_none](data-sources--service_policy--reference--group-003.md#canonical-1322121220312330-2020203201111200-2311220030220133-1230113221000100-1312023332002112-3123221103121002-2311130111202010-3113212203311312): complete subsection reference.

<a id="canonical-1233202013023031-0202311120202333-3012313030031120-2102101110133111-3210123211101232-0231032013100200-2032313323221111-0012303302203311"></a>

<a id="canonical-2031211113001121-3130133032310100-1001212030312010-3330313230132020-3133232121333120-2233122322002333-0203323202123332-2233010230200331"></a>

## max_request_line_size_exceeds property — request_constraints / 123230213121 / 14

Type: `"number"`. Computed.

Exclusive with \[max\_request\_line\_size\_none\].

Upstream description:

Exclusive with \[max\_request\_line\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_line_size_none](data-sources--service_policy--reference--group-003.md#canonical-3303310121220321-1333112201233030-2100110022032131-2321020113020030-3330321231211201-3102301030120222-1121101330213131-3032002032220221): complete subsection reference.

<a id="canonical-0011232231110030-3323303223203001-2113221232332112-3121131310333022-0303130033201111-0021300320323302-3322322332212231-2220102220200113"></a>

<a id="canonical-3012131122111321-1021111122322333-2220130320233230-0103213003313103-2211102331321112-1022303310012131-1203032013120221-2212303310023301"></a>

## max_request_size_exceeds property — request_constraints / 123230213121 / 15

Type: `"number"`. Computed.

Match on the Request Size that exceed this value. Exclusive with \[max\_request\_size\_none\]

Upstream description:

Exclusive with \[max\_request\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65536"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65536"
  }
}
```

- [max_request_size_none](data-sources--service_policy--reference--group-003.md#canonical-0122120303322110-0301111123212331-3221302320132302-3102111331131132-0100122223121322-0112231121010203-0223222202320320-0100221021320223): complete subsection reference.

<a id="canonical-3102321322200020-0201221200020012-2332200321023110-3021313221230132-2200132210221103-3103222112233222-0212001222221231-3120102233223013"></a>

<a id="canonical-2332033102111212-0032132022323332-3003313100321033-1102030321230211-3110003222123313-3110010331322031-1232120011023231-1120133311312211"></a>

## max_url_size_exceeds property — request_constraints / 123230213121 / 16

Type: `"number"`. Computed.

Match on the URL Size that exceed this value. Exclusive with \[max\_url\_size\_none\]

Upstream description:

Exclusive with \[max\_url\_size\_none\]

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "128000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "128000"
  }
}
```

- [max_url_size_none](data-sources--service_policy--reference--group-003.md#canonical-1101220230310312-1222213030110211-1110021021323212-3323213032310020-0003122203330310-0310031302103111-0230311231310300-3102333033111103): complete subsection reference.

<a id="canonical-2021311022003012-2013322321013310-3000123231301133-3312233131232331-0230000103323032-2111300233030220-0000122000233230-0012313323011220"></a>

## Next pages — request_constraints / 123230213121 / 17

- [rule_list.rules.spec.request_constraints.max_cookie_count_none](data-sources--service_policy--reference--group-003.md#canonical-1233023322123212-3210231301202012-3232321110232302-3303331333320223-0000211210010011-3020010300232033-1020132130321000-1121233123310122)
- [rule_list.rules.spec.request_constraints.max_cookie_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-0231212010201010-0223201300213000-2202331323101213-0310120010202330-3012303020202332-0002030121103011-0311200120002022-0130200302313333)
- [rule_list.rules.spec.request_constraints.max_cookie_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-2231100030011131-1130113322112212-2210332232213102-2023102200203201-1033302101013133-3113210120333302-0000232331321022-0013221232231122)
- [rule_list.rules.spec.request_constraints.max_header_count_none](data-sources--service_policy--reference--group-003.md#canonical-0223013130031300-2131021200313021-3200122220002011-1103222320211221-1330021221203101-0303311221020313-2033221331010032-3030233200013100)
- [rule_list.rules.spec.request_constraints.max_header_key_size_none](data-sources--service_policy--reference--group-003.md#canonical-1023233320313320-2030010120301220-0321110122003322-1022010102033132-1330230222320103-0112122133101001-1320222312022112-3023312322023121)
- [rule_list.rules.spec.request_constraints.max_header_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-1222121330131233-3031102030122221-1230031313133222-0031223303313321-1231023011033113-0301121023330213-1022031331323200-2133232122112020)
- [rule_list.rules.spec.request_constraints.max_parameter_count_none](data-sources--service_policy--reference--group-003.md#canonical-2123322111023022-1112333301122232-0103312312133313-0100232123231020-0233102100310221-1300100012131300-0233003200130022-1221121211322313)
- [rule_list.rules.spec.request_constraints.max_parameter_name_size_none](data-sources--service_policy--reference--group-003.md#canonical-3200031212011103-0120221233031230-0321321321201211-3322311002313301-2222121133100311-1212132210011122-3320000223021100-1032333013122123)
- [rule_list.rules.spec.request_constraints.max_parameter_value_size_none](data-sources--service_policy--reference--group-003.md#canonical-1013302000001312-2130301022132020-1111003112120101-3132221102122221-1230011211011133-1321123133323131-2233202312121330-1300323300210221)
- [rule_list.rules.spec.request_constraints.max_query_size_none](data-sources--service_policy--reference--group-003.md#canonical-1322121220312330-2020203201111200-2311220030220133-1230113221000100-1312023332002112-3123221103121002-2311130111202010-3113212203311312)
- [rule_list.rules.spec.request_constraints.max_request_line_size_none](data-sources--service_policy--reference--group-003.md#canonical-3303310121220321-1333112201233030-2100110022032131-2321020113020030-3330321231211201-3102301030120222-1121101330213131-3032002032220221)
- [rule_list.rules.spec.request_constraints.max_request_size_none](data-sources--service_policy--reference--group-003.md#canonical-0122120303322110-0301111123212331-3221302320132302-3102111331131132-0100122223121322-0112231121010203-0223222202320320-0100221021320223)
- [rule_list.rules.spec.request_constraints.max_url_size_none](data-sources--service_policy--reference--group-003.md#canonical-1101220230310312-1222213030110211-1110021021323212-3323213032310020-0003122203330310-0310031302103111-0230311231310300-3102333033111103)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1233023322123212-3210231301202012-3232321110232302-3303331333320223-0000211210010011-3020010300232033-1020132130321000-1121233123310122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001021332011122-0213020210010213-2132200113130011-0130220133020131-0223300201131322-1321031002103212-3103103222233100-3230313201313132"></a>

## rule_list.rules.spec.request_constraints.max_cookie_count_none — max_cookie_count_none / 110123103320 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_count_none

<a id="canonical-1111313123200300-3212220133232203-1300023320321201-0122130333002003-3001130311302312-1000301222001313-1323233131120130-1203222310311223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie count none.

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

<a id="canonical-0231110021133122-3030022103202303-1201203331001011-0133320321000333-3221311132001323-1211000322020331-1131311201001201-0230211220212003"></a>

## Direct properties — max_cookie_count_none / 110123103320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132313102130331-2023123333233010-1100010002201131-0101031310123310-2311101310011100-2300210300333332-3013223023230220-3302020300213312"></a>

## Next pages — max_cookie_count_none / 110123103320 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0231212010201010-0223201300213000-2202331323101213-0310120010202330-3012303020202332-0002030121103011-0311200120002022-0130200302313333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133311221201300-3020220202222300-3330001023030210-3310112133201112-3123210011301113-1301220333332221-0320233210301121-3233301012313020"></a>

## rule_list.rules.spec.request_constraints.max_cookie_key_size_none — max_cookie_key_size_none / 012321023303 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_key_size_none

<a id="canonical-2021221011301022-3011212022323322-3001223110300212-0311230002132131-3310122203230113-3332032132012031-2332021012123322-3201221303223213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie key size none.

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

<a id="canonical-3003031300210320-2020202330133121-0123112211121332-2231213022232133-3121230033333213-2232201112013110-3213223021030131-2222133133131331"></a>

## Direct properties — max_cookie_key_size_none / 012321023303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032133121112303-3002313231203120-3202003301210321-2022320213121202-1212332230231103-3302301332210120-2113211021121333-1111222231333222"></a>

## Next pages — max_cookie_key_size_none / 012321023303 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-2231100030011131-1130113322112212-2210332232213102-2023102200203201-1033302101013133-3113210120333302-0000232331321022-0013221232231122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032333331320013-2003120212223233-1100231300111003-0132200032233120-1312110111002022-1032231031320321-3303323013330022-3001120122332303"></a>

## rule_list.rules.spec.request_constraints.max_cookie_value_size_none — max_cookie_value_size_none / 000211202312 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_cookie_value_size_none

<a id="canonical-0121330310211302-2003222211210310-0212230230301210-3330023103312300-1220132113112132-3313133311130222-0333022021331003-1113123123000333"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max cookie value size none.

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

<a id="canonical-3203301300110120-0111012230221303-1230312122331112-3213231220012003-3000022310012201-2022130231213333-1130012333001212-3222231011101123"></a>

## Direct properties — max_cookie_value_size_none / 000211202312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033112211223310-2333321220112213-2102011212012302-3011321003230030-3311013000131101-0322231021033111-2312013331300222-0201031120222110"></a>

## Next pages — max_cookie_value_size_none / 000211202312 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0223013130031300-2131021200313021-3200122220002011-1103222320211221-1330021221203101-0303311221020313-2033221331010032-3030233200013100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123202332130222-0213313133222300-1133202101232312-0112203020221030-2111132302313002-0002012203311022-0103101302322111-3200213310030203"></a>

## rule_list.rules.spec.request_constraints.max_header_count_none — max_header_count_none / 023222022322 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_count_none

<a id="canonical-1311313201013113-1212123033301013-0201322220033001-2011103102121023-3232133201022011-1303310131012321-1200120011010321-2033301222023130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header count none.

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

<a id="canonical-0020132213133011-1011322133132303-0021331312102100-3113223211323001-2232020312023120-1031203222211223-0202033300212232-0301003220131203"></a>

## Direct properties — max_header_count_none / 023222022322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002123301021031-1001010011101022-3022221010110123-3310212112220233-0303120223022110-1120311201123030-1331113231123330-2230221011330001"></a>

## Next pages — max_header_count_none / 023222022322 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1023233320313320-2030010120301220-0321110122003322-1022010102033132-1330230222320103-0112122133101001-1320222312022112-3023312322023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303101013332102-3322112111311322-1001233132120232-0301201002030112-1322131012032102-1310102021303102-1332022310131012-2132000000200033"></a>

## rule_list.rules.spec.request_constraints.max_header_key_size_none — max_header_key_size_none / 132031220022 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_key_size_none

<a id="canonical-2221002200001123-3212303123011102-1222130123011013-2010123222030023-2330230120132310-1102132303311221-0333210211333013-2220232220220200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header key size none.

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

<a id="canonical-3012201123003033-2221203310202103-2112112313001302-1003301111303222-3133303220222003-0222330100202311-2302130022233032-1223112033130131"></a>

## Direct properties — max_header_key_size_none / 132031220022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202221021020302-2131333313120202-3013120231333022-3311302123111213-0120032020200120-1012323031301201-0321311100310302-1113320102030123"></a>

## Next pages — max_header_key_size_none / 132031220022 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1222121330131233-3031102030122221-1230031313133222-0031223303313321-1231023011033113-0301121023330213-1022031331323200-2133232122112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103230321321031-2123113310133310-3012013033300122-3200232200133230-1222220021330012-2133020310310013-1012332121001130-0012013033323211"></a>

## rule_list.rules.spec.request_constraints.max_header_value_size_none — max_header_value_size_none / 302312121110 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_header_value_size_none

<a id="canonical-2033001132200123-2321030230232231-2212130213013221-2121321112321123-3012121310323302-3023133310122302-3100010120101123-0003230301103103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max header value size none.

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

<a id="canonical-0332223100321313-3330222023113213-1001312210101112-3233223233110230-0133100130110013-2333200203123110-2200321112023311-3110332202231001"></a>

## Direct properties — max_header_value_size_none / 302312121110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020331001210320-0331102223022033-3011003010023211-1201231313132223-1030022301120313-0102022111311112-0330220332132100-1021303202120131"></a>

## Next pages — max_header_value_size_none / 302312121110 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-2123322111023022-1112333301122232-0103312312133313-0100232123231020-0233102100310221-1300100012131300-0233003200130022-1221121211322313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330003032021113-1331121100200212-2230233203010110-3013013311310033-3222103132101123-0311120232311103-1102231232122232-0331112203223113"></a>

## rule_list.rules.spec.request_constraints.max_parameter_count_none — max_parameter_count_none / 003232102030 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_count_none

<a id="canonical-0313322233111221-1210021213200123-3231221212033211-3312310330100321-0201032030100102-0022133123230211-3212130210111222-2333100110021230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max parameter count none.

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

<a id="canonical-0012022020330332-1020103122332212-1003231222213012-0323003313222120-3230222210323220-1121030223120333-0133302121310013-1002112103221222"></a>

## Direct properties — max_parameter_count_none / 003232102030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310113302301330-1011033001122032-2311133012333210-3321223333331332-2223222003011023-3012033221321203-2123033322130102-3323200331033020"></a>

## Next pages — max_parameter_count_none / 003232102030 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3200031212011103-0120221233031230-0321321321201211-3322311002313301-2222121133100311-1212132210011122-3320000223021100-1032333013122123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220002001122132-3010113213132022-0033220113303122-2033323011312330-2212330330131211-2233330022330023-3322023001322231-3330213012333031"></a>

## rule_list.rules.spec.request_constraints.max_parameter_name_size_none — max_parameter_name_size_none / 311012312323 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_name_size_none

<a id="canonical-3101312003323232-0032113110210331-2321113220033213-0023130003312201-2223212030323220-1301323121000213-3320321233033012-3212333220323302"></a>

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

<a id="canonical-1101132231123312-2113010012001122-2231313012213301-1111121130001312-2030131020310213-3201003200320231-0100211012110213-3323030032033200"></a>

## Direct properties — max_parameter_name_size_none / 311012312323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303323211212230-1023221210102132-1233320112020111-0031120200312331-1033003320112022-3330000231333003-1301321011110220-0101131220122212"></a>

## Next pages — max_parameter_name_size_none / 311012312323 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1013302000001312-2130301022132020-1111003112120101-3132221102122221-1230011211011133-1321123133323131-2233202312121330-1300323300210221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202121012312322-2320022013020303-2220311220032110-0122103130203101-2313020033330222-3233200333223320-2011213032330113-0120113132133220"></a>

## rule_list.rules.spec.request_constraints.max_parameter_value_size_none — max_parameter_value_size_none / 211203220031 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_parameter_value_size_none

<a id="canonical-0002210320222020-1300110232313333-3012300310132032-1011303323333311-0100331302223311-2003333101032303-2301210011100320-1231223103102123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max parameter value size none.

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

<a id="canonical-3222002113113312-2010312300210132-1303220310233020-3132331100102020-2201201201000121-1300122303221322-3100220021021332-1131320302111203"></a>

## Direct properties — max_parameter_value_size_none / 211203220031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313122200221302-0110213003110210-2012020030222033-1233213311021112-0220002303122022-2230302323130320-0200210221223020-3322133110210101"></a>

## Next pages — max_parameter_value_size_none / 211203220031 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1322121220312330-2020203201111200-2311220030220133-1230113221000100-1312023332002112-3123221103121002-2311130111202010-3113212203311312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210200230003222-3232030000223122-1021323222111023-3233000233032332-2211212110232310-3100030230113121-0220222220023100-0120210233102321"></a>

## rule_list.rules.spec.request_constraints.max_query_size_none — max_query_size_none / 033332002013 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_query_size_none

<a id="canonical-2302321312102200-0233313322120033-2310132332032032-1001111220102121-3030020302032323-3321112312323322-2331303333100020-2003311223202201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max query size none.

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

<a id="canonical-0320302231233030-1320310101223333-3011203213031301-2220103012012322-0303012113233222-0320011212133011-3102100110313210-0322313132102203"></a>

## Direct properties — max_query_size_none / 033332002013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111112020033002-2201130213133110-3033021011001333-0332223203131202-1032330003113312-2322220233122332-0000312222013213-3301313112221123"></a>

## Next pages — max_query_size_none / 033332002013 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3303310121220321-1333112201233030-2100110022032131-2321020113020030-3330321231211201-3102301030120222-1121101330213131-3032002032220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212131310120211-0213330310122220-0301330120320101-2202030131023233-2102122112222300-3211202322210303-1113033103130032-1303333203130112"></a>

## rule_list.rules.spec.request_constraints.max_request_line_size_none — max_request_line_size_none / 210222212013 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_request_line_size_none

<a id="canonical-3131101223332303-3021032331023101-2232323000321112-0230333123331211-0111103111303200-1323112132023102-0111110012123303-2120320320121113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max request line size none.

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

<a id="canonical-0313123103122012-2320122021032031-0312100212011003-3201322020333122-2122020133112003-1210223000321311-0102312030133130-0333200310310122"></a>

## Direct properties — max_request_line_size_none / 210222212013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231303232221021-0232202332033323-1331110133020122-3322102112220111-1003003330032031-3002200100321133-0001133123003030-1211331000030230"></a>

## Next pages — max_request_line_size_none / 210222212013 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0122120303322110-0301111123212331-3221302320132302-3102111331131132-0100122223121322-0112231121010203-0223222202320320-0100221021320223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212313132321110-2301213130132220-2212101323001000-2000211013033223-3220201010330122-3122313030103021-3232202100322232-1102202321303103"></a>

## rule_list.rules.spec.request_constraints.max_request_size_none — max_request_size_none / 332213130112 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_request_size_none

<a id="canonical-2032300221313213-2323021323221310-2230003003033312-2101131033103331-0210020313003103-3320100033301110-1102301113012302-2211132101112103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for max request size none.

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

<a id="canonical-3012033320023100-1332210101321023-3231102312002220-3120311110331110-1021133033132032-1212212130302013-3030000212312000-3330020110223303"></a>

## Direct properties — max_request_size_none / 332213130112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203302200213302-1000003323202123-3321002120332031-0101001100212133-0012030023023123-1102113213133203-2232201230213200-2221111223332301"></a>

## Next pages — max_request_size_none / 332213130112 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1101220230310312-1222213030110211-1110021021323212-3323213032310020-0003122203330310-0310031302103111-0230311231310300-3102333033111103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010030011221221-2202013131300233-2210122001203012-3100110310013103-2033320203013112-0121310032302122-2302221232031321-0002321111220103"></a>

## rule_list.rules.spec.request_constraints.max_url_size_none — max_url_size_none / 302132320101 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- rule_list.rules.spec.request_constraints.max_url_size_none

<a id="canonical-1313023023203300-0032003021013100-1302123103200301-2121223010312110-2013102002212233-1002000013001322-0121231202230231-0213321321331003"></a>

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

<a id="canonical-1322200110000333-3232300002113210-2203223312100203-2302000310313023-3111020331001231-3110130010221311-2002213331331003-1331110130330320"></a>

## Direct properties — max_url_size_none / 302132320101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021201110110131-0322033201321021-0303120212112110-0330003201220213-0323223332110233-1223200211233303-0201233111102032-3303201201021232"></a>

## Next pages — max_url_size_none / 302132320101 / 4

- [rule_list.rules.spec.request_constraints](data-sources--service_policy--reference--group-002.md#canonical-2110021331103130-3023312121033312-2231302230120211-1103201133030222-0221022130001011-3002013103120132-1321211031031000-0320022212322013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212122133300330-2012112022120131-3312010002330001-3323131003002210-3032033112011110-0131312300031022-2312221131031312-3023211301030302"></a>

## rule_list.rules.spec.segment_policy — segment_policy / 210001101202 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.segment_policy

<a id="canonical-3323232232032330-1122211331303130-1023302210003331-0013122122013120-1002031300033001-2331202230203020-1031113221232130-1002212101031213"></a>

Type: `"single"`. Computed.

Configure source and destination segment for policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dst_segment_choice": "[\"dst_any\",\"dst_segments\",\"intra_segment\"]",
  "x-ves-oneof-field-src_segment_choice": "[\"src_any\",\"src_segments\"]"
}
```

<a id="canonical-2323001201103101-3200323223101122-2020322200303030-3120310313102223-0331000001201321-0310220132023320-3223233031300313-3330131302203223"></a>

## Direct properties — segment_policy / 210001101202 / 3

- [dst_any](data-sources--service_policy--reference--group-003.md#canonical-0031122310033231-2010230020011010-1303321201322031-3320020223231331-2303031122211232-0331300200012303-1021230330130131-2303123000000020): complete subsection reference.

- [dst_segments](data-sources--service_policy--reference--group-003.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123): complete subsection reference.

- [intra_segment](data-sources--service_policy--reference--group-003.md#canonical-0203323232233001-2120013110020322-0221121112113023-2320113023302210-2100223233232320-2310322302003122-3100131032123133-0222202310211010): complete subsection reference.

- [src_any](data-sources--service_policy--reference--group-003.md#canonical-1300223022331330-1313302231013010-3330313313133130-2202101212123102-3202302012303311-2032313311113131-3222330200001323-3020110200313330): complete subsection reference.

- [src_segments](data-sources--service_policy--reference--group-003.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233): complete subsection reference.

<a id="canonical-0323212312011302-2202121110221322-0133321032230003-3231000001212302-2312100212020221-3200100312101133-0233000311212312-2030113101023320"></a>

## Next pages — segment_policy / 210001101202 / 4

- [rule_list.rules.spec.segment_policy.dst_any](data-sources--service_policy--reference--group-003.md#canonical-0031122310033231-2010230020011010-1303321201322031-3320020223231331-2303031122211232-0331300200012303-1021230330130131-2303123000000020)
- [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-003.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123)
- [rule_list.rules.spec.segment_policy.intra_segment](data-sources--service_policy--reference--group-003.md#canonical-0203323232233001-2120013110020322-0221121112113023-2320113023302210-2100223233232320-2310322302003122-3100131032123133-0222202310211010)
- [rule_list.rules.spec.segment_policy.src_any](data-sources--service_policy--reference--group-003.md#canonical-1300223022331330-1313302231013010-3330313313133130-2202101212123102-3202302012303311-2032313311113131-3222330200001323-3020110200313330)
- [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-003.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0031122310033231-2010230020011010-1303321201322031-3320020223231331-2303031122211232-0331300200012303-1021230330130131-2303123000000020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010033220311112-2311213022103022-2100333013131300-3210323330331222-3110332001331231-2223133010013322-2102331121010300-1010233212330332"></a>

## rule_list.rules.spec.segment_policy.dst_any — dst_any / 012301330000 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.dst_any

<a id="canonical-1321312333231211-2012322212003330-1002310223333030-1202100132013202-2030132101123031-1130131223203111-1023332120210333-2122311100310133"></a>

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

<a id="canonical-1011223002021321-1101010312301301-2211300222020222-1201212232220132-0102320213013133-0210313103112102-0122221321001103-1122011230303200"></a>

## Direct properties — dst_any / 012301330000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023301021120211-2000131331300000-0031123302231013-2301231303131110-3100131000003001-2330123321321022-3310030302333121-3322012123122321"></a>

## Next pages — dst_any / 012301330000 / 4

- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032120203313332-0030231301012021-2113002311220303-3313133130301231-3223033122133303-2310132210012213-0321001332320001-1210300220102320"></a>

## rule_list.rules.spec.segment_policy.dst_segments — dst_segments / 100233101021 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.dst_segments

<a id="canonical-0031213203022332-1022112230002233-2203321002120022-1133112331111202-3300311331010213-0022102322032123-1023020002302211-3001021021023002"></a>

Type: `"single"`. Computed.

Configuration parameter for dst segments.

Upstream description:

List of references to Segments.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0321312131112233-0230120330223012-2302001323122001-1221330301133221-3131131030003032-1032113110130210-2110121203121201-0031301010033322"></a>

## Direct properties — dst_segments / 100233101021 / 3

- [segments](data-sources--service_policy--reference--group-003.md#canonical-3201210223131120-2023030303012223-3333023031202332-0122002233013211-3110031303331122-0330302011320032-1333221323213100-3101330210002213): complete subsection reference.

<a id="canonical-2210231223001013-3130221322331200-2231201010112011-3103000330000123-0033130203100210-3133113320302023-1313030302300201-0212230301103130"></a>

## Next pages — dst_segments / 100233101021 / 4

- [rule_list.rules.spec.segment_policy.dst_segments.segments](data-sources--service_policy--reference--group-003.md#canonical-3201210223131120-2023030303012223-3333023031202332-0122002233013211-3110031303331122-0330302011320032-1333221323213100-3101330210002213)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3201210223131120-2023030303012223-3333023031202332-0122002233013211-3110031303331122-0330302011320032-1333221323213100-3101330210002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133201211002022-1220032231310303-0333001321030011-3232300000333330-1012133310131311-2203210220003333-0013203101210322-0033022231122312"></a>

## rule_list.rules.spec.segment_policy.dst_segments.segments — segments / 132021001012 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-003.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123)
- rule_list.rules.spec.segment_policy.dst_segments.segments

<a id="canonical-1002310132133310-0201232300331013-0301123313122333-1131113103202010-3220122322323201-2201223110003021-2303112112230112-1303201201010000"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0313312202320121-3121120233312213-1101223010010122-1010101010033210-2313323032011031-0302203201220103-1321330213311032-0223202110300212"></a>

## Direct properties — segments / 132021001012 / 3

<a id="canonical-1221110230301310-0131201212321122-0113003003312103-0232010031031121-1110302312211233-2030202321211333-0203333032201300-0223232330003130"></a>

<a id="canonical-1210310032010030-2211121313103110-1023103020221220-3331330320110220-1123102212131231-2210021231202320-3001020212021213-1032323103111030"></a>

## name property — segments / 132021001012 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2120202321223210-2122112212030221-2221021330201121-1013023030301202-0300121210330333-0100311130313211-3013202223322221-1012320300222201"></a>

<a id="canonical-2021232113103111-1111202001302222-3213310030032230-1031212032011330-3022201130002220-3120303303311323-1332121032310300-1022303033123221"></a>

## namespace property — segments / 132021001012 / 5

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3011301032103032-1233300121303113-0031122021013133-2220333002203010-2233131313023311-1001333033100330-3032030220010002-3233032100103301"></a>

<a id="canonical-3132303211010111-0011122333112031-3303002223200301-2031312100220023-1233300002212113-1221230133133322-1101320313131310-3201113121100333"></a>

## tenant property — segments / 132021001012 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3003122133113230-1122023320311230-1110300020100002-0031212100133300-1321332001331013-0312230223222321-0013010101133000-3333111121123103"></a>

## Next pages — segments / 132021001012 / 7

- [rule_list.rules.spec.segment_policy.dst_segments](data-sources--service_policy--reference--group-003.md#canonical-0011131210322333-3303310012013300-3023111203210330-0020021112012101-0303233333320023-2333301023133001-0000020301032201-2310201111202123)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0203323232233001-2120013110020322-0221121112113023-2320113023302210-2100223233232320-2310322302003122-3100131032123133-0222202310211010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321133131100300-3310200120210102-1031133231033102-3203313103033103-3222021133131300-3000311300031123-3311013031311110-2130122211223302"></a>

## rule_list.rules.spec.segment_policy.intra_segment — intra_segment / 023110101002 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.intra_segment

<a id="canonical-1230110130222112-0033112022211020-0200023213002330-1223122322211222-2123221023313332-3013022001202203-0001312021201121-0033213231022131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for intra segment.

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

<a id="canonical-1112312130230211-0131202002302120-2131303300123322-0331130032032323-0330203323131311-1021030033222103-2023331230013231-1313301110011132"></a>

## Direct properties — intra_segment / 023110101002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232321113230303-2111110103313320-3223320031223113-1130320012112113-2310010211331102-3212300332022223-1122103103132022-0211321002302121"></a>

## Next pages — intra_segment / 023110101002 / 4

- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1300223022331330-1313302231013010-3330313313133130-2202101212123102-3202302012303311-2032313311113131-3222330200001323-3020110200313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310233132100110-3112013332100332-0332130320030302-2330221012220122-0221112203132113-0321032102210103-1021030001303232-3330213030213231"></a>

## rule_list.rules.spec.segment_policy.src_any — src_any / 001312010322 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.src_any

<a id="canonical-1101103031230102-0230020121302021-2001333033001031-3023002312220232-1302000131223222-3313132130030110-0023031003332310-3132002333101203"></a>

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

<a id="canonical-3310223011110310-0033323113330020-0033202011030021-0300231233310212-3201212213210102-3132212130332032-3113200220110102-3310112023103010"></a>

## Direct properties — src_any / 001312010322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322202300103031-1130021303020311-0002202212113231-1310012102210231-0232302321013021-0300220111102333-0322122113200220-2132011111003301"></a>

## Next pages — src_any / 001312010322 / 4

- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012001022002231-3231221021230311-2200012011213300-3310002110023321-0112222001022303-3030002232302212-2311030323000321-2223202133232102"></a>

## rule_list.rules.spec.segment_policy.src_segments — src_segments / 202020300013 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- rule_list.rules.spec.segment_policy.src_segments

<a id="canonical-2301133103201303-2001023333321101-0132010032100323-1030203123002021-1100021200021233-1320101333332010-2230331010212112-2210232020122310"></a>

Type: `"single"`. Computed.

Configuration parameter for src segments.

Upstream description:

List of references to Segments.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2312121202002132-0333111110320111-0030112320100113-0222112101023203-0102210322203133-0202022131331021-0223113231332113-1210010010330033"></a>

## Direct properties — src_segments / 202020300013 / 3

- [segments](data-sources--service_policy--reference--group-003.md#canonical-1313121023332203-0102232233233030-0032313202313131-1121121011302313-1103332101020131-1030113133012303-3010003230212213-3030323032021001): complete subsection reference.

<a id="canonical-0102302311100100-3101333133130012-2301121220232031-3223130102212233-1030320133203322-1133333102302220-2313322320320121-2112203323101213"></a>

## Next pages — src_segments / 202020300013 / 4

- [rule_list.rules.spec.segment_policy.src_segments.segments](data-sources--service_policy--reference--group-003.md#canonical-1313121023332203-0102232233233030-0032313202313131-1121121011302313-1103332101020131-1030113133012303-3010003230212213-3030323032021001)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1313121023332203-0102232233233030-0032313202313131-1121121011302313-1103332101020131-1030113133012303-3010003230212213-3030323032021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122013102010233-2311332233011031-0100001130311022-0103232202333313-2022100013120320-0123301033210033-3203033101102130-2013201202133323"></a>

## rule_list.rules.spec.segment_policy.src_segments.segments — segments / 120230100210 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.segment_policy](data-sources--service_policy--reference--group-003.md#canonical-0121322232330011-2001200302133213-0020123132212012-0113020311112312-0201321011303010-1302103121002333-1323212012233310-1210231031031002)
- [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-003.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233)
- rule_list.rules.spec.segment_policy.src_segments.segments

<a id="canonical-3010233120011013-0220003003232100-1220130101303213-0232203100213013-1122133000103112-1033220213020101-0210122121030133-1030030213021023"></a>

Type: `"list"`. Computed.

Segments. Select list of segments.

Upstream description:

Select list of segments.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1130001223202333-1200023213022103-0123313301033332-2103210203031021-0221202113013201-3100203110020132-2111320112122022-3331231300131310"></a>

## Direct properties — segments / 120230100210 / 3

<a id="canonical-2320133321111300-3211202300330000-3123020210222102-0112120102130231-2330211230222122-1100212120320010-2221233321323330-1101320131012010"></a>

<a id="canonical-2033220101012103-1101223022030002-1232113032011222-1201230300301122-2233013200223011-0133322013031120-1002312311211311-1010131233000333"></a>

## name property — segments / 120230100210 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1230301122130022-0203020120330023-2330000013103220-1310100333320033-2213100010322130-2002032220100312-0211030220313133-2322310232202113"></a>

<a id="canonical-3021000222020331-3203121231230301-0022300011121313-1133220031303231-0232330302223022-1031302031130302-0022133132120213-1030311230013210"></a>

## namespace property — segments / 120230100210 / 5

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1321133222213232-3313200100232002-3231022123330231-0130302331000310-2133101112123201-1030323223131133-1331313020103033-2312303012113023"></a>

<a id="canonical-3102320101310122-2313222123330121-2032022200122212-1100011312003210-2031012010201112-2311121121210323-0311333020201110-2130101230211132"></a>

## tenant property — segments / 120230100210 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0000300223333131-2001210211333121-0230301330101113-2020331332101102-3031311213310002-0302113001201330-1111210111201221-0313133033102031"></a>

## Next pages — segments / 120230100210 / 7

- [rule_list.rules.spec.segment_policy.src_segments](data-sources--service_policy--reference--group-003.md#canonical-0321011030032012-0311231002031103-3002101310212323-0023110012220020-1300300231320033-2222023130213033-3113130312302023-2330031210112233)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1333003020033103-2303031332133313-1101333013023130-2313101003001221-0111232133300132-2011202312030000-1012231232313131-2333220102021302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223030322130232-2301333120002212-0313121223323103-2211303110300200-1312023032102010-2201223211013232-2003312203321012-2122022111022212"></a>

## rule_list.rules.spec.tls_fingerprint_matcher — tls_fingerprint_matcher / 313010123002 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.tls_fingerprint_matcher

<a id="canonical-3311100032302102-2323202000133230-2023032001333021-1201210301202211-1223001232030332-3002102203101103-0312130010013122-1133130200001323"></a>

Type: `"single"`. Computed.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3021332033023300-1100030030323021-3332330312003302-3300121203132222-2022001123120030-3313222011001212-2121330020012113-3310011313202110"></a>

## Direct properties — tls_fingerprint_matcher / 313010123002 / 3

<a id="canonical-0333121223330310-3322313032201230-0013120222103300-3003203120231123-1332230003322001-0112022203101232-0013231101021231-2321333333332220"></a>

<a id="canonical-3032021033100131-0102302103220132-2001033222300213-2223103212310100-2112223300023313-2331133001233312-1011031233110312-0303210133230320"></a>

## classes property — tls_fingerprint_matcher / 313010123002 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2011231012212202-2021321320303201-0312332010222313-2222120231002131-2330133220203113-0021122300313021-3012121121312031-0223303013010030"></a>

<a id="canonical-1001331330013130-2310133030313213-1121313303123310-0221102003333322-3000232102113212-1233122330211002-3310101031302200-2331021312313133"></a>

## exact_values property — tls_fingerprint_matcher / 313010123002 / 5

Type: `["list", "string"]`. Computed.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200131312313331-0021301110001231-2020033301023203-2020100202230221-0300203113201130-1333021011032201-1312231123332333-0000113031133221"></a>

<a id="canonical-0200112230321223-1033120302113023-1201332112033201-2100301120100100-3032003101232022-0022000030201223-3023002013133023-3012113031211121"></a>

## excluded_values property — tls_fingerprint_matcher / 313010123002 / 6

Type: `["list", "string"]`. Computed.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0010032030222333-0321003311123013-3333000033012212-1111102020110300-2122033201302212-0031101022030030-3032323003113122-1120320030200322"></a>

## Next pages — tls_fingerprint_matcher / 313010123002 / 7

- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3302232320022303-1301323100000232-2222223202031122-2132301121130100-3212113133030300-0223231120230100-3101210030303312-3003333002303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033201132100321-0013130130220213-0030310103320100-2111210110220103-1002100121300022-3332313130300113-1111000001223032-0331221030322022"></a>

## rule_list.rules.spec.user_identity_matcher — user_identity_matcher / 132013222121 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.user_identity_matcher

<a id="canonical-0100212130233102-1310321312010333-2030333021002020-0223013210212033-3201132300013202-3220332031312301-2333203202012002-0211003212223012"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2332121312220322-0000012112030110-2000310123230101-3231310101022030-0302213012331001-0223233120321010-0313232020310120-3012321320012030"></a>

## Direct properties — user_identity_matcher / 132013222121 / 3

<a id="canonical-2233132333222131-0231231101030331-1223112100320122-1323202021302012-0122002132001213-2302112323131103-0232310133211212-3132022030202221"></a>

<a id="canonical-3201022021333102-3333023131330033-2123133111321331-3013312023321231-1133320130111121-3321133222120012-2130210102032011-1331033211323203"></a>

## exact_values property — user_identity_matcher / 132013222121 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-1131312023312021-1230101210102111-2311331301310200-3201232002222003-2202210120300210-2231031302101120-3221231123211010-3233111032100223"></a>

<a id="canonical-0221113223220002-0330111332313013-0302030331312122-0231221301230033-2103302012003012-3300012130230322-3123310213302033-0103300013311030"></a>

## regex_values property — user_identity_matcher / 132013222121 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-2331333321031110-0230310200300001-0001210000013321-0210023321313331-2003023030213231-0001010121221120-2333030030000312-2110222310210000"></a>

## Next pages — user_identity_matcher / 132013222121 / 6

- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011021110231330-0233011111310331-2330303310202331-2121232012310221-1130120220320201-2121302221100023-1223333013223331-3223212020233013"></a>

## rule_list.rules.spec.waf_action — waf_action / 310302120203 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- rule_list.rules.spec.waf_action

<a id="canonical-2201213013321111-0021023302222322-3012030010120000-0313121021213130-0213311111011011-3132001123022010-0013200303120001-0010312221111010"></a>

Type: `"single"`. Computed.

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Upstream description:

Modify App Firewall behavior for a matching request. The modification could either be to entirely
skip firewall processing or to customize the firewall rules to be applied as defined by App Firewall
Rule Control settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_type": "[\"app_firewall_detection_control\",\"none\",\"waf_skip_processing\"]"
}
```

<a id="canonical-0020312313221121-3122022131132110-0012111131223130-3221212030003200-2231312010112310-1231120013112020-2132301230223220-0200111212031010"></a>

## Direct properties — waf_action / 310302120203 / 3

- [app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020): complete subsection reference.

- [none](data-sources--service_policy--reference--group-003.md#canonical-0310123023311333-0022102211121010-0020210033011011-2001230110222211-2210213330032131-3212220310101201-1223020211031333-3230113300311321): complete subsection reference.

- [waf_skip_processing](data-sources--service_policy--reference--group-003.md#canonical-0000112103021210-0222132122221021-1220123212112130-2313130013212221-2010010233130222-3122031033113002-0031013220303221-1020013000133310): complete subsection reference.

<a id="canonical-0212313102131230-1222331231323233-3312231103233201-1110222100333101-3312122230023321-1030331100033232-2130312000322312-2303300213021303"></a>

## Next pages — waf_action / 310302120203 / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- [rule_list.rules.spec.waf_action.none](data-sources--service_policy--reference--group-003.md#canonical-0310123023311333-0022102211121010-0020210033011011-2001230110222211-2210213330032131-3212220310101201-1223020211031333-3230113300311321)
- [rule_list.rules.spec.waf_action.waf_skip_processing](data-sources--service_policy--reference--group-003.md#canonical-0000112103021210-0222132122221021-1220123212112130-2313130013212221-2010010233130222-3122031033113002-0031013220303221-1020013000133310)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222213010011303-2110200332322113-2000020131200211-3122203321311201-0130233323000303-3101233122100120-1032212112321120-1323032011003021"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control — app_firewall_detection_control / 212202132300 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- rule_list.rules.spec.waf_action.app_firewall_detection_control

<a id="canonical-0020213223313321-2120133310320131-2011210000132312-3113301020021323-0302030302133130-0203011310312310-0112222313103012-1001301021303002"></a>

Type: `"single"`. Computed.

Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded
from triggering on the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3103211020130200-2022000020233312-1103203131200102-2012222112210010-3102300113032323-3313200023330321-1230230110031201-2332223111031313"></a>

## Direct properties — app_firewall_detection_control / 212202132300 / 3

- [exclude_attack_type_contexts](data-sources--service_policy--reference--group-003.md#canonical-3323003110031030-2201330032221222-2000010213012021-2131130212001113-2012322012011313-1201333120303103-0000122123120001-1011102131322323): complete subsection reference.

- [exclude_bot_name_contexts](data-sources--service_policy--reference--group-003.md#canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030): complete subsection reference.

- [exclude_signature_contexts](data-sources--service_policy--reference--group-003.md#canonical-2031101212110222-0203330312200223-3230231021002311-3022111302033032-3230123020220100-3001321120131013-0033323300220123-3122012332222033): complete subsection reference.

- [exclude_violation_contexts](data-sources--service_policy--reference--group-003.md#canonical-0100312021032302-3122211211201132-3012220103230032-1212223321010131-0230222232110210-0131112003213120-0132312221300330-3223101002023033): complete subsection reference.

<a id="canonical-3210023230102001-3321031331111221-3110201032001001-3002332031113331-0031021333333230-1301122013230333-1230213301002011-2110233211232300"></a>

## Next pages — app_firewall_detection_control / 212202132300 / 4

- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts](data-sources--service_policy--reference--group-003.md#canonical-3323003110031030-2201330032221222-2000010213012021-2131130212001113-2012322012011313-1201333120303103-0000122123120001-1011102131322323)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts](data-sources--service_policy--reference--group-003.md#canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts](data-sources--service_policy--reference--group-003.md#canonical-2031101212110222-0203330312200223-3230231021002311-3022111302033032-3230123020220100-3001321120131013-0033323300220123-3122012332222033)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts](data-sources--service_policy--reference--group-003.md#canonical-0100312021032302-3122211211201132-3012220103230032-1212223321010131-0230222232110210-0131112003213120-0132312221300330-3223101002023033)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3323003110031030-2201330032221222-2000010213012021-2131130212001113-2012322012011313-1201333120303103-0000122123120001-1011102131322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221221322302021-1111333311102223-3113100332012122-0110210133301303-2201213010310311-1122013132221021-1033310001002021-1032133323101230"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts — exclude_attack_type_contexts / 213320011232 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_attack_type_contexts

<a id="canonical-0313223222013300-2223010213233203-3000112212111221-1300011232121130-1012212112013302-3001330322220022-1201310312333133-3330230221001133"></a>

Type: `"list"`. Computed.

Exclude an entire attack type only in the named context. For migrated per-parameter exceptions,
prefer this over signature-ID exclusions because one payload can trigger several signatures;
unrelated parameters and attack types remain protected.

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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1120232023302023-2133031311132100-0031100100202320-2320013201002321-1001301121021131-3330033020121033-0123332012132030-2231102331330121"></a>

## Direct properties — exclude_attack_type_contexts / 213320011232 / 3

<a id="canonical-3121103133132030-0130203123101231-3223300333120333-0011100233023020-2101231323232111-1112123213130202-3020010302003300-0321200300301321"></a>

<a id="canonical-3322301133220330-1203231312023120-3121012012002303-1210302103331123-2323131112323033-0101332000321031-2202232003111010-3032303202331013"></a>

## context property — exclude_attack_type_contexts / 213320011232 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

Exclusion scope. Use CONTEXT\_PARAMETER with context\_name for one parameter, CONTEXT\_COOKIE for
one cookie, or CONTEXT\_ANY only for an intentionally global scope.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3110231203333213-2100132103102023-2002232113203223-3011311300330011-2220312103212200-1031213011003321-2002323130010303-1322133302203303"></a>

<a id="canonical-0210231112002031-0200110312111033-2021113321032000-3303121233320201-3000330031301221-1013102122312102-0303131022120211-2010212121231320"></a>

## context_name property — exclude_attack_type_contexts / 213320011232 / 5

Type: `"string"`. Computed.

Parameter, cookie, or header name selected by context. For a parameter-scoped WAF exception, set
context to CONTEXT\_PARAMETER and name only the intended parameter.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2012101323113130-2322301023023333-1001022231231003-1131310120203112-2123010102302122-1332001302033030-2332303331031232-0120332211102301"></a>

<a id="canonical-0002131330120003-3201322311302303-3012321203010321-1321300212302203-2001002300020022-0222111111323303-2230113120102231-0303231021333012"></a>

## exclude_attack_type property — exclude_attack_type_contexts / 213320011232 / 6

Type: `"string"`. Computed.

\[Enum:
ATTACK\_TYPE\_NONE|ATTACK\_TYPE\_NON\_BROWSER\_CLIENT|ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS|ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE|ATTACK\_TYPE\_DETECTION\_EVASION|ATTACK\_TYPE\_VULNERABILITY\_SCAN|ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY|ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS|ATTACK\_TYPE\_BUFFER\_OVERFLOW|ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION|ATTACK\_TYPE\_INFORMATION\_LEAKAGE|ATTACK\_TYPE\_DIRECTORY\_INDEXING|ATTACK\_TYPE\_PATH\_TRAVERSAL|ATTACK\_TYPE\_XPATH\_INJECTION|ATTACK\_TYPE\_LDAP\_INJECTION|ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION|ATTACK\_TYPE\_COMMAND\_EXECUTION|ATTACK\_TYPE\_SQL\_INJECTION|ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING|ATTACK\_TYPE\_DENIAL\_OF\_SERVICE|ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK|ATTACK\_TYPE\_SESSION\_HIJACKING|ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING|ATTACK\_TYPE\_FORCEFUL\_BROWSING|ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE|ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD|ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\]
List of all Attack Types ATTACK\_TYPE\_NONE ATTACK\_TYPE\_NON\_BROWSER\_CLIENT
ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE
ATTACK\_TYPE\_DETECTION\_EVASION ATTACK\_TYPE\_VULNERABILITY\_SCAN
ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS..
Possible values are \`ATTACK\_TYPE\_NONE\`, \`ATTACK\_TYPE\_NON\_BROWSER\_CLIENT\`,
\`ATTACK\_TYPE\_OTHER\_APPLICATION\_ATTACKS\`, \`ATTACK\_TYPE\_TROJAN\_BACKDOOR\_SPYWARE\`,
\`ATTACK\_TYPE\_DETECTION\_EVASION\`, \`ATTACK\_TYPE\_VULNERABILITY\_SCAN\`,
\`ATTACK\_TYPE\_ABUSE\_OF\_FUNCTIONALITY\`,
\`ATTACK\_TYPE\_AUTHENTICATION\_AUTHORIZATION\_ATTACKS\`, \`ATTACK\_TYPE\_BUFFER\_OVERFLOW\`,
\`ATTACK\_TYPE\_PREDICTABLE\_RESOURCE\_LOCATION\`, \`ATTACK\_TYPE\_INFORMATION\_LEAKAGE\`,
\`ATTACK\_TYPE\_DIRECTORY\_INDEXING\`, \`ATTACK\_TYPE\_PATH\_TRAVERSAL\`,
\`ATTACK\_TYPE\_XPATH\_INJECTION\`, \`ATTACK\_TYPE\_LDAP\_INJECTION\`,
\`ATTACK\_TYPE\_SERVER\_SIDE\_CODE\_INJECTION\`, \`ATTACK\_TYPE\_COMMAND\_EXECUTION\`,
\`ATTACK\_TYPE\_SQL\_INJECTION\`, \`ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING\`,
\`ATTACK\_TYPE\_DENIAL\_OF\_SERVICE\`, \`ATTACK\_TYPE\_HTTP\_PARSER\_ATTACK\`,
\`ATTACK\_TYPE\_SESSION\_HIJACKING\`, \`ATTACK\_TYPE\_HTTP\_RESPONSE\_SPLITTING\`,
\`ATTACK\_TYPE\_FORCEFUL\_BROWSING\`, \`ATTACK\_TYPE\_REMOTE\_FILE\_INCLUDE\`,
\`ATTACK\_TYPE\_MALICIOUS\_FILE\_UPLOAD\`, \`ATTACK\_TYPE\_GRAPHQL\_PARSER\_ATTACK\`. Defaults to
\`ATTACK\_TYPE\_NONE\`.

Upstream description:

Attack-type enum excluded in this context, for example ATTACK\_TYPE\_CROSS\_SITE\_SCRIPTING. Other
attack types remain enforced.

Receipt-pinned upstream constraints:

```json
{
  "default": "ATTACK_TYPE_NONE",
  "enum": [
    "ATTACK_TYPE_NONE",
    "ATTACK_TYPE_NON_BROWSER_CLIENT",
    "ATTACK_TYPE_OTHER_APPLICATION_ATTACKS",
    "ATTACK_TYPE_TROJAN_BACKDOOR_SPYWARE",
    "ATTACK_TYPE_DETECTION_EVASION",
    "ATTACK_TYPE_VULNERABILITY_SCAN",
    "ATTACK_TYPE_ABUSE_OF_FUNCTIONALITY",
    "ATTACK_TYPE_AUTHENTICATION_AUTHORIZATION_ATTACKS",
    "ATTACK_TYPE_BUFFER_OVERFLOW",
    "ATTACK_TYPE_PREDICTABLE_RESOURCE_LOCATION",
    "ATTACK_TYPE_INFORMATION_LEAKAGE",
    "ATTACK_TYPE_DIRECTORY_INDEXING",
    "ATTACK_TYPE_PATH_TRAVERSAL",
    "ATTACK_TYPE_XPATH_INJECTION",
    "ATTACK_TYPE_LDAP_INJECTION",
    "ATTACK_TYPE_SERVER_SIDE_CODE_INJECTION",
    "ATTACK_TYPE_COMMAND_EXECUTION",
    "ATTACK_TYPE_SQL_INJECTION",
    "ATTACK_TYPE_CROSS_SITE_SCRIPTING",
    "ATTACK_TYPE_DENIAL_OF_SERVICE",
    "ATTACK_TYPE_HTTP_PARSER_ATTACK",
    "ATTACK_TYPE_SESSION_HIJACKING",
    "ATTACK_TYPE_HTTP_RESPONSE_SPLITTING",
    "ATTACK_TYPE_FORCEFUL_BROWSING",
    "ATTACK_TYPE_REMOTE_FILE_INCLUDE",
    "ATTACK_TYPE_MALICIOUS_FILE_UPLOAD",
    "ATTACK_TYPE_GRAPHQL_PARSER_ATTACK"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212331133133203-1033023120202322-3213131211221221-2001223333011221-3110302033010131-0320002131312303-0021022010102120-2212003020010002"></a>

## Next pages — exclude_attack_type_contexts / 213320011232 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3223230210221003-2032232113223131-0333230231131301-2132210130000201-2120221212200230-0200130213332300-2002232320022030-0013000222011030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302002303301211-2200232120012101-0233022022210213-2023002330230231-1212101030120030-2301002023202110-3112311023301011-3132300001333220"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts — exclude_bot_name_contexts / 121111301123 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_bot_name_contexts

<a id="canonical-0100022320130212-0121002013023332-2121022130012331-1013120021320110-1103233103001031-2020212020100033-3013123002120033-3213112223212333"></a>

Type: `"list"`. Computed.

Bot Names to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310230231113121-3020201132113130-1102112103310102-3110322333302322-3110122233322321-0323200122221321-0112122100123201-2030031120332233"></a>

## Direct properties — exclude_bot_name_contexts / 121111301123 / 3

<a id="canonical-1223031212301030-3302120030211111-2311012131320230-3220013222002312-1013232001233222-1110102202001032-1012112103321022-3322203213131131"></a>

<a id="canonical-0131231120323202-1110223113311220-3300111231303223-2021111332212120-2323133023010013-0102010021332012-3102221312031210-1210210313330010"></a>

## bot_name property — exclude_bot_name_contexts / 121111301123 / 4

Type: `"string"`. Computed.

Bot Name. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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

<a id="canonical-0333011030110021-3120333111333231-3120022011122303-2310321000033102-3102031331310111-1221200212330222-0332100021122030-2313103101121112"></a>

## Next pages — exclude_bot_name_contexts / 121111301123 / 5

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-2031101212110222-0203330312200223-3230231021002311-3022111302033032-3230123020220100-3001321120131013-0033323300220123-3122012332222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003103011233303-0220032110331003-0300132320032003-1101023111011113-3131303122320120-2111033020311213-0111032033203130-2331321312221202"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts — exclude_signature_contexts / 210322032130 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_signature_contexts

<a id="canonical-3232303210031312-3033010222203121-0320100002303131-2030300312300110-2123132031302323-3331330031132312-1222222311001223-1320310302201320"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
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
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1221322132120203-2012310132213031-3010120232202200-0032120012010212-2332103210333302-0111120322323213-1212133213010110-0220013113131100"></a>

## Direct properties — exclude_signature_contexts / 210322032130 / 3

<a id="canonical-1332010312113312-1222033200113132-3130310022333322-2022301122032021-3102112122101313-3310200131020332-1012201313030131-1202232330323011"></a>

<a id="canonical-2001322121131230-1111333211022233-1213213102221200-1021211103333103-1000200100000110-0000221203101132-2322132201322230-1312032013301223"></a>

## context property — exclude_signature_contexts / 210322032130 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0020111011133001-3233222011223222-1131010313311201-3230000233201331-1123103231210002-1232303332322002-2301321003132221-3032102300223220"></a>

<a id="canonical-2213011002332320-1111030001330321-0132301120121211-3110031033313003-0313030032123122-3122213032101001-1230313012122200-3113300211102233"></a>

## context_name property — exclude_signature_contexts / 210322032130 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3133010203031111-3030010100300220-1213331210330100-3232011213113231-1303230323033130-1022103132032131-0113032213230231-1212102000021003"></a>

<a id="canonical-0002321310033113-0213111321311202-2112221332101202-2220103213020201-0202102301031202-0103001323033032-1012011232130003-3121000200102213"></a>

## signature_id property — exclude_signature_contexts / 210322032130 / 6

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
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
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

<a id="canonical-1322300200002212-2222330232212021-3012321322230203-2321301330002201-1010200023301230-3020302233313122-3322131111111223-3202213221231132"></a>

## Next pages — exclude_signature_contexts / 210322032130 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0100312021032302-3122211211201132-3012220103230032-1212223321010131-0230222232110210-0131112003213120-0132312221300330-3223101002023033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013320301122022-0202310321022203-2221201232120210-1112011211203321-2220002101112002-1013300111032221-3222022110000031-1102011133233310"></a>

## rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts — exclude_violation_contexts / 332230121203 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- rule_list.rules.spec.waf_action.app_firewall_detection_control.exclude_violation_contexts

<a id="canonical-1221312223200101-1021231322200023-0320023100020323-3313123222113123-3333120013100121-1323222002111003-1103303311030123-1302011122232211"></a>

Type: `"list"`. Computed.

Violations to be excluded for the defined match criteria.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1320110002322220-1201301101020222-0102201300203312-2013200133222223-3002333321033113-1300220030033200-0013103002321232-1203122201022200"></a>

## Direct properties — exclude_violation_contexts / 332230121203 / 3

<a id="canonical-1220123322101023-0120030220103020-1221310033032110-0100132033311030-3133332212032200-2003113302323023-2131110233001010-2002332333311212"></a>

<a id="canonical-2030113131003020-0323211000321332-1000213301000203-1100302121322230-0322030000200123-3130030220321221-2100110300222120-0330132311323110"></a>

## context property — exclude_violation_contexts / 332230121203 / 4

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2331132200110012-2323133030002030-1103203213222131-3010030102110001-1020320021031023-1223303310332231-1130221321311100-3001002230313130"></a>

<a id="canonical-3300023110201113-2003020022201012-3231003200302322-3101101021121012-2130223211030100-1232121133131101-3001000203011303-3030013230000213"></a>

## context_name property — exclude_violation_contexts / 332230121203 / 5

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3321133301001022-3230213101020230-0103202203100010-3231032103230102-1320221302200110-0220231231013003-1000002012123322-0223032212312123"></a>

<a id="canonical-2331013131320212-2011220302103203-3111221221331212-0330211002120030-0311210131312023-1120132131302203-3300320203303113-3130100332203313"></a>

## exclude_violation property — exclude_violation_contexts / 332230121203 / 6

Type: `"string"`. Computed.

\[Enum:
VIOL\_NONE|VIOL\_FILETYPE|VIOL\_METHOD|VIOL\_MANDATORY\_HEADER|VIOL\_HTTP\_RESPONSE\_STATUS|VIOL\_REQUEST\_MAX\_LENGTH|VIOL\_FILE\_UPLOAD|VIOL\_FILE\_UPLOAD\_IN\_BODY|VIOL\_XML\_MALFORMED|VIOL\_JSON\_MALFORMED|VIOL\_ASM\_COOKIE\_MODIFIED|VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS|VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE|VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT|VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST|VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION|VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS|VIOL\_EVASION\_DIRECTORY\_TRAVERSALS|VIOL\_MALFORMED\_REQUEST|VIOL\_EVASION\_MULTIPLE\_DECODING|VIOL\_DATA\_GUARD|VIOL\_EVASION\_APACHE\_WHITESPACE|VIOL\_COOKIE\_MODIFIED|VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS|VIOL\_EVASION\_IIS\_BACKSLASHES|VIOL\_EVASION\_PERCENT\_U\_DECODING|VIOL\_EVASION\_BARE\_BYTE\_DECODING|VIOL\_EVASION\_BAD\_UNESCAPE|VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST|VIOL\_ENCODING|VIOL\_COOKIE\_MALFORMED|VIOL\_GRAPHQL\_FORMAT|VIOL\_GRAPHQL\_MALFORMED|VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\]
List of all supported Violation Types VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER
VIOL\_HTTP\_RESPONSE\_STATUS VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD
VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED
VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS.. Possible values are \`VIOL\_NONE\`,
\`VIOL\_FILETYPE\`, \`VIOL\_METHOD\`, \`VIOL\_MANDATORY\_HEADER\`, \`VIOL\_HTTP\_RESPONSE\_STATUS\`,
\`VIOL\_REQUEST\_MAX\_LENGTH\`, \`VIOL\_FILE\_UPLOAD\`, \`VIOL\_FILE\_UPLOAD\_IN\_BODY\`,
\`VIOL\_XML\_MALFORMED\`, \`VIOL\_JSON\_MALFORMED\`, \`VIOL\_ASM\_COOKIE\_MODIFIED\`,
\`VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE\`,
\`VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT\`, \`VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST\`,
\`VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION\`,
\`VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS\`,
\`VIOL\_EVASION\_DIRECTORY\_TRAVERSALS\`, \`VIOL\_MALFORMED\_REQUEST\`,
\`VIOL\_EVASION\_MULTIPLE\_DECODING\`, \`VIOL\_DATA\_GUARD\`, \`VIOL\_EVASION\_APACHE\_WHITESPACE\`,
\`VIOL\_COOKIE\_MODIFIED\`, \`VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS\`,
\`VIOL\_EVASION\_IIS\_BACKSLASHES\`, \`VIOL\_EVASION\_PERCENT\_U\_DECODING\`,
\`VIOL\_EVASION\_BARE\_BYTE\_DECODING\`, \`VIOL\_EVASION\_BAD\_UNESCAPE\`,
\`VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST\`, \`VIOL\_ENCODING\`,
\`VIOL\_COOKIE\_MALFORMED\`, \`VIOL\_GRAPHQL\_FORMAT\`, \`VIOL\_GRAPHQL\_MALFORMED\`,
\`VIOL\_GRAPHQL\_INTROSPECTION\_QUERY\`. Defaults to \`VIOL\_NONE\`.

Upstream description:

List of all supported Violation Types

VIOL\_NONE VIOL\_FILETYPE VIOL\_METHOD VIOL\_MANDATORY\_HEADER VIOL\_HTTP\_RESPONSE\_STATUS
VIOL\_REQUEST\_MAX\_LENGTH VIOL\_FILE\_UPLOAD VIOL\_FILE\_UPLOAD\_IN\_BODY VIOL\_XML\_MALFORMED
VIOL\_JSON\_MALFORMED VIOL\_ASM\_COOKIE\_MODIFIED VIOL\_HTTP\_PROTOCOL\_MULTIPLE\_HOST\_HEADERS
VIOL\_HTTP\_PROTOCOL\_BAD\_HOST\_HEADER\_VALUE VIOL\_HTTP\_PROTOCOL\_UNPARSABLE\_REQUEST\_CONTENT
VIOL\_HTTP\_PROTOCOL\_NULL\_IN\_REQUEST VIOL\_HTTP\_PROTOCOL\_BAD\_HTTP\_VERSION
VIOL\_HTTP\_PROTOCOL\_CRLF\_CHARACTERS\_BEFORE\_REQUEST\_START
VIOL\_HTTP\_PROTOCOL\_NO\_HOST\_HEADER\_IN\_HTTP\_1\_1\_REQUEST
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_PARAMETERS\_PARSING
VIOL\_HTTP\_PROTOCOL\_SEVERAL\_CONTENT\_LENGTH\_HEADERS
VIOL\_HTTP\_PROTOCOL\_CONTENT\_LENGTH\_SHOULD\_BE\_A\_POSITIVE\_NUMBER
VIOL\_EVASION\_DIRECTORY\_TRAVERSALS VIOL\_MALFORMED\_REQUEST VIOL\_EVASION\_MULTIPLE\_DECODING
VIOL\_DATA\_GUARD VIOL\_EVASION\_APACHE\_WHITESPACE VIOL\_COOKIE\_MODIFIED
VIOL\_EVASION\_IIS\_UNICODE\_CODEPOINTS VIOL\_EVASION\_IIS\_BACKSLASHES
VIOL\_EVASION\_PERCENT\_U\_DECODING VIOL\_EVASION\_BARE\_BYTE\_DECODING VIOL\_EVASION\_BAD\_UNESCAPE
VIOL\_HTTP\_PROTOCOL\_BAD\_MULTIPART\_FORMDATA\_REQUEST\_PARSING
VIOL\_HTTP\_PROTOCOL\_BODY\_IN\_GET\_OR\_HEAD\_REQUEST
VIOL\_HTTP\_PROTOCOL\_HIGH\_ASCII\_CHARACTERS\_IN\_HEADERS VIOL\_ENCODING VIOL\_COOKIE\_MALFORMED
VIOL\_GRAPHQL\_FORMAT VIOL\_GRAPHQL\_MALFORMED VIOL\_GRAPHQL\_INTROSPECTION\_QUERY.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIOL_NONE",
  "enum": [
    "VIOL_NONE",
    "VIOL_FILETYPE",
    "VIOL_METHOD",
    "VIOL_MANDATORY_HEADER",
    "VIOL_HTTP_RESPONSE_STATUS",
    "VIOL_REQUEST_MAX_LENGTH",
    "VIOL_FILE_UPLOAD",
    "VIOL_FILE_UPLOAD_IN_BODY",
    "VIOL_XML_MALFORMED",
    "VIOL_JSON_MALFORMED",
    "VIOL_ASM_COOKIE_MODIFIED",
    "VIOL_HTTP_PROTOCOL_MULTIPLE_HOST_HEADERS",
    "VIOL_HTTP_PROTOCOL_BAD_HOST_HEADER_VALUE",
    "VIOL_HTTP_PROTOCOL_UNPARSABLE_REQUEST_CONTENT",
    "VIOL_HTTP_PROTOCOL_NULL_IN_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_HTTP_VERSION",
    "VIOL_HTTP_PROTOCOL_CRLF_CHARACTERS_BEFORE_REQUEST_START",
    "VIOL_HTTP_PROTOCOL_NO_HOST_HEADER_IN_HTTP_1_1_REQUEST",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_PARAMETERS_PARSING",
    "VIOL_HTTP_PROTOCOL_SEVERAL_CONTENT_LENGTH_HEADERS",
    "VIOL_HTTP_PROTOCOL_CONTENT_LENGTH_SHOULD_BE_A_POSITIVE_NUMBER",
    "VIOL_EVASION_DIRECTORY_TRAVERSALS",
    "VIOL_MALFORMED_REQUEST",
    "VIOL_EVASION_MULTIPLE_DECODING",
    "VIOL_DATA_GUARD",
    "VIOL_EVASION_APACHE_WHITESPACE",
    "VIOL_COOKIE_MODIFIED",
    "VIOL_EVASION_IIS_UNICODE_CODEPOINTS",
    "VIOL_EVASION_IIS_BACKSLASHES",
    "VIOL_EVASION_PERCENT_U_DECODING",
    "VIOL_EVASION_BARE_BYTE_DECODING",
    "VIOL_EVASION_BAD_UNESCAPE",
    "VIOL_HTTP_PROTOCOL_BAD_MULTIPART_FORMDATA_REQUEST_PARSING",
    "VIOL_HTTP_PROTOCOL_BODY_IN_GET_OR_HEAD_REQUEST",
    "VIOL_HTTP_PROTOCOL_HIGH_ASCII_CHARACTERS_IN_HEADERS",
    "VIOL_ENCODING",
    "VIOL_COOKIE_MALFORMED",
    "VIOL_GRAPHQL_FORMAT",
    "VIOL_GRAPHQL_MALFORMED",
    "VIOL_GRAPHQL_INTROSPECTION_QUERY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130333211321313-3132022303312021-0323302023320213-0102231132000312-2300013311331322-0020330211331203-2022220310112032-2313310202130223"></a>

## Next pages — exclude_violation_contexts / 332230121203 / 7

- [rule_list.rules.spec.waf_action.app_firewall_detection_control](data-sources--service_policy--reference--group-003.md#canonical-0312203122000130-1211201300112130-2033131032302133-2131001020021211-1223121201301131-3203110303111102-2022322010100302-0112200311232020)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0310123023311333-0022102211121010-0020210033011011-2001230110222211-2210213330032131-3212220310101201-1223020211031333-3230113300311321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202200232000022-3311023300121030-1111103300101021-2131023012023200-3333020232120212-0002201002130220-1332103311211121-3303321002110111"></a>

## rule_list.rules.spec.waf_action.none — none / 221302032320 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- rule_list.rules.spec.waf_action.none

<a id="canonical-2203311033100230-3111003011320013-1332130103303122-1310101021212132-1320110233120033-3232112031330130-3310330322330033-1002322232233121"></a>

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

<a id="canonical-1202030103202331-1011000121312321-1222100110202112-1101202223103030-3323313123333103-3002113330003203-0013110302101300-0320221031323322"></a>

## Direct properties — none / 221302032320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210333220132321-3212211200220311-1101131110213220-2221323031110102-1012022120330100-0213302031330202-1211103010121113-2331222122320121"></a>

## Next pages — none / 221302032320 / 4

- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-0000112103021210-0222132122221021-1220123212112130-2313130013212221-2010010233130222-3122031033113002-0031013220303221-1020013000133310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331312232001022-2002022101132021-2223321323130013-0003213011210231-3231012211112201-1113201012202200-3311302000011312-1323131232011101"></a>

## rule_list.rules.spec.waf_action.waf_skip_processing — waf_skip_processing / 301220333031 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [rule_list](data-sources--service_policy--reference--group-001.md#canonical-0200021302120322-3330103111223310-0110033233312323-2211001021302330-2000313233321013-1330302203133111-3101023222321313-3302013113202033)
- [rule_list.rules](data-sources--service_policy--reference--group-001.md#canonical-2332210012031020-3130121121203233-2133312023012212-1300213020002321-0210122201111201-2313301011130132-1102100221012102-3320333010302121)
- [rule_list.rules.spec](data-sources--service_policy--reference--group-001.md#canonical-1313133011110120-1202202112012022-0312300012021010-2230300132201133-1123333233212022-2100210133123323-2030333300021222-2110310211021210)
- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- rule_list.rules.spec.waf_action.waf_skip_processing

<a id="canonical-0132120133010122-1021113000121033-3232231102111113-0123102230223213-2130123312110203-1132011330211331-2121311032221321-0313103021211210"></a>

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

<a id="canonical-3220312230311132-3103000212332333-1003231033312013-2123122022301011-0322011020131022-0013200133002202-3001113031021320-0032013330032202"></a>

## Direct properties — waf_skip_processing / 301220333031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321320222033012-2301003030302103-3330223213031133-2230011320010032-1231222011303200-2302300201331233-0021020101333333-0201300232033322"></a>

## Next pages — waf_skip_processing / 301220333031 / 4

- [rule_list.rules.spec.waf_action](data-sources--service_policy--reference--group-003.md#canonical-1321030020301232-0323133310020002-1031312201202122-2313331303112210-3101113223232332-0322221031012323-3201210323110133-2121002301220100)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3021113112310011-2312023300100332-3211231333010200-0121033221312122-0113321130233033-3313001233322031-0000023202030312-0011301313100330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223133022021211-0331222003011300-3233030010103013-1221112021312033-2220321000333201-1231000022021220-3102321000031000-1201202010330112"></a>

## server_name_matcher — server_name_matcher / 101303022312 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- server_name_matcher

<a id="canonical-1031033321123202-3020320322203132-2113112002313103-3203303203110001-3300220113221302-3312210031110132-0122213123201000-2230003120010133"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3030031113123111-2031032223323022-0030001201132210-0310312112012132-2000203100122321-1301011200320333-1101331000111230-2031110321000322"></a>

## Direct properties — server_name_matcher / 101303022312 / 3

<a id="canonical-3102111120301111-0023003010303110-1121230010300330-0330133120331220-3102033131230022-0201123130301130-0033023200330101-0202030222301322"></a>

<a id="canonical-2210331133002021-3221202020013023-3232002002013021-0303303212031230-2322310130311133-2333131302130232-0330122010031200-1321011030221001"></a>

## exact_values property — server_name_matcher / 101303022312 / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

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

<a id="canonical-0033203303333232-2330202232201321-0213131123033320-3010001200021203-3332220323302031-0320023322232000-3231002200233122-3223330123330222"></a>

<a id="canonical-2212011220202131-2213022312221132-2330133320110132-3001113103020232-2130102313300300-0203233212021123-3333321322321202-3132113220101001"></a>

## regex_values property — server_name_matcher / 101303022312 / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

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

<a id="canonical-3301210002232012-0321220202201322-1213320011031122-1330231011300300-0202113130133002-2331122230231233-2100323122302111-2221122000202210"></a>

## Next pages — server_name_matcher / 101303022312 / 6

- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)

<a id="canonical-3122313323222133-3022130222310211-0113113011303232-0203212121203123-0123123113030131-0111331230021010-3121011333120221-2330322213331033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021321310232020-0230020021012020-0100331203222202-0303310230203102-1223001003113313-2322222321212213-3211232333323300-1223033100020130"></a>

## server_selector — server_selector / 103311320200 / 2

Breadcrumbs:

- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- server_selector

<a id="canonical-3031212021313110-0302012231302222-1103201333123221-1112231301032212-3312311023100330-0321020012301232-0331330301220320-3112122133101112"></a>

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

<a id="canonical-1113230023313113-3331331100303013-3110222231013033-1220111001002102-1302300310130312-1300222110302201-3202300023310202-2013203311303211"></a>

## Direct properties — server_selector / 103311320200 / 3

<a id="canonical-0220131332331303-2310002010012121-2302003200221023-1132030230311202-2010221001010303-1122120300033102-3322323013300002-0101103301110200"></a>

<a id="canonical-0202032020131200-3101231332133322-1322220233311000-1332120022132102-3102303330321310-1333022003200010-0331003221103203-2131330102123103"></a>

## expressions property — server_selector / 103311320200 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1201202031202131-3312203031122003-2310020112120121-0030121310011300-3020312012133200-0303111221200303-0033031123333211-0002012231332030"></a>

## Next pages — server_selector / 103311320200 / 5

- [Property reference](data-sources--service_policy--reference--group-001.md#canonical-3312102302111103-0221020013211300-2120323123300113-1101123102331230-1032322121222101-1332231300102132-0233300331323303-0031323122032013)
- [xcsh_service_policy](../data-sources/service_policy.md#canonical-1112022302333333-1130100232232301-3023212310221022-3103031003221200-2030020130230132-1201210321332200-2212311003030121-1321310310032303)
