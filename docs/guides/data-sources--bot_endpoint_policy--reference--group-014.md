---
page_title: "xcsh_bot_endpoint_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_endpoint_policy reference."
---

# xcsh_bot_endpoint_policy reference

<a id="canonical-0231033211201011-2112210302031000-1330222022003321-0003312310131100-2310023302000130-1323112212311010-1202231131303203-3201100001000002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body](data-sources--bot_endpoint_policy--reference--group-013.md#canonical-3002120302230233-2030303202231231-2233223130300111-1113023211313210-2021132012021312-0110301201302331-0333333010001323-1330231032023130)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none](data-sources--bot_endpoint_policy--reference--group-013.md#canonical-2101210231103210-1202312100010221-0300222012132212-3013231112203221-3210210300223200-2101131123302011-3231321132211231-3133103010321211)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2

<a id="canonical-2310022030000202-2132333022211030-3132033222132101-2000101213011302-0332220233113330-3133301020200322-1111332121333110-0013231211122223"></a>

Type: `"list"`. Computed.

Response Body Matcher(s). Response Body Matchers.

<a id="canonical-3122320231011203-3002011103011112-0101320313301100-0020112003202021-0321320300300113-0112232311123031-1302010010311202-3213223101013030"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2`

<a id="canonical-2110113320330022-1123202033233030-2233300012233202-3132030202310330-2301001211300103-2121211321320322-1321133031023233-1332102310303001"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2.case_sensitive` property

Type: `"bool"`. Computed.

Configuration parameter for case sensitive.

<a id="canonical-0100112002120111-2310232221330123-0130303221310302-2332001030001321-3010212211233110-1203121011120113-2102212033132301-3313103030221122"></a>

<a id="canonical-1231133123122031-0123011231210230-3012103020013010-0331213330321302-3230310212111113-2132132210113301-0231000022033121-0003002000123020"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2.not` property

Type: `"bool"`. Computed.

Not(!). Configuration parameter for not

<a id="canonical-2212332321301010-1033021332012113-2021132232110132-2022221132313301-1030212012313203-0330321330102011-2100023031231022-3322013331032200"></a>

<a id="canonical-0323302121003331-0131132330201220-3101201212113113-1033311100102230-2222003000121031-3312022313231111-1223100003013232-0130000332001001"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2.operator` property

Type: `"string"`. Computed.

\[Enum:
RESPONSE\_OPERATOR\_EQUALS\_TO|RESPONSE\_OPERATOR\_CONTAINS|RESPONSE\_OPERATOR\_STARTS\_WITH|RESPONSE\_OPERATOR\_ENDS\_WITH\]
&#8203;- RESPONSE\_OPERATOR\_EQUALS\_TO: EQUALS\_TO value - RESPONSE\_OPERATOR\_CONTAINS: CONTAINS value -
RESPONSE\_OPERATOR\_STARTS\_WITH: STARTS\_WITH value - RESPONSE\_OPERATOR\_ENDS\_WITH: ENDS\_WITH
value. Possible values are \`RESPONSE\_OPERATOR\_EQUALS\_TO\`, \`RESPONSE\_OPERATOR\_CONTAINS\`,
\`RESPONSE\_OPERATOR\_STARTS\_WITH\`, \`RESPONSE\_OPERATOR\_ENDS\_WITH\`. Defaults to
\`RESPONSE\_OPERATOR\_EQUALS\_TO\`.

<a id="canonical-2223001021311001-1203013133222313-2120203211123313-0303322002133130-2203012132012221-3000001112332120-1303023113100001-0211112331231233"></a>

<a id="canonical-3012333011111102-3122313022013233-3110301212313211-1010123123201122-0210223313200130-2032331223031102-0102333330212230-1101301313301230"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_none.response_body_match_v2.value` property

Type: `"string"`. Computed.

Value. Configuration parameter for value

<a id="canonical-2020203200212022-3201103212231112-2102111320330133-2202013321213002-0300210321331003-2003300320102230-3110303012021313-2122222000100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body](data-sources--bot_endpoint_policy--reference--group-013.md#canonical-3002120302230233-2030303202231231-2233223130300111-1113023211313210-2021132012021312-0110301201302331-0333333010001323-1330231032023130)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or

<a id="canonical-3020312010200230-0011313120223132-0320331123232213-3101122010001321-2322321032133303-0022311220120212-2231032122123122-3021013023330333"></a>

Type: `"single"`. Computed.

Response Body Matcher. Response Body matcher Choice.

<a id="canonical-2030121100321230-0131112103012302-2212120031223210-1012323312021122-1201303320202211-0012323201210333-2023011133303133-2023031320220322"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or`

- [response_body_match_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1031301221100012-1320200330222302-3211311132311010-3011133011123000-3013220132022013-2312330112031321-2121111320233210-0021131113221323): complete subsection reference.

<a id="canonical-1031301221100012-1320200330222302-3211311132311010-3011133011123000-3013220132022013-2312330112031321-2121111320233210-0021131113221323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body](data-sources--bot_endpoint_policy--reference--group-013.md#canonical-3002120302230233-2030303202231231-2233223130300111-1113023211313210-2021132012021312-0110301201302331-0333333010001323-1330231032023130)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2020203200212022-3201103212231112-2102111320330133-2202013321213002-0300210321331003-2003300320102230-3110303012021313-2122222000100310)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2

<a id="canonical-0110003310130223-3130021131033033-3322002100230022-1020210322030022-1200221332131132-2031001112331000-1102310202320132-3130110103000222"></a>

Type: `"list"`. Computed.

Response Body Matcher(s). Response Body Matchers.

<a id="canonical-2020103322213332-2031121212331113-0022023023031011-0332323330213203-2300010000222303-3022212310121220-2112211331220232-0030132003100322"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2`

<a id="canonical-2322210123221001-3331113002320202-0232103300023010-0323223121112103-0112323001301021-1021201320021002-3301210113311010-1021001112133232"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2.case_sensitive` property

Type: `"bool"`. Computed.

Configuration parameter for case sensitive.

<a id="canonical-0322010201311020-3231102202022311-1321002303223100-1000031122023002-3021300022230122-0201331333123102-3102203021311232-0123222033233201"></a>

<a id="canonical-2012003223310220-1312022100300012-1232320331212201-3133032323000031-1330221331212110-2330003230111200-1211330131010303-0130001323212233"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2.not` property

Type: `"bool"`. Computed.

Not(!). Configuration parameter for not

<a id="canonical-2011011123000011-2110320112233020-0003300010212201-1333102233121133-0003301012031333-0213100000013010-3212232132121103-0032133112133112"></a>

<a id="canonical-1300123012010032-0033303210212123-1331030012032000-2220223032130233-1201202210133010-3211201102003312-1300322222011322-0003323320120113"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2.operator` property

Type: `"string"`. Computed.

\[Enum:
RESPONSE\_OPERATOR\_EQUALS\_TO|RESPONSE\_OPERATOR\_CONTAINS|RESPONSE\_OPERATOR\_STARTS\_WITH|RESPONSE\_OPERATOR\_ENDS\_WITH\]
&#8203;- RESPONSE\_OPERATOR\_EQUALS\_TO: EQUALS\_TO value - RESPONSE\_OPERATOR\_CONTAINS: CONTAINS value -
RESPONSE\_OPERATOR\_STARTS\_WITH: STARTS\_WITH value - RESPONSE\_OPERATOR\_ENDS\_WITH: ENDS\_WITH
value. Possible values are \`RESPONSE\_OPERATOR\_EQUALS\_TO\`, \`RESPONSE\_OPERATOR\_CONTAINS\`,
\`RESPONSE\_OPERATOR\_STARTS\_WITH\`, \`RESPONSE\_OPERATOR\_ENDS\_WITH\`. Defaults to
\`RESPONSE\_OPERATOR\_EQUALS\_TO\`.

<a id="canonical-0033210330330213-1322023322212221-1111321301020132-2331121212223013-0000221203012121-3103001102013312-2332231131302331-3132223003312123"></a>

<a id="canonical-2221331102113311-2220200013332300-3232202113000132-1011023020223213-2100122110100312-0013130320011032-1221020202122002-3333330302301131"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_body.response_body_or.response_body_match_v2.value` property

Type: `"string"`. Computed.

Value. Configuration parameter for value

<a id="canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code

<a id="canonical-2022231120202123-1110232102301232-3103320021322311-0120033033222203-0030000032333212-3013003322222110-0120220112211013-2232013210231100"></a>

Type: `"single"`. Computed.

Response Code.

<a id="canonical-3233000220122200-0110213023331010-1312130321020212-2320133102130020-0133002332210222-1300321330031233-2330023103122332-1130112223300322"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code`

- [response_code_all](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3131320000021303-0033333133011321-2100320212200232-0222112210123031-0030131301330110-3332012110110210-0100202230231112-3303001220300331): complete subsection reference.

- [response_code_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3021330300323020-2311303203111002-0301333121303301-0032013202310121-2121122113113203-0002231123002202-1122130021010031-1312301211200123): complete subsection reference.

- [response_code_none](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1012003101022030-0313210232011200-1130232020312231-3110333100202032-1312103130221220-1302021200300030-2100133310012031-2033320110010131): complete subsection reference.

- [response_code_or](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3320022111333221-3012032132213201-2111100112120032-3132332303012320-3111101313222033-2131023321100232-0132231122003033-2232331202203321): complete subsection reference.

<a id="canonical-3131320000021303-0033333133011321-2100320212200232-0222112210123031-0030131301330110-3332012110110210-0100202230231112-3303001220300331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_all` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_all

<a id="canonical-1223130111100111-3101312122323112-1010322100300223-0021203303221301-0001030101231231-0131110013021212-2331030033001102-2331201032311233"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021330300323020-2311303203111002-0301333121303301-0032013202310121-2121122113113203-0002231123002202-1122130021010031-1312301211200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and

<a id="canonical-0132032310012101-0012201221020010-2210011320012011-3301211110002131-2001033033202213-1330002331311012-1103311131221020-3211100020311011"></a>

Type: `"single"`. Computed.

Response Code Matcher. Response Header matcher choice.

<a id="canonical-0000330303111012-3233033220200303-0203223202120001-1011230201130321-0212002100121030-0013010232323032-1021232102000232-3022112323123230"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and`

- [response_code_match](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3103200111220132-1110020233012031-0122230223011112-0103002303233131-1331211000000331-3203232131013011-3013321200213300-3023233013200201): complete subsection reference.

<a id="canonical-3103200111220132-1110020233012031-0122230223011112-0103002303233131-1331211000000331-3203232131013011-3013321200213300-3023233013200201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3021330300323020-2311303203111002-0301333121303301-0032013202310121-2121122113113203-0002231123002202-1122130021010031-1312301211200123)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match

<a id="canonical-3232202232121133-2101201011330312-1312203330130231-2211321221111132-0033211001212123-1312121102001031-1231132211231020-3211311331330110"></a>

Type: `"list"`. Computed.

Response Code Matcher(s). Response Code Matchers.

<a id="canonical-2333220231110311-1322213131001332-0231110010002232-0332001333303332-3103101100311333-3221101320231120-3030130032322333-1001223211131230"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match`

<a id="canonical-2120301213220321-3203031332023323-1010112130012000-0331123112100202-1111033220231100-2120111232123323-0311102002331121-2030113311301300"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match.operator` property

Type: `"string"`. Computed.

\[Enum:
CODE\_EQUALS|CODE\_NOT\_EQUAL\_TO|CODE\_LESS\_THAN|CODE\_GREATER\_THAN|CODE\_LESS\_THAN\_OR\_EQUAlS\_TO|CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\]
&#8203;- CODE\_EQUALS: EQUALS value - CODE\_NOT\_EQUAL\_TO: NOT\_EQUAL\_TO value - CODE\_LESS\_THAN:
LESS\_THAN value - CODE\_GREATER\_THAN: GREATER\_THAN value - CODE\_LESS\_THAN\_OR\_EQUAlS\_TO:
LESS\_THAN\_OR\_EQUAlS\_TO value - CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO:
GREATER\_THAN\_OR\_EQUAlS\_TO value. Possible values are \`CODE\_EQUALS\`, \`CODE\_NOT\_EQUAL\_TO\`,
\`CODE\_LESS\_THAN\`, \`CODE\_GREATER\_THAN\`, \`CODE\_LESS\_THAN\_OR\_EQUAlS\_TO\`,
\`CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\`. Defaults to \`CODE\_EQUALS\`.

<a id="canonical-1033302333013013-2233120120223000-1132222203211331-3103323032222121-3302122210321323-1311022311201131-1302132323023213-2030320203133203"></a>

<a id="canonical-1002012212001110-2333132203312102-0110213012211330-2113003322002321-3100311231023302-2211200101322131-0022112010203330-0223220213111213"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_and.response_code_match.value` property

Type: `"number"`. Computed.

Value. Configuration parameter for value

<a id="canonical-1012003101022030-0313210232011200-1130232020312231-3110333100202032-1312103130221220-1302021200300030-2100133310012031-2033320110010131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none

<a id="canonical-1202200021302110-0002310110233310-3201203333302023-0110230131210122-2320310002112111-1330122131021212-3120331302222231-3231000211210313"></a>

Type: `"single"`. Computed.

Response Code Matcher. Response Header matcher choice.

<a id="canonical-1330113212212023-0112332321333120-0001012011311111-0130213001003313-0322003221210131-3001123002210032-1222110021302311-3100110212320303"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none`

- [response_code_match](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-0301001110033233-3211212311303212-0002113331121030-2211033203232330-0011221002231303-2311131200332300-1131111220312300-2233013321031223): complete subsection reference.

<a id="canonical-0301001110033233-3211212311303212-0002113331121030-2211033203232330-0011221002231303-2311131200332300-1131111220312300-2233013321031223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none.response_code_match` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1012003101022030-0313210232011200-1130232020312231-3110333100202032-1312103130221220-1302021200300030-2100133310012031-2033320110010131)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none.response_code_match

<a id="canonical-0212202101200122-3121021302002312-0133010332010332-1133001202131121-3031032311012131-1120201211210030-2303331223010120-3001023121232012"></a>

Type: `"list"`. Computed.

Response Code Matcher(s). Response Code Matchers.

<a id="canonical-2020311203113000-1333031011213203-2113120030323232-1003233133112132-2312022311323113-2123022032011001-1302302301332221-2000203132212001"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none.response_code_match`

<a id="canonical-0023001022222122-2201112233223010-3100231023232211-2202130320103333-0113131200211110-2111103021030112-1300201303022001-3201323203212122"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none.response_code_match.operator` property

Type: `"string"`. Computed.

\[Enum:
CODE\_EQUALS|CODE\_NOT\_EQUAL\_TO|CODE\_LESS\_THAN|CODE\_GREATER\_THAN|CODE\_LESS\_THAN\_OR\_EQUAlS\_TO|CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\]
&#8203;- CODE\_EQUALS: EQUALS value - CODE\_NOT\_EQUAL\_TO: NOT\_EQUAL\_TO value - CODE\_LESS\_THAN:
LESS\_THAN value - CODE\_GREATER\_THAN: GREATER\_THAN value - CODE\_LESS\_THAN\_OR\_EQUAlS\_TO:
LESS\_THAN\_OR\_EQUAlS\_TO value - CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO:
GREATER\_THAN\_OR\_EQUAlS\_TO value. Possible values are \`CODE\_EQUALS\`, \`CODE\_NOT\_EQUAL\_TO\`,
\`CODE\_LESS\_THAN\`, \`CODE\_GREATER\_THAN\`, \`CODE\_LESS\_THAN\_OR\_EQUAlS\_TO\`,
\`CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\`. Defaults to \`CODE\_EQUALS\`.

<a id="canonical-2211030021110131-0001302012233030-0112123303012303-3212112320122103-2022212332310030-2201321320320333-0213022330132220-0213103101022111"></a>

<a id="canonical-0000011300122200-1200213220323003-0000222012320313-2202311222002220-3023133200123112-0012202212000213-1330020113110311-1032330133303121"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_none.response_code_match.value` property

Type: `"number"`. Computed.

Value. Configuration parameter for value

<a id="canonical-3320022111333221-3012032132213201-2111100112120032-3132332303012320-3111101313222033-2131023321100232-0132231122003033-2232331202203321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or

<a id="canonical-3001010321113321-3210023330310313-3000022221101320-0113111222132130-1213201023322020-0310102231110221-3102012002123301-2003013212302333"></a>

Type: `"single"`. Computed.

Response Code Matcher. Response Header matcher choice.

<a id="canonical-0020211310330333-0220031131220001-0300000223332133-0211003301012322-0103200011223212-3323003122201313-3113311231230300-0213313010122202"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or`

- [response_code_match](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-0131220112010111-0100130212213201-0013030211201023-1120210311023100-3003231311030300-2122123231033332-1232221010122002-3302211321320301): complete subsection reference.

<a id="canonical-0131220112010111-0100130212213201-0013030211201023-1120210311023100-3003231311030300-2122123231033332-1232221010122002-3302211321320301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1202202320211232-0300212103222132-3102033132323312-3030312100102301-1211320022220222-0022223230003331-0213121133330312-2020101123331203)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3320022111333221-3012032132213201-2111100112120032-3132332303012320-3111101313222033-2131023321100232-0132231122003033-2232331202203321)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match

<a id="canonical-2120101013032103-0230110012023220-2300101203123211-0001111232133012-2011003100013000-3010223013131231-2103332230103132-2013032331311311"></a>

Type: `"list"`. Computed.

Response Code Matcher(s). Response Code Matchers.

<a id="canonical-0321102202223113-3301211003310222-1320213012022212-0321022102033230-2230013012203033-1203210213333203-1333212310202203-1123333203230302"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match`

<a id="canonical-2000230312131221-2022013033113130-1331212303202212-1310213302031002-3301100222233132-2121323211010233-2020023103311211-0310322002102303"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match.operator` property

Type: `"string"`. Computed.

\[Enum:
CODE\_EQUALS|CODE\_NOT\_EQUAL\_TO|CODE\_LESS\_THAN|CODE\_GREATER\_THAN|CODE\_LESS\_THAN\_OR\_EQUAlS\_TO|CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\]
&#8203;- CODE\_EQUALS: EQUALS value - CODE\_NOT\_EQUAL\_TO: NOT\_EQUAL\_TO value - CODE\_LESS\_THAN:
LESS\_THAN value - CODE\_GREATER\_THAN: GREATER\_THAN value - CODE\_LESS\_THAN\_OR\_EQUAlS\_TO:
LESS\_THAN\_OR\_EQUAlS\_TO value - CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO:
GREATER\_THAN\_OR\_EQUAlS\_TO value. Possible values are \`CODE\_EQUALS\`, \`CODE\_NOT\_EQUAL\_TO\`,
\`CODE\_LESS\_THAN\`, \`CODE\_GREATER\_THAN\`, \`CODE\_LESS\_THAN\_OR\_EQUAlS\_TO\`,
\`CODE\_GREATER\_THAN\_OR\_EQUAlS\_TO\`. Defaults to \`CODE\_EQUALS\`.

<a id="canonical-3300313233233303-1002312231311330-0301221022002002-2001233102313012-3122332301301021-0230003012132010-0020321131321000-2333220323303122"></a>

<a id="canonical-3310101320220210-2032223023031012-0003222123021320-1102000323033222-1023212223132301-0130302300000201-2022012202100111-3212121122322021"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_code.response_code_or.response_code_match.value` property

Type: `"number"`. Computed.

Value. Configuration parameter for value

<a id="canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2

<a id="canonical-3313110311223020-2122120131303201-0023023000333111-0122133301132102-3021032130121232-2303123023122012-3323030011313131-2332331020302100"></a>

Type: `"single"`. Computed.

Configuration parameter for response header v2.

<a id="canonical-2132222330012031-3331121213322002-3011211030100002-1123312101313103-2021233111032232-3310220030033212-3112300020113123-1000211201011011"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2`

- [response_header_all](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1332002301202023-2310120320003032-2113131131023023-0133032301310331-1211312332201222-2212121022023021-0101232033300011-2102330223223311): complete subsection reference.

- [response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121): complete subsection reference.

- [response_header_none](data-sources--bot_endpoint_policy--reference--group-015.md#canonical-3302111211031100-1311032010300101-2101022111113030-0203110303121023-0311202001221122-1012010233203330-1102123333222321-0332031310023120): complete subsection reference.

- [response_header_or](data-sources--bot_endpoint_policy--reference--group-015.md#canonical-2000323010232330-3200312200333103-0331211320033213-2131022132301221-0002231320231100-3113002330223033-1210200020222220-1311232001310322): complete subsection reference.

<a id="canonical-1332002301202023-2310120320003032-2113131131023023-0133032301310331-1211312332201222-2212121022023021-0101232033300011-2102330223223311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_all` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_all

<a id="canonical-2001010023021311-2132033211230002-2032202330032301-1321113221003221-0012020133223020-2020210122132330-0100311213310020-1312122222210210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for response header all.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and

<a id="canonical-0220101230320000-3223001010022332-3211101031302210-1110021103131111-1200132310213133-0102212331212312-3032120012332312-3131303002233230"></a>

Type: `"single"`. Computed.

Configuration parameter for response header and.

<a id="canonical-1100031111321231-1231220201120011-0132032132222300-3011130230121123-3220202312333001-2033033112102022-0221110132322331-1212203112013002"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and`

- [response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311): complete subsection reference.

<a id="canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator

<a id="canonical-1210132111113000-0201311022333223-0331332102200200-2223110003233233-1100112112332231-3222002220030200-3012101302303022-0221302323300130"></a>

Type: `"list"`. Computed.

Configuration parameter for response header operator.

<a id="canonical-1030221233121322-0111230210130133-3132122323331130-3321201330000232-2100021022101000-0131111231130132-2021303321320122-1001201032103220"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator`

- [header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321): complete subsection reference.

<a id="canonical-2011121321230030-1103223330302213-3033200132332133-3320200322112022-1031022012103011-2232003313121313-2210320112112231-0030020113122010"></a>

<a id="canonical-0213102213223123-3011133233123021-3310101320231001-0212323200332012-3100103103111331-2033013310010032-2011013331020333-1010102000320013"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.name` property

Type: `"string"`. Computed.

Header Name. Header Name.

<a id="canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header

<a id="canonical-2313330302033131-0300231313120311-1323230320231030-3113133131300102-2220202133311013-2120321000302101-0320013231012201-0031233122322222"></a>

Type: `"single"`. Computed.

Operator. Operator

<a id="canonical-3132222120201211-0202133213112103-1121220003131322-0113323200301100-1332102133012220-3100133122223013-0231231330111311-1312110001011200"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header`

- [header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1000320013131022-1310323313222220-0321003101121221-2130220000111302-0031021112002330-3113123123311333-1002332203100312-0233112233200001): complete subsection reference.

- [header_anything](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-0103133121331102-3133113022330221-2231010030000130-3303320020132320-1110110030312321-3332330131211230-2113003112330232-0322101023031120): complete subsection reference.

- [header_none](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-0322032111203220-0033211232133113-3333211330322212-0301320300202121-2323223103332101-1310001001331011-3233212133033011-3201022101321023): complete subsection reference.

- [header_not_present](data-sources--bot_endpoint_policy--reference--group-015.md#canonical-2200031303101011-1033020033213002-2102212303313133-2201312311323100-0112200330333103-1102110123321002-1231232323113331-3302031122132211): complete subsection reference.

- [header_or](data-sources--bot_endpoint_policy--reference--group-015.md#canonical-3313122103313312-1212103130001033-0223131201001332-2310030111213101-3221232302300313-2313003120120201-3132203132302213-0311222230011301): complete subsection reference.

<a id="canonical-1000320013131022-1310323313222220-0321003101121221-2130220000111302-0031021112002330-3113123123311333-1002332203100312-0233112233200001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and

<a id="canonical-0121231020022333-0000203103233030-1321211122120221-0303021221333210-1300130013300223-0301000130000302-3023031122300330-2002111103232303"></a>

Type: `"single"`. Computed.

Header Matcher. Response Header matcher Choice.

<a id="canonical-0203002312310321-3111200221111302-0323000200223123-1011332132120312-0230111102322332-1322232230202331-2322220130211332-3111320333231301"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and`

- [response_header_match_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2002033001202233-1020122331103202-1230202102100112-1002312133120111-1021300003010032-3031212210212131-0222202211021120-3023211021101110): complete subsection reference.

<a id="canonical-2002033001202233-1020122331103202-1230202102100112-1002312133120111-1021300003010032-3031212210212131-0222202211021120-3023211021101110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1000320013131022-1310323313222220-0321003101121221-2130220000111302-0031021112002330-3113123123311333-1002332203100312-0233112233200001)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2

<a id="canonical-0130223123300313-0131220021133111-3320021303110021-2132123322212221-3330310222030211-3321021133031223-1221132311300130-0331301131122303"></a>

Type: `"list"`. Computed.

Response Header Matcher(s). Response Header Matchers.

<a id="canonical-1331302333323221-1010203002120210-3203020113013322-0310102100200221-1102033310012002-1101311032203102-1121131033010301-1201032300230130"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2`

<a id="canonical-3100310021311012-2123013011121103-1220311222310023-0012112013210300-0210233223033200-0202313133223301-2113203012322301-1123201310311020"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2.case_sensitive` property

Type: `"bool"`. Computed.

Configuration parameter for case sensitive.

<a id="canonical-0033120232032031-0213112000232030-0223210112322032-0230132132123303-2333010202013122-2002202322111000-1232200003331101-2122010212030202"></a>

<a id="canonical-1302021302110210-3100132323230122-2213031012122200-0132313332200301-2321201312132221-3032000303231013-1232203311301003-3233301211222000"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2.not` property

Type: `"bool"`. Computed.

Not(!). Configuration parameter for not

<a id="canonical-2012202101103003-3111332132310221-1303223030322102-2033210133310211-0110231223200333-0130323102121330-0313023023210001-3322022021012231"></a>

<a id="canonical-3321011200101320-2100331120112320-2210013220020203-2032101211022001-0323313330122110-2223321001203231-3033330313311010-3222023331131000"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2.operator` property

Type: `"string"`. Computed.

\[Enum:
RESPONSE\_OPERATOR\_EQUALS\_TO|RESPONSE\_OPERATOR\_CONTAINS|RESPONSE\_OPERATOR\_STARTS\_WITH|RESPONSE\_OPERATOR\_ENDS\_WITH\]
&#8203;- RESPONSE\_OPERATOR\_EQUALS\_TO: EQUALS\_TO value - RESPONSE\_OPERATOR\_CONTAINS: CONTAINS value -
RESPONSE\_OPERATOR\_STARTS\_WITH: STARTS\_WITH value - RESPONSE\_OPERATOR\_ENDS\_WITH: ENDS\_WITH
value. Possible values are \`RESPONSE\_OPERATOR\_EQUALS\_TO\`, \`RESPONSE\_OPERATOR\_CONTAINS\`,
\`RESPONSE\_OPERATOR\_STARTS\_WITH\`, \`RESPONSE\_OPERATOR\_ENDS\_WITH\`. Defaults to
\`RESPONSE\_OPERATOR\_EQUALS\_TO\`.

<a id="canonical-1213123230132212-1230233203121132-0221133233031332-3000110112301233-1300322110002233-3003232011121011-1213010120311312-3300112011232213"></a>

<a id="canonical-3011000211231231-0223132213123332-0210102212013212-0312220201021330-0210200210202110-1101113110012030-3022301133033030-1332313103220111"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_and.response_header_match_v2.value` property

Type: `"string"`. Computed.

Value. Configuration parameter for value

<a id="canonical-0103133121331102-3133113022330221-2231010030000130-3303320020132320-1110110030312321-3332330131211230-2113003112330232-0322101023031120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_anything` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_anything

<a id="canonical-1002132020101310-2002133122302102-3132200012100331-2312211201102011-3133102003320310-0101320210000023-2102130001302001-0200213232131033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for header anything.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322032111203220-0033211232133113-3333211330322212-0301320300202121-2323223103332101-1310001001331011-3233212133033011-3201022101321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none

<a id="canonical-3031223022020030-2121301113022023-0322121200232233-0111103300103103-0202122000021130-0121233210123332-0032213330100233-2323232010323023"></a>

Type: `"single"`. Computed.

Configuration parameter for header none.

<a id="canonical-2232303122320201-0131311030333211-0000203131310211-2030122130212110-1103312311232022-0011131330220031-1323312233030100-0133323312311313"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none`

- [response_header_match_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1132001303020103-2130203122203102-2020030010011133-1233302232232023-3010323200010032-3300103211201203-3213330110301300-1112110110010233): complete subsection reference.

<a id="canonical-1132001303020103-2130203122203102-2020030010011133-1233302232232023-3010323200010032-3300103211201203-3213330110301300-1112110110010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2` properties

Breadcrumbs:

- [xcsh_bot_endpoint_policy](../data-sources/bot_endpoint_policy.md#canonical-0320011110221201-2100233230203101-3330100033222310-3313200201232233-3120023321332231-0201323030313010-1100103303021130-0133102322333231)
- [Property reference](data-sources--bot_endpoint_policy--reference--group-001.md#canonical-2201213100123012-2230231320333220-2132303223203302-0013322032200310-2032113232303323-2311011301021201-2323010133021002-1332121131020011)
- [endpoint_policy_content](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-0210213102330220-1113322211223022-2131332133130302-2331330321103210-2202002320311112-0131102003032103-3022330010213201-0033101113331032)
- [endpoint_policy_content.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1003332120300123-0311030111211122-1303003012311012-2302120100300322-1002321233323110-3130120012030100-0110031123111213-2333130222330103)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints](data-sources--bot_endpoint_policy--reference--group-004.md#canonical-1132111333013010-2211321310011120-3233013002110322-0333103123021110-3221130131311113-1302021130310201-2001321322333130-2120320222330331)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria](data-sources--bot_endpoint_policy--reference--group-008.md#canonical-2022103320231012-2113101132303022-0012113203100131-2013302331010013-1312310032132013-0133303212133133-0320132302211101-1211322223330302)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success](data-sources--bot_endpoint_policy--reference--group-012.md#canonical-2322030202013023-3001302202021133-3112000310011100-0031011321302223-1102233201020123-0013022001322022-1131120302121312-3023120203313323)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3022210332102010-3303200122333013-0102311003130323-2211203103022001-0102331021233001-2102132331021020-0331133302201322-0010302032211132)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-2230223320220022-2213201131030220-2131103213202103-1213003331211020-2110101201322220-3200312310331033-3310111113220202-0123230121121121)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-3001031211030011-0221333021302221-1003202133002210-0232321213013013-3301121131312021-1232322320022301-3223301333130020-2012031010020311)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-1232032001201103-0301130032030300-0213032101332030-2322130001013123-0011231023300130-2121203312022020-0100132033212132-3302301023212321)
- [endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none](data-sources--bot_endpoint_policy--reference--group-014.md#canonical-0322032111203220-0033211232133113-3333211330322212-0301320300202121-2323223103332101-1310001001331011-3233212133033011-3201022101321023)
- endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2

<a id="canonical-3012220123233311-0031311222312233-1331020133133132-3001302101013113-1002223020130331-0202221233000001-3131110233101103-1312003221322210"></a>

Type: `"list"`. Computed.

Response Header Matcher(s). Response Header Matchers.

<a id="canonical-0132000112212100-3231030133322222-2133021122231233-1301311331220201-0312120011100010-2113311113230230-1103311030213322-3333212032122100"></a>

### Direct properties for `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2`

<a id="canonical-3022112231120030-2321332023012303-1001133030110310-2212202130130021-0122313132312020-2101030120301302-1301100001022033-2332122013233132"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2.case_sensitive` property

Type: `"bool"`. Computed.

Configuration parameter for case sensitive.

<a id="canonical-1130313321302003-3212123131322332-1303321022311321-0301000331012312-1032022112220232-2133002121102300-3110010112000221-1010203011101002"></a>

<a id="canonical-0220333212132221-1203312012312330-1121233002311132-1132202311021131-0031321301112123-0203312222320011-3132022003032320-2323322323203223"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2.not` property

Type: `"bool"`. Computed.

Not(!). Configuration parameter for not

<a id="canonical-1323310313302231-2222311102332323-0010332330323202-3301223313300100-2000302032123113-1201213331223113-0311123023130112-2001113322222200"></a>

<a id="canonical-0133321122210032-3230012333232113-2120003303311231-1322332021012000-3103331030003002-2332013023312211-1211110111122121-2122300101021022"></a>

#### `endpoint_policy_content.protected_mobile_endpoints.protected_mobile_endpoints.transaction_result_criteria.transaction_result_success.response_header_v2.response_header_and.response_header_operator.header.header_none.response_header_match_v2.operator` property

Type: `"string"`. Computed.

\[Enum:
RESPONSE\_OPERATOR\_EQUALS\_TO|RESPONSE\_OPERATOR\_CONTAINS|RESPONSE\_OPERATOR\_STARTS\_WITH|RESPONSE\_OPERATOR\_ENDS\_WITH\]
&#8203;- RESPONSE\_OPERATOR\_EQUALS\_TO: EQUALS\_TO value - RESPONSE\_OPERATOR\_CONTAINS: CONTAINS value -
RESPONSE\_OPERATOR\_STARTS\_WITH: STARTS\_WITH value - RESPONSE\_OPERATOR\_ENDS\_WITH: ENDS\_WITH
value. Possible values are \`RESPONSE\_OPERATOR\_EQUALS\_TO\`, \`RESPONSE\_OPERATOR\_CONTAINS\`,
\`RESPONSE\_OPERATOR\_STARTS\_WITH\`, \`RESPONSE\_OPERATOR\_ENDS\_WITH\`. Defaults to
\`RESPONSE\_OPERATOR\_EQUALS\_TO\`.

<a id="canonical-1321322032332332-2302233033002302-0331011201323021-1202322213300010-1231210002210300-0131121331203121-3311101030030230-0101310221221201"></a>
