---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-1120302011120211-1320332311003030-2133222101332322-3112330130332133-1032101320333221-0230123300313012-1030223132233010-3310113113220203"></a>

## Next pages — validation_all_spec_endpoints / 011222013030 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303332103220211-0001221101013123-1223333133303030-2032212000101010-2130221012021123-2103013232033120-1132200130001002-2100131123330013"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode — fall_through_mode / 303001023030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.fall_through_mode

<a id="canonical-2020203003301001-1222321311112110-3222210120002030-0320221003230333-2320233212321113-2130312210101232-1100102212232010-0102201032010332"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-3133323301012333-3101333132213130-1022321103010211-2011333311233220-1330121213112102-1003223000130203-1210222113233123-2331011301320012"></a>

## Direct properties — fall_through_mode / 303001023030 / 3

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222323213331011-3331012010011322-3023222213303310-1232001033102120-3221113033113323-3331322111100332-1112331012300322-0102132302020030): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023): complete subsection reference.

<a id="canonical-0100010330310202-0222323010320102-2311223333102302-0303003212212330-0020012130320001-1332312131323320-0310200303333002-1203130230131212"></a>

## Next pages — fall_through_mode / 303001023030 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-0222323213331011-3331012010011322-3023222213303310-1232001033102120-3221113033113323-3331322111100332-1112331012300322-0102132302020030)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0222323213331011-3331012010011322-3023222213303310-1232001033102120-3221113033113323-3331322111100332-1112331012300322-0102132302020030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110122202102223-1102213103301213-1222301130032132-3012301122111112-3321110201233122-3332032313333223-1303100013233320-1103012310211222"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow — fall_through_mode_allow / 333220130311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_allow

<a id="canonical-2203333311013003-0213110002311213-0031220323203321-2323321031122121-3210033201220303-0203312301320103-3111333320320111-2311012032011002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-0011002100120332-3000200131113102-3222031233003301-3030001032300100-2230302030122301-3232313202013301-0123211220213031-3233102011112233"></a>

## Direct properties — fall_through_mode_allow / 333220130311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012023333123010-1103013133132212-3201213022022230-3023300121002102-1010231210000310-1030312322312301-3212201032101131-1310220220223110"></a>

## Next pages — fall_through_mode_allow / 333220130311 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323313202320233-3213221121323100-3232330302333112-1210010302312102-1330201010103033-1303322221311231-3201310130101232-2222121032333003"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom — fall_through_mode_custom / 102223112321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom

<a id="canonical-1222013132233021-0211110132133301-2211202013130120-1012200102230330-0113201000032010-3012331303211203-2321113133030212-2200112001221012"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0030312021223220-0322132013221021-1322310232011110-0202301023011313-1331110130332233-2110032301223233-1113210133230022-3010230200303333"></a>

## Direct properties — fall_through_mode_custom / 102223112321 / 3

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222): complete subsection reference.

<a id="canonical-0022230302123103-0133332221222111-0323011010213311-2022112031100300-0123201313112201-0111122203231011-0302331223102313-0023200331231323"></a>

## Next pages — fall_through_mode_custom / 102223112321 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133200010130301-1112022123121000-2230100322313122-2323022213032211-0123002121120001-1131213313212210-2310131231231221-0201120201220011"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — open_api_validation_rules / 102310323001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0221321132110111-0233131232133021-3203013212312302-0302100021230212-1130001332313032-0001210332210223-1211000010021020-1303302231103210"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0010220103220202-1222001202312023-2133032203032021-2233020200310201-1203032033013203-3303313212033121-1012002132322112-1133221121230311"></a>

## Direct properties — open_api_validation_rules / 102310323001 / 3

- [action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-1111001011322323-1010220020321130-2203212102130333-1211333023132120-0203003001300331-3123001222320311-3210103012113302-3122220200033110): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-2211103003120222-2122202132331112-3031300120010000-1302113112000323-1202201123031030-0131112223123102-0310122031103002-1303122220133032): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-2200122231211020-0001023322111131-1011333211022033-3230202211313211-2003233200201223-2231111022311232-1202331222001010-0123331302201202): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213): complete subsection reference.

<a id="canonical-2002302031133210-0211100220211212-1033030032100322-0123322313203312-2132201323033011-0223311221332023-1013023001122310-1022233212312023"></a>

<a id="canonical-0110311210331002-0330230303201111-0322230101102320-3301301233203121-1203312303110031-0102200131100331-3113131111330022-0203103210101301"></a>

## api_group property — open_api_validation_rules / 102310323001 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-2323303311132223-2331001321333203-2311132012331000-2130121031121303-1330231020213331-1130110123030231-3000103230030231-3310002110212122"></a>

<a id="canonical-2202210203231321-0220101110202001-3323311201101331-0231113123133123-0322031012321100-2021331021100313-1122322123020123-1002311233320000"></a>

## base_path property — open_api_validation_rules / 102310323001 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-3020032022020013-0023320022000021-2001032112013230-3013123000022110-0301133312020233-2211132210021022-0000030000302100-2113132002130101): complete subsection reference.

<a id="canonical-2030021201202203-3022013002022023-0132013013033003-1312030002320120-2113230032213321-3312203011021311-2322130032113202-3031301301221012"></a>

## Next pages — open_api_validation_rules / 102310323001 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-1111001011322323-1010220020321130-2203212102130333-1211333023132120-0203003001300331-3123001222320311-3210103012113302-3122220200033110)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-2211103003120222-2122202132331112-3031300120010000-1302113112000323-1202201123031030-0131112223123102-0310122031103002-1303122220133032)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-2200122231211020-0001023322111131-1011333211022033-3230202211313211-2003233200201223-2231111022311232-1202331222001010-0123331302201202)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-3020032022020013-0023320022000021-2001032112013230-3013123000022110-0301133312020233-2211132210021022-0000030000302100-2113132002130101)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1111001011322323-1010220020321130-2203212102130333-1211333023132120-0203003001300331-3123001222320311-3210103012113302-3122220200033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111120021031303-2303033331201200-2012000122010312-2033021103101001-2003220132100003-0020220032331213-3212330012031302-0111303230330230"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — action_block / 230302102030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-0202220202021311-3313312322000000-1132221113111002-2103210312011020-2022102302232020-3300311120220113-0332011031201332-2132300213303102"></a>

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

<a id="canonical-2113210221000003-3113321102200130-3302132112112330-3313012123100001-3233310223032022-3020320120223202-1002221023033133-2302112110001132"></a>

## Direct properties — action_block / 230302102030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321101011001021-3332033233333311-1321122323302231-0000212302132030-0010323031332310-3301213100033010-1321023232303310-2333310222210130"></a>

## Next pages — action_block / 230302102030 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2211103003120222-2122202132331112-3031300120010000-1302113112000323-1202201123031030-0131112223123102-0310122031103002-1303122220133032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223132003302130-3310010221012331-2030132321132002-0200202230320133-1112112001303202-0033313120202323-3101200222032210-2330210200000310"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — action_report / 330123203232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-1310112100131121-0321220210022000-1012330231100123-2211101101130233-0301031323221322-2313313111210222-2110002012211113-1311330000312312"></a>

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

<a id="canonical-2022000331230310-3311211031200131-3110130333022023-2011113200111220-3213000223032302-2001013202120310-2201011132132133-3112133332102302"></a>

## Direct properties — action_report / 330123203232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110020031230210-1111032200110331-2300302322300211-1210010312232223-2320200010033200-0022100112002230-1122231300023321-0311103311312003"></a>

## Next pages — action_report / 330123203232 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2200122231211020-0001023322111131-1011333211022033-3230202211313211-2003233200201223-2231111022311232-1202331222001010-0123331302201202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201312303200020-1112132110230021-0233131030331032-3301112232302012-3122323203200101-2003022212031211-1033020303213112-2031031023111303"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — action_skip / 033300031231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-2031222113203133-0221013323103201-0033130300020022-1311301010131000-3323002123313302-2102131110331000-1012033123032102-3031013001201133"></a>

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

<a id="canonical-0220321100320213-0232002223103122-1323122221201031-3033030011133033-3033002212030233-0222113132003112-1113323223222320-2031002200111023"></a>

## Direct properties — action_skip / 033300031231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203230333102203-2000032113220122-3202221100332310-1031122321331120-2000002223213121-0003013103002010-1210320200113323-0232313133012102"></a>

## Next pages — action_skip / 033300031231 / 4

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0012223332023113-1301002030011201-3211230101223131-0303322232110312-2323130313210031-3123013220200023-0323220203200223-2333333023201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111321103311103-3310032103213310-2233133022002220-2013221113300310-2300021121122110-2001302032220332-3302332000333031-0221311131202223"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_endpoint / 222120230213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-2000200310202003-2323022131000210-0112312312122320-1221221101313212-2301111203202011-2230130301222230-3020102120110220-2322130313333102"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2330200212310300-1323223212303332-3132211110230110-2223211110313110-0002031310120323-2101312001230022-2013023020310213-1133300113233102"></a>

## Direct properties — api_endpoint / 222120230213 / 3

<a id="canonical-0012102010033303-1300312320330112-1200231301230202-3121212130200103-1223031203131113-3021030323010013-1020030112030110-0311122333330332"></a>

<a id="canonical-2320023332101330-3012002210132003-3112033122200211-0301100002312110-2332013211022131-0122310202210211-0012113212102101-1212103223022020"></a>

## methods property — api_endpoint / 222120230213 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1102001330211100-2320102222003221-0232332110202133-2310230033303100-0130332300103100-2132022032202123-1220002330210232-1201332232323233"></a>

<a id="canonical-2122103203333232-2032302323322113-2012031203312302-3120030213013303-0200200201120301-3300023002312300-2311303000032301-1102230231021231"></a>

## path property — api_endpoint / 222120230213 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0032100012001331-3012320123030330-0110320201003120-2311221103000010-3032110220121331-3321031330230220-3211230320123000-2121321212130232"></a>

## Next pages — api_endpoint / 222120230213 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3020032022020013-0023320022000021-2001032112013230-3013123000022110-0301133312020233-2211132210021022-0000030000302100-2113132002130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012023230010231-1333001032200221-3013121000001322-1121100130022221-3032300312100330-1100210202101003-2331302332232301-3133322131220103"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — metadata / 322213000122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-0230222100032102-3023101100200313-3203333130032202-2000021113101010-3212103100201012-1010120321113313-1311031012132312-1320013131320133)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-2202211132302310-1313322223011013-0103021332331333-2330010033320330-2313232011113213-3032002302222302-0331210310102210-2323330332002023)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-0021302312030332-0303201132233001-3230312010323201-3110203320230233-3300110333031120-2132000002312131-2020313030320032-1313211230132131"></a>

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

<a id="canonical-2031312120333230-0222002223322133-2311200320111311-3003200231210303-1333020001211002-1030133320223002-3132320013101211-2031122010031332"></a>

## Direct properties — metadata / 322213000122 / 3

<a id="canonical-2310110312203300-2330001223032212-0230212100032002-0120121220033030-0131130320223310-3122121100310000-2331132113303121-0231120021231321"></a>

<a id="canonical-0323211310123112-1111212211310200-1201221113020222-3323132030312203-3200313332122111-0200130222121032-3010113333022022-0131121303110111"></a>

## description_spec property — metadata / 322213000122 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0212013303002210-3200020233333023-2330031031113023-2101131202001102-3231231122101031-0200200011330312-2223213321313031-1312111201110132"></a>

<a id="canonical-0210220311111211-1321333300212212-0120111000201200-2023231321103103-2132101000022013-3333213333300231-2033231110011202-2231222330010011"></a>

## name property — metadata / 322213000122 / 5

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

<a id="canonical-0232110132313101-0032130233221210-3111103300122121-1312200122023030-2031021131100113-3130120321232133-1323232311233313-0011120121231202"></a>

## Next pages — metadata / 322213000122 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-2220110120022313-2010003222013113-0103032211212231-1002123202123032-1302012013021303-3132102133203300-0303233312110111-1120301230301222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330233000003022-0220200000212130-0013233001301022-2033123003212000-1323222000312223-1003313313322132-1331232012302002-3322231212220112"></a>

## api_specification.validation_all_spec_endpoints.settings — settings / 111311113331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-0333213303321311-0313303122100312-3320131121213203-1130020332003202-0111223011101032-1312321122020213-2230110300011032-2113022120000010"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-2220101032100021-0333232121212110-0212321120201210-1302000211202021-0133123011122302-2103302301311200-0130032303032020-2323203123221012"></a>

## Direct properties — settings / 111311113331 / 3

- [oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-3113232302202210-3022000123021322-1231130111032320-1011021230232011-3200213202302112-1323010132133111-0201032010333213-1120210212002131): complete subsection reference.

- [oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1013232123013112-1211132010230200-1000112320320002-1100331013131201-0221220102322102-1002203213221302-3001313030000221-2301312000201220): complete subsection reference.

- [property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012): complete subsection reference.

- [property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-3220123201332000-0331112323303301-2101330223310203-0302311120123132-2333320103213120-2120010031023120-0103202000232213-1332331112110303): complete subsection reference.

<a id="canonical-2113230000110011-0313203102201021-0230230330131120-0320330321102302-2220301101002303-3103023220022230-3320130321021123-1030113001132002"></a>

## Next pages — settings / 111311113331 / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-3113232302202210-3022000123021322-1231130111032320-1011021230232011-3200213202302112-1323010132133111-0201032010333213-1120210212002131)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1013232123013112-1211132010230200-1000112320320002-1100331013131201-0221220102322102-1002203213221302-3001313030000221-2301312000201220)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](data-sources--http_loadbalancer--reference--group-009.md#canonical-3220123201332000-0331112323303301-2101330223310203-0302311120123132-2333320103213120-2120010031023120-0103202000232213-1332331112110303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3113232302202210-3022000123021322-1231130111032320-1011021230232011-3200213202302112-1323010132133111-0201032010333213-1120210212002131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221101300221323-0011233312120001-3103122303030320-1100000110002020-1032131000231311-2332133130022231-0130002210220023-3202103023000031"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — oversized_body_fail_validation / 233300203303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-1110122223200132-0133023212132023-3202200303101103-2212103302001011-3133310122322132-2033110230323023-2132101312030320-0023320232102313"></a>

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

<a id="canonical-3001233003332322-3121011021113121-3230121030033203-0002300233220301-2013020213103011-2311001303132002-0130221212102302-2201233132233013"></a>

## Direct properties — oversized_body_fail_validation / 233300203303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311330333220320-0330300212313031-0030021132100233-3313202302031232-0022222010312300-2121032112322200-1213222213122210-3301033122313223"></a>

## Next pages — oversized_body_fail_validation / 233300203303 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1013232123013112-1211132010230200-1000112320320002-1100331013131201-0221220102322102-1002203213221302-3001313030000221-2301312000201220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310133102321222-3331221001322013-2230333302202210-1331020223300311-2012121010130110-2131021302200313-3203022101230210-2320331100103101"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — oversized_body_skip_validation / 110032003331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-0011210313111112-0202133011013010-0000102202112200-0303300310213203-0303332132232231-0131011000323303-1302101232222321-0101101223120023"></a>

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

<a id="canonical-0023112322232131-3220032211031223-0230231211101120-3213131013322010-3133203213301020-3033201131312301-1033211100122323-1001133322230120"></a>

## Direct properties — oversized_body_skip_validation / 110032003331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331202223001222-0303122103032121-3333233111331332-1200101030120302-3323033331323013-2303203001333231-0003323301132020-2033321231111001"></a>

## Next pages — oversized_body_skip_validation / 110032003331 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023110133103232-2032031213032200-2012303121023310-0122232031231221-3002102303133132-0321020333101213-2131223112132301-2101210120223222"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — property_validation_settings_custom / 132300200003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-0001011221301303-2110113101100112-1330311223232331-3120233232020232-0100320001323300-2332313003103310-0231101112303002-2321320210202333"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3023331302232202-1202023100223031-2022030312221132-3212233302102303-3133232103221333-1013102211202321-0331001013321101-2321133111123131"></a>

## Direct properties — property_validation_settings_custom / 132300200003 / 3

- [query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330): complete subsection reference.

<a id="canonical-0232300030032013-3130121030102212-2212020202210001-0202311303321231-1232022123033331-1220032303211012-0223223020110012-1110023331302013"></a>

## Next pages — property_validation_settings_custom / 132300200003 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303010110001203-1131013221003320-0303101121001112-0201201303131213-0000322322232303-1101020012332030-0232310231031300-1300312220103202"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — query_parameters / 320121023011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-2323332022220100-0331232002302100-1110122130222310-2002320100233221-0201003330213310-1123320133231112-0001222303301322-1332122222113220"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-3120122212131221-0300233213110022-0220013001332312-3331103313001232-1301011112331201-2322003113032020-2233311010033132-1332102121302330"></a>

## Direct properties — query_parameters / 320121023011 / 3

- [allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-1333210300113120-0122203221120030-3311320212031000-2231122010213203-0323211101012030-0102110001203211-1333332122221230-1233323301120303): complete subsection reference.

- [disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-0330321311010311-3133200223221032-1331111221033122-3102223010112303-2120123010133320-0121132231021010-0002230330113011-0323230113232020): complete subsection reference.

<a id="canonical-3311121001112313-1320020221221220-3310232230133232-1120331103202132-0210320120110313-2312231322211233-1132320020120312-2333202320323123"></a>

## Next pages — query_parameters / 320121023011 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-1333210300113120-0122203221120030-3311320212031000-2231122010213203-0323211101012030-0102110001203211-1333332122221230-1233323301120303)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-0330321311010311-3133200223221032-1331111221033122-3102223010112303-2120123010133320-0121132231021010-0002230330113011-0323230113232020)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1333210300113120-0122203221120030-3311320212031000-2231122010213203-0323211101012030-0102110001203211-1333332122221230-1233323301120303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130012121012201-1122233103220110-2201112133023323-1323300221323022-0011033131221212-3223103032130013-3002113200303113-3131131322111132"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — allow_additional_parameters / 133332300123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-2231001101313222-3033100113331011-2232102101131120-2022013213112300-1003000311311302-1002000233312233-2202313223002013-3333202121011010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

<a id="canonical-1122010302032211-3302130202331313-1310330303001121-3210332130031312-2210321133112013-3031312022232320-0110030331200031-2132322012110001"></a>

## Direct properties — allow_additional_parameters / 133332300123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133100132011010-2123332123301233-3113023112310010-1212301031101103-1110012301312012-1010110133230220-3013030112110011-3301220003302212"></a>

## Next pages — allow_additional_parameters / 133332300123 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0330321311010311-3133200223221032-1331111221033122-3102223010112303-2120123010133320-0121132231021010-0002230330113011-0323230113232020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122321031232332-0003310320233220-1120120030201303-2201303011330223-1112222320320311-3111032013331320-1011022330211032-1222010011202031"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — disallow_additional_parameters / 020112232313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0021130223123320-0323003003013222-1200302203020303-3010121231033121-1002300321002331-2031002112220101-3333310102013320-0102113202032012)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-3202032012013033-0001122011332023-1001332231300030-0031113232113133-3000033211323013-3031313220000213-3223323022330011-3223323023003000"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

<a id="canonical-3333213231113330-1323001222321131-0121303331202123-2030133330121221-2213123123120121-1022013121202031-3332203232302000-3301200323110030"></a>

## Direct properties — disallow_additional_parameters / 020112232313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231003331201212-0312331002031002-0110123030222003-0020213131012023-3202130322022302-1111033020020111-3120032203302320-2131132031300123"></a>

## Next pages — disallow_additional_parameters / 020112232313 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--http_loadbalancer--reference--group-009.md#canonical-2031321030012203-0311201200131332-1020330023013102-2103023020001223-2302000032100230-0302212311103233-3330102010223030-2330323002332330)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3220123201332000-0331112323303301-2101330223310203-0302311120123132-2333320103213120-2120010031023120-0103202000232213-1332331112110303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321132210221031-2122332021233302-3030000301130132-0022223333123313-2223310121000000-2201303120223322-3031132302101330-2002001132333302"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — property_validation_settings_default / 301323202331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-2031230212122022-0101121101230231-3130132032332111-3212132000002021-2020101310023223-3132033201011202-0122213133130130-0222110120102102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-3021100133302113-1211132223313211-0032220331121020-0012122101132310-3220133021011210-1213132111220110-2201101333122232-2102013323320111"></a>

## Direct properties — property_validation_settings_default / 301323202331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112130320310123-2111100310200311-2321123230112122-0323220013002212-2133311333033301-2010310102011223-1311123331323110-1321100333023110"></a>

## Next pages — property_validation_settings_default / 301323202331 / 4

- [api_specification.validation_all_spec_endpoints.settings](data-sources--http_loadbalancer--reference--group-009.md#canonical-2011221130012223-1333201113103310-2100022022010202-3212231300332011-3322100101003030-3121030013312030-3311031000331113-1320313320033033)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223232121213220-0130312121011221-3003202122033313-2122201002302133-3322202103110023-1111120033212012-3210321111222332-1011032100033021"></a>

## api_specification.validation_all_spec_endpoints.validation_mode — validation_mode / 133213321133 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-2332132223223231-2223132301000320-3333121011113200-3231330300110130-1210013120122321-2310122202110312-2013211011032033-1332233133220032"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-1020302130113222-3100111220212321-2032302230301113-1310233031111303-2331333211132111-1030133022123102-3103211022110012-3212323212032200"></a>

## Direct properties — validation_mode / 133213321133 / 3

- [response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203): complete subsection reference.

- [skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132203122300201-2033033330013231-0011311333222001-3332223323003210-0223010111120230-0301121301032132-2233203322301220-2002213320322312): complete subsection reference.

- [skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-0301113120020203-0120232321323230-0300313133221220-2222202200111002-2030333130000033-3030031211203331-0003112022013001-0003302203031320): complete subsection reference.

- [validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131): complete subsection reference.

<a id="canonical-2221321102231031-2000330131211033-0012002223210033-2033310131332021-3300323230020103-3131232302222301-3223203213101121-1222220322021223"></a>

## Next pages — validation_mode / 133213321133 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132203122300201-2033033330013231-0011311333222001-3332223323003210-0223010111120230-0301121301032132-2233203322301220-2002213320322312)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-0301113120020203-0120232321323230-0300313133221220-2222202200111002-2030333130000033-3030031211203331-0003112022013001-0003302203031320)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212123013123233-2231333002003112-0201112303111303-0310020213100223-0233231020123320-3202112212023302-3201211002312203-1323213000103003"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active — response_validation_mode_active / 133200300300 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-2212232323103202-3313303212012333-0131311213231112-0012132220301231-3322322020210132-3121031023101201-0333120211302223-0210132322100310"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-3321303231030113-1031230013030002-0213101032003310-2231301323021330-1023332201113103-2311123321000031-0310200303022302-0311003311320312"></a>

## Direct properties — response_validation_mode_active / 133200300300 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3133312121132100-2220000100320020-2013322310103113-0230201130123330-1230221221013112-1312233012312312-1300132302112031-2131003021220133): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-0332012230010002-3213030320323032-1331200221132313-2332212311221301-0321010222031121-0002202320313002-0332131222103113-1320030200031221): complete subsection reference.

<a id="canonical-2300132123102120-0030321203031231-3011330332333223-0030032303221011-3231103320103322-3300302113233200-0032100101022322-3310010322330002"></a>

<a id="canonical-0213021333033200-3220013330020032-3300111002103113-3301330323131300-1021120222003120-1200023012323231-1032111201211011-3211311332022332"></a>

## response_validation_properties property — response_validation_mode_active / 133200300300 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202122012302121-0321001300233330-0233112122113320-3320023131222003-2131120030131233-1131221110110122-1012330223101320-2131000123303213"></a>

## Next pages — response_validation_mode_active / 133200300300 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3133312121132100-2220000100320020-2013322310103113-0230201130123330-1230221221013112-1312233012312312-1300132302112031-2131003021220133)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-0332012230010002-3213030320323032-1331200221132313-2332212311221301-0321010222031121-0002202320313002-0332131222103113-1320030200031221)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3133312121132100-2220000100320020-2013322310103113-0230201130123330-1230221221013112-1312233012312312-1300132302112031-2131003021220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320321010310231-3103200331131222-2231012123313122-2003233100331232-0001120301200031-2013121232333323-1313333110033310-0130130023112303"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block — enforcement_block / 031030112113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-1213201023312212-2112132203123330-0313323321122123-2000330120030202-0013333222132331-3313133213212302-3102301010132331-2013313303311032"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3132221033023301-3231301030333003-1210202201232231-1203330313333033-3032223112321201-1213222022031111-1103312202302323-0200130020323322"></a>

## Direct properties — enforcement_block / 031030112113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322320311000313-1332113232132112-3210210233220213-0301313111230300-2000231313010102-3011333133013221-0103030030012330-2223321112132030"></a>

## Next pages — enforcement_block / 031030112113 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0332012230010002-3213030320323032-1331200221132313-2332212311221301-0321010222031121-0002202320313002-0332131222103113-1320030200031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330330203122121-0023120110133330-2301123310122121-3202322301200223-1310331323232322-1032232130233213-1323013032223031-1033213300210301"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report — enforcement_report / 020211010223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-0221012213111000-3212031303322213-1203321233333201-3323200333203010-1313013001010221-2313033113322013-2320300302313131-2003230220020203"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3001132103302210-2332301123322302-2122013010320031-1313211302332013-3102311101132321-3200312111201113-3131110100110203-2032001303120233"></a>

## Direct properties — enforcement_report / 020211010223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311020130023131-3203302213222022-0321223220123013-1113010322101123-1220101132321311-1303233233311003-3120100011232230-2013001311323222"></a>

## Next pages — enforcement_report / 020211010223 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2113031011323331-3212210113000210-3333311022023330-3332212101322302-2220002302303110-3323202311113323-0232011200212010-1201310003021203)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3132203122300201-2033033330013231-0011311333222001-3332223323003210-0223010111120230-0301121301032132-2233203322301220-2002213320322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000001312211011-1330331312313102-2313310002223010-2303101321001201-0320111021322331-1003130101002321-3123013101010331-3212100022033200"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation — skip_response_validation / 133022110031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-1321022011123031-1033321031211332-3003121112223222-0333223212030313-0030022232331331-3020002323223202-2310013332313323-3011113030100222"></a>

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

<a id="canonical-0022221331130003-3332201312101031-2203032203202201-2020121012112131-0003000111332022-0201102112323233-1332011100101113-2221032323133331"></a>

## Direct properties — skip_response_validation / 133022110031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302312333020002-3132133033210130-0233113331111113-2200030031100002-0323000130111303-1101113300133301-2231120203031311-0112312232002211"></a>

## Next pages — skip_response_validation / 133022110031 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0301113120020203-0120232321323230-0300313133221220-2222202200111002-2030333130000033-3030031211203331-0003112022013001-0003302203031320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313311200331123-1011220233023232-0123101320120203-3332012021122132-0222230020211302-0011110010203102-3112333123310032-0033301013022002"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_validation — skip_validation / 313110311120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-3322231322213301-1020321311302223-1022022113100221-3120220123222231-0103221031132233-0020112322213130-2020331123210001-1233012221300312"></a>

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

<a id="canonical-1313100313232302-2201302013301013-2132221012030101-0133210331123001-1123023133110223-2011001021323320-3202122222013331-3102032210133210"></a>

## Direct properties — skip_validation / 313110311120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202311121331032-2121233011333300-0332221211010033-1013202332030330-1202212303123121-3323101030230302-3103210322122021-3300121230211201"></a>

## Next pages — skip_validation / 313110311120 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220011033321012-1331332230021002-2121023301111221-0222313122012012-3321112103032013-0122333123321112-1301021001323100-0132322121123021"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active — validation_mode_active / 131110223331 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-1222230200100133-0230332211023322-0300213031011211-2123011232112112-3112302112100121-2110010300220102-1122031101033302-3210022300031021"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-1221013302200333-0202330231200130-3213023312223112-0310312332023020-0101133212003330-0331021321120021-3302233132200200-0012321020003121"></a>

## Direct properties — validation_mode_active / 131110223331 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3000232033312222-3220022322231321-2212102002220222-3211201111222322-0133233300232002-2133333311312331-1122013032102123-2210312021023001): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3331103210000311-1011321023313011-0132201220120311-2010002300033130-2302121022123222-0202132331000202-2320201210220130-0213023012231213): complete subsection reference.

<a id="canonical-0013301023023212-3133232032102021-0100333011302200-3102032021133132-2211000130030330-0331131012121131-1032232213232032-2101233023310133"></a>

<a id="canonical-0233230311332300-3131133331300130-1123201320222103-1020232331223312-3201222110130323-3213123300003322-3322220332223330-3322030100330302"></a>

## request_validation_properties property — validation_mode_active / 131110223331 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0102203002330120-0331232131100020-3220302312012103-2332120313022123-1311233133003113-0023212003130103-0020033221130212-0031012023210202"></a>

## Next pages — validation_mode_active / 131110223331 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3000232033312222-3220022322231321-2212102002220222-3211201111222322-0133233300232002-2133333311312331-1122013032102123-2210312021023001)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3331103210000311-1011321023313011-0132201220120311-2010002300033130-2302121022123222-0202132331000202-2320201210220130-0213023012231213)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3000232033312222-3220022322231321-2212102002220222-3211201111222322-0133233300232002-2133333311312331-1122013032102123-2210312021023001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033002200122233-2100322221212202-0301130100021023-2232222122010320-0111101011323020-0030032111001323-3313102031321331-1313033211121123"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block — enforcement_block / 302132113202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-0033033132310010-0033103331310012-0023232021330133-3223322133323211-3112220222111032-2022103202310301-0321030212330130-1130020002021213"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1201202123000111-2001233000031030-1223313003332112-0222033231322110-0330103321130112-3013002311132202-1230233232330010-2112200012200002"></a>

## Direct properties — enforcement_block / 302132113202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111222132231020-1031122000231202-0330102132313111-2210122011110310-3212200220330001-0332332002010030-0000023323233322-1222232331121303"></a>

## Next pages — enforcement_block / 302132113202 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3331103210000311-1011321023313011-0132201220120311-2010002300033130-2302121022123222-0202132331000202-2320201210220130-0213023012231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101332110233003-0111311003033313-2110332222303112-1011202012310123-2211203130321202-2210033011120202-1321200123213310-2102102031011332"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report — enforcement_report / 013232333320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_all_spec_endpoints](data-sources--http_loadbalancer--reference--group-008.md#canonical-0001013203130012-1332302230013000-1300020131332301-0211102220001303-1212202231222233-2321322223223303-3130120323321213-2010013230023321)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1133202312201202-3123311122032101-3131021112212112-1230131313011331-2101330211131231-1023013332200022-3012303330322101-2320330022221030)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-1132310122233223-1300323121301100-2020102022330301-0311110112113010-0212220301303000-3232023202022012-1003130201120132-3300031210200330"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0022203011111212-2013030213012302-0010320200122320-3131201121232122-2322110221000332-3312013232200123-1222333201313333-1102020322000201"></a>

## Direct properties — enforcement_report / 013232333320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223012303201022-1002223131020120-0001132200002222-2203331111031033-1023113132130030-2321132033330331-2311133231120223-1011232131012233"></a>

## Next pages — enforcement_report / 013232333320 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2233200001301320-3303223031022002-0323313113022133-1312023121023002-1213312202233011-2321333132103132-2121223303133121-0121210010133131)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033023323212213-0231200213001302-3020311303112011-0222213301103131-2120130310111310-1213013012021311-1310021322000313-0123101133000001"></a>

## api_specification.validation_custom_list — validation_custom_list / 021321210030 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- api_specification.validation_custom_list

<a id="canonical-3101132312231321-3223323020302011-1221032111111101-2012333233023122-1333331130232323-3013122202333212-0102323010000300-0023000021312223"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-0123331211300231-3000111032333001-0302123123032211-2233021001133121-2300311221131212-2100333301220303-1203032021010312-1300000302101133"></a>

## Direct properties — validation_custom_list / 021321210030 / 3

- [fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111): complete subsection reference.

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231): complete subsection reference.

- [settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-1112200302330330-0330323032110031-3333202023000202-3101020301022323-2022212332013123-0100001232113333-1301210303031331-2322220133323302): complete subsection reference.

<a id="canonical-3001000121223113-0133030223000222-2011333220011001-1211211200020322-1021223313331132-0202300330210033-1031331122101123-1210023330123210"></a>

## Next pages — validation_custom_list / 021321210030 / 4

- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.settings](data-sources--http_loadbalancer--reference--group-010.md#canonical-1112200302330330-0330323032110031-3333202023000202-3101020301022323-2022212332013123-0100001232113333-1301210303031331-2322220133323302)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313220233202003-0010000112101233-2001122213301323-0103331133212131-3301331121223032-3200002022210122-1123121220321021-3220201110132233"></a>

## api_specification.validation_custom_list.fall_through_mode — fall_through_mode / 111313311310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-0331300023303112-1113002331201200-2233031330023223-0022213102332232-1033320010020010-3121112100303020-1310023303300020-1123022212103322"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-3220131301232103-2012233002211211-0233023020122221-2111231102123112-2103313203103030-2123233031322101-3231120322131203-0030131121022133"></a>

## Direct properties — fall_through_mode / 111313311310 / 3

- [fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-2021002022311320-3312333031031200-2230023210200122-3110221031110210-1132133211000120-2033323313231001-2022331211323130-2332030020002120): complete subsection reference.

- [fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233): complete subsection reference.

<a id="canonical-3131232311111233-3223320010300001-0131201232120001-1020103130130201-3233211203302202-1332000021012231-1223020010320220-1130123221200211"></a>

## Next pages — fall_through_mode / 111313311310 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](data-sources--http_loadbalancer--reference--group-009.md#canonical-2021002022311320-3312333031031200-2230023210200122-3110221031110210-1132133211000120-2033323313231001-2022331211323130-2332030020002120)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2021002022311320-3312333031031200-2230023210200122-3110221031110210-1132133211000120-2033323313231001-2022331211323130-2332030020002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033131210121300-1303022221022012-3032222013222300-3003212130022033-2320232002033110-1013311001001223-0231233331321131-0010113103022110"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow — fall_through_mode_allow / 032212203203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-2003031030233012-0313301022023321-3100000102330032-2013113223002303-0003033103200022-1020212201123100-1013201220002010-1323220133332011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-1101333202103033-2323030110231311-2010000002131113-1121100030121233-3101212231021232-2312010123223223-2000323000132020-2213202100132310"></a>

## Direct properties — fall_through_mode_allow / 032212203203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133311003232313-1311032110122203-3212213223001311-0033103333003000-0222011000331123-0333212130302223-3220000211000122-2111300333020011"></a>

## Next pages — fall_through_mode_allow / 032212203203 / 4

- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302001331311032-1330222102123332-3100012221012013-0322311322122232-0310203132211333-0002132020213201-3330112233000111-3122101321100010"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom — fall_through_mode_custom / 010332323223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-2320231033113033-0332333222102123-3222313000032232-1001213011022220-0010233233002123-0002013210133232-0011220113001322-2112200033122031"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000002330030311-2322022030003022-1323233221023032-2002132102003223-2220213012231203-2213221013133300-2211302012100302-1030303322201100"></a>

## Direct properties — fall_through_mode_custom / 010332323223 / 3

- [open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232): complete subsection reference.

<a id="canonical-3223322030332023-2110323011331023-2100031011113332-1032113120321313-3033322321311112-0310130333313330-1233213201331010-2010322033020213"></a>

## Next pages — fall_through_mode_custom / 010332323223 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303011210033032-0121020131221002-1322000233002212-3033012303033311-1020002320030222-1021120223212012-0102301131132223-3031301133210101"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — open_api_validation_rules / 331120233333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0330231301121223-0022222023331222-0001220202322223-1300232110332030-3131112302203100-2301001110213320-3300132023033202-3210122223000001"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-0232110211302301-2332323002100332-0030022113130233-2201232023220033-2212331111101131-3113001131130112-2222020212300321-0031120230212020"></a>

## Direct properties — open_api_validation_rules / 331120233333 / 3

- [action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-0111101003232101-3131100310020233-2010312231212110-2313002113010123-1333322033101130-3121111300312130-2322013132223311-3111133312323323): complete subsection reference.

- [action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-1110203031310300-1221132322033110-3232321323312231-0130233322323301-2121010221023323-3200012322000200-2213023201121021-1132202121220001): complete subsection reference.

- [action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-0333203022231331-3123012100031001-3222301122310110-2201220333210122-0121121233130032-3203331100330231-2000101033132133-1100022230213010): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-1221313211210230-1332320300101232-0201112130123232-1312013223130000-3102212012111222-1021232200203020-0131102013121221-3303113003212302): complete subsection reference.

<a id="canonical-3303201020220210-3020030211132101-1302232221133320-0221021300233011-1122002310123012-0111121213222121-1101133133211201-3131311033211100"></a>

<a id="canonical-3223232033321301-1321023023130111-1112200011322130-1331120001000222-0000001000221000-1000211101332223-1300111033101102-0000121013110123"></a>

## api_group property — open_api_validation_rules / 331120233333 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-1112120120013312-3332232321000213-0313120011233203-2303123320310222-3232303133100231-3101302333001232-0032200323103333-0310300321323212"></a>

<a id="canonical-0030213332323110-2300133301200203-1213002000312131-2110002330333313-3010211200311130-3002232230233232-1322320212012133-3102232223113120"></a>

## base_path property — open_api_validation_rules / 331120233333 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-3230303110032112-1233220022032312-0230100031022321-3220321200100002-3120330332220313-3032320112122231-1220303221310301-2113322030112220): complete subsection reference.

<a id="canonical-0212121023322133-1311321120112121-1312123112002200-2122311123100302-2212012011201233-0131223202121202-0122333302221212-1302223232222100"></a>

## Next pages — open_api_validation_rules / 331120233333 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-0111101003232101-3131100310020233-2010312231212110-2313002113010123-1333322033101130-3121111300312130-2322013132223311-3111133312323323)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-1110203031310300-1221132322033110-3232321323312231-0130233322323301-2121010221023323-3200012322000200-2213023201121021-1132202121220001)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](data-sources--http_loadbalancer--reference--group-009.md#canonical-0333203022231331-3123012100031001-3222301122310110-2201220333210122-0121121233130032-3203331100330231-2000101033132133-1100022230213010)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-1221313211210230-1332320300101232-0201112130123232-1312013223130000-3102212012111222-1021232200203020-0131102013121221-3303113003212302)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-3230303110032112-1233220022032312-0230100031022321-3220321200100002-3120330332220313-3032320112122231-1220303221310301-2113322030112220)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0111101003232101-3131100310020233-2010312231212110-2313002113010123-1333322033101130-3121111300312130-2322013132223311-3111133312323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232013121032300-1312102313200110-0001130121300222-2200021122023201-2021332020030302-3230313110011023-0120022322122212-3312303131002222"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — action_block / 331011201332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-3000232303110113-2321011002333013-0213332021003332-3120011000333100-1132122213132202-0120001221103123-0121230121133233-3032233322312111"></a>

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

<a id="canonical-2303132122110113-3123100103131102-0323112323232001-2120212330002200-3202013102123201-1021303102330321-0012223011200111-2202200103032133"></a>

## Direct properties — action_block / 331011201332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321120212310301-1112303213230032-3321203022011230-0020302201212023-3031331013103200-1311023010330231-3200031032031020-1313320222212112"></a>

## Next pages — action_block / 331011201332 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1110203031310300-1221132322033110-3232321323312231-0130233322323301-2121010221023323-3200012322000200-2213023201121021-1132202121220001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322021320101133-3213231132101233-1110322101120002-2032031202121310-2220310330121003-2002132220000113-0201111120132333-1033112130322202"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — action_report / 333321102120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-3021102031331012-0310213012333032-1002113230010230-2330232213313331-2233013320230303-1122201021130101-1103203002102011-1121202302212100"></a>

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

<a id="canonical-2320333020230310-3101023320033032-3331132221031233-3333110233300200-1123330200011102-0310200031012001-1322012120321230-3202210211310132"></a>

## Direct properties — action_report / 333321102120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012121012102103-3120022113021312-1230223110222002-1202120230323322-2200101103112110-3210313333122013-1210201123002321-0300010303223312"></a>

## Next pages — action_report / 333321102120 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0333203022231331-3123012100031001-3222301122310110-2201220333210122-0121121233130032-3203331100330231-2000101033132133-1100022230213010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200213100012213-0321223300331023-0210202003100333-1200133123301220-0302130321133011-0333103211232323-1010003220312203-2030321111022233"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — action_skip / 312332213210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-3223302331130102-3030210213111130-2000012121220013-1213011202003011-0332020001033123-0103033022110012-3323323313111131-0030101020301230"></a>

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

<a id="canonical-2001222303020330-2212000131333332-0001123330323110-0111321023112233-1220230023120001-2212002001313100-0001102011301330-3200013131321123"></a>

## Direct properties — action_skip / 312332213210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110123032332323-1023120003331311-0321010212203220-1303202122010022-1301322311320303-1322020020001130-1321333110303202-3020013133110000"></a>

## Next pages — action_skip / 312332213210 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1221313211210230-1332320300101232-0201112130123232-1312013223130000-3102212012111222-1021232200203020-0131102013121221-3303113003212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132223031223113-3033201210232132-1201113102202300-3010003012011011-1201112103203111-2312112221232121-1031230030211113-0222123102030013"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_endpoint / 033032012323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-0200002303313220-2131332132131311-0210300313232010-2033332010313322-0322302101033132-3310103103013310-3230032221231121-2311303211311321"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302120320011320-0203321011221220-2122023301000200-0021123011103212-1112131101332003-1203023213001331-3320103022111200-1303211022012310"></a>

## Direct properties — api_endpoint / 033032012323 / 3

<a id="canonical-3331312010023102-0111121023100133-2132120333111300-0110332032213120-3031012110232202-1110300233300113-0023232231131313-1031231031002313"></a>

<a id="canonical-3320013001113111-0212110103130013-2213202223123321-2321311201221101-0122223130121030-1303031331102311-2112320310110231-0310013213333101"></a>

## methods property — api_endpoint / 033032012323 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320230102111333-2131110111320203-1330032010012032-3103200131221101-1231131233133121-3033000233022010-3032311213021200-1022203012011011"></a>

<a id="canonical-3220311332022232-1113103022131221-3301122021210111-3212223130332222-3332033233002231-0203203312211210-3211312332303323-3023200300022322"></a>

## path property — api_endpoint / 033032012323 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1231223213031033-2212320220203302-2223033131201133-1200330221133222-1021330321230102-1223120131331033-0130220131131212-2313213102221033"></a>

## Next pages — api_endpoint / 033032012323 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3230303110032112-1233220022032312-0230100031022321-3220321200100002-3120330332220313-3032320112122231-1220303221310301-2113322030112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230122312323001-0010120002211113-2210230332112132-1330223110010012-2022201322213231-2032321030133000-0021300202211320-0113101021023020"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — metadata / 210321322011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-3132010312131133-0030101021312113-2023232301112001-3332133002113310-1033002033301331-1011333132022120-1333022201133213-1111001101130111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--http_loadbalancer--reference--group-009.md#canonical-0322000233010300-0013113010211213-1020233000302100-2031110322310112-3030120232101221-2222121111311102-0333220012213113-1111232013032233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-1001222021000000-0033130301003100-3311333100312232-1022203301123302-2100130331100130-1112031332031011-0320112322120300-1220211001022331"></a>

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

<a id="canonical-3230131101133100-0110110333312213-1030022111231300-0030131332201312-2103330111332033-0323110230133300-3003330002032312-2300021311101023"></a>

## Direct properties — metadata / 210321322011 / 3

<a id="canonical-3311001202333231-2212311323212122-1002232020333303-2201130031021123-0332210322102020-3322032111102312-3331110023323002-2133022113100001"></a>

<a id="canonical-3213102100031321-3330021303203102-3312331231112221-2122211310302101-1133032330221122-0301301122312013-2301303203101210-0301231133231212"></a>

## description_spec property — metadata / 210321322011 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1022313300131111-2311001120013001-0100222230020311-2130202130213123-3113311011203303-1311121022030100-2312020331013103-3133130113123131"></a>

<a id="canonical-2320123331231030-2323001320130121-0033210233223001-0203231202202000-2012331221132332-1321012030103301-0012102013330031-2131011200003231"></a>

## name property — metadata / 210321322011 / 5

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

<a id="canonical-0122133112131131-0100012133131203-3110130013300301-0223223030331313-3320301131332110-2031113133031113-2232021021332300-0121132131312123"></a>

## Next pages — metadata / 210321322011 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-1231010231221122-2012213003001310-1311230132021210-1113323200101111-3321332320003210-2133203023300220-2212200031302002-2102013123012232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311233320031213-0202233011320021-1312300021300011-2121333233213122-1133200333023033-2332121011333131-3132302210312210-3101032113233130"></a>

## api_specification.validation_custom_list.open_api_validation_rules — open_api_validation_rules / 021112302032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-1110001301230112-2123330332212031-0202113201310003-0323011323310302-2010122303221121-0000210102100230-2311112201330220-0302111021103321"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2200330211223032-0010323102220012-0133031111212230-0131000222310332-3231103002231001-0012002301313021-0211323101213232-2233111303003013"></a>

## Direct properties — open_api_validation_rules / 021112302032 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-009.md#canonical-1321210020020023-3201033300010220-3312223333032320-1133120322203211-3231112333001210-3321333112001111-3313112013022321-0330103323210130): complete subsection reference.

- [api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-3111131311321033-0323003321030103-1020210303022013-0222020003001322-3123032033301332-2302103003201310-3220030033323200-2303001110311313): complete subsection reference.

<a id="canonical-1333013313330123-2023011112120312-2120213111233003-1322122331211322-1030003021323323-3033223301100200-2312310203020020-0021112211101201"></a>

<a id="canonical-0333203202101110-3032223113210221-3302213303200221-3212011133213203-0332023032311333-2203231212110101-0131230132100302-0213323202331300"></a>

## api_group property — open_api_validation_rules / 021112302032 / 4

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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

<a id="canonical-2332311022313223-3010023300300202-1201122011030213-2020000103133312-1332302120010101-0332121122323300-2333211010200223-2003322302101022"></a>

<a id="canonical-2222213023310222-1221301131120203-0312000321220311-1020230300131232-3102120212211300-2212022201003003-1220310301021103-2113031303230213"></a>

## base_path property — open_api_validation_rules / 021112302032 / 5

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-2302033232300200-2232033301133101-3231212111213310-2311332120222103-0311333111002321-2300020210313230-1030031130311121-0311322301301301): complete subsection reference.

<a id="canonical-3310102013332001-0010320031100311-1223000310122102-1232321113000320-0031310002313011-3030030300212031-2301020312212223-2302300232010201"></a>

<a id="canonical-0212122330331303-3130003013320303-0331203330030230-0223220330020000-1113031032313031-0322130212212202-3131321013310032-2130100201013003"></a>

## specific_domain property — open_api_validation_rules / 021112302032 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130): complete subsection reference.

<a id="canonical-3012320220111201-2223131232303220-0312223322103220-1033320301312201-2331120203021331-1121020231002020-0333213320101020-2332210300000310"></a>

## Next pages — open_api_validation_rules / 021112302032 / 7

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](data-sources--http_loadbalancer--reference--group-009.md#canonical-1321210020020023-3201033300010220-3312223333032320-1133120322203211-3231112333001210-3321333112001111-3313112013022321-0330103323210130)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](data-sources--http_loadbalancer--reference--group-009.md#canonical-3111131311321033-0323003321030103-1020210303022013-0222020003001322-3123032033301332-2302103003201310-3220030033323200-2303001110311313)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](data-sources--http_loadbalancer--reference--group-009.md#canonical-2302033232300200-2232033301133101-3231212111213310-2311332120222103-0311333111002321-2300020210313230-1030031130311121-0311322301301301)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1321210020020023-3201033300010220-3312223333032320-1133120322203211-3231112333001210-3321333112001111-3313112013022321-0330103323210130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121222020233101-3222112033122303-3303032023202230-0133131211312313-1230222002001202-2131001303032331-2220020131223333-1011130100220311"></a>

## api_specification.validation_custom_list.open_api_validation_rules.any_domain — any_domain / 230331303221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-1200302021211333-0312023213200123-3202221230332302-1221303002030323-0123011012223333-0201220030330222-3031023233132123-0203312233210200"></a>

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

<a id="canonical-0301300110013301-0013122122112130-1332130302332231-3030330123333032-3022110230021003-3323133213122122-3213122322331133-0222212111122001"></a>

## Direct properties — any_domain / 230331303221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210333312333231-1200330320333000-0322223212021030-3200212313101131-0100230023300222-2230012321030313-3231330132300123-1310010200032201"></a>

## Next pages — any_domain / 230331303221 / 4

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3111131311321033-0323003321030103-1020210303022013-0222020003001322-3123032033301332-2302103003201310-3220030033323200-2303001110311313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310022101200222-3023203112003101-1300112133330121-2211330311120330-0022033110311331-2320001121300320-3112223313333330-2312130022333302"></a>

## api_specification.validation_custom_list.open_api_validation_rules.api_endpoint — api_endpoint / 110002213001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-0212110223100002-3332022130100310-0323313012120020-2121211223323103-1023001212122213-3011211112102031-1302001112232301-3220223033111300"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301000323022201-1113213122123212-1012220011111232-3232112202000323-2013332120021333-0330133331102313-0011000202320023-0303301013103231"></a>

## Direct properties — api_endpoint / 110002213001 / 3

<a id="canonical-0200002020133230-1032033232030302-0332213000330213-0101221200233100-0021300203012303-3003331301203312-3310023001331101-1331021110132033"></a>

<a id="canonical-2212033312000013-2113021321111233-3132133333201132-3030023200123212-3202323011131003-2133013230211333-2033121002233002-2110321321232231"></a>

## methods property — api_endpoint / 110002213001 / 4

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1301222000200000-1003221031312112-1131221032132311-3201322322110200-0301123230102312-1111112212121211-2102013200013122-2022301200222110"></a>

<a id="canonical-0310100103000131-3103131100312133-1310000023221112-0223201010123312-0223320112023100-3023301123333022-0001121120032122-2101100013103222"></a>

## path property — api_endpoint / 110002213001 / 5

Type: `"string"`. Computed.

Path. Path to be matched.

Upstream description:

Path to be matched.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0303332110321320-1130233311132021-3302220110331232-1333032333212213-2032130121020021-1121212201310312-3201131210003303-1031121300112313"></a>

## Next pages — api_endpoint / 110002213001 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2302033232300200-2232033301133101-3231212111213310-2311332120222103-0311333111002321-2300020210313230-1030031130311121-0311322301301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022230031232101-1021203120021320-1131012133031222-2100121103333331-2122232221120222-3203212010210320-2132032321332230-3330130030223011"></a>

## api_specification.validation_custom_list.open_api_validation_rules.metadata — metadata / 220021230310 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-2321231322330032-3210030232000101-3103032100030113-1132100322010332-2221120110100110-0003111121031231-3220011320310103-3103232023202212"></a>

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

<a id="canonical-2222331323312123-0310230310231011-0220023212331110-0223321130030113-3321000213221333-2213230210012012-0302101131201001-3311313312313313"></a>

## Direct properties — metadata / 220021230310 / 3

<a id="canonical-1232323331212233-0011330212313200-1002012322031203-1033003023221111-2210030110320223-0302223312133103-2322223003132333-0201320122103032"></a>

<a id="canonical-0033130113110130-0333003110333303-0201000022220023-3113013003030221-1332102323203130-1021200031202001-0311121033302110-1133132213100000"></a>

## description_spec property — metadata / 220021230310 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-1012311132330032-1221013002203233-1320023010303202-2000333021010233-1100220011332302-3310311022212231-2120321313300213-1201022222130110"></a>

<a id="canonical-3120232331312103-1010301013300230-3320113033221003-1021332332002100-0203333120131032-1201220001221230-2203020020113013-2030011021022011"></a>

## name property — metadata / 220021230310 / 5

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

<a id="canonical-0101312230203321-1002032032201211-2203200332130303-2230130013220120-0203123200122321-2203011222313321-3331201222003322-0232322120233122"></a>

## Next pages — metadata / 220021230310 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111002121123312-1000221221220312-2221031022121332-1333233030313132-3120322030023011-2110020000220231-2021001100122321-1221213303030000"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode — validation_mode / 120202123120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-0302201132232310-1221020030010201-3130120320020313-2113323021310320-2123102012301111-0300102000133113-1200123211012223-1112112130113322"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-2203103001030210-0130011021021021-1211231130111020-1101310101023001-0203231212223031-0123000301120301-0132332232033021-0102210223101220"></a>

## Direct properties — validation_mode / 120202123120 / 3

- [response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031): complete subsection reference.

- [skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1302200011232010-0013222133002230-3100231310013020-2321011001222332-3201031112223021-2201303112123111-1212023111213232-2023100332100000): complete subsection reference.

- [skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-2310311021130102-0323202312023333-2022330123233221-3301023113303131-2211001110233320-0012111321303123-0201101211000023-1102020022202311): complete subsection reference.

- [validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2023111302021331-1113332023011333-1212300333110001-2021301221200110-2332102322202321-3022002032232312-2132200003302221-3300120113013013): complete subsection reference.

<a id="canonical-0032010122113120-0103333233232200-3132121013321231-3112033201300213-2313222100320223-3213331232013100-1110310122031212-2201001030130230"></a>

## Next pages — validation_mode / 120202123120 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-1302200011232010-0013222133002230-3100231310013020-2321011001222332-3201031112223021-2201303112123111-1212023111213232-2023100332100000)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](data-sources--http_loadbalancer--reference--group-009.md#canonical-2310311021130102-0323202312023333-2022330123233221-3301023113303131-2211001110233320-0012111321303123-0201101211000023-1102020022202311)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-2023111302021331-1113332023011333-1212300333110001-2021301221200110-2332102322202321-3022002032232312-2132200003302221-3300120113013013)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001032121122221-2312233020212120-3011210203000011-0213001221323320-0121102330022112-3302133030210123-0012213313122210-1212012131110332"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active — response_validation_mode_active / 221303030011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-3031021030132220-3330001303203000-0132011311033133-1232000122102301-3001020033012223-3201230003032311-2300033010000132-3323301022123301"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-2120211202022122-1003121211321021-3323212022221322-0303001321300222-1133122331332210-2221130122213011-1033001323033121-0313313131202232"></a>

## Direct properties — response_validation_mode_active / 221303030011 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3320130123230202-3231311203001231-3303010223103300-3130302212231003-1333311003000212-0300023020012013-3332131301111031-1322233220212031): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3112223332010020-3213123032002132-1223032023011320-3111123213311031-1310312100002032-0121231000300302-3111231321332221-2300223322013330): complete subsection reference.

<a id="canonical-1200110030113333-1002032222222331-1232230000322330-2122232322003232-3320131223333303-2100221232033320-1012322111001311-1213112211111221"></a>

<a id="canonical-3123202303131311-3113302231313111-2200330313021321-0212302102231303-3310013013230123-3311110101231331-2323102300030031-1311022200112003"></a>

## response_validation_properties property — response_validation_mode_active / 221303030011 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2303023110331332-1113031030003312-1321132033032331-0003030030033310-1331213022332003-2020102230112001-3202311200132123-2231012333103133"></a>

## Next pages — response_validation_mode_active / 221303030011 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-3320130123230202-3231311203001231-3303010223103300-3130302212231003-1333311003000212-0300023020012013-3332131301111031-1322233220212031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-009.md#canonical-3112223332010020-3213123032002132-1223032023011320-3111123213311031-1310312100002032-0121231000300302-3111231321332221-2300223322013330)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3320130123230202-3231311203001231-3303010223103300-3130302212231003-1333311003000212-0300023020012013-3332131301111031-1322233220212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233000020332100-3010110323102103-3123223332002333-2001031100132012-1131313121023313-3030021212003312-2213100320211102-3201130223200310"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block — enforcement_block / 121133021130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-0133002200131100-1323100032101333-1120310130322011-3311022221211023-2111300203010323-3312013132102101-0321132220332101-0301321132201322"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1000310211112132-3230110121211211-3322211301301003-0222013311221101-2123013023201130-1021003302221133-3121211233100133-0220131331102331"></a>

## Direct properties — enforcement_block / 121133021130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313213210230030-2113101323103303-3232120110213213-0203211203002322-2001203330311201-3201112310332100-2100213321320113-3012301311201300"></a>

## Next pages — enforcement_block / 121133021130 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3112223332010020-3213123032002132-1223032023011320-3111123213311031-1310312100002032-0121231000300302-3111231321332221-2300223322013330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231211323221230-2202332011330230-3111000232201021-1233320321233212-1301130221332333-0313303110302331-1130230313201212-1123230233002030"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report — enforcement_report / 111200300113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-1010030200031311-1031221333221231-0201211201300032-0110121303332131-0210123333212133-0232002232232003-3001321301232220-3210031121130020"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0332222010220213-2121220020312122-2310232110032111-1213103221202313-1321123013311322-0230303021110222-3012003023120203-0221022013100332"></a>

## Direct properties — enforcement_report / 111200300113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320300201000020-0320200133003210-3320232221311000-2010210333032231-0112122210303233-1213023123031110-3331023023230200-3300112202232210"></a>

## Next pages — enforcement_report / 111200300113 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--http_loadbalancer--reference--group-009.md#canonical-0120020103203202-0132212232303202-3222222013030311-0312000132223313-1100302133132022-3221021220022333-2122321120330001-1310223303122031)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1302200011232010-0013222133002230-3100231310013020-2321011001222332-3201031112223021-2201303112123111-1212023111213232-2023100332100000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121031330031020-1132321312212113-0331302311023333-0303210133202102-0230221112301321-0303202003020301-2202111213230202-2222021320013222"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation — skip_response_validation / 101020202121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-3301321131131022-1001102300333302-3313130021132220-0000321113223132-2332130103101002-3212302030031032-2101023202002210-3321310133120310"></a>

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

<a id="canonical-1213021011210012-0332113002202200-3113313010232112-2201310230233212-0323330221013010-2202002001222211-3100300021323313-3012021102331330"></a>

## Direct properties — skip_response_validation / 101020202121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320013123230120-0123000023222303-1131112213002031-0122231123010123-2111111133102200-2300030321010321-1021331101103133-1231132120301031"></a>

## Next pages — skip_response_validation / 101020202121 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2310311021130102-0323202312023333-2022330123233221-3301023113303131-2211001110233320-0012111321303123-0201101211000023-1102020022202311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130012211133032-0101120301312012-3301302332310100-2121022030320300-1333223333221002-1300233231133001-0121301002033011-0011233031101310"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation — skip_validation / 211102232022 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-2210322201022103-0303330210321332-3311302201022130-2112311203001103-0321130301133110-1223113030230201-2212312130333321-3112331312230322"></a>

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

<a id="canonical-0202001131001311-1103211132111300-2002220230312020-2201331122132120-3321030011023310-3100100333212310-1231033013231202-1303012202123222"></a>

## Direct properties — skip_validation / 211102232022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030100201012300-0203011212310002-2111023203332000-2010013013301101-1230202011303030-1102012100300120-3223123211120032-1032102130012201"></a>

## Next pages — skip_validation / 211102232022 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2023111302021331-1113332023011333-1212300333110001-2021301221200110-2332102322202321-3022002032232312-2132200003302221-3300120113013013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131330113030303-0213310120011023-2022110122202331-2110021323103132-0320001001221320-2211222102032320-2020211213320201-1021120220003032"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active — validation_mode_active / 033033302212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [api_specification](data-sources--http_loadbalancer--reference--group-008.md#canonical-3213330003311032-3113000030321002-3232132031132230-0311031121330011-0010303120101320-2333133223203001-3333321323020113-1022221301231303)
- [api_specification.validation_custom_list](data-sources--http_loadbalancer--reference--group-009.md#canonical-3130033232020110-3313122322230220-0013010322121330-1010221001222212-0303122303321323-0321120302100110-2311221301023003-3212020130112020)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--http_loadbalancer--reference--group-009.md#canonical-3101013010013321-1130302112100013-2323101210330031-2232333201223122-2323122330331133-1111203123230212-3131211313011201-0212322203302231)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-3131300132013203-1003223122111330-0331102021113120-0003312332212132-1033132021233333-3021000231202133-0302111123030113-2002130121121101"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-0130200232133110-2302332323303001-1123131112002000-0133212011332032-3003032233111310-1320002112000233-1122232103020213-0313103132301223"></a>

## Direct properties — validation_mode_active / 033033302212 / 3

- [enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-0210311001303313-1001312311331032-0332330130130132-2302301203111020-2221123121321123-3131013112220322-0233023022223210-0323221023323211): complete subsection reference.

- [enforcement_report](data-sources--http_loadbalancer--reference--group-010.md#canonical-0313111020222232-3133322131310122-0021013232011122-2231222330233303-0221031112230210-0111132310012211-0222123233123330-0203223310020311): complete subsection reference.

<a id="canonical-1310010130131320-2131312000223332-0303210013332021-2212030102103330-1023221111310033-2301110121333310-2011202013100322-2031320330122010"></a>

<a id="canonical-3130120103222231-1133322131122030-0301111232200322-0101303030212122-1002011022123323-2022013023212203-3233010202232303-1221011200133113"></a>

## request_validation_properties property — validation_mode_active / 033033302212 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0202101230002301-3133321330131333-3122223313330322-1232312232030222-1003003121311012-1323122132103030-3032310330310101-3110201300022122"></a>

## Next pages — validation_mode_active / 033033302212 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](data-sources--http_loadbalancer--reference--group-009.md#canonical-0210311001303313-1001312311331032-0332330130130132-2302301203111020-2221123121321123-3131013112220322-0233023022223210-0323221023323211)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](data-sources--http_loadbalancer--reference--group-010.md#canonical-0313111020222232-3133322131310122-0021013232011122-2231222330233303-0221031112230210-0111132310012211-0222123233123330-0203223310020311)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--http_loadbalancer--reference--group-009.md#canonical-1300302231130211-2212323210111301-0203032301002113-1220131330213112-2303310322203133-1220333313321113-0230020333101032-3022030323011130)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0210311001303313-1001312311331032-0332330130130132-2302301203111020-2221123121321123-3131013112220322-0233023022223210-0323221023323211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
